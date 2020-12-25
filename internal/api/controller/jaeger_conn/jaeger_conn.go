package jaeger_conn

import (
	"context"

	"github.com/xinliangnote/go-gin-api/internal/api/model/proto/listen"
	"github.com/xinliangnote/go-gin-api/internal/api/model/proto/read"
	"github.com/xinliangnote/go-gin-api/internal/api/model/proto/speak"
	"github.com/xinliangnote/go-gin-api/internal/api/model/proto/write"
	"github.com/xinliangnote/go-gin-api/internal/api/util/grpc_client"
	"github.com/xinliangnote/go-gin-api/internal/api/util/request"
	"github.com/xinliangnote/go-gin-api/internal/api/util/response"

	"github.com/gin-gonic/gin"
)

func JaegerTest(c *gin.Context) {

	// 调用 gRPC 服务
	conn := grpc_client.CreateServiceListenConn(c)
	grpcListenClient := listen.NewListenClient(conn)
	resListen, _ := grpcListenClient.ListenData(context.Background(), &listen.Request{Name: "listen"})

	// 调用 gRPC 服务
	conn = grpc_client.CreateServiceSpeakConn(c)
	grpcSpeakClient := speak.NewSpeakClient(conn)
	resSpeak, _ := grpcSpeakClient.SpeakData(context.Background(), &speak.Request{Name: "speak"})

	// 调用 gRPC 服务
	conn = grpc_client.CreateServiceReadConn(c)
	grpcReadClient := read.NewReadClient(conn)
	resRead, _ := grpcReadClient.ReadData(context.Background(), &read.Request{Name: "read"})

	// 调用 gRPC 服务
	conn = grpc_client.CreateServiceWriteConn(c)
	grpcWriteClient := write.NewWriteClient(conn)
	resWrite, _ := grpcWriteClient.WriteData(context.Background(), &write.Request{Name: "write"})

	defer conn.Close()

	// 调用 HTTP 服务
	resHttpGet := ""
	_, err := request.HttpGet("http://localhost:9905/sing", c)
	if err == nil {
		resHttpGet = "[HttpGetOk]"
	}

	// 业务处理...

	msg := resListen.Message + "-" +
		resSpeak.Message + "-" +
		resRead.Message + "-" +
		resWrite.Message + "-" +
		resHttpGet

	utilGin := response.Gin{Ctx: c}
	utilGin.Response(1, msg, nil)
}
