package grpcclient

import (
	"context"
	"runtime/debug"
	"time"

	"github.com/xinliangnote/go-gin-api/pkg/time_parse"
	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"github.com/golang/protobuf/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	// Trace ID
	TraceID = "trace-id"
)

// ClientInterceptor the client's interceptor
type ClientInterceptor struct {
	sign  Sign
	trace *trace.Trace
	grpc  *trace.Grpc
}

// NewClientInterceptor create a client interceptor
func NewClientInterceptor(sign Sign, trace *trace.Trace, grpc *trace.Grpc) *ClientInterceptor {
	return &ClientInterceptor{
		sign:  sign,
		trace: trace,
		grpc:  grpc,
	}
}

// UnaryInterceptor a interceptor for client unary operations
func (c *ClientInterceptor) UnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
	ts := time.Now()
	var invokerErr error

	defer func() { // double recover for safety
		if p := recover(); p != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "got double panic err: %+v,Stack: %s", p, debug.Stack())
			return
		}
	}()

	defer func() {
		if p := recover(); p != nil {
			debug.PrintStack()
			err = status.Errorf(codes.Internal, "got panic err: %+v,Stack: %s", p, debug.Stack())
			return
		}

		var mapReq, mapReply map[string]interface{}
		mapReq, err = ProtoMessage2Map(req.(proto.Message))
		if err != nil {
			return
		}

		mapReply, err = ProtoMessage2Map(reply.(proto.Message))
		if err != nil {
			return
		}

		meta, _ := metadata.FromOutgoingContext(ctx)

		if c.trace != nil {
			c.grpc.Timestamp = time_parse.CSTLayoutString()
			c.grpc.Addr = cc.Target()
			c.grpc.Method = method
			c.grpc.Meta = meta
			c.grpc.Request = mapReq
			c.grpc.Response = mapReply
			c.grpc.CostSeconds = time.Since(ts).Seconds()

			if invokerErr != nil {
				statusErr, ok := status.FromError(invokerErr)
				if ok {
					c.grpc.Code = statusErr.Code().String()
					c.grpc.Message = statusErr.Message()
				}
			}

			c.trace.AppendGRPC(c.grpc)
		}
	}()

	if c.sign != nil {
		var raw string
		if req != nil {
			if raw, err = ProtoMessage2JSON(req.(proto.Message)); err != nil {
				return
			}
		}

		var signature string
		if signature, err = c.sign([]byte(raw)); err != nil {
			return
		}

		meta, _ := metadata.FromOutgoingContext(ctx)
		if meta == nil {
			meta = make(metadata.MD)
		}

		meta.Set(ProxyAuthorization, signature)
		ctx = metadata.NewOutgoingContext(ctx, meta)
	}

	if c.trace != nil {
		meta, _ := metadata.FromOutgoingContext(ctx)
		if meta == nil {
			meta = make(metadata.MD)
		}

		meta.Set(TraceID, c.trace.ID())
		ctx = metadata.NewOutgoingContext(ctx, meta)
	}

	invokerErr = invoker(ctx, method, req, reply, cc, opts...)
	if invokerErr != nil {
		return invokerErr
	}

	return nil
}
