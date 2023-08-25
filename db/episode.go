package db

import "database/sql"

type Episode struct {
	Id       int
	Index    int
	Title    *string
	AltTitle *string
	Stream   *string
}

func (Episode) scan(rows *sql.Rows) (dbObject, error) {
	var e Episode
	err := rows.Scan(&e.Id, &e.Index, &e.Title, &e.AltTitle, &e.Stream)
	return e, err
}

func EpisodesFromAnimeId(animeId int) ([]Episode, error) {
	episodes, err := query[Episode](`
	select e.id, e.index_of, e.title, e.alt_title, es.stream from episodes e
	left join episode_streams es on es.episode_id = e.id
	where e.anime_id = $1
	order by e.index_of`, animeId)
	return episodes, err
}
