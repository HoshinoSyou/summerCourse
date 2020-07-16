package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"summerCourse/service"
)

// 查询商品
func SelectGoods(ctx *gin.Context) {
	goods := service.SelectGoods()
	ctx.JSON(http.StatusOK, gin.H{
		"status": 200,
		"info":   "success",
		"data": struct {
			Goods []service.Goods `json:"goods"`
		}{goods},
	})
}

// 添加商品
func AddGoods(ctx *gin.Context) {
	name := ctx.PostForm("name")
	_price := ctx.PostForm("price")
	_num := ctx.PostForm("num")
	price, _ := strconv.Atoi(_price)
	num, _ := strconv.Atoi(_num)
	err := service.AddGoods(name, price, num)
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"info":   "failure",
			"msg":    "添加商品失败！" + err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"status": 200,
			"info":   "success",
			"msg":    "添加商品成功！",
		})
	}
}
