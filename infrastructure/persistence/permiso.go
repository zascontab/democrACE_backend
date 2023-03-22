package persistence

import (
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"gorm.io/gorm"
)

var permisoRepository repository.PermisoRepository

type PermisoRepository struct {
	db *gorm.DB
}

func NewPermisoRepository(db *gorm.DB) repository.PermisoRepository {
	if permisoRepository == nil {
		permisoRepository = &PermisoRepository{db: db}
	}
	return permisoRepository
}

func (r *PermisoRepository) Get(id uint) (*domain.Permiso, error) {
	return Get[domain.Permiso](r.db, id)
}

func (r *PermisoRepository) GetAll(f *domain.PermisoFilter) ([]*domain.Permiso, error) {
	return GetAll[domain.Permiso](r.db, f, r.BuildQuery)
}

func (r *PermisoRepository) BuildQuery(q *gorm.DB, f *domain.PermisoFilter) *gorm.DB {
	if f != nil {
		if f.IDs != nil {
			q = q.Where("id IN ?", *f.IDs)
		}
		if f.Limit != nil {
			q = q.Limit(*f.Limit)
		}
		if f.Offset != nil {
			q = q.Offset(*f.Offset)
		}
		if f.Nombre != nil {
			q = q.Where("nombre LIKE ?", "%"+*f.Nombre+"%")
		}
		if f.GruposIDs != nil {
			q = q.Table("permisos").Joins("inner join grupo_permisos on grupo_permisos.permiso_id = permisos.id").Where("grupo_permisos.grupo_id IN (?)", *f.GruposIDs).Group("permisos.id")
		}
		if f.UsuariosIDs != nil {
			q = q.Table("permisos").Joins("inner join usuario_permisos on usuario_permisos.permiso_id = permisos.id").Where("usuario_permisos.usuario_id IN (?)", *f.UsuariosIDs).Group("permisos.id")
		}
	}
	return q
}
