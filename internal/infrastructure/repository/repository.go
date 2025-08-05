package repository

type Repository interface {
	Save(model any) error
	FindById(id int) (any, error)
	FindAll() ([]any, error)
	DeleteById(id int) error
}
