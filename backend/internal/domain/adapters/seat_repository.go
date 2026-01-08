package adapters

import (
	"context"
	"errors"

	"github.com/qairlines/internal/domain/entities"
)

var ErrSeatNotFound = errors.New("seat not found")

type ISeatRepository interface {
	GetSeatByID(ctx context.Context, seatID int64) (*entities.Seat, error)
}
