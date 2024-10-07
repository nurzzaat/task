package repository

import (
	"context"
	"database/sql"

	"github.com/nurzzaat/task/internal/models"
)

type songRepository struct {
	db *sql.DB
}

func NewSongRepository(db *sql.DB) models.SongRepository {
	return &songRepository{db: db}
}

func (sr *songRepository) CreateSong(c context.Context, songRequest models.SongRequest) (int, error) {
	var id int
	query := `INSERT INTO songs(group_name, name, release_date, lyric, link) VALUES ($1, $2, $3 , $4 ,$5) returning id;`
	if err := sr.db.QueryRow(query, songRequest.Group, songRequest.Name, songRequest.ReleaseDate, songRequest.Lyric, songRequest.Link).Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
func (sr *songRepository) UpdateSong(c context.Context, songID uint, songRequest models.SongRequest) error {
	query := `UPDATE songs SET group_name=$1, name=$2, release_date=$3, lyric=$4, link=$5 WHERE id = $6 `
	if _, err := sr.db.Exec(query, songRequest.Group, songRequest.Name, songRequest.ReleaseDate, songRequest.Lyric, songRequest.Link, songID); err != nil {
		return err
	}
	return nil
}
func (sr *songRepository) DeleteSong(c context.Context, songID uint) error {
	query := `DELETE FROM songs WHERE id =$1`
	if _, err := sr.db.Exec(query, songID); err != nil {
		return err
	}
	return nil
}
func (sr *songRepository) GetAll(c context.Context, properties models.Properties) ([]models.SongResponse, int, error) {
	var count int
	songs := []models.SongResponse{}
	countQuery := `SELECT COUNT(id)
			FROM songs 
			WHERE release_date >= $1 
			AND release_date <= $2 
			AND group_name ILIKE $3 AND name ILIKE $4 AND lyric ILIKE $5 AND link ILIKE $6`
	if err := sr.db.QueryRow(countQuery, properties.From, properties.To, properties.Group, properties.Song, properties.Lyric, properties.Link).Scan(&count); err != nil {
		return songs, count, err
	}
	query := `SELECT id , group_name, name, release_date, lyric, link 
			FROM songs 
			WHERE release_date >= $1 
			AND release_date <= $2 
			AND group_name ILIKE $3 AND name ILIKE $4 AND lyric ILIKE $5 AND link ILIKE $6
			ORDER BY id desc
			LIMIT $7 OFFSET $8;`
	rows, err := sr.db.Query(query, properties.From, properties.To, properties.Group, properties.Song, properties.Lyric, properties.Link, properties.Size, properties.Page)
	if err != nil {
		return songs, count, err
	}
	for rows.Next() {
		song := models.SongResponse{}
		if err := rows.Scan(&song.ID, &song.Group, &song.Name, &song.ReleaseDate, &song.Lyric, &song.Link); err != nil {
			return songs, count, err
		}
		songs = append(songs, song)
	}
	return songs, count, err
}
func (sr *songRepository) GetByID(c context.Context, songID uint) (models.SongResponse, error) {
	song := models.SongResponse{}
	query := `SELECT id , group_name, name, release_date, lyric, link 
			FROM songs 
			WHERE id =$1`
	if err := sr.db.QueryRow(query, songID).Scan(&song.ID, &song.Group, &song.Name, &song.ReleaseDate, &song.Lyric, &song.Link); err != nil {
		return song, err
	}
	return song, nil
}
