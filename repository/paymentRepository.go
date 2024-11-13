package repository

import (
	"cinema/model"
	"database/sql"
	"errors"
	"log"
)

type PaymentRepository interface {
	GetAllPaymentMethods() ([]model.PaymentMethod, error)
	ProcessPayment(bookingID int, paymentMethod string) error
}

type paymentRepository struct {
	db *sql.DB
}

func NewPaymentRepository(db *sql.DB) PaymentRepository {
	return &paymentRepository{db}
}

func (r *paymentRepository) GetAllPaymentMethods() ([]model.PaymentMethod, error) {
    query := `SELECT id, method_name FROM payment_methods`
    rows, err := r.db.Query(query)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var methods []model.PaymentMethod
    for rows.Next() {
        var method model.PaymentMethod
        if err := rows.Scan(&method.ID, &method.MethodName); err != nil {
            log.Println("Error scanning row:", err)
            return nil, err
        }
        methods = append(methods, method)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
    return methods, nil
}

func (r *paymentRepository) ProcessPayment(bookingID int, paymentMethod string) error {
	query := `UPDATE bookings SET status = 'completed' WHERE id = $1`
	result, err := r.db.Exec(query, bookingID)
	if err != nil {
		return err
	}
	if rows, _ := result.RowsAffected(); rows == 0 {
		return errors.New("no booking found to update")
	}
	return nil
}
