package repository

import (
	"cinema/model"
	"database/sql"
	"time"
)

type BookingRepository interface {
	CreateBooking(booking *model.Booking) error
	GetBookingByID(id int) (*model.Booking, error)
	IsSeatBooked(seatID int, date time.Time, time time.Time) (bool, error)
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) CreateBooking(booking *model.Booking) error {
	query := `INSERT INTO bookings (user_id, cinema_id, seat_id, booking_date, booking_time, payment_method, status) 
              VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`
	return r.db.QueryRow(query, booking.UserID, booking.CinemaID, booking.SeatID, booking.BookingDate, booking.BookingTime, booking.PaymentMethod, booking.Status).Scan(&booking.ID)
}

func (r *bookingRepository) GetBookingByID(id int) (*model.Booking, error) {
	query := `SELECT id, user_id, cinema_id, seat_id, booking_date, booking_time, payment_method, status, created_at FROM bookings WHERE id = $1`
	booking := &model.Booking{}
	err := r.db.QueryRow(query, id).Scan(&booking.ID, &booking.UserID, &booking.CinemaID, &booking.SeatID, &booking.BookingDate, &booking.BookingTime, &booking.PaymentMethod, &booking.Status, &booking.CreatedAt)
	if err != nil {
		return nil, err
	}
	return booking, nil
}

func (r *bookingRepository) IsSeatBooked(seatID int, date time.Time, time time.Time) (bool, error) {
	query := `SELECT COUNT(*) FROM bookings WHERE seat_id = $1 AND booking_date = $2 AND booking_time = $3 AND status = 'booked'`
	var count int
	err := r.db.QueryRow(query, seatID, date, time).Scan(&count)
	return count > 0, err
}
