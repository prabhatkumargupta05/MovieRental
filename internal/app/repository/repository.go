package repository

type Repository interface {
	GetEndPoint() (string, error)
}

type repository struct{}

func NewRepository() (Repository, error) {
	return &repository{}, nil
}

func (repo repository) GetEndPoint() (string, error) {
	_ = repo
	return "helloworld", nil
}
