package grpcclient

import (
	"context"
	"time"

	"github.com/xinliangnote/go-gin-api/pkg/trace"

	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/resolver"
)

var (
	defaultDialTimeout = time.Second * 2
)

type Trace = trace.T

type Option func(*option)

type option struct {
	credential      credentials.TransportCredentials
	keepalive       *keepalive.ClientParameters
	resolverBuilder resolver.Builder
	dialTimeout     time.Duration
	sign            Sign
	trace           *trace.Trace
	grpc            *trace.Grpc
}

// WithCredential setup credential for tls
func WithCredential(credential credentials.TransportCredentials) Option {
	return func(opt *option) {
		opt.credential = credential
	}
}

// WithKeepAlive setup keepalive parameters
func WithKeepAlive(keepalive *keepalive.ClientParameters) Option {
	return func(opt *option) {
		opt.keepalive = keepalive
	}
}

// WithDialTimeout setup the dial timeout
func WithDialTimeout(timeout time.Duration) Option {
	return func(opt *option) {
		opt.dialTimeout = timeout
	}
}

// WithSign setup the signature handler
func WithSign(sign Sign) Option {
	return func(opt *option) {
		opt.sign = sign
	}
}

// WithTrace setup trace info
func WithTrace(t Trace) Option {
	return func(opt *option) {
		if t != nil {
			opt.trace = t.(*trace.Trace)
			opt.grpc = new(trace.Grpc)
		}
	}
}

func New(target string, options ...Option) (*grpc.ClientConn, error) {
	if target == "" {
		return nil, errors.New("target required")
	}

	opt := new(option)
	for _, f := range options {
		f(opt)
	}

	kacp := defaultKeepAlive
	if opt.keepalive != nil {
		kacp = opt.keepalive
	}

	dialTimeout := defaultDialTimeout
	if opt.dialTimeout > 0 {
		dialTimeout = opt.dialTimeout
	}

	clientInterceptor := NewClientInterceptor(opt.sign, opt.trace, opt.grpc)

	dialOptions := []grpc.DialOption{
		grpc.WithBlock(),
		grpc.WithKeepaliveParams(*kacp),
		grpc.WithUnaryInterceptor(clientInterceptor.UnaryInterceptor),
	}

	if opt.credential == nil {
		dialOptions = append(dialOptions, grpc.WithInsecure())
	} else {
		dialOptions = append(dialOptions, grpc.WithTransportCredentials(opt.credential))
	}

	ctx, cancel := context.WithTimeout(context.Background(), dialTimeout)
	defer cancel()

	conn, err := grpc.DialContext(ctx, target, dialOptions...)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return conn, nil
}
