package service

import (
	"github.com/jinzhu/gorm"
	"log"
	"summerCourse/model"
)

type Goods struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
	Num   int    `json:"num"`
}

// 添加商品
func AddGoods(name string, price int, num int) error {
	// TODO
	goods := model.Goods{
		Model: gorm.Model{},
		Name:  name,
		Price: price,
		Num:   num,
	}
	err := goods.AddGoods()
	return err
}

// 将从数据库查询的所有good信息继承制作Goods切片并返回切片
func SelectGoods() (goods []Goods) {
	_goods, err := model.SelectGoods()
	if err != nil {
		log.Printf("Error get goods info. Error: %s", err)
	}
	for _, v := range _goods {
		good := Goods{
			ID:    v.ID,
			Name:  v.Name,
			Price: v.Price,
			Num:   v.Num,
		}
		goods = append(goods, good)
	}
	return goods
}
