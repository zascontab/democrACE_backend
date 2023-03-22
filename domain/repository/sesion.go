package repository

import "github.com/sonderkevin/governance/domain"

type SesionRepository interface {
	Get(uint) (*domain.Sesion, error)
	GetByUsuario(uint) ([]*domain.Sesion, error)
	Save(*domain.Sesion) error
	Update(*domain.Sesion) error
	Delete(uint, string) error
	RemoveExpiredSessions()
}
