package controllers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/marifsulaksono/eticketing-test-mkp/entities"
	"github.com/marifsulaksono/eticketing-test-mkp/services"
	"github.com/marifsulaksono/eticketing-test-mkp/utils"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthController struct {
	UserService services.UserService
}

func NewAuthController(us services.UserService) *AuthController {
	return &AuthController{UserService: us}
}

func (a *AuthController) RegsiterHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx  = r.Context()
		user entities.User
	)

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if err := a.UserService.CreateUser(ctx, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Register berhasil")
}

func (a *AuthController) LoginHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	// Basic Authentication
	username, password, ok := r.BasicAuth()
	if !ok {
		http.Error(w, "BasicAuth required", http.StatusUnauthorized)
		log.Println("BasicAuth required")
		return
	}

	// Check user validation on database
	var userLogin entities.User
	userLogin, err := a.UserService.GetUserByUsername(ctx, username)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Username tidak ditemukan", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Check password validation
	if err := bcrypt.CompareHashAndPassword([]byte(userLogin.Password), []byte(password)); err != nil {
		http.Error(w, "Username/password tidak sesuai", http.StatusUnauthorized)
		return
	}

	token, err := utils.GenerateToken(userLogin)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(token)
}
