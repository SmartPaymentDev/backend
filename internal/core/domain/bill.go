package domain

type Bill struct {
	CUSTID         int
	BILLCD         string
	BILLAC         string
	BILLNM         string
	BILLAM         int
	FLPART         string
	PAIDST         string
	PAIDDT         string
	NOREFF         string
	FSTSBolehBayar string
	FUrutan        int
	FTGLTagihan    string
	FIDBANK        string
	FRecID         string
	AA             int
	BTA            string
	TRANSNO        string
	CreateID       string
	PAIDDT_ACTUAL  string
	VPaymentDetail []VPaymentDetail
}
