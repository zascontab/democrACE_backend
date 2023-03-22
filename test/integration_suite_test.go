package test

import (
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/sonderkevin/governance/application"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/infrastructure/persistence"
	"github.com/sonderkevin/governance/server"
	"github.com/stretchr/testify/suite"
	"golang.org/x/exp/slices"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Config struct {
	PORT                         string
	DB_HOST                      string
	DB_PORT                      string
	DB_USER                      string
	DB_PASSWORD                  string
	DB_NAME                      string
	SECRET                       string
	ADMIN                        string
	ADMIN_EMAIL                  string
	ADMIN_PASSWORD               string
	TEMP_TOKEN_EXP               string
	REFRESH_TOKEN_EXP            string
	EMAIL_SENDER                 string
	EMAIL_SENDER_PASSWORD        string
	EMAIL_HOST                   string
	EMAIL_PORT                   string
	ADMIN_DEFAULT_PERMISOS       []string
	FUNCIONARIO_DEFAULT_PERMISOS []string
	INVITADO_DEFAULT_PERMISOS    []string
}

type Index struct {
	Name    string
	Table   string
	Columns []string
}

type IntTestSuite struct {
	suite.Suite
	config *Config
	db     *gorm.DB
	// actividadService *application.ActividadService
	authService     *application.AuthService
	programaService *application.ProgramaService
	usuarioService  *application.UsuarioService
	usuarios        []*domain.Usuario
}

func TestIntTestSuite(t *testing.T) {
	suite.Run(t, &IntTestSuite{})
}

func (s *IntTestSuite) SetupSuite() {
	loadConfig(s)
	connectDB(s)
	setupDatabase(s)
	initServices(s)
}

func (s *IntTestSuite) TearDownSuite() {
	tearDownDatabase(s)
}

func (s *IntTestSuite) AfterTest(suiteName, testName string) {
	clearDatabase(s)
}

func loadConfig(s *IntTestSuite) {
	s.config = &Config{}
	s.config.PORT = getEnv(s, "PORT")
	s.config.DB_HOST = getEnv(s, "DB_HOST")
	s.config.DB_PORT = getEnv(s, "DB_PORT")
	s.config.DB_USER = getEnv(s, "DB_USER")
	s.config.DB_PASSWORD = getEnv(s, "DB_PASSWORD")
	s.config.DB_NAME = getEnv(s, "DB_NAME")
	s.config.SECRET = getEnv(s, "SECRET")
	s.config.ADMIN = getEnv(s, "ADMIN")
	s.config.ADMIN_EMAIL = getEnv(s, "ADMIN_EMAIL")
	s.config.ADMIN_PASSWORD = getEnv(s, "ADMIN_PASSWORD")
	s.config.TEMP_TOKEN_EXP = getEnv(s, "TEMP_TOKEN_EXP")
	s.config.REFRESH_TOKEN_EXP = getEnv(s, "REFRESH_TOKEN_EXP")
	s.config.EMAIL_SENDER = getEnv(s, "EMAIL_SENDER")
	s.config.EMAIL_SENDER_PASSWORD = getEnv(s, "EMAIL_SENDER_PASSWORD")
	s.config.EMAIL_HOST = getEnv(s, "EMAIL_HOST")
	s.config.EMAIL_PORT = getEnv(s, "EMAIL_PORT")
	s.config.ADMIN_DEFAULT_PERMISOS = []string{"all"}
	s.config.FUNCIONARIO_DEFAULT_PERMISOS = []string{"me"}
	s.config.INVITADO_DEFAULT_PERMISOS = []string{"me"}
}

func getEnv(s *IntTestSuite, key string) string {
	value, ok := os.LookupEnv(key)
	if !ok {
		s.FailNow(fmt.Sprintf("%s environment variable not set", key))
	}
	return value
}

func connectDB(s *IntTestSuite) {
	s.T().Log("connecting to the database")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", s.config.DB_HOST, s.config.DB_PORT, s.config.DB_USER, s.config.DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(dsn))
	s.Require().Nil(err)
	s.db = db

	err = s.db.Exec(fmt.Sprintf("CREATE DATABASE %s", s.config.DB_NAME)).Error
	s.Require().Nil(err)

	dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", s.config.DB_HOST, s.config.DB_PORT, s.config.DB_USER, s.config.DB_PASSWORD, s.config.DB_NAME)
	db, err = gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{
			Logger:                                   logger.Default.LogMode(logger.Silent),
			DisableForeignKeyConstraintWhenMigrating: true,
		},
	)
	s.Require().Nil(err)

	s.T().Log("connected to the database!")
	s.db = db
}

