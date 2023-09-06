package db

type Filter struct {
	Id    int
	Name  string
	Count int
}

type Filters struct {
	Title   string
	Filters []Filter
}

func Filtes(title string) (*Filters, error) {
	rows, err := DB.Query(`
	select count(ai.anime_id), i.info, i.id  from infos i
	join anime_infos ai ON ai.info_id = i.id 
	where i.type_id = (select it.id  from info_types it where it.type_of = $1)
	group by i.id
	order by count(ai.anime_id) desc`, title)
	if err != nil {
		return nil, err
	}
	var fl Filters
	fl.Title = title
	for rows.Next() {
		var f Filter
		err := rows.Scan(&f.Count, &f.Name, &f.Id)
		if err != nil {
			return &fl, err
		}
		fl.Filters = append(fl.Filters, f)
	}
	return &fl, nil
}

func Types() (*Filters, error) {
	rows, err := DB.Query(`
	select count(a.id), at2.type_of  from  anime_types at2 
	join animes a on a.type_id = at2.id
	group by at2.id 
	order by  count(a.id) desc`)
	if err != nil {
		return nil, err
	}
	var fl Filters
	fl.Title = "types"
	for rows.Next() {
		var f Filter
		err := rows.Scan(&f.Count, &f.Name)
		if err != nil {
			return &fl, err
		}
		fl.Filters = append(fl.Filters, f)
	}
	return &fl, nil
}

func Seasons() (*Filters, error) {
	rows, err := DB.Query(`
	select count(a.id), s.season  from seasons s 
	join animes a on s.id = a.season_id 
	group by s.id 
	order by s.value desc`)
	if err != nil {
		return nil, err
	}
	var fl Filters
	fl.Title = "seasons"
	for rows.Next() {
		var f Filter
		err := rows.Scan(&f.Count, &f.Name)
		if err != nil {
			return &fl, err
		}
		fl.Filters = append(fl.Filters, f)
	}
	return &fl, nil
}
