package util

import "github.com/gin-gonic/gin"

type Gin struct {
	Ctx *gin.Context
}

type response struct {
	Code     int         `json:"code"`
	Message  string      `json:"msg"`
	Data     interface{} `json:"data"`
}

func (g *Gin)Response(code int, msg string, data interface{}) {
	g.Ctx.JSON(200, response{
		Code    : code,
		Message : msg,
		Data    : data,
	})
	return
}
