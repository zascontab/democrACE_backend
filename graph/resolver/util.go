package resolver

import (
	"context"
	"errors"

	"github.com/sonderkevin/governance/application"
	"github.com/sonderkevin/governance/db"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/graph/model"
	"github.com/sonderkevin/governance/infrastructure/persistence"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"golang.org/x/exp/slices"
)

func IsAuthorized(ctx context.Context, permiso string) bool {
	// Everyone is authorized
	if slices.Contains([]string{"login", "registrar", "getTempToken", "enviarCodigo"}, permiso) {
		return true
	}

	u, err := GetUsuarioFromCtx(ctx)
	if err != nil {
		return false
	}

	if len(u.Permisos) > 0 {
		permisos := []string{}
		for _, p := range u.Permisos {
			permisos = append(permisos, p.Nombre)
		}
		return slices.ContainsFunc(permisos, func(p string) bool { return p == "all" || p == permiso })
	}

	if u.GrupoID != 0 {
		repo := persistence.NewPermisoRepository(db.New())

		ids := []uint{u.GrupoID}
		pp, err := repo.GetAll(&domain.PermisoFilter{GruposIDs: &ids})
		if err != nil {
			return false
		}

		permisos := []string{}
		for _, p := range pp {
			permisos = append(permisos, p.Nombre)
		}

		if len(permisos) > 0 {
			return slices.ContainsFunc(permisos, func(p string) bool { return p == "all" || p == permiso })
		}
	}

	return false
}

func GetUsuarioFromCtx(ctx context.Context) (*domain.Usuario, error) {
	u, ok := ctx.Value("usuario").(*domain.Usuario)
	if !ok {
		return nil, errors.New("access denied")
	}
	return u, nil
}

func IsAdmin(ctx context.Context) bool {
	u, err := GetUsuarioFromCtx(ctx)
	if err != nil {
		return false
	}

	// TODO: Grupo Admin id
	return u.GrupoID == 1
}

func IsResponsable(s *application.UsuarioService, ctx context.Context, i *model.UsuarioInput) bool {
	if IsAdmin(ctx) {
		return true
	}

	rr, err := s.GetAll(i)
	if err != nil {
		return false
	}

	if !ContainsID(ctx, rr) {
		return false
	}

	return true
}

func SameID(ctx context.Context, id string) bool {
	u, err := GetUsuarioFromCtx(ctx)
	if err != nil {
		return false
	}

	uintid, err := application.DecryptAndConvert(id)
	if err != nil {
		return false
	}

	return u.ID == uintid
}

func ContainsID(ctx context.Context, responsables []*model.Usuario) bool {
	u, err := GetUsuarioFromCtx(ctx)
	if err != nil {
		return false
	}

	var ids []string
	for _, usuario := range responsables {
		ids = append(ids, usuario.ID)
	}

	uintids, err := application.DecryptAndConvertArray(ids)
	if err != nil {
		return false
	}

	return slices.Contains(uintids, u.ID)
}

func Error(err string) error {
	return &gqlerror.Error{
		Message: err,
	}
}
