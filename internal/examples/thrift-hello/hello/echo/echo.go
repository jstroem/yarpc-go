// Code generated by thriftrw v1.29.2. DO NOT EDIT.
// @generated

// Copyright (c) 2022 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package echo

import (
	errors "errors"
	fmt "fmt"
	multierr "go.uber.org/multierr"
	stream "go.uber.org/thriftrw/protocol/stream"
	thriftreflect "go.uber.org/thriftrw/thriftreflect"
	wire "go.uber.org/thriftrw/wire"
	zapcore "go.uber.org/zap/zapcore"
	strings "strings"
)

type EchoRequest struct {
	Message string `json:"message,required"`
	Count   int16  `json:"count,required"`
}

// ToWire translates a EchoRequest struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *EchoRequest) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueI16(v.Count), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a EchoRequest struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a EchoRequest struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v EchoRequest
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *EchoRequest) FromWire(w wire.Value) error {
	var err error

	messageIsSet := false
	countIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI16 {
				v.Count, err = field.Value.GetI16(), error(nil)
				if err != nil {
					return err
				}
				countIsSet = true
			}
		}
	}

	if !messageIsSet {
		return errors.New("field Message of EchoRequest is required")
	}

	if !countIsSet {
		return errors.New("field Count of EchoRequest is required")
	}

	return nil
}

// Encode serializes a EchoRequest struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a EchoRequest struct could not be encoded.
func (v *EchoRequest) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Message); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TI16}); err != nil {
		return err
	}
	if err := sw.WriteInt16(v.Count); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a EchoRequest struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a EchoRequest struct could not be generated from the wire
// representation.
func (v *EchoRequest) Decode(sr stream.Reader) error {

	messageIsSet := false
	countIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Message, err = sr.ReadString()
			if err != nil {
				return err
			}
			messageIsSet = true
		case fh.ID == 2 && fh.Type == wire.TI16:
			v.Count, err = sr.ReadInt16()
			if err != nil {
				return err
			}
			countIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !messageIsSet {
		return errors.New("field Message of EchoRequest is required")
	}

	if !countIsSet {
		return errors.New("field Count of EchoRequest is required")
	}

	return nil
}

// String returns a readable string representation of a EchoRequest
// struct.
func (v *EchoRequest) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	fields[i] = fmt.Sprintf("Count: %v", v.Count)
	i++

	return fmt.Sprintf("EchoRequest{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this EchoRequest match the
// provided EchoRequest.
//
// This function performs a deep comparison.
func (v *EchoRequest) Equals(rhs *EchoRequest) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message == rhs.Message) {
		return false
	}
	if !(v.Count == rhs.Count) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EchoRequest.
func (v *EchoRequest) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message", v.Message)
	enc.AddInt16("count", v.Count)
	return err
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *EchoRequest) GetMessage() (o string) {
	if v != nil {
		o = v.Message
	}
	return
}

// GetCount returns the value of Count if it is set or its
// zero value if it is unset.
func (v *EchoRequest) GetCount() (o int16) {
	if v != nil {
		o = v.Count
	}
	return
}

type EchoResponse struct {
	Message string `json:"message,required"`
	Count   int16  `json:"count,required"`
}

// ToWire translates a EchoResponse struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *EchoResponse) ToWire() (wire.Value, error) {
	var (
		fields [2]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	w, err = wire.NewValueString(v.Message), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 1, Value: w}
	i++

	w, err = wire.NewValueI16(v.Count), error(nil)
	if err != nil {
		return w, err
	}
	fields[i] = wire.Field{ID: 2, Value: w}
	i++

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

