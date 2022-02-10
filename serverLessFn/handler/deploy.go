package handler

import "net/http"

func deploy(fn interface{}) {

	// todo call functionsFramework

	fnHTTP, ok := fn.(func(http.ResponseWriter, *http.Request))
	if !ok {
		panic("expected function to have signature func(http.ResponseWriter, *http.Request)")
	}

	fnHTTP(nil, nil)
}

// WrapDeploy definition gets auto-generated
func WrapDeploy() {
}
