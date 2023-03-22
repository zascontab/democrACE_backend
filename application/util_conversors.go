package application

import (
	"time"

	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/graph/model"
)

func ConvertUsuario(usuario *domain.Usuario) (*model.Usuario, error) {
	if usuario == nil {
		return nil, nil
	}

	id, err := ConvertAndEncrypt(usuario.ID)
	if err != nil {

		return nil, err
	}

	return &model.Usuario{
		ID:            id,
		Status:        usuario.Status,
		Nombres:       usuario.Nombres,
		Apellidos:     usuario.Apellidos,
		NombreUsuario: usuario.NombreUsuario,
		Email:         usuario.Email,
	}, nil
}

func ConvertPointersToStrings(pointers []*string) []string {
	if len(pointers) == 0 {
		return nil
	}

	strings := make([]string, 0, len(pointers))
	for _, p := range pointers {
		if p != nil {
			strings = append(strings, *p)
		}
	}

	return strings
}

func FormatFecha(f *time.Time) string {
	return f.Format("2006-01-02")
}

/* func FormatFechas(fechas []*domain.Date) []*string {
	var ff []*string
	for _, f := range fechas {
		formatted := f.Date.Format("2006-01-02")
		ff = append(ff, &formatted)
	}
	return ff
} */

func ParseFecha(f string) (time.Time, error) {
	return time.Parse("2006-01-02", f)
}

func ParseFechas(fechas []*string) ([]time.Time, error) {
	fpString := ConvertPointersToStrings(fechas)
	var fp []time.Time
	for _, f := range fpString {
		parsed, err := ParseFecha(f)
		if err != nil {

			return nil, err
		}
		fp = append(fp, parsed)
	}
	return fp, nil
}
