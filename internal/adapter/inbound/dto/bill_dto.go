package dto

import "smart-payment-dev-be/internal/core/domain"

type BillDto struct{}

type FilterBill struct {
	CustId    string
	Token     string
	YearMonth string
	PaidSt    string
	Page      int
	PerPage   int
}

type Bill struct {
	CUSTID          int              `json:"custid"`
	BILLCD          string           `json:"billcd"`
	BILLAC          string           `json:"billac"`
	BILLNM          string           `json:"billnm"`
	BILLAM          int              `json:"billam"`
	FLPART          string           `json:"flpart"`
	PAIDST          string           `json:"paidst"`
	PAIDDT          string           `json:"paiddt"`
	NOREFF          string           `json:"noreff"`
	FSTSBolehBayar  string           `json:"fsts_boleh_bayar"`
	FUrutan         int              `json:"f_urutan"`
	FTGLTagihan     string           `json:"ftgl_tagihan"`
	FIDBANK         string           `json:"fidbank"`
	FRecID          string           `json:"f_rec_id"`
	AA              int              `json:"aa"`
	BTA             string           `json:"bta"`
	TRANSNO         string           `json:"transno"`
	CreateID        string           `json:"create_id"`
	PAIDDT_ACTUAL   string           `json:"paiddt_actual"`
	VPaymentDetails []VPaymentDetail `json:"v_payment_details"`
}

type BillResponses struct {
	Bills []Bill `json:"bills"`
	Total int    `json:"total"`
}

func (n BillDto) GetBillTransferOut(bills []domain.Bill, total int) BillResponses {
	var vPaymentDetailDto VPaymentDetailDto
	var billList []Bill

	for _, data := range bills {

		billList = append(billList, Bill{
			CUSTID:          data.CUSTID,
			BILLCD:          data.BILLCD,
			BILLAC:          data.BILLAC,
			BILLNM:          data.BILLNM,
			BILLAM:          data.BILLAM,
			FLPART:          data.FLPART,
			PAIDST:          data.PAIDST,
			PAIDDT:          data.PAIDDT,
			NOREFF:          data.NOREFF,
			FSTSBolehBayar:  data.FSTSBolehBayar,
			FUrutan:         data.FUrutan,
			FTGLTagihan:     data.FTGLTagihan,
			FIDBANK:         data.FIDBANK,
			FRecID:          data.FRecID,
			AA:              data.AA,
			BTA:             data.BTA,
			TRANSNO:         data.TRANSNO,
			CreateID:        data.CreateID,
			PAIDDT_ACTUAL:   data.PAIDDT_ACTUAL,
			VPaymentDetails: vPaymentDetailDto.GetVPaymentDetailTransferOut(data.VPaymentDetail),
		})
	}

	return BillResponses{
		Bills: billList,
		Total: total,
	}

}
