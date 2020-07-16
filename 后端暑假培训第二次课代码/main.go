package main

import (
	"github.com/gin-gonic/gin"
	"summerCourse/controller"
	"summerCourse/model"
	"summerCourse/service"
)

func main() {
	model.InitDB()
	service.InitService()
	service.PutMap()
	r := gin.Default()
	r.GET("/getGoods", controller.SelectGoods)
	r.POST("/order", controller.MakeOrder)
	r.POST("/addGoods", controller.AddGoods)
	r.Run(":8080")
}