func setupDatabase(s *IntTestSuite) {
	s.T().Log("setting up database")

	err := s.db.AutoMigrate(&domain.Permiso{}, &domain.Grupo{}, &domain.Usuario{})
	s.Nil(err)
	s.db.DisableForeignKeyConstraintWhenMigrating = false
	err = s.db.AutoMigrate(&domain.Permiso{}, &domain.Grupo{}, &domain.Usuario{})
	s.Nil(err)

	createIndexes(s)
	seedDb(s)
}

func initServices(s *IntTestSuite) {
	// actividadRepository := persistence.NewActividadRepository(s.db)
	archivoRepository := persistence.NewArchivoRepository(s.db)
	programaRepository := persistence.NewProgramaRepository(s.db)
	proyectoRepository := persistence.NewProyectoRepository(s.db)
	permisoRepository := persistence.NewPermisoRepository(s.db)
	usuarioRepository := persistence.NewUsuarioRepository(s.db)
	s.authService = &application.AuthService{UR: usuarioRepository}
	s.programaService = &application.ProgramaService{R: programaRepository, AR: archivoRepository, PR: proyectoRepository, UR: usuarioRepository}
	s.usuarioService = &application.UsuarioService{R: usuarioRepository, PR: permisoRepository}
}

func seedDb(s *IntTestSuite) {
	es := server.GenerateExecutableSchema(s.db)
	permisos := makePermisos(es.Schema())

	// Create Permisos
	var pp []domain.Permiso
	if s.db.Find(&pp).RowsAffected == 0 {
		for _, p := range permisos {
			pp = append(pp, domain.Permiso{Nombre: p})
		}
		s.db.Create(&pp)
	}

	// Create Grupos
	gg := []domain.Grupo{{Nombre: "Admin"}, {Nombre: "Funcionario"}, {Nombre: "Invitado"}}
	if s.db.Where("nombre IN (?)", []string{"Admin", "Funcionario", "Invitado"}).Find(&[]domain.Grupo{}).RowsAffected != 3 {
		// Assign default permisos to grupos
		for _, p := range pp {
			p := p
			if slices.Contains(s.config.ADMIN_DEFAULT_PERMISOS, p.Nombre) {
				gg[0].Permisos = append(gg[0].Permisos, &p)
			}
			if slices.Contains(s.config.FUNCIONARIO_DEFAULT_PERMISOS, p.Nombre) {
				gg[1].Permisos = append(gg[1].Permisos, &p)
			}
			if slices.Contains(s.config.INVITADO_DEFAULT_PERMISOS, p.Nombre) {
				gg[2].Permisos = append(gg[2].Permisos, &p)
			}
		}
		s.db.Create(&gg)
	}

	uu := []*domain.Usuario{
		{
			Status:        true,
			Verificado:    true,
			NombreUsuario: s.config.ADMIN,
			Email:         s.config.ADMIN_EMAIL,
			Password:      application.HashPassword(s.config.ADMIN_PASSWORD),
			GrupoID:       gg[0].ID,
		},
		{
			Status:             true,
			Verificado:         true,
			CodigoVerificacion: "",
			Nombres:            "nombres1",
			Apellidos:          "apellidos1",
			NombreUsuario:      "nombreusuario1",
			Email:              "email1",
			Password:           application.HashPassword("password1"),
			GrupoID:            gg[2].ID,
		},
		{
			Status:             true,
			Verificado:         false,
			CodigoVerificacion: "",
			Nombres:            "nombres2",
			Apellidos:          "apellidos2",
			NombreUsuario:      "nombreusuario2",
			Email:              "email2",
			Password:           application.HashPassword("password2"),
			GrupoID:            gg[2].ID,
		},
	}

	s.db.Create(&uu)
	s.usuarios = uu

	CreateProgramas(s)
	CreateProyectos(s)
	CreateActividades(s)
	CreateEjecuciones(s)
}

