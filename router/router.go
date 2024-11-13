package router

import (
	"cinema/database"
	"cinema/handler"
	middleware_auth "cinema/middleware"
	"cinema/repository"
	"cinema/service"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"net/http"
)

func NewRouter() http.Handler {
	r := chi.NewRouter()
	db := database.NewPostgresDB()

	userRepo := repository.NewAuthRepositoryDb(db)
	srv := service.NewAuthService(userRepo)

	cinemaRepo := repository.NewCinemaRepositoryDb(db)
	srvCnms := service.NewCinemaService(cinemaRepo)

	bookingRepo := repository.NewBookingRepository(db)
	srvBooking := service.NewBookingService(bookingRepo, userRepo)

	paymentRepo := repository.NewPaymentRepository(db)
	srvPay := service.NewPaymentService(paymentRepo)

	h := handler.NewAuthHandler(srv)
	c := handler.NewCinemaHandler(srvCnms)
	b := handler.NewBookingHandler(srvBooking)
	p := handler.NewPaymentHandler(srvPay)

	r.Use(middleware.Logger)

	r.Group(func(r chi.Router) {
		r.Route("/api", func(r chi.Router) {
			r.Post("/register", h.RegisterHandler)
			r.Post("/login", h.LoginHandler)
			r.Post("/logout", h.LogoutHandler)

		})
	})

	r.Group(func(r chi.Router) {
		r.Route("/api/cinemas", func(r chi.Router) {
			r.Use(middleware_auth.AuthMiddleware(srv))
			r.Get("/", c.GetAllCinemas)
			r.Get("/{cinemaId}", c.GetCinemaByID)
			r.Get("/{cinemaId}/seats", c.GetSeats)
			r.Post("/booking", b.BookSeat)
			r.Get("/payment-methods", p.GetPaymentMethods)
			r.Post("/pay", p.ProcessPayment)

		})
	})

	return r
}
