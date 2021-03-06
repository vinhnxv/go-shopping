// Code generated by protoc-gen-micro. DO NOT EDIT.
// source: shipping.proto

package shipping

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	context "context"
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
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
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Shipping service

type ShippingService interface {
	GetShippingCost(ctx context.Context, in *ShippingCostRequest, opts ...client.CallOption) (*ShippingCostResponse, error)
	MarkItemShipped(ctx context.Context, in *MarkShippedRequest, opts ...client.CallOption) (*MarkShippedResponse, error)
	GetShippingStatus(ctx context.Context, in *ShippingStatusRequest, opts ...client.CallOption) (*ShippingStatusResponse, error)
}

type shippingService struct {
	c    client.Client
	name string
}

func NewShippingService(name string, c client.Client) ShippingService {
	if c == nil {
		c = client.NewClient()
	}
	if len(name) == 0 {
		name = "shipping"
	}
	return &shippingService{
		c:    c,
		name: name,
	}
}

func (c *shippingService) GetShippingCost(ctx context.Context, in *ShippingCostRequest, opts ...client.CallOption) (*ShippingCostResponse, error) {
	req := c.c.NewRequest(c.name, "Shipping.GetShippingCost", in)
	out := new(ShippingCostResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) MarkItemShipped(ctx context.Context, in *MarkShippedRequest, opts ...client.CallOption) (*MarkShippedResponse, error) {
	req := c.c.NewRequest(c.name, "Shipping.MarkItemShipped", in)
	out := new(MarkShippedResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *shippingService) GetShippingStatus(ctx context.Context, in *ShippingStatusRequest, opts ...client.CallOption) (*ShippingStatusResponse, error) {
	req := c.c.NewRequest(c.name, "Shipping.GetShippingStatus", in)
	out := new(ShippingStatusResponse)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Shipping service

type ShippingHandler interface {
	GetShippingCost(context.Context, *ShippingCostRequest, *ShippingCostResponse) error
	MarkItemShipped(context.Context, *MarkShippedRequest, *MarkShippedResponse) error
	GetShippingStatus(context.Context, *ShippingStatusRequest, *ShippingStatusResponse) error
}

func RegisterShippingHandler(s server.Server, hdlr ShippingHandler, opts ...server.HandlerOption) error {
	type shipping interface {
		GetShippingCost(ctx context.Context, in *ShippingCostRequest, out *ShippingCostResponse) error
		MarkItemShipped(ctx context.Context, in *MarkShippedRequest, out *MarkShippedResponse) error
		GetShippingStatus(ctx context.Context, in *ShippingStatusRequest, out *ShippingStatusResponse) error
	}
	type Shipping struct {
		shipping
	}
	h := &shippingHandler{hdlr}
	return s.Handle(s.NewHandler(&Shipping{h}, opts...))
}

type shippingHandler struct {
	ShippingHandler
}

func (h *shippingHandler) GetShippingCost(ctx context.Context, in *ShippingCostRequest, out *ShippingCostResponse) error {
	return h.ShippingHandler.GetShippingCost(ctx, in, out)
}

func (h *shippingHandler) MarkItemShipped(ctx context.Context, in *MarkShippedRequest, out *MarkShippedResponse) error {
	return h.ShippingHandler.MarkItemShipped(ctx, in, out)
}

func (h *shippingHandler) GetShippingStatus(ctx context.Context, in *ShippingStatusRequest, out *ShippingStatusResponse) error {
	return h.ShippingHandler.GetShippingStatus(ctx, in, out)
}
