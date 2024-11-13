package model

import "time"

type Booking struct {
    ID            int       `json:"bookingId"`
    UserID        int       `json:"-"`
    SeatID        int       `json:"seatId"`
    CinemaID      int       `json:"cinemaId"`
    BookingDate   time.Time `json:"date"`
    BookingTime   time.Time `json:"time"`
    PaymentMethod string    `json:"paymentMethod"`
    Status        string    `json:"-"`
    CreatedAt     time.Time `json:"-"`
}
