package dto

import "smart-payment-dev-be/internal/core/domain"

type TransactionDTO struct{}

type Transaction struct {
	Urut      int    `json:"urut"`
	CUSTID    int    `json:"custid"`
	METODE    string `json:"metode"`
	TRXDATE   string `json:"trxdate"`
	NOREFF    string `json:"noreff"`
	FIDBANK   string `json:"fidbank"`
	KDCHANNEL string `json:"kdchannel"`
	DEBET     int    `json:"debet"`
	KREDIT    int    `json:"kredit"`
	REFFBANK  string `json:"reffbank"`
	TRANSNO   string `json:"transno"`
}

type VSaldoVa struct {
	CUSTID int    `json:"custid"`
	NOCUST string `json:"nocust"`
	NMCUST string `json:"nmcust"`
	NUM2ND string `json:"num_2_nd"`
	STCUST string `json:"stcust"`
	CODE01 string `json:"code_01"`
	DESC01 string `json:"desc_01"`
	CODE02 string `json:"code_02"`
	DESC02 string `json:"desc_02"`
	CODE03 string `json:"code_03"`
	DESC03 string `json:"desc_03"`
	CODE04 string `json:"code_04"`
	DESC04 string `json:"desc_04"`
	GENUS  string `json:"genus"`
	SALDO  int    `json:"saldo"`
}

type TransactionDetail struct {
	Urut      int    `json:"urut"`
	CUSTID    int    `json:"custid"`
	METODE    string `json:"metode"`
	TRXDATE   string `json:"trxdate"`
	NOREFF    string `json:"noreff"`
	FIDBANK   string `json:"fidbank"`
	KDCHANNEL string `json:"kdchannel"`
	DEBET     int    `json:"debet"`
	KREDIT    int    `json:"kredit"`
	REFFBANK  string `json:"reffbank"`
	TRANSNO   string `json:"transno"`
}

type VSaldoVaCashless struct {
	CUSTID int    `json:"custid"`
	NOCUST string `json:"nocust"`
	NMCUST string `json:"nmcust"`
	NUM2ND string `json:"num_2_nd"`
	STCUST string `json:"stcust"`
	CODE01 string `json:"code_01"`
	CODE02 string `json:"code_02"`
	DESC02 string `json:"desc_02"`
	CODE03 string `json:"code_03"`
	DESC03 string `json:"desc_03"`
	CODE04 string `json:"code_04"`
	DESC04 string `json:"desc_04"`
	GENUS  string `json:"genus"`
	SALDO  int    `json:"saldo"`
}

type FilterTransaction struct {
	Token   string
	CustId  string
	From    string
	To      string
	Page    int
	PerPage int
}

type FilterTransactionDetail struct {
	Token   string
	CustId  string
	From    string
	To      string
	Page    int
	PerPage int
}

type TransactionRequest struct {
	Amount int `json:"amount"`
}

type VSaldoVaResponse struct {
	VSaldoVa VSaldoVa `json:"v_saldo_va"`
}

type VSaldoVaCashlessResponse struct {
	VSaldoVaCashless VSaldoVaCashless `json:"v_saldo_va_cashless"`
}

type TransactionDetailResponses struct {
	TransactionDetails []TransactionDetail `json:"transaction_details"`
	Total              int                 `json:"total"`
}

type TransactionResponses struct {
	Transactions []Transaction `json:"transactions"`
	Total        int           `json:"total"`
}

func (t *TransactionDTO) ReqTransformIn(req TransactionRequest) domain.Transaction {
	return domain.Transaction{
		KREDIT:  req.Amount,
		NOREFF:  "MOBILE",
		FIDBANK: "1140002",
	}
}

func (t *TransactionDTO) GetTransactionsTransformOut(res []domain.Transaction, total int) interface{} {
	var list []Transaction

	for _, data := range res {
		list = append(list, Transaction{
			Urut:      data.Urut,
			CUSTID:    data.CUSTID,
			METODE:    data.METODE,
			TRXDATE:   data.TRXDATE,
			NOREFF:    data.NOREFF,
			FIDBANK:   data.FIDBANK,
			KDCHANNEL: data.KDCHANNEL,
			DEBET:     data.DEBET,
			KREDIT:    data.KREDIT,
			REFFBANK:  data.REFFBANK,
			TRANSNO:   data.TRANSNO,
		})
	}

	return TransactionResponses{
		Transactions: list,
		Total:        total,
	}
}

func (t *TransactionDTO) GetVSaldoVaTransformOut(res domain.VSaldoVa) interface{} {
	return VSaldoVaResponse{
		VSaldoVa: VSaldoVa{
			CUSTID: res.CUSTID,
			NOCUST: res.NOCUST,
			NMCUST: res.NMCUST,
			NUM2ND: res.NUM2ND,
			STCUST: res.STCUST,
			CODE01: res.CODE01,
			DESC01: res.DESC01,
			CODE02: res.CODE02,
			DESC02: res.DESC02,
			CODE03: res.CODE03,
			DESC03: res.DESC03,
			CODE04: res.CODE04,
			DESC04: res.DESC04,
			GENUS:  res.GENUS,
			SALDO:  res.SALDO,
		},
	}
}

func (t *TransactionDTO) GetTransactionDetailsTransformOut(res []domain.TransactionDetail, total int) interface{} {
	var list []TransactionDetail

	for _, data := range res {
		list = append(list, TransactionDetail{
			Urut:      data.Urut,
			CUSTID:    data.CUSTID,
			METODE:    data.METODE,
			TRXDATE:   data.TRXDATE,
			NOREFF:    data.NOREFF,
			FIDBANK:   data.FIDBANK,
			KDCHANNEL: data.KDCHANNEL,
			KREDIT:    data.KREDIT,
			DEBET:     data.DEBET,
			REFFBANK:  data.REFFBANK,
			TRANSNO:   data.TRANSNO,
		})
	}

	return TransactionDetailResponses{
		TransactionDetails: list,
		Total:              total,
	}
}

func (t *TransactionDTO) GetVSaldoVaCashlessTransformOut(res domain.VSaldoVaCashless) interface{} {
	return VSaldoVaCashlessResponse{
		VSaldoVaCashless: VSaldoVaCashless{
			CUSTID: res.CUSTID,
			NOCUST: res.NOCUST,
			NMCUST: res.NMCUST,
			NUM2ND: res.NUM2ND,
			STCUST: res.STCUST,
			CODE01: res.CODE01,
			CODE02: res.CODE02,
			DESC02: res.DESC02,
			CODE03: res.CODE03,
			DESC03: res.DESC03,
			CODE04: res.CODE04,
			DESC04: res.DESC04,
			GENUS:  res.GENUS,
			SALDO:  res.SALDO,
		},
	}
}
