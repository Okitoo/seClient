package seClient


type Market_data struct {
	Asks []Order_public `json:"asks"`
	Bids   []Order_public `json:"bids"`
	Market Market `json:"market"`
}