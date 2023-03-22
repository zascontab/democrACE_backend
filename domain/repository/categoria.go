package repository

import "github.com/sonderkevin/governance/domain"

type CategoriaRepository interface {
	Get(uint) (*domain.Categoria, error)
	GetAll(*domain.CategoriaFilter) ([]*domain.Categoria, error)
	Save(*domain.Categoria) error
	Update(*domain.Categoria) error
	Delete(uint) error
}
