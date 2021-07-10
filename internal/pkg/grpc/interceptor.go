package grpc

import (
	"context"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/xinliangnote/go-gin-api/internal/pkg/code"
	"github.com/xinliangnote/go-gin-api/internal/pkg/core"
	"github.com/xinliangnote/go-gin-api/internal/pkg/notify"
	"github.com/xinliangnote/go-gin-api/pkg/errno"
	"github.com/xinliangnote/go-gin-api/pkg/p"
	"github.com/xinliangnote/go-gin-api/pkg/time_parse"
	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"github.com/golang/protobuf/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	TraceID = "trace-id"
)

type clientWithContextKeyType struct{ name string }

var ClientWithContextKey = clientWithContextKeyType{"_client_with_context"}

// ClientInterceptor the client's interceptor
type ClientInterceptor struct {
	sign Sign
}

// NewClientInterceptor create a client interceptor
func NewClientInterceptor(sign Sign) *ClientInterceptor {
	return &ClientInterceptor{
		sign: sign,
	}
}

// UnaryInterceptor a interceptor for client unary operations
func (c *ClientInterceptor) UnaryInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	var (
		invokerErr  error
		ts          = time.Now()
		coreContext = ctx.Value(ClientWithContextKey).(core.Context)
	)

	defer func() { // double recover for safety
		if err := recover(); err != nil {
			stackInfo := string(debug.Stack())
			coreContext.Logger().Error("UnaryInterceptor got double panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo))
			coreContext.AbortWithError(errno.NewError(
				http.StatusInternalServerError,
				code.ServerError,
				code.Text(code.ServerError)),
			)
			notify.OnPanicNotify(coreContext, err, stackInfo)
		}
	}()

	defer func() {
		if err := recover(); err != nil {
			stackInfo := string(debug.Stack())
			coreContext.Logger().Error("UnaryInterceptor got panic", zap.String("panic", fmt.Sprintf("%+v", err)), zap.String("stack", stackInfo))
			coreContext.AbortWithError(errno.NewError(
				http.StatusInternalServerError,
				code.ServerError,
				code.Text(code.ServerError)),
			)
			notify.OnPanicNotify(coreContext, err, stackInfo)
		}

		if coreContext.Trace() != nil {
			var mapReq, mapReply map[string]interface{}
			mapReq, err := ProtoMessage2Map(req.(proto.Message))
			if err != nil {
				p.Println("req ProtoMessage2Map err", err, p.WithTrace(coreContext.Trace()))
			}

			mapReply, err = ProtoMessage2Map(reply.(proto.Message))
			if err != nil {
				p.Println("reply ProtoMessage2Map err", err, p.WithTrace(coreContext.Trace()))
			}

			meta, _ := metadata.FromOutgoingContext(ctx)

			gRPCTrace := new(trace.Grpc)
			gRPCTrace.Timestamp = time_parse.CSTLayoutString()
			gRPCTrace.Addr = cc.Target()
			gRPCTrace.Method = method
			gRPCTrace.Meta = meta
			gRPCTrace.Request = mapReq
			gRPCTrace.Response = mapReply
			gRPCTrace.CostSeconds = time.Since(ts).Seconds()

			if invokerErr != nil {
				statusErr, ok := status.FromError(invokerErr)
				if ok {
					gRPCTrace.Code = statusErr.Code().String()
					gRPCTrace.Message = statusErr.Message()
				}
			}

			coreContext.Trace().AppendGRPC(gRPCTrace)
		}
	}()

	if c.sign != nil {
		var (
			raw       string
			signature string
			err       error
		)

		if req != nil {
			if raw, err = ProtoMessage2JSON(req.(proto.Message)); err != nil {
				return err
			}
		}

		if signature, err = c.sign([]byte(raw)); err != nil {
			return err
		}

		meta, _ := metadata.FromOutgoingContext(ctx)
		if meta == nil {
			meta = make(metadata.MD)
		}

		meta.Set(ProxyAuthorization, signature)
		ctx = metadata.NewOutgoingContext(ctx, meta)
	}

	if coreContext.Trace() != nil {
		meta, _ := metadata.FromOutgoingContext(ctx)
		if meta == nil {
			meta = make(metadata.MD)
		}

		meta.Set(TraceID, coreContext.Trace().ID())
		ctx = metadata.NewOutgoingContext(ctx, meta)
	}

	invokerErr = invoker(ctx, method, req, reply, cc, opts...)
	return invokerErr
}
