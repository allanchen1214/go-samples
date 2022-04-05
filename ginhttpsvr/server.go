package main

import (
	"github.com/allanchen1214/bingo/log"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	r := newRouter()

	if err := r.Run(":8080"); err != nil {
		log.Fatal("start svr failed", zap.Error(err))
	}
}

func newRouter() *gin.Engine {
	r := gin.Default()
	RegisterOrderRouter(r)
	return r
}
