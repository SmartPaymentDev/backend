package transaction

import (
	"context"
	"fmt"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/port/outbound/registry"
	"smart-payment-dev-be/shared/util"
	"strconv"
)

type TransactionService struct {
	repositoryRegistry registry.RepositoryRegistry
}

func NewTransactionService(repositoryRegistry registry.RepositoryRegistry) *TransactionService {
	return &TransactionService{
		repositoryRegistry: repositoryRegistry,
	}
}

func (t *TransactionService) GetTransactionByCustId(ctx context.Context, filter dto.FilterTransaction) (interface{}, error) {
	var transactionDTO dto.TransactionDTO

	repo := t.repositoryRegistry.GetTransactionRepository()

	idCust, err := util.ParseJwtToken(filter.Token)

	if err != nil {
		return nil, err
	}

	filter.CustId = idCust

	res, err := repo.GetTransactionByCustId(ctx, filter)
	if err != nil {
		return nil, err
	}

	count, err := repo.GetCountTransactionByCustId(ctx, filter)
	if err != nil {
		return nil, err
	}

	return transactionDTO.GetTransactionsTransformOut(res, count), nil
}

func (t *TransactionService) GetVSaldoVaByCustId(ctx context.Context, token string) (interface{}, error) {
	var transactionDTO dto.TransactionDTO

	repo := t.repositoryRegistry.GetTransactionRepository()

	idCust, err := util.ParseJwtToken(token)
	if err != nil {
		return nil, err
	}

	res, err := repo.GetVSaldoVaByCustId(ctx, idCust)
	if err != nil {
		return nil, err
	}

	return transactionDTO.GetVSaldoVaTransformOut(res), nil
}

func (t *TransactionService) GetVSaldoVaCashlessByCustId(ctx context.Context, token string) (interface{}, error) {
	var transactionDTO dto.TransactionDTO

	repo := t.repositoryRegistry.GetTransactionRepository()

	idCust, err := util.ParseJwtToken(token)
	if err != nil {
		return nil, err
	}

	res, err := repo.GetVSaldoVaCashlessByCustId(ctx, idCust)
	if err != nil {
		return nil, err
	}

	return transactionDTO.GetVSaldoVaCashlessTransformOut(res), nil
}

func (t *TransactionService) GetTransactionDetailByCustId(ctx context.Context, filter dto.FilterTransactionDetail) (interface{}, error) {
	var transactionDTO dto.TransactionDTO

	repo := t.repositoryRegistry.GetTransactionRepository()

	idCust, err := util.ParseJwtToken(filter.Token)

	if err != nil {
		return nil, err
	}

	filter.CustId = idCust

	res, err := repo.GetTransactionDetailsByCustId(ctx, filter)
	if err != nil {
		return nil, err
	}

	count, err := repo.GetCountTransactionDetailByCustId(ctx, filter)
	if err != nil {
		return nil, err
	}

	return transactionDTO.GetTransactionDetailsTransformOut(res, count), nil
}

func (t *TransactionService) CreateTransactionByCustId(ctx context.Context, token string, req dto.TransactionRequest) error {
	var transactionDTO dto.TransactionDTO

	repo := t.repositoryRegistry.GetTransactionRepository()

	idCust, err := util.ParseJwtToken(token)

	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(idCust)

	urut, _ := repo.GetCountTransactionNow(ctx)

	res := transactionDTO.ReqTransformIn(req)
	res.TRANSNO = util.TransactonNo(urut)
	res.CUSTID = id

	// topup data
	res.METODE = "TOP UP CASHLESS"
	res.DEBET = 0
	fmt.Println("T no ", res.TRANSNO)
	err = repo.CreateTransactionByCustId(ctx, res)
	if err != nil {
		return err
	}

	// admin fee
	res.METODE = "ADMIN FEE"
	res.KREDIT = 0
	res.DEBET = 1000 // only 1rb

	err = repo.CreateTransactionByCustId(ctx, res)
	if err != nil {
		return err
	}

	return nil
}
