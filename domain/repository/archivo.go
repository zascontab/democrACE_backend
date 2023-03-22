package repository

import "github.com/sonderkevin/governance/domain"

type ArchivoRepository interface {
	Get(uint) (*domain.Archivo, error)
	GetAll(*domain.ArchivoFilter) ([]*domain.Archivo, error)
	Save(*domain.Archivo) error
	Update(*domain.Archivo) error
	Delete(uint) error
}
