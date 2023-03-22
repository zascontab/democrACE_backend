package db

import (
	"fmt"

	"gorm.io/gorm"
)

type Index struct {
	Name     string
	Table    string
	Columns  []string
	Unscoped bool
}

func CreateIndexes() {
	indexes := []Index{

		{
			Name:    "ux_categorias_nombre_tipo",
			Table:   "categoria",
			Columns: []string{"nombre", "tipo"},
		},
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
	CheckAndCreateIndexes(indexes)
}

func CheckAndCreateIndexes(idxs []Index) {
	for _, idx := range idxs {
		CheckAndCreateIndex(idx)
	}
}

func CheckAndCreateIndex(idx Index) {
	if SelectIndex(idx.Name).RowsAffected == 0 {
		if err := CreateIndex(idx).Error; err != nil {
			panic("failed to create index: " + idx.Name)
		}
	}
}

func SelectIndex(name string) *gorm.DB {
	return db.Exec("SELECT FROM pg_indexes WHERE indexname = ?", name)
}

func CreateIndex(idx Index) *gorm.DB {
	q := fmt.Sprintf("CREATE UNIQUE INDEX %s ON %s(%s)", idx.Name, idx.Table, ParseColumns(idx.Columns))
	if !idx.Unscoped {
		q += " WHERE deleted_at IS NULL"
	}
	return db.Exec(q)
}

func ParseColumns(cc []string) string {
	out := ""
	for i, c := range cc {
		if i > 0 {
			out += ", "
		}
		out += c
	}
	return out
}
