package dataconv

import "time"

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
