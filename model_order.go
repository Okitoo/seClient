package seClient

type Order struct {
	Id         int64  `json:"id"`
	Market_id  int64  `json:"market_id"`
	Price      uint64 `json:"price"` //price per alt in base currency
	Amount     uint64 `json:"amount"`
	Total      uint64 `json:"total"` // total price for the amount of alts in the order in base currency
	Order_type int    `json:"order_type"`
	Valid      bool   `json:"valid"`
	Ts         int64  `json:"ts"`
}

type Order_public struct {
	MarketID  int `json:"market_id"`
	Price     int `json:"price"`
	Amount    int `json:"amount"`
	Total     int `json:"total"`
	OrderType int `json:"order_type"`
}

const (
	ORDER_TYPE_BUY = iota
	ORDER_TYPE_SELL
)