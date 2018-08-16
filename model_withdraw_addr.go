package seClient

type Address_withdraw struct {
	Id        int64  `json:"id"`
	Title     string `json:"title"`
	Coin_code string `json:"coin_code"`
	Address   string `json:"address"` // URL to get last block
}
