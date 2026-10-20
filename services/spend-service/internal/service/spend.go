package service

import (
	"context"
	"fmt"

	"github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/generated"
	"github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/repository"
)

type SpendService struct {
	repo repository.SpendRepository
}

func NewSpendService(repo repository.SpendRepository) *SpendService {
	return &SpendService{repo: repo}
}

func (s *SpendService) GetSummary(
	ctx context.Context,
	organizationID int64,
) (generated.SpendSummary, error) {

	const currency = "USD"

	result, err := s.repo.GetSummary(ctx, organizationID, currency)
	if err != nil {
		return generated.SpendSummary{}, fmt.Errorf(
			"get spend summary: %w", err,
		)
	}

	return generated.SpendSummary{
		Currency: result.Currency,
		TotalCost: fmt.Sprintf(
			"%d.%02d",
			result.TotalCostCents/100,
			result.TotalCostCents%100,
		),
		ResourceCount: result.ResourceCount,
	}, nil
}
