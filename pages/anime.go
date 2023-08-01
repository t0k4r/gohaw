package pages

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"golang.org/x/sync/errgroup"
)

type info struct {
	Id     int
	Title  string
	Values []infoVal
}
type infoVal struct {
	Id    int
	Value string
}

type anime struct {
	Id          int
	Title       string
	Description string
	MalUrl      string
	Cover       string
	TypeOfId    *int
	TypeOf      *string
	SeasonId    *int
	Season      *string
	Infos       []info
	Episodes    []episode
}

func Anime(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.NoContent(400)
	}
	row := DB.QueryRow(`
	select a.id, a.title, a.description, a.mal_url, a.cover, t.id, t.type_of, s.id, s.season from animes a
	left join seasons s on s.id = a.season_id
	left join anime_types t on t.id = a.type_id
	where a.id = $1`, id)
	if row.Err() != nil {
		return row.Err()
	}
	var a anime
	err = row.Scan(&a.Id, &a.Title, &a.Description, &a.MalUrl, &a.Cover, &a.TypeOfId, &a.TypeOf, &a.SeasonId, &a.Season)
	if err != nil {
		return err
	}
	g := new(errgroup.Group)
	g.Go(appendInfo(&a, "genres"))
	g.Go(appendInfo(&a, "themes"))
	g.Go(appendInfo(&a, "demographics"))
	g.Go(appendInfo(&a, "studios"))
	g.Go(appendInfo(&a, "producers"))
	g.Go(appendEpisodes(&a))
	err = g.Wait()
	fmt.Println(a)
	if err != nil {
		return err
	}
	if isHx(c.Request()) {
		return c.Render(200, "Anime", a)
	}
	return c.Render(200, "pageAnime.html", a)
}

func appendInfo(a *anime, infoType string) func() error {
	return func() error {
		inf, err := getInfo(a.Id, infoType)
		if err != nil {
			return err
		}
		a.Infos = append(a.Infos, inf)
		return nil
	}
}

func getInfo(animeId int, infoType string) (info, error) {
	var inf info

	inf.Title = strings.ToUpper(string(infoType[0])) + infoType[1:]

	rows, err := DB.Query(`
	select it.id, i.id ,i.info from anime_infos ai
	join infos i on i.id = ai.info_id
	join info_types it on it.id = i.type_id
	where ai.anime_id = $1 and it.type_of = $2
	`, animeId, infoType)
	if err != nil {
		return inf, err
	}
	for rows.Next() {
		var i infoVal
		err = rows.Scan(&inf.Id, &i.Id, &i.Value)
		if err != nil {
			return inf, err
		}
		inf.Values = append(inf.Values, i)
	}
	return inf, nil
}
