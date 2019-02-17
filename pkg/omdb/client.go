package omdb

type Client interface {
	ByTitleBulk(titles ...string) (map[string]Movie, error)
	ByTitle(title string) (Movie, bool, error)
}

type NoOpClient struct{}

func (NoOpClient) ByTitleBulk(titles ...string) (map[string]Movie, error) {
	out := make(map[string]Movie)

	for _, t := range titles {
		out[t] = Movie{}
	}

	return out, nil
}

func (NoOpClient) ByTitle(title string) (Movie, bool, error) {
	return Movie{}, false, nil
}
