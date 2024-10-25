package bill

import (
	"context"
	"smart-payment-dev-be/internal/adapter/inbound/dto"
	"smart-payment-dev-be/internal/core/domain"
	"smart-payment-dev-be/internal/core/port/outbound/registry"
	"smart-payment-dev-be/shared/util"
)

type BillService struct {
	repositoryRegistry registry.RepositoryRegistry
}

func NewBillService(repositoryRegistry registry.RepositoryRegistry) *BillService {
	return &BillService{
		repositoryRegistry: repositoryRegistry,
	}
}

func (b *BillService) GetBillByCustId(ctx context.Context, filter dto.FilterBill) (interface{}, error) {
	var billDto dto.BillDto
	var bills []domain.Bill

	repo := b.repositoryRegistry.GetBillRepository()
	repoVPaymentDetail := b.repositoryRegistry.GetVPaymentDetailRepository()

	idCust, err := util.ParseJwtToken(filter.Token)

	if err != nil {
		return nil, err
	}

	filter.CustId = idCust

	dataBil, err := repo.GetBillByCustId(ctx, filter)
	if err != nil {
		return nil, err
	}

	for _, bill := range dataBil {
		resp, err := repoVPaymentDetail.GetVPaymentDetailByNoCust(ctx, bill.CUSTID, bill.BILLCD)
		if err != nil {
			return nil, err
		}
		bill.VPaymentDetail = resp
		bills = append(bills, bill)
	}

	count, _ := repo.GetCountBillByCustId(ctx, filter)
	resp := billDto.GetBillTransferOut(bills, count)

	return resp, nil
}
