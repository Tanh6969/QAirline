package adapters

import (
	"context"

	"github.com/qairlines/internal/domain/entities"
)

type IHealthRepository interface {
	GetHealth(ctx context.Context) (entities.Health, error)
}
