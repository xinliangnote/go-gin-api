package grpc

import (
	"time"

	"google.golang.org/grpc/keepalive"
)

var (
	keepAlive = &keepalive.ClientParameters{
		Time:                10 * time.Second,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}
)