// FromWire deserializes a EchoResponse struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a EchoResponse struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v EchoResponse
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *EchoResponse) FromWire(w wire.Value) error {
	var err error

	messageIsSet := false
	countIsSet := false

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TBinary {
				v.Message, err = field.Value.GetString(), error(nil)
				if err != nil {
					return err
				}
				messageIsSet = true
			}
		case 2:
			if field.Value.Type() == wire.TI16 {
				v.Count, err = field.Value.GetI16(), error(nil)
				if err != nil {
					return err
				}
				countIsSet = true
			}
		}
	}

	if !messageIsSet {
		return errors.New("field Message of EchoResponse is required")
	}

	if !countIsSet {
		return errors.New("field Count of EchoResponse is required")
	}

	return nil
}

// Encode serializes a EchoResponse struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a EchoResponse struct could not be encoded.
func (v *EchoResponse) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TBinary}); err != nil {
		return err
	}
	if err := sw.WriteString(v.Message); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 2, Type: wire.TI16}); err != nil {
		return err
	}
	if err := sw.WriteInt16(v.Count); err != nil {
		return err
	}
	if err := sw.WriteFieldEnd(); err != nil {
		return err
	}

	return sw.WriteStructEnd()
}

// Decode deserializes a EchoResponse struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a EchoResponse struct could not be generated from the wire
// representation.
func (v *EchoResponse) Decode(sr stream.Reader) error {

	messageIsSet := false
	countIsSet := false

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TBinary:
			v.Message, err = sr.ReadString()
			if err != nil {
				return err
			}
			messageIsSet = true
		case fh.ID == 2 && fh.Type == wire.TI16:
			v.Count, err = sr.ReadInt16()
			if err != nil {
				return err
			}
			countIsSet = true
		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	if !messageIsSet {
		return errors.New("field Message of EchoResponse is required")
	}

	if !countIsSet {
		return errors.New("field Count of EchoResponse is required")
	}

	return nil
}

// String returns a readable string representation of a EchoResponse
// struct.
func (v *EchoResponse) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [2]string
	i := 0
	fields[i] = fmt.Sprintf("Message: %v", v.Message)
	i++
	fields[i] = fmt.Sprintf("Count: %v", v.Count)
	i++

	return fmt.Sprintf("EchoResponse{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this EchoResponse match the
// provided EchoResponse.
//
// This function performs a deep comparison.
func (v *EchoResponse) Equals(rhs *EchoResponse) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !(v.Message == rhs.Message) {
		return false
	}
	if !(v.Count == rhs.Count) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of EchoResponse.
func (v *EchoResponse) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	enc.AddString("message", v.Message)
	enc.AddInt16("count", v.Count)
	return err
}

// GetMessage returns the value of Message if it is set or its
// zero value if it is unset.
func (v *EchoResponse) GetMessage() (o string) {
	if v != nil {
		o = v.Message
	}
	return
}

// GetCount returns the value of Count if it is set or its
// zero value if it is unset.
func (v *EchoResponse) GetCount() (o int16) {
	if v != nil {
		o = v.Count
	}
	return
}

// ThriftModule represents the IDL file used to generate this package.
var ThriftModule = &thriftreflect.ThriftModule{
	Name:     "echo",
	Package:  "go.uber.org/yarpc/internal/examples/thrift-hello/hello/echo",
	FilePath: "echo.thrift",
	SHA1:     "619e87984e33d89e2f63c235bf6b0c1003898504",
	Raw:      rawIDL,
}

const rawIDL = "service Hello {\n    EchoResponse echo(1:EchoRequest echo)\n}\n\nstruct EchoRequest {\n    1: required string message;\n    2: required i16 count;\n}\n\nstruct EchoResponse {\n    1: required string message;\n    2: required i16 count;\n}\n"

// Hello_Echo_Args represents the arguments for the Hello.echo function.
//
// The arguments for echo are sent and received over the wire as this struct.
type Hello_Echo_Args struct {
	Echo *EchoRequest `json:"echo,omitempty"`
}

