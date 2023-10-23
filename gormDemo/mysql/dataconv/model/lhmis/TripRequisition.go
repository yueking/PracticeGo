package lhmis

type YwTripRequisition struct {
	Base
	PaymentInfo
	R_ID string

	CAUSE         string
	STARTDATETIME string
	ENDDATETIME   string
}
