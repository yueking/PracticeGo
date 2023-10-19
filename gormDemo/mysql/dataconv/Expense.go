package dataconv

type Expense struct {
	Base
	RequestInfo
	PaymentInfo

	ExpenseId   string
	Description string
	//DescriptionAudit string
}

func (Expense) TableName() string {
	return "expense"
}
