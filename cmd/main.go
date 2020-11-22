package main

import (
	"fmt"
	"pili-apiserver/pkg/dao"
)

func main() {
	// router := gin.Default()
	// router.GET("/", func(c *gin.Context) {
	// 	c.String(http.StatusOK, "test")
	// })
	// router.Run()

	dao := dao.SliceDao{}
	dao.Init()
	fmt.Println(dao.Get(1))
}
