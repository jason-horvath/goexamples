package config

var Repo *Repository

type Repository struct {
	App *AppConfig
}

func CreateRepo(a *AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}
