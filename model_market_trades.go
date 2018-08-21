package seClient

type Market_trade struct {
	History []Trade_public `json:"history"`
	Market  Market `json:"market"`
}
