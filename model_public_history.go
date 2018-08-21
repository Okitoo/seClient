package seClient

type Trade_public struct {
	ID        int `json:"id"`
	MarketID  int `json:"market_id"`
	OrderType int `json:"order_type"`
	Price     int `json:"price"`
	Amount    int `json:"amount"`
	Time      int `json:"time"`
}
