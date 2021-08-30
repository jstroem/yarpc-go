// Code generated by thriftrw-plugin-yarpc
// @generated

package nameclient

import (
	context "context"
	wire "go.uber.org/thriftrw/wire"
	yarpc "go.uber.org/yarpc"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	extends "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/extends"
	reflect "reflect"
)

// Interface is a client for the Name service.
type Interface interface {
	Name(
		ctx context.Context,
		opts ...yarpc.CallOption,
	) (string, error)
}

// New builds a new client for the Name service.
//
// 	client := nameclient.New(dispatcher.ClientConfig("name"))
func New(c transport.ClientConfig, opts ...thrift.ClientOption) Interface {
	return client{
		c: thrift.New(thrift.Config{
			Service:      "Name",
			ClientConfig: c,
		}, opts...),
		nwc: thrift.NewNoWire(thrift.Config{
			Service:      "Name",
			ClientConfig: c,
		}, opts...),
	}
}

func init() {
	yarpc.RegisterClientBuilder(
		func(c transport.ClientConfig, f reflect.StructField) Interface {
			return New(c, thrift.ClientBuilderOptions(c, f)...)
		},
	)
}

type client struct {
	c   thrift.Client
	nwc thrift.NoWireClient
}

func (c client) Name(
	ctx context.Context,
	opts ...yarpc.CallOption,
) (success string, err error) {

	var result extends.Name_Name_Result
	args := extends.Name_Name_Helper.Args()

	if c.nwc != nil && c.nwc.Enabled() {
		if err = c.nwc.Call(ctx, args, &result, opts...); err != nil {
			return
		}
	} else {
		var body wire.Value
		if body, err = c.c.Call(ctx, args, opts...); err != nil {
			return
		}

		if err = result.FromWire(body); err != nil {
			return
		}
	}

	success, err = extends.Name_Name_Helper.UnwrapResponse(&result)
	return
}
