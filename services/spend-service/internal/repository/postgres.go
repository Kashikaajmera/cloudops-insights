package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresRepository struct {
	pool *pgxpool.Pool
}

func NewPostgresRepository(
	pool *pgxpool.Pool,
) *PostgresRepository {
	return &PostgresRepository{
		pool: pool,
	}
}

func (r *PostgresRepository) GetSummary(
	ctx context.Context,
	organizationID int64,
	currency string,
) (SpendSummary, error) {

	const query = `
        SELECT
            COALESCE(SUM(cost_cents), 0),
            COUNT(*)
        FROM cost_records
        WHERE organization_id = $1
          AND currency = $2
    `

	var result SpendSummary

	err := r.pool.QueryRow(
		ctx,
		query,
		organizationID,
		currency,
	).Scan(
		&result.TotalCostCents,
		&result.ResourceCount,
	)

	if err != nil {
		return SpendSummary{}, fmt.Errorf(
			"query spend summary: %w", err,
		)
	}

	result.Currency = currency

	return result, nil
}
