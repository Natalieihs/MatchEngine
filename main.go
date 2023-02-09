package main

import (
	"fmt"
	"sort"
)

// Order 结构体用于存储订单信息
type Order struct {
	price float64 // 订单价格
	size  int     // 订单数量
	side  string  // 订单类型，bid代表买单，ask代表卖单
}

// 实现String方法，方便在输出时打印订单信息
func (o Order) String() string {
	return fmt.Sprintf("%.2f-%d-%s", o.price, o.size, o.side)
}

// OrderBook 结构体用于存储订单簿信息
type OrderBook struct {
	bids []Order // 买单列表
	asks []Order // 卖单列表
}

// 实现String方法，方便在输出时打印订单簿信息
func (ob OrderBook) String() string {
	return fmt.Sprintf("Bids: %v\nAsks: %v", ob.bids, ob.asks)
}

// AddOrder 方法用于向订单簿中添加订单
func (ob *OrderBook) AddOrder(order Order) {
	// 判断订单类型，如果是买单，则将其加入bids列表中
	if order.side == "buy" {
		ob.bids = append(ob.bids, order)
		// 对bids列表按照价格从高到低进行排序
		sort.Slice(ob.bids, func(i, j int) bool {
			return ob.bids[i].price > ob.bids[j].price
		})
	} else {
		// 否则，说明是卖单，将其加入asks列表中
		ob.asks = append(ob.asks, order)
		// 对asks列表按照价格从低到高进行排序
		sort.Slice(ob.asks, func(i, j int) bool {
			return ob.asks[i].price < ob.asks[j].price
		})
	}
}

// Match 函数用于撮合
// Match 函数用于撮合
func (ob *OrderBook) Match() {
	// 循环遍历asks列表，取出第一个卖单
	for i := 0; i < len(ob.asks); i++ {
		ask := ob.asks[i]
		// 循环遍历bids列表，取出第一个买单
		for j := 0; j < len(ob.bids); j++ {
			bid := ob.bids[j]
			// 如果买单价格大于等于卖单价格，说明可以撮合
			if bid.price >= ask.price {
				// 计算撮合数量
				size := bid.size
				if ask.size < bid.size {
					size = ask.size
				}
				// 输出撮合信息
				fmt.Printf("Match: %d @ %.2f\n", size, ask.price)
				// 如果卖单数量等于撮合数量，将其从asks列表中删除
				if ask.size == size {
					ob.asks = append(ob.asks[:i], ob.asks[i+1:]...)
					i--
				} else {
					// 否则，说明卖单数量大于撮合数量，将其数量减去撮合数量
					ob.asks[i].size -= size
				}
				// 如果买单数量等于撮合数量，将其从bids列表中删除
				if bid.size == size {
					ob.bids = append(ob.bids[:j], ob.bids[j+1:]...)
					j--
				} else {
					// 否则，说明买单数量大于撮合数量，将其数量减去撮合数量
					ob.bids[j].size -= size
				}
				break
			} else {
				// 如果买单价格小于卖单价格，说明撮合不成功，退出循环
				break
			}
		}
	}
}

// 创建订单
func main() {
	// 创建订单簿
	ob := &OrderBook{}

	ob.AddOrder(Order{
		price: 1.0,
		size:  100,
		side:  "buy",
	})

	ob.AddOrder(Order{
		price: 2.0,
		size:  200,
		side:  "buy",
	})

	ob.AddOrder(Order{
		price: 3.0,
		size:  300,
		side:  "buy",
	})
	ob.AddOrder(Order{
		price: 2.0,
		size:  200,
		side:  "sell",
	})
	ob.AddOrder(Order{
		price: 3.0,
		size:  300,
		side:  "sell",
	})

	// 输出当前订单簿信息
	ob.String()
	// 撮合
	ob.Match()
	// 输出撮合后的订单簿信息
	ob.String()
}
