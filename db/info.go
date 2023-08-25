package db

type Infos struct {
	TypeId    int
	TypeTitle string
	Values    []Info
}

type Info struct {
	Id    int
	Value string
}

func InfosOfTypeFromAnimeId(animeId int, typeOf string) (*Infos, error) {
	rows, err := DB.Query(`
	select it.id, i.id ,i.info from anime_infos ai
	join infos i on i.id = ai.info_id
	join info_types it on it.id = i.type_id
	where ai.anime_id = $1 and it.type_of = $2
	`, animeId, typeOf)
	if err != nil {
		return nil, err
	}
	var inf Infos
	inf.TypeTitle = typeOf
	for rows.Next() {
		var i Info
		err = rows.Scan(&inf.TypeId, &i.Id, &i.Value)
		if err != nil {
			return nil, err
		}
		inf.Values = append(inf.Values, i)
	}
	return &inf, nil
}
