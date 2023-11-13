package db

import (
	"database/sql"

	"github.com/t0k4r/qb"
)

type Info struct {
	TypeId int
	TypeOf string
	Id     int
	Value  string
}

func SelectInfo() *qb.QSelect {
	return qb.
		Select("infos i").
		Cols("it.id, it.type_of, i.id, i.info").
		Join("info_types it", "it.id = i.type_id")
}

func (i Info) Scan(rows *sql.Rows) (qb.Selectable, error) {
	return i, rows.Scan(&i.TypeId, &i.TypeOf, &i.Id, &i.Value)
}

type OrderedInfo struct {
	Id     int
	TypeOf string
	Values []OrderedInfoValue
}

type OrderedInfoValue struct {
	Id    int
	Value string
}

func OrderInfos(infos []Info) []OrderedInfo {
	var oinfos []OrderedInfo
	for _, i := range infos {
		broken := false
		for _, oi := range oinfos {
			if oi.Id == i.TypeId {
				oi.Values = append(oi.Values, OrderedInfoValue{Id: i.Id, Value: i.Value})
				broken = true
				break
			}
		}
		if !broken {
			oinfos = append(oinfos, OrderedInfo{
				Id:     i.TypeId,
				TypeOf: i.TypeOf,
				Values: []OrderedInfoValue{{Id: i.Id, Value: i.Value}},
			})
		}
	}
	return oinfos
}

func InfosFromAnimeId(id int) ([]OrderedInfo, error) {
	infos, err := Query[Info](SelectInfo().Join("anime_infos ai", "ai.info_id = i.id").Where("ai.anime_id = $1"), id)
	if err != nil {
		return nil, err
	}
	return OrderInfos(infos), nil
}

// type Infos struct {
// 	TypeId    int
// 	TypeTitle string
// 	Values    []Info
// }

// type Info struct {
// 	Id    int
// 	Value string
// }

// func InfosOfTypeFromAnimeId(animeId int, typeOf string) (*Infos, error) {
// 	rows, err := DB.Query(`
// 	select it.id, i.id ,i.info from anime_infos ai
// 	join infos i on i.id = ai.info_id
// 	join info_types it on it.id = i.type_id
// 	where ai.anime_id = $1 and it.type_of = $2
// 	`, animeId, typeOf)
// 	if err != nil {
// 		return nil, err
// 	}
// 	var inf Infos
// 	inf.TypeTitle = typeOf
// 	for rows.Next() {
// 		var i Info
// 		err = rows.Scan(&inf.TypeId, &i.Id, &i.Value)
// 		if err != nil {
// 			return nil, err
// 		}
// 		inf.Values = append(inf.Values, i)
// 	}
// 	return &inf, nil
// }
