package db

import (
	"fmt"

	"github.com/joho/godotenv"
	"github.com/sonderkevin/governance/application"
	"github.com/sonderkevin/governance/config"
	"github.com/sonderkevin/governance/domain"
	"golang.org/x/exp/slices"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

func New() *gorm.DB {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}

	if db != nil {
		return db
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DB_HOST, config.DB_PORT, config.DB_USER, config.DB_PASSWORD, config.DB_NAME)

	var err error
	db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Error),
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	if err != nil {
		panic("failed to connect database" + err.Error())
	}

	fmt.Println("Connected to the database!")

	return db
}

func Migrate(permisos []string) {
	if db == nil {
		return
	}

	// Auto-Migrate
	db.AutoMigrate(&domain.Sesion{}, &domain.Usuario{},
		&domain.Categoria{}, &domain.Archivo{}, &domain.Permiso{},
		&domain.Grupo{},
	)
	db.DisableForeignKeyConstraintWhenMigrating = false
	db.AutoMigrate(&domain.Sesion{}, &domain.Usuario{},
		&domain.Categoria{}, &domain.Archivo{}, &domain.Permiso{},
		&domain.Grupo{},
	)

	CreateIndexes()

	// Create Permisos
	var pp []domain.Permiso
	if db.Find(&pp).RowsAffected == 0 {
		for _, p := range permisos {
			pp = append(pp, domain.Permiso{Nombre: p})
		}
		db.Create(&pp)
	}

	// Create Grupos
	gg := []domain.Grupo{{Nombre: "Admin"}, {Nombre: "Funcionario"}, {Nombre: "Invitado"}}
	if db.Where("nombre IN (?)", []string{"Admin", "Funcionario", "Invitado"}).Find(&[]domain.Grupo{}).RowsAffected != 3 {
		// Assign default permisos to grupos
		for _, p := range pp {
			p := p
			if slices.Contains(config.ADMIN_DEFAULT_PERMISOS, p.Nombre) {
				gg[0].Permisos = append(gg[0].Permisos, &p)
			}
			if slices.Contains(config.FUNCIONARIO_DEFAULT_PERMISOS, p.Nombre) {
				gg[1].Permisos = append(gg[1].Permisos, &p)
			}
			if slices.Contains(config.INVITADO_DEFAULT_PERMISOS, p.Nombre) {
				gg[2].Permisos = append(gg[2].Permisos, &p)
			}
		}
		db.Create(&gg)
	}

	// Create Admin
	if err := db.First(&domain.Usuario{}, "nombre_usuario = ?", config.ADMIN).Error; err == gorm.ErrRecordNotFound {
		db.Create(&domain.Usuario{Status: true, Verificado: true, NombreUsuario: config.ADMIN, Email: config.ADMIN_EMAIL, Password: application.HashPassword(config.ADMIN_PASSWORD), GrupoID: gg[0].ID})
	}

	// Create Categoria Padre
	if err := db.First(&domain.Categoria{}, "nombre = ?", "Padre").Error; err == gorm.ErrRecordNotFound {
		db.Create(&domain.Categoria{Nombre: "Padre", Tipo: "Padre", PadreID: 1})
	}
}

func Close() {
	DB, _ := db.DB()
	DB.Close()
}
