package repository

import "context"

type SpendSummary struct {
	TotalCostCents int64
	ResourceCount  int
	Currency       string
}

type SpendRepository interface {
	GetSummary(
		ctx context.Context,
		organizationID int64,
		currency string,
	) (SpendSummary, error)
}
