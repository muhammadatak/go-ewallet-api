package dto

type TransferInquiryReq struct {
	AccountNumber string  `json:"account_number"`
	Amount        float64 `json:"amount"`
}

type TransferInquiryRes struct {
	InquiryKey string `json:"inquiry_key"`
}

type TransferExecuteReq struct {
	InquiryKey string `json:"inquiry_key"`
	PIN        string `json:"pin"`
}
