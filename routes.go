package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/marifsulaksono/eticketing-test-mkp/controllers"
	"github.com/marifsulaksono/eticketing-test-mkp/middleware"
	"github.com/marifsulaksono/eticketing-test-mkp/repositories"
	"github.com/marifsulaksono/eticketing-test-mkp/services"
	"gorm.io/gorm"
)

func InitRoutes(db *gorm.DB) *mux.Router {
	// user dependency injection
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	// auth dependency injection
	authController := controllers.NewAuthController(userService)

	// showtime dependency injection
	showtimeRepo := repositories.NewShowtimeRepository(db)
	showtimeServ := services.NewShowtimeService(showtimeRepo)
	showCont := controllers.NewShowtimeController(showtimeServ)

	r := mux.NewRouter()

	r.HandleFunc("/api/register", authController.RegsiterHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/login", authController.LoginHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/showtimes", middleware.AuthMiddleware(showCont.GetShowtimesHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/showtimes/{id}", middleware.AuthMiddleware(showCont.GetShowtimeByIdHandler)).Methods(http.MethodGet)
	r.HandleFunc("/api/showtimes/{id}", middleware.AuthMiddleware(showCont.UpdateShowtimesHandler)).Methods(http.MethodPut)
	r.HandleFunc("/api/showtimes/{id}", middleware.AuthMiddleware(showCont.DeleteShowtimesHandler)).Methods(http.MethodDelete)

	return r
}