func CreateProgramas(s *IntTestSuite) {
	pp := []*domain.Programa{
		{
			Nombre:            "programa1",
			Anio:              2023,
			Monto:             1000.00,
			MetaPrograma:      "meta1",
			IndicadorPrograma: "indicador1",
			FechaInicio:       time.Now(),
			FechaFin:          time.Now().AddDate(0, 5, 0),
		},
		{
			Nombre:            "programa2",
			Anio:              2023,
			Monto:             2000.00,
			MetaPrograma:      "meta2",
			IndicadorPrograma: "indicador2",
			FechaInicio:       time.Now().AddDate(0, 1, 2),
			FechaFin:          time.Now().AddDate(0, 3, 0),
		},
		{
			Nombre:            "programa3",
			Anio:              2023,
			Monto:             3000.00,
			MetaPrograma:      "meta3",
			IndicadorPrograma: "indicador3",
			FechaInicio:       time.Now().AddDate(0, 0, 1),
			FechaFin:          time.Now().AddDate(0, 6, 1),
		},
	}

	s.db.Create(&pp)
}

func CreateProyectos(s *IntTestSuite) {
	pp := []*domain.Proyecto{
		{
			Nombre:       "Proyecto1",
			Descripcion:  strPtr("Descripción del proyecto1."),
			Meta:         "Meta del proyecto1",
			Indicador:    "Indicador del proyecto1",
			Estado:       0,
			Posicion:     1,
			Presupuesto:  100,
			Cancelado:    false,
			FechaInicio:  time.Now(),
			FechaFin:     time.Now().AddDate(0, 3, 0),
			ProgramaID:   1,
			Responsables: s.usuarios[:1],
		},
		{
			Nombre:       "Proyecto2",
			Descripcion:  strPtr("Descripción del proyecto2."),
			Meta:         "Meta del proyecto2",
			Indicador:    "Indicador del proyecto2",
			Estado:       1,
			Posicion:     2,
			Presupuesto:  500,
			Cancelado:    false,
			FechaInicio:  time.Now(),
			FechaFin:     time.Now().AddDate(0, 6, 0),
			ProgramaID:   1,
			Responsables: s.usuarios[1:],
		},
		{
			Nombre:         "Proyecto C",
			Descripcion:    strPtr("Descripción del proyecto C."),
			Meta:           "Meta del proyecto C",
			Indicador:      "Indicador del proyecto C",
			Estado:         0,
			Posicion:       3,
			Presupuesto:    3000,
			Cancelado:      true,
			RazonCancelado: strPtr("Razón del proyecto C cancelado."),
			FechaInicio:    time.Now(),
			FechaFin:       time.Now().AddDate(1, 0, 0),
			ProgramaID:     3,
			Responsables:   s.usuarios[:1],
		},
	}
	s.db.Create(&pp)
}

func CreateActividades(s *IntTestSuite) {
	aa := []*domain.Actividad{
		{
			Nombre:                 "Actividad1",
			Descripcion:            strPtr("Descripción de la actividad1."),
			Estado:                 0,
			PresupuestoPlanificado: 50,
			FechasPlanificadas:     []*domain.Date{{Date: time.Now().AddDate(0, 0, 1)}, {Date: time.Now().AddDate(0, 0, 2)}, {Date: time.Now().AddDate(0, 0, 3)}},
			Cumplida:               false,
			ProyectoID:             1,
			Responsables:           s.usuarios[:2],
		},
		{
			Nombre:                 "Actividad2",
			Descripcion:            strPtr("Descripción de la actividad2."),
			Estado:                 1,
			PresupuestoPlanificado: 200,
			FechasPlanificadas:     []*domain.Date{{Date: time.Now().AddDate(0, 0, 4)}, {Date: time.Now().AddDate(0, 0, 5)}, {Date: time.Now().AddDate(0, 0, 6)}},
			Cumplida:               false,
			ProyectoID:             2,
			Responsables:           s.usuarios[2:3],
		},
		{
			Nombre:                 "Actividad C",
			Descripcion:            strPtr("Descripción de la actividad C."),
			Estado:                 0,
			PresupuestoPlanificado: 1000,
			FechasPlanificadas:     []*domain.Date{{Date: time.Now().AddDate(0, 0, 7)}, {Date: time.Now().AddDate(0, 0, 8)}, {Date: time.Now().AddDate(0, 0, 9)}},
			Cumplida:               false,
			ProyectoID:             2,
			Responsables:           s.usuarios[1:3],
		},
	}
	s.db.Create(&aa)
}

