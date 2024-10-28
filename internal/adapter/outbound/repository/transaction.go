package repository

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/domain"
	"smart-payment-dev-be/shared/util"
)

type TransactionRepository struct {
	db DBExecutor
}

func NewTransactionRepository(db DBExecutor) *TransactionRepository {
	return &TransactionRepository{
		db: db,
	}
}

func (t *TransactionRepository) GetTransactionByCustId(ctx context.Context, filter dto.FilterTransaction) ([]domain.Transaction, error) {
	var transactions []domain.Transaction
	sql := `SELECT urut, COALESCE(CUSTID, 0), COALESCE(METODE, ''), COALESCE(TRXDATE, ''), COALESCE(NOREFF, ''), COALESCE(FIDBANK, ''), COALESCE(KDCHANNEL, ''), COALESCE(DEBET, 0), COALESCE(KREDIT, 0), COALESCE(REFFBANK, ''), COALESCE(TRANSNO, '') FROM sccttran WHERE CUSTID = ? ORDER BY urut DESC LIMIT ? OFFSET ?`

	if filter.Page != 0 {
		filter.Page = (filter.PerPage*(filter.Page-1) + 1) - 1
	}

	args := []interface{}{filter.CustId, filter.PerPage, filter.Page}

	rows, err := t.db.QueryxContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		transaction := domain.Transaction{}

		if err = rows.Scan(&transaction.Urut, &transaction.CUSTID, &transaction.METODE, &transaction.TRXDATE, &transaction.NOREFF, &transaction.FIDBANK, &transaction.KDCHANNEL, &transaction.DEBET, &transaction.KREDIT, &transaction.REFFBANK, &transaction.TRANSNO); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (t *TransactionRepository) GetCountTransactionByCustId(ctx context.Context, filter dto.FilterTransaction) (int, error) {
	sql := `SELECT count(urut) FROM sccttran WHERE CUSTID = ?`

	args := []interface{}{filter.CustId}

	count := 0
	err := t.db.QueryRowxContext(ctx, sql, args...).Scan(&count)

	return count, err
}

func (t *TransactionRepository) GetVSaldoVaByCustId(ctx context.Context, custId string) (domain.VSaldoVa, error) {
	var saldoVa domain.VSaldoVa
	sql := `SELECT CUSTID, COALESCE(CODE01, ''),COALESCE(CODE02, ''), COALESCE(DESC02, ''), COALESCE(CODE03, ''), COALESCE(DESC03, ''), COALESCE(CODE04, ''), COALESCE(DESC04, ''), COALESCE(NOCUST, ''), COALESCE(NMCUST, ''), COALESCE(NUM2ND, ''), COALESCE(STCUST, ''), COALESCE(GENUS, ''), COALESCE(SALDO, '')  FROM v_saldo_va WHERE CUSTID = ?`
	args := []interface{}{custId}

	err := t.db.QueryRowxContext(ctx, sql, args...).Scan(&saldoVa.CUSTID, &saldoVa.CODE01, &saldoVa.CODE02, &saldoVa.DESC02, &saldoVa.CODE03, &saldoVa.DESC03, &saldoVa.CODE04, &saldoVa.DESC04, &saldoVa.NOCUST, &saldoVa.NMCUST, &saldoVa.NUM2ND, &saldoVa.STCUST, &saldoVa.GENUS, &saldoVa.SALDO)
	if err != nil {
		return saldoVa, err
	}

	return saldoVa, nil
}

func (t *TransactionRepository) GetVSaldoVaCashlessByCustId(ctx context.Context, custId string) (domain.VSaldoVaCashless, error) {
	var saldoVa domain.VSaldoVaCashless
	sql := `SELECT CUSTID, COALESCE(CODE01, ''),COALESCE(CODE02, ''), COALESCE(DESC02, ''), COALESCE(CODE03, ''), COALESCE(DESC03, ''), COALESCE(CODE04, ''), COALESCE(DESC04, ''), COALESCE(NOCUST, ''), COALESCE(NMCUST, ''), COALESCE(NUM2ND, ''), COALESCE(STCUST, ''), COALESCE(GENUS, ''), COALESCE(SALDO, '')  FROM v_saldo_cashless_1va WHERE CUSTID = ?`
	args := []interface{}{custId}

	err := t.db.QueryRowxContext(ctx, sql, args...).Scan(&saldoVa.CUSTID, &saldoVa.CODE01, &saldoVa.CODE02, &saldoVa.DESC02, &saldoVa.CODE03, &saldoVa.DESC03, &saldoVa.CODE04, &saldoVa.DESC04, &saldoVa.NOCUST, &saldoVa.NMCUST, &saldoVa.NUM2ND, &saldoVa.STCUST, &saldoVa.GENUS, &saldoVa.SALDO)
	if err != nil {
		return saldoVa, err
	}

	return saldoVa, nil
}

func (t *TransactionRepository) GetTransactionDetailsByCustId(ctx context.Context, filter dto.FilterTransactionDetail) ([]domain.TransactionDetail, error) {
	var transactions []domain.TransactionDetail
	sql := `SELECT urut, COALESCE(CUSTID, 0), COALESCE(METODE, ''), COALESCE(TRXDATE, ''), COALESCE(NOREFF, ''), COALESCE(FIDBANK, ''), COALESCE(KDCHANNEL, ''), COALESCE(DEBET, 0), COALESCE(KREDIT, 0), COALESCE(REFFBANK, ''), COALESCE(TRANSNO, '') FROM sccttran_cashless WHERE CUSTID = ? ORDER BY urut DESC LIMIT ? OFFSET ?`

	if filter.Page != 0 {
		filter.Page = (filter.PerPage*(filter.Page-1) + 1) - 1
	}

	args := []interface{}{filter.CustId, filter.PerPage, filter.Page}

	rows, err := t.db.QueryxContext(ctx, sql, args...)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		transaction := domain.TransactionDetail{}

		if err = rows.Scan(&transaction.Urut, &transaction.CUSTID, &transaction.METODE, &transaction.TRXDATE, &transaction.NOREFF, &transaction.FIDBANK, &transaction.KDCHANNEL, &transaction.DEBET, &transaction.KREDIT, &transaction.REFFBANK, &transaction.TRANSNO); err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}

func (t *TransactionRepository) GetCountTransactionDetailByCustId(ctx context.Context, filter dto.FilterTransactionDetail) (int, error) {
	sql := `SELECT count(urut) FROM sccttran_cashless WHERE CUSTID = ?`

	args := []interface{}{filter.CustId}

	count := 0
	err := t.db.QueryRowxContext(ctx, sql, args...).Scan(&count)

	return count, err
}

func (t *TransactionRepository) CreateTransactionByCustId(ctx context.Context, req domain.Transaction) error {
	query := `INSERT INTO sccttran_cashless (CUSTID, METODE, TRXDATE, NOREFF, FIDBANK, KDCHANNEL, DEBET, KREDIT, REFFBANK, TRANSNO) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`

	args := []interface{}{req.CUSTID, req.METODE, util.GetCurrentDate(), req.NOREFF, req.FIDBANK, req.KDCHANNEL, req.DEBET, req.KREDIT, req.REFFBANK, req.TRANSNO}

	_, err := t.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) CreateTransactionTransByCustId(ctx context.Context, req domain.Transaction) error {
	query := `INSERT INTO sccttran (CUSTID, METODE, TRXDATE, NOREFF, KDCHANNEL, DEBET, KREDIT, TRANSNO) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`

	args := []interface{}{req.CUSTID, req.METODE, util.GetCurrentDate(), req.NOREFF, req.KDCHANNEL, req.DEBET, req.KREDIT, req.TRANSNO}

	_, err := t.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (t *TransactionRepository) GetCountTransactionNow(ctx context.Context) (int, error) {
	sql := `SELECT count(urut) FROM sccttran_cashless WHERE DATE_FORMAT(TRXDATE, '%Y-%m-%d') = ? AND METODE = 'TOP UP CASHLESS'`

	args := []interface{}{util.GetCurrentDate().Format("2006-01-02")}

	count := 0
	err := t.db.QueryRowxContext(ctx, sql, args...).Scan(&count)

	return count + 1, err
}

func (t *TransactionRepository) GetCountTransactionTransNow(ctx context.Context) (int, error) {
	sql := `SELECT count(urut) FROM sccttran WHERE DATE_FORMAT(TRXDATE, '%Y-%m-%d') = ? AND METODE = 'TOP UP CASHLESS'`

	args := []interface{}{util.GetCurrentDate().Format("2006-01-02")}

	count := 0
	err := t.db.QueryRowxContext(ctx, sql, args...).Scan(&count)

	return count + 1, err
}
