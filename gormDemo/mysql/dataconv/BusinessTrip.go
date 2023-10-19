package dataconv

type BusinessTrip struct {
	Base
	RequestInfo
	PaymentInfo
	TripId      string
	Description string
}