func CreateEjecuciones(s *IntTestSuite) {
	ee := []*domain.Ejecucion{
		{
			Fecha:        time.Now().AddDate(0, 0, -1),
			Presupuesto:  20,
			ActividadID:  1,
			Responsables: s.usuarios[:2],
		},
		{
			Fecha:        time.Now().AddDate(0, 0, -2),
			Presupuesto:  10,
			ActividadID:  1,
			Responsables: s.usuarios[2:3],
		},
		{
			Fecha:        time.Now().AddDate(0, 0, -3),
			Presupuesto:  20,
			ActividadID:  2,
			Responsables: s.usuarios[1:3],
		},
	}

	s.db.Create(&ee)
}

func clearDatabase(s *IntTestSuite) {
	s.T().Log("cleaning database")
	s.db.AllowGlobalUpdate = true
	s.db.Where("id NOT IN (?)", []uint{1, 2, 3}).Delete(&domain.Usuario{})

	s.db.Where("id = 1").Updates(&domain.Usuario{Password: application.HashPassword(s.config.ADMIN_PASSWORD)})
	s.db.Where("id = 2").Updates(&domain.Usuario{Password: application.HashPassword("password1")})
	s.db.Where("id = 3").Updates(&domain.Usuario{Password: application.HashPassword("password2")})
}

func tearDownDatabase(s *IntTestSuite) {
	s.T().Log("tearing down database")
	d, _ := s.db.DB()
	d.Close()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", s.config.DB_HOST, s.config.DB_PORT, s.config.DB_USER, s.config.DB_PASSWORD)
	db, err := gorm.Open(postgres.Open(dsn))
	s.Require().Nil(err)
	s.db = db
	err = s.db.Exec(fmt.Sprintf("DROP DATABASE %s", s.config.DB_NAME)).Error

	if err != nil {
		s.FailNowf("unable to drop database", err.Error())
	}
}

func createIndexes(s *IntTestSuite) {
	indexes := []Index{
		// {
		// 	Name:    "ux_actividads_nombre_proyecto_id",
		// 	Table:   "actividads",
		// 	Columns: []string{"nombre", "proyecto_id"},
		// },
		// {
		// 	Name:    "ux_categorias_nombre_tipo",
		// 	Table:   "categoria",
		// 	Columns: []string{"nombre", "tipo"},
		// },
		{
			Name:    "ux_grupos_nombre",
			Table:   "grupos",
			Columns: []string{"nombre"},
		},
		{
			Name:    "ux_permisos_nombre",
			Table:   "permisos",
			Columns: []string{"nombre"},
		},
		// {
		// 	Name:    "ux_programas_nombre",
		// 	Table:   "programas",
		// 	Columns: []string{"nombre"},
		// },
		// {
		// 	Name:    "ux_proyectos_nombre_programa_id",
		// 	Table:   "proyectos",
		// 	Columns: []string{"nombre", "programa_id"},
		// },
		{
			Name:    "ux_usuarios_nombre_usuario",
			Table:   "usuarios",
			Columns: []string{"nombre_usuario"},
		},
		{
			Name:    "ux_usuarios_email",
			Table:   "usuarios",
			Columns: []string{"email"},
		},
	}
	checkAndCreateIndexes(s, indexes)
}

func checkAndCreateIndexes(s *IntTestSuite, idxs []Index) {
	for _, idx := range idxs {
		checkAndCreateIndex(s, idx.Name, idx.Table, idx.Columns...)
	}
}

func checkAndCreateIndex(s *IntTestSuite, name, table string, columns ...string) {
	if selectIndex(s, name).RowsAffected == 0 {
		if err := createIndex(s, name, table, columns...).Error; err != nil {
			panic("failed to create index: " + name)
		}
	}
}

func selectIndex(s *IntTestSuite, name string) *gorm.DB {
	return s.db.Exec("SELECT FROM pg_indexes WHERE indexname = ?", name)
}

func createIndex(s *IntTestSuite, name, table string, columns ...string) *gorm.DB {
	return s.db.Exec(fmt.Sprintf("CREATE UNIQUE INDEX %s ON %s(%s) WHERE deleted_at IS NULL", name, table, parseColumns(columns)))
}
