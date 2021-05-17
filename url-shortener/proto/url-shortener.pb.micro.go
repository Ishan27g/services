// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: proto/url-shortener.proto

package urlshortener

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

// Api Endpoints for UrlShortener service

func NewUrlShortenerEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for UrlShortener service

type UrlShortenerService interface {
	Shorten(ctx context.Context, in *ShortenRequest, opts ...client.CallOption) (*ShortenResponse, error)
	List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error)
	Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error)
}

type urlShortenerService struct {
	c    client.Client
	name string
}

func NewUrlShortenerService(name string, c client.Client) UrlShortenerService {
	return &urlShortenerService{
		c:    c,
		name: name,
	}
}

func (c *urlShortenerService) Shorten(ctx context.Context, in *ShortenRequest, opts ...client.CallOption) (*ShortenResponse, error) {
	req := c.c.NewRequest(c.name, "UrlShortener.Shorten", in)
	out := new(ShortenResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerService) List(ctx context.Context, in *ListRequest, opts ...client.CallOption) (*ListResponse, error) {
	req := c.c.NewRequest(c.name, "UrlShortener.List", in)
	out := new(ListResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *urlShortenerService) Get(ctx context.Context, in *GetRequest, opts ...client.CallOption) (*GetResponse, error) {
	req := c.c.NewRequest(c.name, "UrlShortener.Get", in)
	out := new(GetResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for UrlShortener service

type UrlShortenerHandler interface {
	Shorten(context.Context, *ShortenRequest, *ShortenResponse) error
	List(context.Context, *ListRequest, *ListResponse) error
	Get(context.Context, *GetRequest, *GetResponse) error
}

func RegisterUrlShortenerHandler(s server.Server, hdlr UrlShortenerHandler, opts ...server.HandlerOption) error {
	type urlShortener interface {
		Shorten(ctx context.Context, in *ShortenRequest, out *ShortenResponse) error
		List(ctx context.Context, in *ListRequest, out *ListResponse) error
		Get(ctx context.Context, in *GetRequest, out *GetResponse) error
	}
	type UrlShortener struct {
		urlShortener
	}
	h := &urlShortenerHandler{hdlr}
	return s.Handle(s.NewHandler(&UrlShortener{h}, opts...))
}

type urlShortenerHandler struct {
	UrlShortenerHandler
}

func (h *urlShortenerHandler) Shorten(ctx context.Context, in *ShortenRequest, out *ShortenResponse) error {
	return h.UrlShortenerHandler.Shorten(ctx, in, out)
}

func (h *urlShortenerHandler) List(ctx context.Context, in *ListRequest, out *ListResponse) error {
	return h.UrlShortenerHandler.List(ctx, in, out)
}

func (h *urlShortenerHandler) Get(ctx context.Context, in *GetRequest, out *GetResponse) error {
	return h.UrlShortenerHandler.Get(ctx, in, out)
}
