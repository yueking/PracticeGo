package dataconv

import "time"

type ExpenseNode struct {
	ExpenseNodeId string
	ExpenseId     string
	Base
	BaseNode
	Description    string
	OccurrenceTime time.Time
}
