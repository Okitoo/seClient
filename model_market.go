package seClient

type Market struct {
	ID        int64    `json:"id"`
	BaseCoin  string `json:"base_coin"`
	AltCoin   string `json:"alt_coin"`
	MakerFees int    `json:"maker_fees"`
	TakerFees int    `json:"taker_fees"`
	Volume    struct {
			  Base int64 `json:"Base"`
			  Alt  int64 `json:"Alt"`
		  } `json:"volume"`
	HighLow struct {
			  High int `json:"High"`
			  Low  int `json:"Low"`
		  } `json:"high_low"`
	Change24H     int    `json:"change_24h"`
	Market        string `json:"market"`
	LastSellPrice int    `json:"last_sell_price"`
}
