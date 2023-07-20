package pages

import "github.com/labstack/echo/v4"

type filter struct {
	Name  string
	Count int
}

type filterList struct {
	Title   string
	Filters []filter
}

func FilterList(title string) func(echo.Context) error {
	return func(c echo.Context) error {
		rows, err := DB.Query(`
		select count(ai.anime_id), i.info  from infos i
		join anime_infos ai ON ai.info_id = i.id 
		where i.type_id = (select it.id  from info_types it where it.type_of = $1)
		group by i.id
		order by count(ai.anime_id) desc`, title)
		if err != nil {
			return err
		}
		var fl filterList
		fl.Title = title
		for rows.Next() {
			var f filter
			err := rows.Scan(&f.Count, &f.Name)
			if err != nil {
				return err
			}
			fl.Filters = append(fl.Filters, f)
		}
		if isHx(c.Request()) {
			return c.Render(200, "FilterList", fl)
		}
		return c.Render(200, "pageFliterList.html", fl)
	}
}

func Types(c echo.Context) error {
	rows, err := DB.Query(`
	select count(a.id), at2.type_of  from  anime_types at2 
	join animes a on a.type_id = at2.id
	group by at2.id 
	order by  count(a.id) desc`)
	if err != nil {
		return err
	}
	var fl filterList
	fl.Title = "types"
	for rows.Next() {
		var f filter
		err := rows.Scan(&f.Count, &f.Name)
		if err != nil {
			return err
		}
		fl.Filters = append(fl.Filters, f)
	}
	if isHx(c.Request()) {
		return c.Render(200, "FilterList", fl)
	}
	return c.Render(200, "pageFliterList.html", fl)
}

func Seasons(c echo.Context) error {
	rows, err := DB.Query(`
	select count(a.id), s.season  from seasons s 
	join animes a on s.id = a.season_id 
	group by s.id 
	order by s.value desc`)
	if err != nil {
		return err
	}
	var fl filterList
	fl.Title = "seasons"
	for rows.Next() {
		var f filter
		err := rows.Scan(&f.Count, &f.Name)
		if err != nil {
			return err
		}
		fl.Filters = append(fl.Filters, f)
	}
	if isHx(c.Request()) {
		return c.Render(200, "FilterList", fl)
	}
	return c.Render(200, "pageFliterList.html", fl)
}
