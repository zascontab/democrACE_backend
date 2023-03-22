package persistence

import (
	"strconv"
	"strings"

	"gorm.io/gorm"
)

func Get[T any](db *gorm.DB, id uint) (*T, error) {
	var m T
	if err := db.First(&m, id).Error; err != nil {
		return nil, err
	}

	return &m, nil
}

func GetAll[D any, F any](q *gorm.DB, f *F, buildQuery func(*gorm.DB, *F) *gorm.DB) ([]*D, error) {
	var oo []*D

	q = buildQuery(q, f)

	if err := q.Find(&oo).Error; err != nil {
		return nil, err
	}

	return oo, nil
}

func Delete[T any](db *gorm.DB, model T, id uint) error {
	return db.Where("id = ?", id).Delete(&model).Error
}

func InterpretRangoPresupuesto(q *gorm.DB, rangoPresupuesto string) error {
	rango := strings.Split(rangoPresupuesto, ";")
	for _, r := range rango {
		r = strings.TrimSpace(r)
		if strings.HasPrefix(r, ">=") {
			value, err := strconv.ParseFloat(strings.TrimLeft(r, ">="), 64)
			if err != nil {
				return err
			}
			q = q.Where("presupuesto_planificado >= ?", value)
		} else if strings.HasPrefix(r, "<=") {
			value, err := strconv.ParseFloat(strings.TrimLeft(r, "<-"), 64)
			if err != nil {
				return err
			}
			q = q.Where("presupuesto_planificado <= ?", value)
		}
	}
	return nil
}
