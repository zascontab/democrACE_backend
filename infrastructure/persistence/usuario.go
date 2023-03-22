package persistence

import (
	"errors"

	"github.com/sonderkevin/governance/config"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"gorm.io/gorm"
)

var usuarioRepository repository.UsuarioRepository

type UsuarioRepository struct {
	db *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) repository.UsuarioRepository {
	if usuarioRepository == nil {
		usuarioRepository = &UsuarioRepository{db: db}
	}
	return usuarioRepository
}

func (r *UsuarioRepository) Get(id uint) (*domain.Usuario, error) {
	return Get[domain.Usuario](r.db, id)
}

func (r *UsuarioRepository) GetByEmail(email string) (*domain.Usuario, error) {
	var usuario *domain.Usuario
	if err := r.db.First(&usuario, "email = ?", email).Error; err != nil {

		return nil, err
	}
	return usuario, nil
}

func (r *UsuarioRepository) GetAll(f *domain.UsuarioFilter) ([]*domain.Usuario, error) {
	return GetAll[domain.Usuario](r.db, f, r.BuildQuery)
}

func (r *UsuarioRepository) Save(usuario *domain.Usuario) error {
	if err := r.db.Create(usuario).Error; err != nil {
		return err
	}
	return nil
}

func (r *UsuarioRepository) Update(updates *domain.Usuario) error {
	u := domain.Usuario{}
	if err := r.db.Where("id = ?", updates.ID).Find(&u).Error; err != nil {
		return err
	}
	if u.NombreUsuario == config.ADMIN {
		return errors.New("no puede editar usuario " + config.ADMIN)
	}
	if updates.Permisos != nil {
		if err := r.db.Model(updates).Association("Permisos").Replace(updates.Permisos); err != nil {
			return err
		}
	}
	
	if err := r.db.Save(updates).Error; err != nil {
		return err
	}
	return nil
}

func (r *UsuarioRepository) Delete(id uint) error {
	u := domain.Usuario{}
	if err := r.db.Where("id = ?", id).Find(&u).Error; err != nil {
		return err
	}
	if u.NombreUsuario == config.ADMIN {
		return errors.New("no puede eliminar usuario " + config.ADMIN)
	}

	return r.db.Where("id = ?", id).Delete(&u).Error
}

func (r *UsuarioRepository) BuildQuery(q *gorm.DB, f *domain.UsuarioFilter) *gorm.DB {
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
		if f.Status != nil {
			q = q.Where("status = ?", *f.Status)
		}
		if f.Verificado != nil {
			q = q.Where("verificado = ?", *f.Verificado)
		}
		if f.Nombres != nil {
			q = q.Where("nombres LIKE ?", "%"+*f.Nombres+"%")
		}
		if f.Apellidos != nil {
			q = q.Where("apellidos LIKE ?", "%"+*f.Apellidos+"%")
		}
		if f.NombreUsuario != nil {
			q = q.Where("nombre_usuario = ?", *f.NombreUsuario)
		}
		if f.FotoPerfil != nil {
			q = q.Where("foto_perfil = ?", *f.FotoPerfil)
		}
		if f.Email != nil {
			q = q.Where("email = ?", *f.Email)
		}
		if f.GrupoID != nil {
			q = q.Where("grupo_id = ?", *f.GrupoID)
		}
		if f.PermisosIDs != nil && len(*f.PermisosIDs) > 0 {
			q = q.Table("usuarios").Joins("inner join usuario_permisos on usuario_permisos.usuario_id = usuarios.id").Where("usuario_permisos.permiso_id IN (?)", *f.PermisosIDs).Group("usuarios.id")
		}
		
	}
	return q
}
