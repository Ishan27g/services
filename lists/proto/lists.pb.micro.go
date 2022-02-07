// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/lists.proto

package lists

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/micro/v3/service/api"
	client "github.com/micro/micro/v3/service/client"
	server "github.com/micro/micro/v3/service/server"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// Reference imports to suppress errors if they are not otherwise used.
var _ api.Endpoint
var _ context.Context
var _ client.Option
var _ server.Option

// Api Endpoints for Lists service

func NewListsEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Lists service

type ListsService interface {
	New(ctx context.Context, in *List, opts ...client.CallOption) (*Response, error)
	Update(ctx context.Context, in *List, opts ...client.CallOption) (*Response, error)
	Delete(ctx context.Context, in *GetOptions, opts ...client.CallOption) (*Response, error)
	Get(ctx context.Context, in *GetOptions, opts ...client.CallOption) (*AllLists, error)
}

type listsService struct {
	c    client.Client
	name string
}

func NewListsService(name string, c client.Client) ListsService {
	return &listsService{
		c:    c,
		name: name,
	}
}

func (c *listsService) New(ctx context.Context, in *List, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Lists.New", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Update(ctx context.Context, in *List, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Lists.Update", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Delete(ctx context.Context, in *GetOptions, opts ...client.CallOption) (*Response, error) {
	req := c.c.NewRequest(c.name, "Lists.Delete", in)
	out := new(Response)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *listsService) Get(ctx context.Context, in *GetOptions, opts ...client.CallOption) (*AllLists, error) {
	req := c.c.NewRequest(c.name, "Lists.Get", in)
	out := new(AllLists)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Lists service

type ListsHandler interface {
	New(context.Context, *List, *Response) error
	Update(context.Context, *List, *Response) error
	Delete(context.Context, *GetOptions, *Response) error
	Get(context.Context, *GetOptions, *AllLists) error
}

func RegisterListsHandler(s server.Server, hdlr ListsHandler, opts ...server.HandlerOption) error {
	type lists interface {
		New(ctx context.Context, in *List, out *Response) error
		Update(ctx context.Context, in *List, out *Response) error
		Delete(ctx context.Context, in *GetOptions, out *Response) error
		Get(ctx context.Context, in *GetOptions, out *AllLists) error
	}
	type Lists struct {
		lists
	}
	h := &listsHandler{hdlr}
	return s.Handle(s.NewHandler(&Lists{h}, opts...))
}

type listsHandler struct {
	ListsHandler
}

func (h *listsHandler) New(ctx context.Context, in *List, out *Response) error {
	return h.ListsHandler.New(ctx, in, out)
}

func (h *listsHandler) Update(ctx context.Context, in *List, out *Response) error {
	return h.ListsHandler.Update(ctx, in, out)
}

func (h *listsHandler) Delete(ctx context.Context, in *GetOptions, out *Response) error {
	return h.ListsHandler.Delete(ctx, in, out)
}

func (h *listsHandler) Get(ctx context.Context, in *GetOptions, out *AllLists) error {
	return h.ListsHandler.Get(ctx, in, out)
}
