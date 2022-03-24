## 算法题

项目中的 `get_arr_same_elem.go` 文件中，测试文件见 test `get_arr_same_elem_test.go`

## 设计题

> Notice:
> 1. 该程序运行与 Go 1.18 ，其他版本暂未测试
> 2. 配置文件通过 proto 来管理，避免每次修改配置文件都需要修改 config.go 文件，而是使用 proto 自动生成
> 3. 规则配置可以通过 `app/order/service/configs/config.yaml` 中的 `rules.milk_tea.rule` 来配置
> 4. 验证规则新增了两个自定义验证规则：can_addition（可以添加的配料列表，无则不需要填写该规则）、cant_addition（不能同时添加的配料列表，无则不需要填写该规则）
> 5. 抽象成了产品、产品类型、配料、订单、订单Item，产品、产品类型、配料如下：![dspAlL-J8iBhe](http://img.ra6n.cn/space/dspAlL-J8iBhe.png)
### 入口

```shell
$ ./app/order/service/cmd/server/order-service 
INFO msg=config loaded: config.yaml format: yaml
1. 椰果奶茶
配料：红豆,价格：￥12.00
2. 椰果奶茶
配料：芋圆,价格：￥12.00
3. 西⽶奶茶
配料：红豆,价格：￥12.00
4. 美式咖啡
配料：芝⼠奶盖,价格：￥16.00

总价为：￥52.00
```

> 异常执行情况可以通过 `main.go:example()` 中注释部分验证

### 问题

> RPC 调用  `order.service.v1.Order.AddOrder` 添加数据后会累计

### 待优化

- [ ] 引入 wire
- [ ] 目前的添加单个饮品是累加的，即 GRCP 执行完 `order.service.v1.Order.AddOrder` 后数据为清理，下次在执行会饮品会累加
- [ ] 新增单个饮品的实例化待优化，即 `NewOrderItem()` 



