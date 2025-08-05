package repository

type Repository interface {
	Create(model any) error
	GetByID(id uint) (any, error)
	GetAll() (any, error)
	Update(model any) error
	Delete(id uint) error
}
