package dataconv

import "time"

// 9
type Base struct {
	Id          string `gorm:"primarykey"`
	CreatedDate time.Time
	ModifyDate  time.Time
	DeletedDate time.Time
	CreatedBy   string
	ModifyBy    string
	DeletedBy   string
	Deleted     MyBool
	version     uint
}

// 11+2
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
	Agree           string
	HasCross        MyBool
	HasRead         MyBool
}

// 7
type PaymentInfo struct {
	CountMoney         float64
	Payment            float64
	UnPayment          float64
	PaymentPID         string
	PaymentPerson      string
	PaymentDate        time.Time
	PaymentDescription string
}

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
