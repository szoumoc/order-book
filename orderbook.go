package main

import "fmt"

type OrderType int

const (
	GoodTillCancel OrderType = iota
	FillOrKill
)

type Side int

const (
	BUY Side = iota
	SELL
)

type OrderID int64
type Quantity uint64
type Price int64

type LevelInfo struct {
	Price    Price
	Quantity Quantity
}

type LevelInfoList []LevelInfo

type OrderbookLevelInfos struct {
	bids LevelInfoList
	asks LevelInfoList
}

func NewOrderbookLevelInfos(bids, asks LevelInfoList) *OrderbookLevelInfos {
	return &OrderbookLevelInfos{
		bids: bids,
		asks: asks,
	}
}

func (o *OrderbookLevelInfos) GetBids() LevelInfoList {
	return o.bids
}

func (o *OrderbookLevelInfos) GetAsks() LevelInfoList {
	return o.asks
}

type Order struct {
	id                OrderID
	side              Side
	orderType         OrderType
	price             Price
	initialQuantity   Quantity
	remainingQuantity Quantity
}

func NewOrder(id OrderID, side Side, orderType OrderType, price Price, quantity Quantity) *Order {
	return &Order{
		id:                id,
		side:              side,
		orderType:         orderType,
		price:             price,
		initialQuantity:   quantity,
		remainingQuantity: quantity,
	}
}

func (o *Order) GetID() OrderID {
	return o.id
}

func (o *Order) GetSide() Side {
	return o.side
}

func (o *Order) GetType() OrderType {
	return o.orderType
}

func (o *Order) GetPrice() Price {
	return o.price
}

func (o *Order) GetInitialQuantity() Quantity {
	return o.initialQuantity
}
func (o *Order) GetRemainingQuantity() Quantity {
	return o.remainingQuantity
}

func (o *Order) GetFilledQuantity() Quantity {
	return o.initialQuantity - o.remainingQuantity
}
func (o *Order) Fill(quantity Quantity) {
	if quantity > o.remainingQuantity {
		panic("Fill quantity exceeds remaining quantity")
	}
	o.remainingQuantity -= quantity
}

type OrderList []*Order
type OrderListMap []OrderList

type OrderModify struct {
	orderID  OrderID
	price    Price
	side     Side
	quantity Quantity
}

func NewOrderModify(orderID OrderID, price Price, side Side, quantity Quantity) *OrderModify {
	return &OrderModify{
		orderID:  orderID,
		price:    price,
		side:     side,
		quantity: quantity,
	}
}

func (o *OrderModify) GetOrderID() OrderID {
	return o.orderID
}

func (o *OrderModify) GetPrice() Price {
	return o.price
}

func (o *OrderModify) GetSide() Side {
	return o.side
}

func (o *OrderModify) GetQuantity() Quantity {
	return o.quantity
}

func (o *OrderModify) ToOrderPointer(orderType OrderType) *Order {
	return NewOrder(o.orderID, o.side, orderType, o.price, o.quantity)
}

type TradeInfo struct {
	orderID  OrderID
	price    Price
	quantity Quantity
}

type Trade struct {
	bidTradeInfo *TradeInfo
	askTradeInfo *TradeInfo
}

func NewTrade(bidTradeInfo, askTradeInfo *TradeInfo) *Trade {
	return &Trade{
		bidTradeInfo: bidTradeInfo,
		askTradeInfo: askTradeInfo,
	}
}

func main() {
	fmt.Println("Order Book")
}
