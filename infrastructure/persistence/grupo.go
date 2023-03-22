package persistence

import (
	"errors"

	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"gorm.io/gorm"
)

var grupoRepository repository.GrupoRepository

type GrupoRepository struct {
	db *gorm.DB
}

func NewGrupoRepository(db *gorm.DB) repository.GrupoRepository {
	if grupoRepository == nil {
		grupoRepository = &GrupoRepository{db: db}
	}
	return grupoRepository
}

func (r *GrupoRepository) Get(id uint) (*domain.Grupo, error) {
	return Get[domain.Grupo](r.db, id)
}

func (r *GrupoRepository) GetAll(f *domain.GrupoFilter) ([]*domain.Grupo, error) {
	return GetAll[domain.Grupo](r.db, f, r.BuildQuery)
}

func (r *GrupoRepository) Save(grupo *domain.Grupo) error {
	if err := r.db.Create(grupo).Error; err != nil {
		return err
	}
	return nil
}

func (r *GrupoRepository) Update(updates *domain.Grupo) error {
	g := domain.Grupo{}
	if err := r.db.Where("id = ?", updates.ID).Find(&g).Error; err != nil {
		return err
	}
	if g.Nombre == "Admin" {
		return errors.New("no puede eliminar grupo Admin")
	}

	if updates.Permisos != nil {
		if err := r.db.Model(updates).Association("Permisos").Replace(updates.Permisos); err != nil {
			return err
		}
	}
	if updates.Usuarios != nil {
		if err := r.db.Model(updates).Association("Usuarios").Replace(updates.Usuarios); err != nil {
			return err
		}
	}
	if err := r.db.Save(updates).Error; err != nil {
		return err
	}
	return nil
}

func (r *GrupoRepository) Delete(id uint) error {
	g := domain.Grupo{}
	if err := r.db.Where("id = ?", id).Find(&g).Error; err != nil {
		return err
	}
	if g.Nombre == "Admin" {
		return errors.New("no puede eliminar grupo Admin")
	}

	return r.db.Where("id = ?", id).Delete(&g).Error
}

func (r *GrupoRepository) BuildQuery(q *gorm.DB, f *domain.GrupoFilter) *gorm.DB {
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
		if f.Descripcion != nil {
			q = q.Where("descripcion LIKE ?", "%"+*f.Descripcion+"%")
		}
		if f.PermisoID != nil {
			q = q.Table("grupos").Joins("inner join grupo_permisos on grupo_permisos.grupo_id = grupos.id").Where("grupo_permisos.permiso_id = ?", *f.PermisoID).Group("grupos.id")
		}
		if f.UsuarioID != nil {
			q = q.Table("grupos").Joins("inner join usuarios on usuarios.grupo_id = grupos.id").Where("usuarios.id = ?", *f.UsuarioID).Group("grupos.id")
		}
	}
	return q
}
