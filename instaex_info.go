package seClient

type Instaex_info struct {
	MinDeposit  int             `json:"min_deposit"`
	MinConfs    int             `json:"min_confs"`
	MaxDeposit  int             `json:"max_deposit"`
	WithdrawFee int             `json:"withdraw_fee"`
	Outputs     map[int64]int64 `json:"outputs"`
}
