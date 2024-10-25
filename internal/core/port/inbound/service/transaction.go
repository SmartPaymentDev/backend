package service

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
)

type TransactionService interface {
	GetTransactionByCustId(ctx context.Context, filter dto.FilterTransaction) (interface{}, error)
	GetVSaldoVaByCustId(ctx context.Context, token string) (interface{}, error)
	GetVSaldoVaCashlessByCustId(ctx context.Context, token string) (interface{}, error)
	GetTransactionDetailByCustId(ctx context.Context, filter dto.FilterTransactionDetail) (interface{}, error)
	CreateTransactionByCustId(ctx context.Context, token string, req dto.TransactionRequest) error
}
