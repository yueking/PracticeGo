package lhmis

import "time"

type YwRequisition struct {
	R_ID string `json:"id"`

	REQUESTDATE        time.Time
	DESCRIPTION        string
	REQUEST_PERSON     string
	REQUEST_PID        string
	STATUS             string
	TYPE               string
	COUNT_MONEY        float64
	PAYMENT            float64
	UNPAYMENT          float64
	PAYMENTPERSON      string
	PAYMENTPID         string
	PAYMENTDATE        time.Time
	GROUPNAME          string
	LEVELSIGN          int
	DEL                string
	CREATEDATE         time.Time
	CREATEPID          string
	MODIFYDATE         time.Time
	MODIFYPID          string
	PAYMENTDESCRIPTION string
	REGION             string
	OWNER              string
	AUDITDATE          time.Time
	DESCRIPTIONAUDIT   string
}
