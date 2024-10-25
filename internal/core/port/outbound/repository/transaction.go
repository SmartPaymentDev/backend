package repository

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/domain"
)

type TransactionRepository interface {
	GetTransactionByCustId(ctx context.Context, transaction dto.FilterTransaction) ([]domain.Transaction, error)
	GetCountTransactionByCustId(ctx context.Context, filter dto.FilterTransaction) (int, error)
	GetVSaldoVaByCustId(ctx context.Context, custId string) (domain.VSaldoVa, error)
	GetVSaldoVaCashlessByCustId(ctx context.Context, custId string) (domain.VSaldoVaCashless, error)
	GetTransactionDetailsByCustId(ctx context.Context, filter dto.FilterTransactionDetail) ([]domain.TransactionDetail, error)
	GetCountTransactionDetailByCustId(ctx context.Context, filter dto.FilterTransactionDetail) (int, error)
	CreateTransactionByCustId(ctx context.Context, req domain.Transaction) error
}
