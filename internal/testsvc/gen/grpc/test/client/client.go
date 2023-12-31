// Code generated by goa v3.9.1, DO NOT EDIT.
//
// test gRPC client
//
// Command:
// $ goa gen goa.design/clue/internal/testsvc/design

package client

import (
	"context"

	testpb "goa.design/clue/internal/testsvc/gen/grpc/test/pb"
	test "goa.design/clue/internal/testsvc/gen/test"
	goagrpc "goa.design/goa/v3/grpc"
	goa "goa.design/goa/v3/pkg"
	"google.golang.org/grpc"
)

// Client lists the service endpoint gRPC clients.
type Client struct {
	grpccli testpb.TestClient
	opts    []grpc.CallOption
}

// GrpcStreamClientStream implements the test.GrpcStreamClientStream interface.
type GrpcStreamClientStream struct {
	stream testpb.Test_GrpcStreamClient
}

// NewClient instantiates gRPC client for all the test service servers.
func NewClient(cc *grpc.ClientConn, opts ...grpc.CallOption) *Client {
	return &Client{
		grpccli: testpb.NewTestClient(cc),
		opts:    opts,
	}
}

// GrpcMethod calls the "GrpcMethod" function in testpb.TestClient interface.
func (c *Client) GrpcMethod() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildGrpcMethodFunc(c.grpccli, c.opts...),
			EncodeGrpcMethodRequest,
			DecodeGrpcMethodResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// GrpcStream calls the "GrpcStream" function in testpb.TestClient interface.
func (c *Client) GrpcStream() goa.Endpoint {
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		inv := goagrpc.NewInvoker(
			BuildGrpcStreamFunc(c.grpccli, c.opts...),
			nil,
			DecodeGrpcStreamResponse)
		res, err := inv.Invoke(ctx, v)
		if err != nil {
			return nil, goa.Fault(err.Error())
		}
		return res, nil
	}
}

// Recv reads instances of "testpb.GrpcStreamResponse" from the "grpc_stream"
// endpoint gRPC stream.
func (s *GrpcStreamClientStream) Recv() (*test.Fields, error) {
	var res *test.Fields
	v, err := s.stream.Recv()
	if err != nil {
		return res, err
	}
	return NewFields(v), nil
}

// Send streams instances of "testpb.GrpcStreamStreamingRequest" to the
// "grpc_stream" endpoint gRPC stream.
func (s *GrpcStreamClientStream) Send(res *test.Fields) error {
	v := NewProtoGrpcStreamStreamingRequest(res)
	return s.stream.Send(v)
}

func (s *GrpcStreamClientStream) Close() error {
	// Close the send direction of the stream
	return s.stream.CloseSend()
}
