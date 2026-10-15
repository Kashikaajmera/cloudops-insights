
package service

import (
    "fmt"

    "github.com/kashika-ajmera/cloudops-insights/services/spend-service/internal/generated"
)

type SpendService struct{}

func NewSpendService() *SpendService {
    return &SpendService{}
}

func (s *SpendService) GetSummary() generated.SpendSummary {
    costsInCents := []int64{15025, 10050, 20000}

    var total int64
    for _, cost := range costsInCents {
        total += cost
    }

    return generated.SpendSummary{
        Currency:      "USD",
        TotalCost:     fmt.Sprintf("%d.%02d", total/100, total%100),
        ResourceCount: len(costsInCents),
    }
}
