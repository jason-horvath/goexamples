package config

var Repo *Repository

type Repository struct {
	App *AppConfig
}

// CreateRepo - returns a new repository for configuration
func CreateRepo(a *AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

// SetRepo - Takes a Repository and sets it to Repo
func SetRepo(r *Repository) {
	Repo = r
}
