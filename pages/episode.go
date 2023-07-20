package pages

type episode struct {
	Id       int
	Index    int
	Title    *string
	AltTitle *string
	Stream   *string
}

func appendEpisodes(a *anime) func() error {
	return func() error {
		episodes, err := getEpisodes(a.Id)
		if err != nil {
			return err
		}
		a.Episodes = episodes
		return nil
	}
}

func getEpisodes(animeId int) ([]episode, error) {
	var episodes []episode
	rows, err := DB.Query(`
	select e.id, e.index_of, e.title, e.alt_title, es.stream from episodes e
	left join episode_streams es on es.episode_id = e.id
	where e.anime_id = $1
	order by e.index_of`, animeId)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var e episode
		err := rows.Scan(&e.Id, &e.Index, &e.Title, &e.AltTitle, &e.Stream)
		if err != nil {
			return episodes, err
		}
		episodes = append(episodes, e)
	}
	return episodes, nil
}
