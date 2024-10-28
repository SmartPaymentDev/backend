package domain

type Transaction struct {
	Urut      int
	CUSTID    int
	METODE    string
	TRXDATE   string
	NOREFF    string
	FIDBANK   string
	KDCHANNEL string
	DEBET     int
	KREDIT    int
	REFFBANK  string
	TRANSNO   string
}

type VSaldoVa struct {
	CUSTID int
	NOCUST string
	NMCUST string
	NUM2ND string
	STCUST string
	CODE01 string
	DESC01 string
	CODE02 string
	DESC02 string
	CODE03 string
	DESC03 string
	CODE04 string
	DESC04 string
	GENUS  string
	SALDO  int
}

type TransactionDetail struct {
	Urut      int
	CUSTID    int
	METODE    string
	TRXDATE   string
	NOREFF    string
	FIDBANK   string
	KDCHANNEL string
	DEBET     int
	KREDIT    int
	REFFBANK  string
	TRANSNO   string
}

type VSaldoVaCashless struct {
	CUSTID int
	NOCUST string
	NMCUST string
	NUM2ND string
	STCUST string
	CODE01 string
	CODE02 string
	DESC02 string
	CODE03 string
	DESC03 string
	CODE04 string
	DESC04 string
	GENUS  string
	SALDO  int
}
