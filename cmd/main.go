package main

import (
	"pili-apiserver/pkg/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	h := handler.Handler{}
	h.Init()
	router.GET("/role/:id", h.Get)
	router.GET("/role", h.List)
	router.POST("/role", h.Save)
	router.PUT("/role/:id", h.Update)
	router.DELETE("/role/:id", h.Delete)
	router.Run()
}
