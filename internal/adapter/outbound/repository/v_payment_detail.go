package repository

import (
	"context"
	"smart-payment-dev-be/internal/core/domain"
)

type VPaymentDetailRepository struct {
	db DBExecutor
}

func NewVPaymentDetailRepository(db DBExecutor) *VPaymentDetailRepository {
	return &VPaymentDetailRepository{
		db: db,
	}
}

func (b *VPaymentDetailRepository) GetVPaymentDetailByNoCust(ctx context.Context, custId int, billCd string) ([]domain.VPaymentDetail, error) {
	var vPaymentDetails []domain.VPaymentDetail
	sql := `SELECT COALESCE(AA, 0), COALESCE(KodePost, ''), COALESCE(BILLAM, ''), COALESCE(CUSTID, 0), COALESCE(FID, ''), COALESCE(tahun, ''), COALESCE(periode, ''), COALESCE(BILLCD, ''), COALESCE(NamaAkun, '') FROM v_sia_pembayaran_detail WHERE CUSTID = ? AND BILLCD = ?`

	args := []interface{}{custId, billCd}

	rows, err := b.db.QueryxContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		vPaymentDetail := domain.VPaymentDetail{}

		if err = rows.Scan(&vPaymentDetail.AA, &vPaymentDetail.KodePost, &vPaymentDetail.BILLAM, &vPaymentDetail.CUSTID, &vPaymentDetail.FID, &vPaymentDetail.Tahun, &vPaymentDetail.Periode, &vPaymentDetail.BILLCD, &vPaymentDetail.NamaAkun); err != nil {
			return nil, err
		}
		vPaymentDetails = append(vPaymentDetails, vPaymentDetail)
	}

	return vPaymentDetails, nil
}
