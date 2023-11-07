package db

import (
	"database/sql"

	"github.com/t0k4r/qb"
)

type Episode struct {
	Id       int
	Index    int
	Title    *string
	AltTitle *string
	Stream   *string
}

func SelectEpisode() *qb.QSelect {
	return qb.
		Select("episodes e").
		Cols("e.id", "e.index_of", "e.title", "e.alt_title", "es.stream").
		LJoin("episode_streams es", "es.episode_id = e.id")
}

func (e Episode) Scan(rows *sql.Rows) (qb.Selectable, error) {
	return e, rows.Scan(&e.Id, &e.Index, &e.Title, &e.AltTitle, &e.Stream)

}

func EpisodesFromAnimeId(animeId int) ([]Episode, error) {
	return qb.Query[Episode](
		SelectEpisode().Where("e.anime_id = $1").OrderBy("e.index_of"), DB, animeId)
}
