package booking

import (
	"context"
	"errors"

	"github.com/internal/domain/adapters"
	"github.com/internal/domain/entities"
)

type IGetBookingUseCase interface {
	Execute(ctx context.Context, bookingID int64) (entities.Booking, []entities.Ticket, []entities.Ticket, error)
}

type GetBookingUseCase struct {
	bookingRepository adapters.IBookingRepository
}

func NewGetBookingUseCase(bookingRepository adapters.IBookingRepository) *GetBookingUseCase {
	return &GetBookingUseCase{
		bookingRepository: bookingRepository,
	}
}

func (u *GetBookingUseCase) Execute(ctx context.Context, bookingID int64) (entities.Booking, []entities.Ticket, []entities.Ticket, error) {
	booking, departureTickets, returnTickets, err := u.bookingRepository.GetBookingByID(ctx, bookingID)
	if err != nil {
		if errors.Is(err, adapters.ErrBookingNotFound) {
			return entities.Booking{}, nil, nil, adapters.ErrBookingNotFound
		}
		return entities.Booking{}, nil, nil, err
	}
	return booking, departureTickets, returnTickets, nil
}
