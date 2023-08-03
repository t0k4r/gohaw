package db

import "database/sql"

type Anime struct {
	Id          int
	Title       string
	Description string
	MalUrl      string
	Cover       string
	TypeOfId    *int
	TypeOf      *string
	SeasonId    *int
	Season      *string
}

func AnimeById(DB *sql.DB, id int) (*Anime, error) {
	row := DB.QueryRow(`
	select a.id, a.title, a.description, a.mal_url, a.cover, t.id, t.type_of, s.id, s.season from animes a
	left join seasons s on s.id = a.season_id
	left join anime_types t on t.id = a.type_id
	where a.id = $1`, id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var a Anime
	err := row.Scan(&a.Id, &a.Title, &a.Description, &a.MalUrl, &a.Cover, &a.TypeOfId, &a.TypeOf, &a.SeasonId, &a.Season)
	return &a, err
}
