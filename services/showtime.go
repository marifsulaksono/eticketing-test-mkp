package services

import (
	"context"

	"github.com/marifsulaksono/eticketing-test-mkp/entities"
	"github.com/marifsulaksono/eticketing-test-mkp/repositories"
)

type showtimeService struct {
	Repo repositories.ShowtimeRepository
}

type ShowtimeService interface {
	InsertNewShowtime(ctx context.Context, s entities.Showtime) error
	GetAllShowtimes(ctx context.Context) ([]entities.ShowtimeResponse, error)
	GetShowtimeById(ctx context.Context, id int) (entities.ShowtimeResponse, error)
	UpdateShowtime(ctx context.Context, id int, showtime entities.Showtime) error
	DeleteShowtime(ctx context.Context, id int) error
}

func NewShowtimeService(r repositories.ShowtimeRepository) *showtimeService {
	return &showtimeService{Repo: r}
}

func (st *showtimeService) InsertNewShowtime(ctx context.Context, showtime entities.Showtime) error {
	return st.Repo.InsertNewShowtime(ctx, showtime)
}

func (st *showtimeService) GetAllShowtimes(ctx context.Context) ([]entities.ShowtimeResponse, error) {
	return st.Repo.GetAllShowtimes(ctx)
}

func (st *showtimeService) GetShowtimeById(ctx context.Context, id int) (entities.ShowtimeResponse, error) {
	return st.Repo.GetShowtimeById(ctx, id)
}

func (st *showtimeService) UpdateShowtime(ctx context.Context, id int, showtime entities.Showtime) error {
	return st.Repo.UpdateShowtime(ctx, id, showtime)
}

func (st *showtimeService) DeleteShowtime(ctx context.Context, id int) error {
	return st.Repo.DeleteShowtime(ctx, id)
}