// ToWire translates a Hello_Echo_Args struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Hello_Echo_Args) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Echo != nil {
		w, err = v.Echo.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 1, Value: w}
		i++
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _EchoRequest_Read(w wire.Value) (*EchoRequest, error) {
	var v EchoRequest
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Hello_Echo_Args struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Hello_Echo_Args struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Hello_Echo_Args
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Hello_Echo_Args) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 1:
			if field.Value.Type() == wire.TStruct {
				v.Echo, err = _EchoRequest_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	return nil
}

// Encode serializes a Hello_Echo_Args struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Hello_Echo_Args struct could not be encoded.
func (v *Hello_Echo_Args) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Echo != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 1, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.Echo.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	return sw.WriteStructEnd()
}

func _EchoRequest_Decode(sr stream.Reader) (*EchoRequest, error) {
	var v EchoRequest
	err := v.Decode(sr)
	return &v, err
}

// Decode deserializes a Hello_Echo_Args struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Hello_Echo_Args struct could not be generated from the wire
// representation.
func (v *Hello_Echo_Args) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 1 && fh.Type == wire.TStruct:
			v.Echo, err = _EchoRequest_Decode(sr)
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	return nil
}

// String returns a readable string representation of a Hello_Echo_Args
// struct.
func (v *Hello_Echo_Args) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Echo != nil {
		fields[i] = fmt.Sprintf("Echo: %v", v.Echo)
		i++
	}

	return fmt.Sprintf("Hello_Echo_Args{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Hello_Echo_Args match the
// provided Hello_Echo_Args.
//
// This function performs a deep comparison.
func (v *Hello_Echo_Args) Equals(rhs *Hello_Echo_Args) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Echo == nil && rhs.Echo == nil) || (v.Echo != nil && rhs.Echo != nil && v.Echo.Equals(rhs.Echo))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Hello_Echo_Args.
func (v *Hello_Echo_Args) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Echo != nil {
		err = multierr.Append(err, enc.AddObject("echo", v.Echo))
	}
	return err
}

// GetEcho returns the value of Echo if it is set or its
// zero value if it is unset.
func (v *Hello_Echo_Args) GetEcho() (o *EchoRequest) {
	if v != nil && v.Echo != nil {
		return v.Echo
	}

	return
}

