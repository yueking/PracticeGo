package lhmis

import "time"

type Base struct {
	DEL        string
	CREATEDATE time.Time
	CREATEPID  string
	MODIFYDATE time.Time
	MODIFYPID  string
}

type PaymentInfo struct {
	COUNT_MONEY        float64
	PAYMENT            float64
	UNPAYMENT          float64
	PAYMENTPERSON      string
	PAYMENTPID         string
	PAYMENTDATE        time.Time
	PAYMENTDESCRIPTION string
}

type RequestInfo struct {
	REQUESTDATE    time.Time
	REQUEST_PERSON string
	REQUEST_PID    string

	STATUS      string
	TYPE        string
	DESCRIPTION string

	GROUPNAME string
	LEVELSIGN int
	REGION    string
	OWNER     string

	AUDITDATE        time.Time
	DESCRIPTIONAUDIT string
}
