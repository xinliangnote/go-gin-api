package jaeger_conn

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/model/proto/listen"
	"go-gin-api/app/model/proto/read"
	"go-gin-api/app/model/proto/speak"
	"go-gin-api/app/model/proto/write"
	"go-gin-api/app/util"
	"go-gin-api/app/util/grpc_client"
)

func JaegerTest(c *gin.Context) {

	// 调用 gRPC 服务
	grpcListenClient := listen.NewListenClient(grpc_client.CreateServiceListenConn())
	resListen, _ := grpcListenClient.ListenData(context.Background(), &listen.Request{Name: "listen"})

	// 调用 gRPC 服务
	grpcSpeakClient := speak.NewSpeakClient(grpc_client.CreateServiceSpeakConn())
	resSpeak, _ := grpcSpeakClient.SpeakData(context.Background(), &speak.Request{Name: "speak"})

	// 调用 gRPC 服务
	grpcReadClient := read.NewReadClient(grpc_client.CreateServiceReadConn())
	resRead, _ := grpcReadClient.ReadData(context.Background(), &read.Request{Name: "read"})

	// 调用 gRPC 服务
	grpcWriteClient := write.NewWriteClient(grpc_client.CreateServiceWriteConn())
	resWrite, _ := grpcWriteClient.WriteData(context.Background(), &write.Request{Name: "write"})

	// 调用 HTTP 服务
	resHttpGet := ""
	_, err := util.HttpGet("http://localhost:9905/sing")
	if err == nil {
		resHttpGet = "[HttpGetOk]"
	}

	// 业务处理...

	msg := resListen.Message + "-" +
		   resSpeak.Message + "-" +
		   resRead.Message + "-" +
		   resWrite.Message + "-" +
		   resHttpGet


	utilGin := util.Gin{Ctx:c}
	utilGin.Response(1, msg, nil)
}
