package seClient

type Balance struct {
	Coin_code string `json:"coin"`
	Available int64  `json:"available"`
	Pending   int64  `json:"pending"`
	InOrders  int64  `json:"orders"`
}
