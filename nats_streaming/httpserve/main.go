package main

import (
	"github.com/gin-gonic/gin"
	// "log"
	"httpserve/service"
)

func main() {
	router := gin.Default()

	router.POST("/push", service.PushMsg)
	// router.GET("/pull", service.PullMsg)

	router.Run(":80")
}
