package repository

import "github.com/sonderkevin/governance/domain"

type UsuarioRepository interface {
	Get(uint) (*domain.Usuario, error)
	GetByEmail(string) (*domain.Usuario, error)
	GetAll(*domain.UsuarioFilter) ([]*domain.Usuario, error)
	Save(*domain.Usuario) error
	Update(*domain.Usuario) error
	Delete(uint) error
}
