package repository

type Repository interface {
	GetEndPoint() (string, error)
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo repository) GetEndPoint() (string, error) {
	_ = repo
	return "helloworld", nil
}
