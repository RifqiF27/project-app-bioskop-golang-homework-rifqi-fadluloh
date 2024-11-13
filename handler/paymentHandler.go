package handler

import (
	"cinema/model"
	"cinema/service"
	"cinema/utils"
	"encoding/json"
	"net/http"
	"strconv"
)

type PaymentHandler struct {
	paymentService *service.PaymentService
}

func NewPaymentHandler(paymentService *service.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService}
}

func (h *PaymentHandler) GetPaymentMethods(w http.ResponseWriter, r *http.Request) {

	paymentMethods, err := h.paymentService.GetPaymentMethods()
	if err != nil {
		utils.SendJSONResponse(w, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	utils.SendJSONResponse(w, http.StatusOK, "", paymentMethods)
}

func (h *PaymentHandler) ProcessPayment(w http.ResponseWriter, r *http.Request) {
	var req struct {
		BookingID      string               `json:"bookingId"`
		PaymentMethod  string               `json:"paymentMethod"`
		PaymentDetails model.PaymentDetails `json:"paymentDetails"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	bookingID, err := strconv.Atoi(req.BookingID)
	if err != nil {
		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	err = h.paymentService.ProcessPayment(bookingID, req.PaymentMethod)
	if err != nil {

		utils.SendJSONResponse(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	response := map[string]interface{}{
		"statusCode":    200,
		"message":       "Payment successful.",
		"transactionId": "txn" + req.BookingID,
		"bookingId":     req.BookingID,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
