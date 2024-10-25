package repository

import (
	"context"
	"smart-payment-dev-be/internal/core/domain"
)

type UserRepository struct {
	db DBExecutor
}

func NewUserRepository(db DBExecutor) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) GetUserByNoCust(ctx context.Context, nocust string) (domain.User, error) {
	var user domain.User
	args := []interface{}{nocust}
	sql := `SELECT CUSTID, COALESCE(NOCUST, ''),COALESCE(NMCUST, ''), COALESCE(NUM2ND, ''), COALESCE(STCUST, ''), COALESCE(CODE01, ''), COALESCE(DESC01, ''), COALESCE(CODE02, ''), COALESCE(DESC02, ''), COALESCE(CODE03, ''), COALESCE(DESC03, ''), COALESCE(CODE04, ''), COALESCE(DESC04, ''), COALESCE(CODE05, ''), COALESCE(DESC05, ''), COALESCE(TOTPAY, ''), COALESCE(GENUS, ''), COALESCE(LastUpdate, ''), COALESCE(NO_WA, ''), COALESCE(MOBILEPASSWORD, '') FROM scctcust WHERE NOCUST = ?`

	err := u.db.QueryRowxContext(ctx, sql, args...).Scan(&user.CUSTID, &user.NOCUST, &user.NMCUST, &user.NUM2ND, &user.STCUST, &user.CODE01, &user.DESC01, &user.CODE02, &user.DESC02, &user.CODE03, &user.DESC03, &user.CODE04, &user.DESC04, &user.CODE05, &user.DESC05, &user.TOTPAY, &user.GENUS, &user.LastUpdate, &user.NO_WA, &user.MOBILEPASSWORD)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (u *UserRepository) GetUserByCustId(ctx context.Context, custId string) (domain.User, error) {
	var user domain.User
	sql := `SELECT CUSTID, COALESCE(NOCUST, ''),COALESCE(NMCUST, ''), COALESCE(NUM2ND, ''), COALESCE(STCUST, ''), COALESCE(CODE01, ''), COALESCE(DESC01, ''), COALESCE(CODE02, ''), COALESCE(DESC02, ''), COALESCE(CODE03, ''), COALESCE(DESC03, ''), COALESCE(CODE04, ''), COALESCE(DESC04, ''), COALESCE(CODE05, ''), COALESCE(DESC05, ''), COALESCE(TOTPAY, ''), COALESCE(GENUS, ''), COALESCE(LastUpdate, ''), COALESCE(NO_WA, ''), COALESCE(MOBILEPASSWORD, '') FROM scctcust WHERE CUSTID = ?`
	args := []interface{}{custId}

	err := u.db.QueryRowxContext(ctx, sql, args...).Scan(&user.CUSTID, &user.NOCUST, &user.NMCUST, &user.NUM2ND, &user.STCUST, &user.CODE01, &user.DESC01, &user.CODE02, &user.DESC02, &user.CODE03, &user.DESC03, &user.CODE04, &user.DESC04, &user.CODE05, &user.DESC05, &user.TOTPAY, &user.GENUS, &user.LastUpdate, &user.NO_WA, &user.MOBILEPASSWORD)
	if err != nil {
		return user, err
	}

	return user, nil
}
