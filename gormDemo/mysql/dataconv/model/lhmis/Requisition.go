package lhmis

type YwRequisition struct {
	Base
	PaymentInfo
	RequestInfo
	R_ID string `json:"id"`
}
