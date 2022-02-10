package handler

import (
	"bytes"
	"context"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"log"
	"strconv"

	serverLessFn "serverLessFn/proto"
)

type ServerLessFn struct{}

func (e *ServerLessFn) Deploy(ctx context.Context, fn *serverLessFn.Function, response *serverLessFn.FunctionResponse) error {
	const filename = "ast/hello/HelloWorld.go"
	const fnName = "HelloWorld"
	ReWrite(filename, fnName)

	response.Url = "/call" // todo functionsFramework integration

	return nil
}

func ReWrite(fileName string, fnName string) {
	set := token.NewFileSet()
	node, err := parser.ParseFile(set, fileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}
	validateParams := func(node *ast.File) (bool, string) {
		formatNode := func(node ast.Node) string {
			buf := new(bytes.Buffer)
			_ = format.Node(buf, token.NewFileSet(), node)
			return buf.String()
		}
		valid := false
		pkgName := ""
		ast.Inspect(node, func(n ast.Node) bool {
			switch ret := n.(type) {
			case *ast.FuncDecl:
				params := ret.Type.Params.List
				// we are only interested in functions with exactly 2 parameters
				if len(params) == 2 {
					firstParameterIsW := formatNode(params[0].Names[0]) == "w" && formatNode(params[0].Type) == "http.ResponseWriter"
					secondParameterIsR := formatNode(params[1].Names[0]) == "r" && formatNode(params[1].Type) == "*http.Request"
					if firstParameterIsW && secondParameterIsR {
						valid = true
					}
				}
			case *ast.File:
				pkgName = ret.Name.String()
			}
			return true
		})
		return valid, pkgName
	}
	valid, pkgName := validateParams(node)
	if valid {
		fmt.Println(valid, pkgName)
		rewriteDeployDotGo(pkgName, fnName)
	}
}

func rewriteDeployDotGo(pkgName string, fnName string) {
	const modFileName = "ast/deploy.go"

	set := token.NewFileSet()
	node, err := parser.ParseFile(set, modFileName, nil, parser.ParseComments)
	if err != nil {
		fmt.Println(err.Error())
		log.Fatal(err)
	}

	for i := 0; i < len(node.Decls); i++ {
		d := node.Decls[i]
		switch d.(type) {
		case *ast.GenDecl:
			dd := d.(*ast.GenDecl)
			if dd.Tok == token.IMPORT {
				// Add the new import
				iSpec := &ast.ImportSpec{Path: &ast.BasicLit{Value: strconv.Quote("ast/" + pkgName)}}
				dd.Specs = append(dd.Specs, iSpec)
			}
		case *ast.FuncDecl:
			if d.(*ast.FuncDecl).Name.String() == "WrapDeploy" {
				// Add the new function call with relevant package
				newCallStmt := &ast.ExprStmt{
					X: &ast.CallExpr{
						Fun: &ast.Ident{
							Name: "deploy",
						},
						Args: []ast.Expr{
							&ast.BasicLit{
								Kind:  token.STRING,
								Value: pkgName + `.` + fnName,
							},
						},
					},
				}
				d.(*ast.FuncDecl).Body.List = append([]ast.Stmt{newCallStmt}, d.(*ast.FuncDecl).Body.List...)
			}
		}
	}

	// Sort the imports
	ast.SortImports(set, node)

	// Generate the code
	out, err := GenerateFile(set, node)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(string(out))
	err = ioutil.WriteFile(modFileName, out, 0777)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
}

func GenerateFile(fset *token.FileSet, file *ast.File) ([]byte, error) {
	var output []byte
	buffer := bytes.NewBuffer(output)
	if err := printer.Fprint(buffer, fset, file); err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}
