package db

import (
	"database/sql"
	"time"
)

type Anime struct {
	Id          int
	Title       string
	Description string
	MalUrl      string
	Cover       *string
	TypeOfId    *int
	TypeOf      *string
	SeasonId    *int
	Season      *string
}

func (Anime) scan(rows *sql.Rows) (dbObject, error) {
	var a Anime
	err := rows.Scan(&a.Id, &a.Title, &a.Description, &a.MalUrl, &a.Cover, &a.TypeOfId, &a.TypeOf, &a.SeasonId, &a.Season)
	return a, err
}

func AnimeFromId(id int) (*Anime, error) {
	animes, err := query[Anime](`
	select a.id, a.title, a.description, a.mal_url, a.cover, t.id, t.type_of, s.id, s.season from animes a
	left join seasons s on s.id = a.season_id
	left join anime_types t on t.id = a.type_id
	where a.id = $1`, id)
	if len(animes) != 0 {
		return &animes[0], err
	}
	return nil, err
}

func AnimesNow(limit int) ([]Anime, error) {
	animes, err := query[Anime](`
	select a.id, a.title, a.description, a.mal_url, a.cover, t.id, t.type_of, s.id, s.season from animes a
	left join seasons s on s.id = a.season_id
	left join anime_types t on t.id = a.type_id
	where s.value < $1
	order by s.value desc
	limit $2`, time.Now(), limit)
	return animes, err
}

func AnimesFromInfoId(infoId int) ([]Anime, error) {
	animes, err := query[Anime](`
	select a.id, a.title, a.description, a.mal_url, a.cover, t.id, t.type_of, s.id, s.season from animes a
	left join anime_infos ai on ai.anime_id = a.id
	left join seasons s on s.id = a.season_id
	left join anime_types t on t.id = a.type_id
	where ai.info_id = $1`, infoId)
	return animes, err
}

func AnimesFromTypeId(typeId int) ([]Anime, error) {
	animes, err := query[Anime](`
	select a.id, a.title, a.description, a.mal_url, a.cover, t.id, t.type_of, s.id, s.season from animes a
	left join seasons s on s.id = a.season_id
	left join anime_types t on t.id = a.type_id
	where a.type_id = $1`, typeId)
	return animes, err
}

func AnimesFromSeasonId(seasonId int) ([]Anime, error) {
	animes, err := query[Anime](`
	select a.id, a.title, a.description, a.mal_url, a.cover, t.id, t.type_of, s.id, s.season from animes a
	left join seasons s on s.id = a.season_id
	left join anime_types t on t.id = a.type_id
	where a.season_id = $1`, seasonId)
	return animes, err
}
