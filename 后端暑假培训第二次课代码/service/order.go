package service

import (
	"log"
	"summerCourse/model"
)

// order
func MakeOrder(userId string, goodsId uint, num int) {

	order := model.Order{
		UserID:  userId,
		GoodsID: goodsId,
		Num:     num,
	}
	err := order.MakeOrder()
	if err != nil {
		log.Printf("Error make an order. Error: %s", err)
	}
	log.Println("success")
}
