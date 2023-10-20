package lhmis

import "time"

type YwTripRequisition struct {
	R_ID               string
	REQUESTDATE        time.Time
	DESCRIPTION        string
	REQUEST_PERSON     string
	REQUEST_PID        string
	STATUS             string
	TYPE               string
	CAUSE              string
	STARTDATETIME      string
	ENDDATETIME        string
	COUNT_MONEY        float64
	PAYMENT            float64
	UNPAYMENT          float64
	PAYMENTPERSON      string
	PAYMENTPID         string
	PAYMENTDATE        time.Time
	PAYMENTDESCRIPTION string
	GROUPNAME          string
	LEVELSIGN          int
	REGION             string
	DEL                string
	CREATEDATE         time.Time
	CREATEPID          string
	MODIFYDATE         time.Time
	MODIFYPID          string
	OWNER              string
	AUDITDATE          time.Time
	DESCRIPTIONAUDIT   string
}
