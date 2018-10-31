package seClient

type Instaex struct {
	ID                int            `json:"id"`
	InstaID           string         `json:"insta_id"`
	DepositAddress    string         `json:"deposit_address"`
	DepositCoin       string         `json:"deposit_coin"`
	DepositBalance    int            `json:"deposit_balance"`
	MarketName        string         `json:"market_name"`
	OrderType         int            `json:"order_type"`
	Status            int            `json:"status"`
	WithdrawAddresses map[string]int `json:"withdraw_addresses"`
	WithdrawCoin      string         `json:"withdraw_coin"`
	WithdrawBalance   int            `json:"withdraw_balance"`
	CreatedAt         int            `json:"created_at"`
	LastUsedAt        int            `json:"last_used_at"`
}
