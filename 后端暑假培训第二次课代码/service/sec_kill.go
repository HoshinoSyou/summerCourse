package service

import (
	"log"
	"summerCourse/model"
	"sync"
	"time"
)

type User struct {
	UserId  string
	GoodsId uint
}

var OrderChan = make(chan User, 1024)

var ItemMap = make(map[uint]*Item)

type Item struct {
	ID        uint   // 商品id
	Name      string // 名字
	Total     int    // 商品总量
	Left      int    // 商品剩余数量
	IsSoldOut bool   // 是否售罄
	leftCh    chan int
	sellCh    chan int
	done      chan struct{}
	Lock      sync.Mutex
}

// TODO 写一个定时任务，每天定时从数据库加载数据到Map
func PutMap() {
	// 用wg完成等待，即定时任务
	var wg sync.WaitGroup
	wg.Add(100)
	// 每进行一次更新前等待一秒
	time.Sleep(time.Second)
	go func() {
		goods, err := model.SelectGoods()
		if err != nil {
			log.Println("Select goods failed. Error:" + err.Error())
			return
		}
		for k, good := range goods {
			item := &Item{
				ID:        good.ID,
				Name:      good.Name,
				Total:     good.Num,
				Left:      good.Num,
				IsSoldOut: false,
				leftCh:    nil,
				sellCh:    nil,
				done:      nil,
				Lock:      sync.Mutex{},
			}
			ItemMap[uint(k)] = item
		}
		wg.Done()
	}()
	wg.Wait()
}

// 初始化map
func initMap() {
	item := &Item{
		ID:        1,
		Name:      "测试",
		Total:     100,
		Left:      100,
		IsSoldOut: false,
		leftCh:    make(chan int),
		sellCh:    make(chan int),
	}
	ItemMap[item.ID] = item
}

// 通过商品ID获取商品信息
func getItem(itemId uint) *Item {
	return ItemMap[itemId]
}

// 从OrderChan获取user数据，以user中的GoodsId属性查询商品信息，并调用SecKilling方法生成订单
func order() {
	for {
		user := <-OrderChan
		item := getItem(user.GoodsId)
		item.SecKilling(user.UserId)
	}
}

func (item *Item) SecKilling(userId string) {
	// 对item加锁并延时解锁
	item.Lock.Lock()
	defer item.Lock.Unlock()
	// 等价
	// var lock = make(chan struct{}, 1}
	// lock <- struct{}{}
	// defer func() {
	// 		<- lock
	// }
	if item.IsSoldOut {
		return
	}
	item.BuyGoods(1)
	// 生成订单
	MakeOrder(userId, item.ID, 1)

}

// 定时下架
func (item *Item) OffShelve() {
	beginTime := time.Now()
	// 获取第二天时间
	//nextTime := beginTime.Add(time.Hour * 24)
	// 计算次日零点，即商品下架的时间
	//offShelveTime := time.Date(nextTime.Year(), nextTime.Month(), nextTime.Day(), 0, 0, 0, 0, nextTime.Location())
	offShelveTime := beginTime.Add(time.Second * 5)
	timer := time.NewTimer(offShelveTime.Sub(beginTime))
	// 删除map里的要下架的item
	<-timer.C
	delete(ItemMap, item.ID)
	close(item.done)

}

// 出售商品
func (item *Item) SalesGoods() {
	for {
		select {
		case num := <-item.sellCh:
			if item.Left -= num; item.Left <= 0 {
				item.IsSoldOut = true
			} // 如果卖完了就将是否售罄属性改为true
		case item.leftCh <- item.Left:
		case <-item.Done():
			log.Println("我自闭了")
			return
		}
	}
}

// 商品已售完
func (item *Item) Done() <-chan struct{} {
	if item.done == nil {
		item.done = make(chan struct{})
	}
	d := item.done
	return d
}

// 检测商品是否售罄
func (item *Item) Monitor() {
	go item.SalesGoods()
}

// 获取剩余库存
func (item *Item) GetLeft() int {
	var left int
	left = <-item.leftCh
	return left
}

// 购买商品
func (item *Item) BuyGoods(num int) {
	item.sellCh <- num
}

// 初始化服务
func InitService() {
	initMap()
	for _, item := range ItemMap {
		item.Monitor()
		go item.OffShelve()
	}
	for i := 0; i < 10; i++ {
		go order()
	}
}
