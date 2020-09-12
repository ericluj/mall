// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: product.proto

package product

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	api "github.com/micro/go-micro/v2/api"
	client "github.com/micro/go-micro/v2/client"
	server "github.com/micro/go-micro/v2/server"
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

// Api Endpoints for Product service

func NewProductEndpoints() []*api.Endpoint {
	return []*api.Endpoint{}
}

// Client API for Product service

type ProductService interface {
	CreateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error)
	QueryProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error)
	UpdateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error)
}

type productService struct {
	c    client.Client
	name string
}

func NewProductService(name string, c client.Client) ProductService {
	return &productService{
		c:    c,
		name: name,
	}
}

func (c *productService) CreateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Product.CreateProduct", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) QueryProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Product.QueryProduct", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *productService) UpdateProduct(ctx context.Context, in *ProductRequest, opts ...client.CallOption) (*ProductResponse, error) {
	req := c.c.NewRequest(c.name, "Product.UpdateProduct", in)
	out := new(ProductResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Product service

type ProductHandler interface {
	CreateProduct(context.Context, *ProductRequest, *ProductResponse) error
	QueryProduct(context.Context, *ProductRequest, *ProductResponse) error
	UpdateProduct(context.Context, *ProductRequest, *ProductResponse) error
}

func RegisterProductHandler(s server.Server, hdlr ProductHandler, opts ...server.HandlerOption) error {
	type product interface {
		CreateProduct(ctx context.Context, in *ProductRequest, out *ProductResponse) error
		QueryProduct(ctx context.Context, in *ProductRequest, out *ProductResponse) error
		UpdateProduct(ctx context.Context, in *ProductRequest, out *ProductResponse) error
	}
	type Product struct {
		product
	}
	h := &productHandler{hdlr}
	return s.Handle(s.NewHandler(&Product{h}, opts...))
}

type productHandler struct {
	ProductHandler
}

func (h *productHandler) CreateProduct(ctx context.Context, in *ProductRequest, out *ProductResponse) error {
	return h.ProductHandler.CreateProduct(ctx, in, out)
}

func (h *productHandler) QueryProduct(ctx context.Context, in *ProductRequest, out *ProductResponse) error {
	return h.ProductHandler.QueryProduct(ctx, in, out)
}

func (h *productHandler) UpdateProduct(ctx context.Context, in *ProductRequest, out *ProductResponse) error {
	return h.ProductHandler.UpdateProduct(ctx, in, out)
}