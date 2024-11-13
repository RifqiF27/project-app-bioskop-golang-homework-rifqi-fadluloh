package handler

import (
	"cinema/model"
	"cinema/service"
	"cinema/utils"
	"encoding/json"
	"net/http"
	"time"
)

type BookingHandler struct {
	Service service.BookingService
}

func NewBookingHandler(service service.BookingService) *BookingHandler {
	return &BookingHandler{Service: service}
}

func (h *BookingHandler) BookSeat(w http.ResponseWriter, r *http.Request) {
	var req struct {
		CinemaID      int    `json:"cinemaId"`
		SeatID        int    `json:"seatId"`
		Date          string `json:"date"`
		BookingTime   string `json:"time"`
		PaymentMethod string `json:"paymentMethod"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	userID, ok := r.Context().Value("userID").(int)
	if !ok {
		utils.SendJSONResponse(w, http.StatusUnauthorized, "Unauthorized access", nil)
		return
	}

	bookingDate, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid date format", nil)
		return
	}
	bookingTime, err := time.Parse("15:04", req.BookingTime)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, "Invalid time format", nil)
		return
	}

	booking := model.Booking{
		UserID:        userID,
		CinemaID:      req.CinemaID,
		SeatID:        req.SeatID,
		BookingDate:   bookingDate,
		BookingTime:   bookingTime,
		PaymentMethod: req.PaymentMethod,
	}

	createdBooking, err := h.Service.BookSeat(booking)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusConflict, err.Error(), nil)
		return
	}

	utils.SendJSONResponse(w, http.StatusCreated, "Booking confirmed.", createdBooking)
}
