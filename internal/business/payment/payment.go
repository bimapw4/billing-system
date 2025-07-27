package payment

import (
	"billing/internal/common"
	"billing/internal/entity"
	"billing/internal/presentations"
	"billing/internal/repositories"
	"billing/pkg/meta"
	"context"
	"time"
)

type Payment interface {
	Payment(ctx context.Context, payload *entity.Payment) (*presentations.RespPayment, error)
	List(ctx context.Context, m *meta.Params) ([]presentations.Payment, error)
	IsDelinquent(ctx context.Context, loanID string) (bool, []int, error)
	GetOutstanding(ctx context.Context, loanID string) (int, error)
	Status(ctx context.Context, loanID string) (*presentations.PaymentStatus, error)
}

type business struct {
	repo *repositories.Repository
}

func NewBusiness(repo *repositories.Repository) Payment {
	return &business{
		repo: repo,
	}
}

func (b *business) Payment(ctx context.Context, payload *entity.Payment) (*presentations.RespPayment, error) {

	loans, err := b.repo.Loans.Detail(ctx, payload.LoanID)
	if err != nil {
		return nil, err
	}

	if loans.BorrowerID != payload.BorrowerID {
		return nil, presentations.ErrLoanNotFound
	}

	totalPayment := loans.Principal + int((float64(loans.Principal) * (float64(loans.InterestRate) / 100)))
	installments := totalPayment / loans.TotalWeeks

	weeks := common.GetWeekNumber(loans.StartDate, payload.Date)

	isDeliquent, unpaidWeeks, err := b.IsDelinquent(ctx, loans.ID)
	if err != nil {
		return nil, err
	}

	unpaidWeeks = append(unpaidWeeks, weeks)

	fixinstallment := len(unpaidWeeks) * installments

	if fixinstallment != payload.Payment {
		return &presentations.RespPayment{
			Deliquent:       isDeliquent,
			DeliquentAmount: fixinstallment,
		}, presentations.ErrPaymentInstallment
	}

	for _, v := range unpaidWeeks {
		data := presentations.RespPayment{
			Payment: presentations.Payment{
				Week:      v,
				LoanID:    payload.LoanID,
				Paid:      payload.Payment / len(unpaidWeeks),
				IsPaid:    true,
				PaidAt:    &payload.Date,
				UpdatedAt: time.Now(),
			},
		}

		err = b.repo.Payments.UpdatePayment(ctx, data.Payment)
		if err != nil {
			return nil, err
		}
	}

	return nil, nil
}

func (b *business) List(ctx context.Context, m *meta.Params) ([]presentations.Payment, error) {
	list, err := b.repo.Payments.List(ctx, m)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (b *business) IsDelinquent(ctx context.Context, loanID string) (bool, []int, error) {

	loans, err := b.repo.Loans.Detail(ctx, loanID)
	if err != nil {
		return false, nil, err
	}

	weeks := common.GetWeekNumber(loans.StartDate, time.Now())

	unpaidWeek := []int{}
	for i := 1; i <= weeks; i++ {

		payment, err := b.repo.Payments.FindByWeeksAndLoanID(ctx, i, loanID)
		if err != nil {
			return false, nil, err
		}

		if !payment.IsPaid {
			unpaidWeek = append(unpaidWeek, payment.Week)
		}
	}

	if len(unpaidWeek) >= 2 {
		return true, unpaidWeek, nil
	}

	return false, unpaidWeek, nil
}

func (b *business) GetOutstanding(ctx context.Context, loanID string) (int, error) {

	loans, err := b.repo.Loans.Detail(ctx, loanID)
	if err != nil {
		return 0, err
	}

	totalToPay := loans.Principal + int((float64(loans.Principal) * (float64(loans.InterestRate) / 100)))

	weeklyPayment := totalToPay / loans.TotalWeeks

	_, countUpaidWeek, err := b.IsDelinquent(ctx, loanID)
	if err != nil {
		return 0, err
	}

	outstanding := len(countUpaidWeek) * weeklyPayment

	return outstanding, nil
}

func (b *business) Status(ctx context.Context, loanID string) (*presentations.PaymentStatus, error) {

	isDeliquent, deliquentweek, err := b.IsDelinquent(ctx, loanID)
	if err != nil {
		return nil, err
	}

	outstanding, err := b.GetOutstanding(ctx, loanID)
	if err != nil {
		return nil, err
	}

	result := presentations.PaymentStatus{
		IsDelinquent:  isDeliquent,
		DeliquentWeek: deliquentweek,
		Outstanding:   outstanding,
	}

	return &result, nil
}
