package dto

import "smart-payment-dev-be/internal/core/domain"

type VPaymentDetailDto struct{}

type VPaymentDetail struct {
	AA       int    `json:"aa"`
	KodePost string `json:"kode_post"`
	BILLAM   int    `json:"billam"`
	CUSTID   int    `json:"custid"`
	FID      string `json:"fid"`
	Tahun    string `json:"tahun"`
	Periode  string `json:"periode"`
	BILLCD   string `json:"billcd"`
	NamaAkun string `json:"nama_akun"`
}

func (r *VPaymentDetailDto) GetVPaymentDetailTransferOut(vPayemntDetails []domain.VPaymentDetail) []VPaymentDetail {
	var listPayment []VPaymentDetail

	for _, data := range vPayemntDetails {
		listPayment = append(listPayment, VPaymentDetail{
			AA:       data.AA,
			KodePost: data.KodePost,
			BILLAM:   data.BILLAM,
			CUSTID:   data.CUSTID,
			FID:      data.FID,
			Tahun:    data.Tahun,
			Periode:  data.Periode,
			BILLCD:   data.BILLCD,
			NamaAkun: data.NamaAkun,
		})
	}

	return listPayment
}
