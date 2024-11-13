package model

type PaymentMethod struct {
	ID         int    `json:"id"`
	MethodName string `json:"method_name"`
}

type PaymentDetails struct {
	CardNumber string `json:"cardNumber"`
	ExpiryDate string `json:"expiryDate"`
	CVV        string `json:"cvv"`
}
