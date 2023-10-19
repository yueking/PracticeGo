package dataconv

import "time"

type BaseNode struct {
	Owner       string
	OwnerRegion string
	Region      string
	Stage       string
	Status      string
	Hid         string
	Product     string
	Agree       string
	Payment     string
}
type ExpenseNode struct {
	Base
	Description    string
	ExpenseNodeId  string
	Money          float64
	Number         int64
	OccurrenceTime time.Time
	Type           string
	ExpenseId      string
}
