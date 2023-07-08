package pages

import (
	"html/template"
	"log"
	"net/http"
)

type Info struct {
	Count int
	Name  string
}

type TemplInfo struct {
	Title string
	Items []Info
}

func getTemplInfo(info string) TemplInfo {
	rows, err := DB.Query(`
	select count(ai.anime_id), i.value  from infos i
	join anime_infos ai ON ai.info_id = i.id 
	where i.type_id = (select it.id  from info_types it where it.name_of = $1)
	group by i.id`, info)
	if err != nil {
		panic(err)
	}
	var templ TemplInfo
	templ.Title = info
	for rows.Next() {
		var inf Info
		err := rows.Scan(&inf.Count, &inf.Name)
		if err != nil {
			panic(err)
		}
		templ.Items = append(templ.Items, inf)
	}
	return templ
}

func infoList(info string) func(http.ResponseWriter, *http.Request) {
	templ, err := template.ParseFiles("templates/Layout.html", "templates/InfoList.html")
	if err != nil {
		log.Panic(err)
	}
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("content-type", "text/html")
		if isHx(r) {
			err := templ.ExecuteTemplate(w, "Main", getTemplInfo(info))
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := templ.Execute(w, getTemplInfo(info))
			if err != nil {
				log.Panic(err)
			}
		}
	}
}
