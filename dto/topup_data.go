package dto

type TopUpReq struct {
	Amount float64 `json:"amount"`
	UserId int64   `json:"-"`
}

type TopUpRes struct {
	SnapUrl string `json:"snap_url"`
}
