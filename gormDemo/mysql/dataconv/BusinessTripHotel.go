package dataconv

type BusinessTripHotel struct {
	HotelId string
	TripId  string
	Base
	BaseNode
	Place       string
	Companion   string
	Days        int
	CaterNumber int
	CaterMoney  float64
}
