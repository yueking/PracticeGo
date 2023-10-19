package dataconv

import (
	"time"
)

type Base struct {
	Id string `gorm:"primarykey"`

	CreatedDate time.Time
	ModifyDate  time.Time
	DeletedDate time.Time

	CreatedBy string
	ModifyBy  string
	DeletedBy string

	Deleted MyBool
	version uint
}
type RequestInfo struct {
	RequestDate     time.Time
	RequestPerson   string
	RequestPersonID string
	GroupName       string
	region          string
	OwnerRegion00   string
	OwnerRegion41   string
	OwnerRegion51   string
	OwnerRegion32   string
	Status          string
	//Agree           string
}
type PaymentInfo struct {
	CountMoney    float64
	Payment       float64
	UnPayment     float64
	PaymentPID    string
	PaymentPerson string
	PaymentDate   time.Time
}
type Expense struct {
	Base
	RequestInfo
	PaymentInfo

	Description string
	Id          string
	ExpenseId   string
}

func (Expense) TableName() string {
	return "expense"
}
