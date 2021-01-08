package main

import (
	"fmt"
	"strconv"
)

/*
接口 interface
接口就是函数签名的集合，当一个类型定义了接口中的所有函数，称它实现了该接口。
*/

// 定义一个商品 Good 接口
type Good interface {
	// 结算账户
	settleAccount() int
	// 订单信息
	orderInfo() string
}

// 定义一个手机 Phone 结构体并实现函数
type Phone struct {
	name     string
	quantity int
	price    int
}

func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}

func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" + phone.name +
		"，计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

// 定义一个赠品 FreeGift 结构体并实现函数
type FreeGift struct {
	name     string
	quantity int
	price    int
}

func (gift FreeGift) settleAccount() int {
	return 0
}

func (gift FreeGift) orderInfo() string {
	return "您要购买" + strconv.Itoa(gift.quantity) + "个" + gift.name +
		"，计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

// 计算订单金额
func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}

func main() {
	iPhone := Phone{
		name:     "iPhone12",
		quantity: 1,
		price:    10000,
	}

	earPhones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	// 使用 Good 切片作为购物车
	goods := []Good{iPhone, earPhones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需支付 %d 元\n", allPrice)

	/*
		您要购买1个iPhone12，计：10000元
		您要购买1个耳机，计：0元
		该订单总共需支付 10000 元
	*/
}
