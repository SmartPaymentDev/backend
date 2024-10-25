package repository

import (
	"context"
	"smart-payment-dev-be/internal/core/domain"
)

type VPaymentDetailRepository interface {
	GetVPaymentDetailByNoCust(ctx context.Context, custId int, billCd string) ([]domain.VPaymentDetail, error)
}
