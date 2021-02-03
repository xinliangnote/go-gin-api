package grpclient

import (
	"time"

	"google.golang.org/grpc/keepalive"
)

var (
	defaultKeepAlive = &keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}
)
