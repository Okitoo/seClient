package seClient

type Address_deposit struct {
	Coin_code string `json:"coin_code"`
	Address   string `json:"address"` // URL to get last block
	Main      bool   `json:"main"`    // URL to get last block
}
