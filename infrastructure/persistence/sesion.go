package persistence

import (
	"fmt"
	"time"

	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"gorm.io/gorm"
)

var sesionRepository repository.SesionRepository

type SesionRepository struct {
	db *gorm.DB
}

func NewSesionRepository(db *gorm.DB) repository.SesionRepository {
	if sesionRepository == nil {
		sesionRepository = &SesionRepository{db: db}
	}
	return sesionRepository
}

func (r *SesionRepository) Get(id uint) (*domain.Sesion, error) {
	s := domain.Sesion{}
	if err := r.db.First(&s, id).Error; err != nil {
		return nil, err
	}
	return &s, nil
}

func (r *SesionRepository) GetByUsuario(id uint) ([]*domain.Sesion, error) {
	ss := []*domain.Sesion{}
	if err := r.db.Where("usuario_id = ?", id).Find(&ss).Error; err != nil {
		return nil, err
	}
	return ss, nil
}

func (r *SesionRepository) Save(sesion *domain.Sesion) error {
	if err := r.db.Create(sesion).Error; err != nil {
		return err
	}
	return nil
}

func (r *SesionRepository) Update(u *domain.Sesion) error {
	if err := r.db.Model(&domain.Sesion{}).Where("usuario_id = ? AND dispositivo = ?", u.UsuarioID, u.Dispositivo).Update("exp", u.Exp).Error; err != nil {
		return err
	}
	return nil
}

func (r *SesionRepository) Delete(usuarioID uint, device string) error {
	return r.db.Where("usuario_id = ? AND dispositivo = ?", usuarioID, device).Delete(&domain.Sesion{}).Error
}

func (r *SesionRepository) RemoveExpiredSessions() {
	t := time.Now()
	if err := r.db.Where("exp < ?", t).Delete(&domain.Sesion{}).Error; err != nil {
		fmt.Printf("error deleting sessions: %s", err.Error())
	}
}
