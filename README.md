# summerCourse
## 后端暑假培训第二次课代码
给第二次课代码写上的注释（自己的理解），并完成带有TODO标签的内容（即实现后台添加商品到数据库，并把数据库的信息同步到ItemMap）
#### TODO1：
添加商品
- 新建“/addGoods”路径的路由
- 在controller层中新建AddGoods函数获取输入的商品信息，并传给service层的AddGoods函数，得到返回的err进行判断
- 在service层的AddGoods函数中，以获取的数据声明goods结构体变量，并传给goods.AddGoods方法，得到其返回的err并返回
#### TODO2：
写一个定时任务，每天定时从数据库加载数据到Map
- 利用sync.WaitGroup完成定时更新数据
- 每次进行goroutine前等待一秒
- 调用model层的goods.SelectGoods方法获取数据库中所有商品信息，即Goods切片
- 遍历切片，以Goods结构体的成员声明item结构体（item结构体的其他不存在于Goods结构体的成员均赋予默认值）
- 以遍历Goods切片的键k为ItemMap的键插入k-item
- 在goroutine内调用wg.Done方法，并在外调用wg.Wait方法
