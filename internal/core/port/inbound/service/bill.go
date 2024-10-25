package service

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
)

type BillService interface {
	GetBillByCustId(ctx context.Context, filter dto.FilterBill) (interface{}, error)
}
