package repository

import (
	"context"
	"fmt"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/domain"
)

type BillRepository struct {
	db DBExecutor
}

func NewBillRepository(db DBExecutor) *BillRepository {
	return &BillRepository{
		db: db,
	}
}

func (b *BillRepository) GetBillByCustId(ctx context.Context, filter dto.FilterBill) ([]domain.Bill, error) {
	var bills []domain.Bill
	sql := `SELECT CUSTID, COALESCE(BILLCD, ''), COALESCE(BILLAC, ''), COALESCE(BILLNM, ''), COALESCE(BILLAM, 0), COALESCE(FLPART, ''), COALESCE(PAIDST, ''), COALESCE(PAIDDT, ''), COALESCE(NOREFF, ''), COALESCE(FSTSBolehBayar, ''), COALESCE(FUrutan, 0), COALESCE(FTGLTagihan, ''), COALESCE(FIDBANK, ''), COALESCE(FRecID, ''), COALESCE(AA, 0), COALESCE(BTA, ''), COALESCE(TRANSNO, ''), COALESCE(CreateID, ''), COALESCE(PAIDDT_ACTUAL, '') FROM scctbill WHERE CUSTID = ?`

	if filter.PaidSt != "" {
		sql += fmt.Sprintf(` AND PAIDST = '%s'`, filter.PaidSt)
	}

	if filter.Page != 0 {
		filter.Page = (filter.PerPage*(filter.Page-1) + 1) - 1
	}

	if filter.PaidSt != "1" {
		sql += " ORDER BY PAIDDT DESC"
	} else {
		sql += " ORDER BY FUrutan ASC"
	}

	sql += " LIMIT ? OFFSET ?"
	args := []interface{}{filter.CustId, filter.PerPage, filter.Page}

	rows, err := b.db.QueryxContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		bill := domain.Bill{}

		if err = rows.Scan(&bill.CUSTID, &bill.BILLCD, &bill.BILLAC, &bill.BILLNM, &bill.BILLAM, &bill.FLPART, &bill.PAIDST, &bill.PAIDDT, &bill.NOREFF, &bill.FSTSBolehBayar, &bill.FUrutan, &bill.FTGLTagihan, &bill.FIDBANK, &bill.FRecID, &bill.AA, &bill.BTA, &bill.TRANSNO, &bill.CreateID, &bill.PAIDDT_ACTUAL); err != nil {
			return nil, err
		}
		bills = append(bills, bill)
	}

	return bills, nil
}

func (b *BillRepository) GetCountBillByCustId(ctx context.Context, filter dto.FilterBill) (int, error) {
	sql := `SELECT count(CUSTID) FROM scctbill WHERE CUSTID = ?`

	if filter.PaidSt != "" {
		sql += fmt.Sprintf(` AND PAIDST = '%s'`, filter.PaidSt)
	}

	args := []interface{}{filter.CustId}

	count := 0
	err := b.db.QueryRowxContext(ctx, sql, args...).Scan(&count)

	return count, err
}
