package repositories

import (
	"context"

	"github.com/marifsulaksono/eticketing-test-mkp/entities"
	"gorm.io/gorm"
)

type showtimeRepository struct {
	DB *gorm.DB
}

type ShowtimeRepository interface {
	InsertNewShowtime(ctx context.Context, showtime entities.Showtime) error
	GetAllShowtimes(ctx context.Context) ([]entities.Showtime, error)
	GetShowtimeById(ctx context.Context, id int) (entities.Showtime, error)
	UpdateShowtime(ctx context.Context, id int, showtime entities.Showtime) error
	DeleteShowtime(ctx context.Context, id int) error
}

func NewShowtimeRepository(db *gorm.DB) *showtimeRepository {
	return &showtimeRepository{DB: db}
}

func (st *showtimeRepository) InsertNewShowtime(ctx context.Context, showtime entities.Showtime) error {
	return st.DB.Create(&showtime).Error
}

func (st *showtimeRepository) GetAllShowtimes(ctx context.Context) ([]entities.Showtime, error) {
	var showtimes []entities.Showtime
	err := st.DB.Preload("Branch").Preload("Stage").Preload("Movie").Find(&showtimes).Error
	return showtimes, err
}

func (st *showtimeRepository) GetShowtimeById(ctx context.Context, id int) (entities.Showtime, error) {
	var showtime entities.Showtime
	err := st.DB.Where("id = ?", id).Preload("Branch").Preload("Stage").Preload("Movie").First(&showtime).Error
	return showtime, err
}

func (st *showtimeRepository) UpdateShowtime(ctx context.Context, id int, showtime entities.Showtime) error {
	return st.DB.Model(entities.Showtime{}).Where("id = ?", id).Updates(showtime).Error
}

func (st *showtimeRepository) DeleteShowtime(ctx context.Context, id int) error {
	return st.DB.Where("id = ?", id).Delete(&entities.Showtime{}).Error
}
