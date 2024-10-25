package repository

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/domain"
)

type BillRepository interface {
	GetBillByCustId(ctx context.Context, filter dto.FilterBill) ([]domain.Bill, error)
	GetCountBillByCustId(ctx context.Context, filter dto.FilterBill) (int, error)
}
