package service

import (
	"cinema/model"
	"cinema/repository"
	"errors"
)

type PaymentService struct {
	paymentRepo repository.PaymentRepository
}

func NewPaymentService(repo repository.PaymentRepository) *PaymentService {
	return &PaymentService{repo}
}

func (s *PaymentService) GetPaymentMethods() ([]model.PaymentMethod, error) {
	return s.paymentRepo.GetAllPaymentMethods()
}

func (s *PaymentService) ProcessPayment(bookingID int, paymentMethod string) error {
	if paymentMethod == "" {
        return errors.New("payment method is required")
    }
	
	return s.paymentRepo.ProcessPayment(bookingID, paymentMethod)
}