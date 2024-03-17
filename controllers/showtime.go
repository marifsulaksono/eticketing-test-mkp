package controllers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/marifsulaksono/eticketing-test-mkp/entities"
	"github.com/marifsulaksono/eticketing-test-mkp/services"
	"gorm.io/gorm"
)

type ShowtimeController struct {
	Service services.ShowtimeService
}

func NewShowtimeController(s services.ShowtimeService) *ShowtimeController {
	return &ShowtimeController{Service: s}
}

func (st *ShowtimeController) InsertNewShowtimeHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		showtime entities.Showtime
	)

	err := json.NewDecoder(r.Body).Decode(&showtime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	err = st.Service.InsertNewShowtime(ctx, showtime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Data berhasil disimpan")
}

func (st *ShowtimeController) GetShowtimesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	showtimes, err := st.Service.GetAllShowtimes(ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(showtimes)
}

func (st *ShowtimeController) GetShowtimeByIdHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	showtime, err := st.Service.GetShowtimeById(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			http.Error(w, "Data tidak ditemukan", http.StatusNotFound)
			return
		}

		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(showtime)
}

func (st *ShowtimeController) UpdateShowtimesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx      = r.Context()
		showtime entities.Showtime
	)

	err := json.NewDecoder(r.Body).Decode(&showtime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = st.Service.UpdateShowtime(ctx, id, showtime)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Data berhasil diubah")
}

func (st *ShowtimeController) DeleteShowtimesHandler(w http.ResponseWriter, r *http.Request) {
	var (
		ctx = r.Context()
	)

	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = st.Service.DeleteShowtime(ctx, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Data berhasil dihapus")
}
