package main

import (
	"fmt"
	"pili-apiserver/pkg/handler"
)

func main() {
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "test")
	// })
	// router.Run()

	h := handler.Handler{}
	h.Init()
	fmt.Println(h.Get(1))
}
