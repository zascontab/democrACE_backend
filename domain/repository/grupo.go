package repository

import "github.com/sonderkevin/governance/domain"

type GrupoRepository interface {
	Get(uint) (*domain.Grupo, error)
	GetAll(*domain.GrupoFilter) ([]*domain.Grupo, error)
	Save(*domain.Grupo) error
	Update(*domain.Grupo) error
	Delete(uint) error
}
