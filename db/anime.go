package db

import (
	"database/sql"
	"time"

	"github.com/t0k4r/qb"
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

func SelectAnime() *qb.QSelect {
	return qb.
		Select("animes a").
		Cols("a.id", "a.title", "a.description", "a.mal_url", "a.cover", "t.id", "t.type_of", "s.id", "s.season").
		LJoin("seasons s", "s.id = a.season_id").
		LJoin("anime_types t", "t.id = a.type_id")
}

func (a Anime) Scan(rows *sql.Rows) (qb.Selectable, error) {
	return a, rows.Scan(&a.Id, &a.Title, &a.Description, &a.MalUrl, &a.Cover, &a.TypeOfId, &a.TypeOf, &a.SeasonId, &a.Season)
}

func AnimeFromId(id int) (*Anime, error) {
	return firstOrNil(
		qb.Query[Anime](
			SelectAnime().Where(" a.id = $1"), DB, id))
}

func AnimesNow(limit int) ([]Anime, error) {
	return qb.Query[Anime](
		SelectAnime().Where("s.value < $1").OrderBy("s.value desc").Limit("$2"), DB, time.Now(), limit)
}

func AnimesFromInfoId(infoId int) ([]Anime, error) {
	return qb.Query[Anime](
		SelectAnime().LJoin("anime_infos ai", "ai.anime_id = a.id").Where("ai.info_id = $1"), DB, infoId)
}

func AnimesFromTypeId(typeId int) ([]Anime, error) {
	return qb.Query[Anime](
		SelectAnime().Where("a.type_id = $1"), DB, typeId)
}

func AnimesFromSeasonId(seasonId int) ([]Anime, error) {
	return qb.Query[Anime](
		SelectAnime().Where("a.season_id = $1"), DB, seasonId)
}

func firstOrNil[T any](items []T, err error) (*T, error) {
	var i *T = nil
	if len(items) > 0 {
		i = &items[0]
	}
	return i, err
}
