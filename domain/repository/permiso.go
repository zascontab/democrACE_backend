package repository

import "github.com/sonderkevin/governance/domain"

type PermisoRepository interface {
	Get(uint) (*domain.Permiso, error)
	GetAll(*domain.PermisoFilter) ([]*domain.Permiso, error)
}