// IsSetEcho returns true if Echo is not nil.
func (v *Hello_Echo_Args) IsSetEcho() bool {
	return v != nil && v.Echo != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the arguments.
//
// This will always be "echo" for this struct.
func (v *Hello_Echo_Args) MethodName() string {
	return "echo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Call for this struct.
func (v *Hello_Echo_Args) EnvelopeType() wire.EnvelopeType {
	return wire.Call
}

// Hello_Echo_Helper provides functions that aid in handling the
// parameters and return values of the Hello.echo
// function.
var Hello_Echo_Helper = struct {
	// Args accepts the parameters of echo in-order and returns
	// the arguments struct for the function.
	Args func(
		echo *EchoRequest,
	) *Hello_Echo_Args

	// IsException returns true if the given error can be thrown
	// by echo.
	//
	// An error can be thrown by echo only if the
	// corresponding exception type was mentioned in the 'throws'
	// section for it in the Thrift file.
	IsException func(error) bool

	// WrapResponse returns the result struct for echo
	// given its return value and error.
	//
	// This allows mapping values and errors returned by
	// echo into a serializable result struct.
	// WrapResponse returns a non-nil error if the provided
	// error cannot be thrown by echo
	//
	//   value, err := echo(args)
	//   result, err := Hello_Echo_Helper.WrapResponse(value, err)
	//   if err != nil {
	//     return fmt.Errorf("unexpected error from echo: %v", err)
	//   }
	//   serialize(result)
	WrapResponse func(*EchoResponse, error) (*Hello_Echo_Result, error)

	// UnwrapResponse takes the result struct for echo
	// and returns the value or error returned by it.
	//
	// The error is non-nil only if echo threw an
	// exception.
	//
	//   result := deserialize(bytes)
	//   value, err := Hello_Echo_Helper.UnwrapResponse(result)
	UnwrapResponse func(*Hello_Echo_Result) (*EchoResponse, error)
}{}

func init() {
	Hello_Echo_Helper.Args = func(
		echo *EchoRequest,
	) *Hello_Echo_Args {
		return &Hello_Echo_Args{
			Echo: echo,
		}
	}

	Hello_Echo_Helper.IsException = func(err error) bool {
		switch err.(type) {
		default:
			return false
		}
	}

	Hello_Echo_Helper.WrapResponse = func(success *EchoResponse, err error) (*Hello_Echo_Result, error) {
		if err == nil {
			return &Hello_Echo_Result{Success: success}, nil
		}

		return nil, err
	}
	Hello_Echo_Helper.UnwrapResponse = func(result *Hello_Echo_Result) (success *EchoResponse, err error) {

		if result.Success != nil {
			success = result.Success
			return
		}

		err = errors.New("expected a non-void result")
		return
	}

}

// Hello_Echo_Result represents the result of a Hello.echo function call.
//
// The result of a echo execution is sent and received over the wire as this struct.
//
// Success is set only if the function did not throw an exception.
type Hello_Echo_Result struct {
	// Value returned by echo after a successful execution.
	Success *EchoResponse `json:"success,omitempty"`
}

// ToWire translates a Hello_Echo_Result struct into a Thrift-level intermediate
// representation. This intermediate representation may be serialized
// into bytes using a ThriftRW protocol implementation.
//
// An error is returned if the struct or any of its fields failed to
// validate.
//
//   x, err := v.ToWire()
//   if err != nil {
//     return err
//   }
//
//   if err := binaryProtocol.Encode(x, writer); err != nil {
//     return err
//   }
func (v *Hello_Echo_Result) ToWire() (wire.Value, error) {
	var (
		fields [1]wire.Field
		i      int = 0
		w      wire.Value
		err    error
	)

	if v.Success != nil {
		w, err = v.Success.ToWire()
		if err != nil {
			return w, err
		}
		fields[i] = wire.Field{ID: 0, Value: w}
		i++
	}

	if i != 1 {
		return wire.Value{}, fmt.Errorf("Hello_Echo_Result should have exactly one field: got %v fields", i)
	}

	return wire.NewValueStruct(wire.Struct{Fields: fields[:i]}), nil
}

func _EchoResponse_Read(w wire.Value) (*EchoResponse, error) {
	var v EchoResponse
	err := v.FromWire(w)
	return &v, err
}

// FromWire deserializes a Hello_Echo_Result struct from its Thrift-level
// representation. The Thrift-level representation may be obtained
// from a ThriftRW protocol implementation.
//
// An error is returned if we were unable to build a Hello_Echo_Result struct
// from the provided intermediate representation.
//
//   x, err := binaryProtocol.Decode(reader, wire.TStruct)
//   if err != nil {
//     return nil, err
//   }
//
//   var v Hello_Echo_Result
//   if err := v.FromWire(x); err != nil {
//     return nil, err
//   }
//   return &v, nil
func (v *Hello_Echo_Result) FromWire(w wire.Value) error {
	var err error

	for _, field := range w.GetStruct().Fields {
		switch field.ID {
		case 0:
			if field.Value.Type() == wire.TStruct {
				v.Success, err = _EchoResponse_Read(field.Value)
				if err != nil {
					return err
				}

			}
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Hello_Echo_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// Encode serializes a Hello_Echo_Result struct directly into bytes, without going
// through an intermediary type.
//
// An error is returned if a Hello_Echo_Result struct could not be encoded.
func (v *Hello_Echo_Result) Encode(sw stream.Writer) error {
	if err := sw.WriteStructBegin(); err != nil {
		return err
	}

	if v.Success != nil {
		if err := sw.WriteFieldBegin(stream.FieldHeader{ID: 0, Type: wire.TStruct}); err != nil {
			return err
		}
		if err := v.Success.Encode(sw); err != nil {
			return err
		}
		if err := sw.WriteFieldEnd(); err != nil {
			return err
		}
	}

	count := 0
	if v.Success != nil {
		count++
	}

	if count != 1 {
		return fmt.Errorf("Hello_Echo_Result should have exactly one field: got %v fields", count)
	}

	return sw.WriteStructEnd()
}

func _EchoResponse_Decode(sr stream.Reader) (*EchoResponse, error) {
	var v EchoResponse
	err := v.Decode(sr)
	return &v, err
}

// Decode deserializes a Hello_Echo_Result struct directly from its Thrift-level
// representation, without going through an intemediary type.
//
// An error is returned if a Hello_Echo_Result struct could not be generated from the wire
// representation.
func (v *Hello_Echo_Result) Decode(sr stream.Reader) error {

	if err := sr.ReadStructBegin(); err != nil {
		return err
	}

	fh, ok, err := sr.ReadFieldBegin()
	if err != nil {
		return err
	}

	for ok {
		switch {
		case fh.ID == 0 && fh.Type == wire.TStruct:
			v.Success, err = _EchoResponse_Decode(sr)
			if err != nil {
				return err
			}

		default:
			if err := sr.Skip(fh.Type); err != nil {
				return err
			}
		}

		if err := sr.ReadFieldEnd(); err != nil {
			return err
		}

		if fh, ok, err = sr.ReadFieldBegin(); err != nil {
			return err
		}
	}

	if err := sr.ReadStructEnd(); err != nil {
		return err
	}

	count := 0
	if v.Success != nil {
		count++
	}
	if count != 1 {
		return fmt.Errorf("Hello_Echo_Result should have exactly one field: got %v fields", count)
	}

	return nil
}

// String returns a readable string representation of a Hello_Echo_Result
// struct.
func (v *Hello_Echo_Result) String() string {
	if v == nil {
		return "<nil>"
	}

	var fields [1]string
	i := 0
	if v.Success != nil {
		fields[i] = fmt.Sprintf("Success: %v", v.Success)
		i++
	}

	return fmt.Sprintf("Hello_Echo_Result{%v}", strings.Join(fields[:i], ", "))
}

// Equals returns true if all the fields of this Hello_Echo_Result match the
// provided Hello_Echo_Result.
//
// This function performs a deep comparison.
func (v *Hello_Echo_Result) Equals(rhs *Hello_Echo_Result) bool {
	if v == nil {
		return rhs == nil
	} else if rhs == nil {
		return false
	}
	if !((v.Success == nil && rhs.Success == nil) || (v.Success != nil && rhs.Success != nil && v.Success.Equals(rhs.Success))) {
		return false
	}

	return true
}

// MarshalLogObject implements zapcore.ObjectMarshaler, enabling
// fast logging of Hello_Echo_Result.
func (v *Hello_Echo_Result) MarshalLogObject(enc zapcore.ObjectEncoder) (err error) {
	if v == nil {
		return nil
	}
	if v.Success != nil {
		err = multierr.Append(err, enc.AddObject("success", v.Success))
	}
	return err
}

// GetSuccess returns the value of Success if it is set or its
// zero value if it is unset.
func (v *Hello_Echo_Result) GetSuccess() (o *EchoResponse) {
	if v != nil && v.Success != nil {
		return v.Success
	}

	return
}

// IsSetSuccess returns true if Success is not nil.
func (v *Hello_Echo_Result) IsSetSuccess() bool {
	return v != nil && v.Success != nil
}

// MethodName returns the name of the Thrift function as specified in
// the IDL, for which this struct represent the result.
//
// This will always be "echo" for this struct.
func (v *Hello_Echo_Result) MethodName() string {
	return "echo"
}

// EnvelopeType returns the kind of value inside this struct.
//
// This will always be Reply for this struct.
func (v *Hello_Echo_Result) EnvelopeType() wire.EnvelopeType {
	return wire.Reply
}
