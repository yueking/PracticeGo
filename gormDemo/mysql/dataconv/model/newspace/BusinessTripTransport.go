package newspace

import "time"

type BusinessTripTransport struct {
	TransportId string
	TripId      string
	Base
	BaseNode
	startDate  time.Time
	endDate    time.Time
	startPlace string
	endPlace   string
}
