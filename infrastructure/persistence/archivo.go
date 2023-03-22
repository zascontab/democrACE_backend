package persistence

import (
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"gorm.io/gorm"
)

var archivoRepository repository.ArchivoRepository

type ArchivoRepository struct {
	db *gorm.DB
}

func NewArchivoRepository(db *gorm.DB) repository.ArchivoRepository {
	if archivoRepository == nil {
		archivoRepository = &ArchivoRepository{db: db}
	}
	return archivoRepository
}

func (r *ArchivoRepository) Get(id uint) (*domain.Archivo, error) {
	return Get[domain.Archivo](r.db, id)
}

func (r *ArchivoRepository) GetAll(f *domain.ArchivoFilter) ([]*domain.Archivo, error) {
	return GetAll[domain.Archivo](r.db, f, r.BuildQuery)
}

func (r *ArchivoRepository) Save(archivo *domain.Archivo) error {
	if err := r.db.Create(&archivo).Error; err != nil {
		return err
	}

	return nil
}

func (r *ArchivoRepository) Update(updates *domain.Archivo) error {
	if err := r.db.Updates(updates).Error; err != nil {
		return err
	}

	return nil
}

func (r *ArchivoRepository) Delete(id uint) error {
	return Delete(r.db, domain.Archivo{}, id)
}

func (r *ArchivoRepository) BuildQuery(q *gorm.DB, f *domain.ArchivoFilter) *gorm.DB {
	if f != nil {
		if f.IDs != nil && len(*f.IDs) > 0 {
			q = q.Where("id IN (?)", *f.IDs)
		}
		if f.Limit != nil {
			q = q.Limit(*f.Limit)
		}
		if f.Offset != nil {
			q = q.Offset(*f.Offset)
		}
		if f.Nombre != nil && *f.Nombre != "" {
			q = q.Where("nombre LIKE ?", "%"+*f.Nombre+"%")
		}
		if f.ProgramaID != nil && *f.ProgramaID != 0 {
			q = q.Where("programa_id = ?", *f.ProgramaID)
		}
		if f.ProyectoID != nil && *f.ProyectoID != 0 {
			q = q.Where("proyecto_id = ?", *f.ProyectoID)
		}
		if f.ActividadID != nil && *f.ActividadID != 0 {
			q = q.Where("actividad_id = ?", *f.ActividadID)
		}
		if f.EjecucionID != nil && *f.EjecucionID != 0 {
			q = q.Where("ejecucion_id = ?", *f.EjecucionID)
		}
	}
	return q
}
