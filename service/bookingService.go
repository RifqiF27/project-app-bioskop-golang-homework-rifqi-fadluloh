package service

import (
	"cinema/model"
	"cinema/repository"
	"errors"
	"log"
)

type BookingService interface {
	BookSeat(booking model.Booking) (*model.Booking, error)
}
type bookingService struct {
	Repo     repository.BookingRepository
	userRepo repository.AuthRepository
}

func NewBookingService(repo repository.BookingRepository, user repository.AuthRepository) BookingService {
	return &bookingService{Repo: repo, userRepo: user}
}

func (s *bookingService) BookSeat(booking model.Booking) (*model.Booking, error) {

	booked, err := s.Repo.IsSeatBooked(booking.SeatID, booking.BookingDate, booking.BookingTime)
	if err != nil {
		return nil, err
	}
	if booked {
		return nil, errors.New("seat already booked")
	}

	booking.Status = "pending"
	err = s.Repo.CreateBooking(&booking)
	if err != nil {

		log.Printf("Failed to create booking: %v", err)
		return nil, errors.New("failed to create booking: " + err.Error())
	}
	return &booking, nil
}
