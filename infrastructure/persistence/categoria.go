package persistence

import (
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"gorm.io/gorm"
)

var categoriaRepository repository.CategoriaRepository

type CategoriaRepository struct {
	db *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) repository.CategoriaRepository {
	if categoriaRepository == nil {
		categoriaRepository = &CategoriaRepository{db: db}
	}
	return categoriaRepository
}

func (r *CategoriaRepository) Get(id uint) (*domain.Categoria, error) {
	return Get[domain.Categoria](r.db, id)
}

func (r *CategoriaRepository) GetAll(f *domain.CategoriaFilter) ([]*domain.Categoria, error) {
	return GetAll[domain.Categoria](r.db, f, r.BuildQuery)
}

func (r *CategoriaRepository) Save(category *domain.Categoria) error {
	if err := r.db.Create(category).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoriaRepository) Update(updates *domain.Categoria) error {
	if err := r.db.Save(updates).Error; err != nil {
		return err
	}
	return nil
}

func (r *CategoriaRepository) Delete(id uint) error {
	o := domain.Categoria{}
	o.ID = id
	tx := r.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Delete(&o).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Delete(&domain.Categoria{PadreID: id}).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (r *CategoriaRepository) BuildQuery(q *gorm.DB, f *domain.CategoriaFilter) *gorm.DB {
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
		if f.Tipo != nil {
			q = q.Where("tipo LIKE ?", "%"+*f.Tipo+"%")
		}
		if f.PadreID != nil {
			q = q.Where("padre_id = ?", *f.PadreID)
		}
	}
	return q
}
