// Code generated by thriftrw-plugin-yarpc
// @generated

package storeserver

import (
	context "context"
	stream "go.uber.org/thriftrw/protocol/stream"
	wire "go.uber.org/thriftrw/wire"
	transport "go.uber.org/yarpc/api/transport"
	thrift "go.uber.org/yarpc/encoding/thrift"
	atomic "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic"
	readonlystoreserver "go.uber.org/yarpc/encoding/thrift/thriftrw-plugin-yarpc/internal/tests/atomic/readonlystoreserver"
	yarpcerrors "go.uber.org/yarpc/yarpcerrors"
)

// Interface is the server-side interface for the Store service.
type Interface interface {
	readonlystoreserver.Interface

	CompareAndSwap(
		ctx context.Context,
		Request *atomic.CompareAndSwap,
	) error

	Forget(
		ctx context.Context,
		Key *string,
	) error

	Increment(
		ctx context.Context,
		Key *string,
		Value *int64,
	) error
}

// New prepares an implementation of the Store service for
// registration.
//
// 	handler := StoreHandler{}
// 	dispatcher.Register(storeserver.New(handler))
func New(impl Interface, opts ...thrift.RegisterOption) []transport.Procedure {
	h := handler{impl}
	service := thrift.Service{
		Name: "Store",
		Methods: []thrift.Method{

			thrift.Method{
				Name: "compareAndSwap",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Unary,
					Unary:  thrift.UnaryHandler(h.CompareAndSwap),
					NoWire: compareandswap_NoWireHandler{impl},
				},
				Signature:    "CompareAndSwap(Request *atomic.CompareAndSwap)",
				ThriftModule: atomic.ThriftModule,
			},

			thrift.Method{
				Name: "forget",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Oneway,
					Oneway: thrift.OnewayHandler(h.Forget),
					NoWire: forget_NoWireHandler{impl},
				},
				Signature:    "Forget(Key *string)",
				ThriftModule: atomic.ThriftModule,
			},

			thrift.Method{
				Name: "increment",
				HandlerSpec: thrift.HandlerSpec{

					Type:   transport.Unary,
					Unary:  thrift.UnaryHandler(h.Increment),
					NoWire: increment_NoWireHandler{impl},
				},
				Signature:    "Increment(Key *string, Value *int64)",
				ThriftModule: atomic.ThriftModule,
			},
		},
	}

	procedures := make([]transport.Procedure, 0, 3)

	procedures = append(
		procedures,
		readonlystoreserver.New(
			impl,
			append(
				opts,
				thrift.Named("Store"),
			)...,
		)...,
	)
	procedures = append(procedures, thrift.BuildProcedures(service, opts...)...)
	return procedures
}

type handler struct{ impl Interface }

type yarpcErrorNamer interface{ YARPCErrorName() string }

type yarpcErrorCoder interface{ YARPCErrorCode() *yarpcerrors.Code }

func (h handler) CompareAndSwap(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args atomic.Store_CompareAndSwap_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'Store' procedure 'CompareAndSwap': %w", err)
	}

	appErr := h.impl.CompareAndSwap(ctx, args.Request)

	hadError := appErr != nil
	result, err := atomic.Store_CompareAndSwap_Helper.WrapResponse(appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}

	return response, err
}

func (h handler) Forget(ctx context.Context, body wire.Value) error {
	var args atomic.Store_Forget_Args
	if err := args.FromWire(body); err != nil {
		return err
	}

	return h.impl.Forget(ctx, args.Key)
}

func (h handler) Increment(ctx context.Context, body wire.Value) (thrift.Response, error) {
	var args atomic.Store_Increment_Args
	if err := args.FromWire(body); err != nil {
		return thrift.Response{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode Thrift request for service 'Store' procedure 'Increment': %w", err)
	}

	appErr := h.impl.Increment(ctx, args.Key, args.Value)

	hadError := appErr != nil
	result, err := atomic.Store_Increment_Helper.WrapResponse(appErr)

	var response thrift.Response
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}

	return response, err
}

type compareandswap_NoWireHandler struct{ impl Interface }

func (h compareandswap_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args atomic.Store_CompareAndSwap_Args
		rw   stream.ResponseWriter
		err  error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)
	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'Store' procedure 'CompareAndSwap': %w", err)
	}

	appErr := h.impl.CompareAndSwap(ctx, args.Request)

	hadError := appErr != nil
	result, err := atomic.Store_CompareAndSwap_Helper.WrapResponse(appErr)
	response := thrift.NoWireResponse{ResponseWriter: rw}
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}
	return response, err

}

type forget_NoWireHandler struct{ impl Interface }

func (h forget_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args atomic.Store_Forget_Args

		err error
	)

	if _, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args); err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'Store' procedure 'Forget': %w", err)
	}

	return thrift.NoWireResponse{}, h.impl.Forget(ctx, args.Key)

}

type increment_NoWireHandler struct{ impl Interface }

func (h increment_NoWireHandler) HandleNoWire(ctx context.Context, nwc *thrift.NoWireCall) (thrift.NoWireResponse, error) {
	var (
		args atomic.Store_Increment_Args
		rw   stream.ResponseWriter
		err  error
	)

	rw, err = nwc.RequestReader.ReadRequest(ctx, nwc.EnvelopeType, nwc.Reader, &args)
	if err != nil {
		return thrift.NoWireResponse{}, yarpcerrors.InvalidArgumentErrorf(
			"could not decode (via no wire) Thrift request for service 'Store' procedure 'Increment': %w", err)
	}

	appErr := h.impl.Increment(ctx, args.Key, args.Value)

	hadError := appErr != nil
	result, err := atomic.Store_Increment_Helper.WrapResponse(appErr)
	response := thrift.NoWireResponse{ResponseWriter: rw}
	if err == nil {
		response.IsApplicationError = hadError
		response.Body = result
		if namer, ok := appErr.(yarpcErrorNamer); ok {
			response.ApplicationErrorName = namer.YARPCErrorName()
		}
		if extractor, ok := appErr.(yarpcErrorCoder); ok {
			response.ApplicationErrorCode = extractor.YARPCErrorCode()
		}
		if appErr != nil {
			response.ApplicationErrorDetails = appErr.Error()
		}
	}
	return response, err

}
