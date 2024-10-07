package models

import "context"

type SongResponse struct {
	ID          uint   `json:"songId"`
	Group       string `json:"group"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
	Lyric        string `json:"text"`
	Link        string `json:"link"`
}

type SongRequest struct {
	Group       string `json:"group"`
	Name        string `json:"name"`
	ReleaseDate string `json:"releaseDate"`
	Lyric        string `json:"text"`
	Link        string `json:"link"`
}

type SongRepository interface {
	CreateSong(c context.Context, songRequest SongRequest) (int, error)
	UpdateSong(c context.Context, songID uint, songRequest SongRequest) error
	DeleteSong(c context.Context, songID uint) error
	GetAll(c context.Context, properties Properties) ([]SongResponse,int, error)
	GetByID(c context.Context, songID uint) (SongResponse, error)
}
