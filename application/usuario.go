package application

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"github.com/sonderkevin/governance/graph/model"
)

type UsuarioService struct {
	R  repository.UsuarioRepository
	PR repository.PermisoRepository
	L  *dataloader.Loader
}

func (s *UsuarioService) Loader() *dataloader.Loader {
	if s.L == nil {
		s.L = dataloader.NewBatchedLoader(s.GetBatch)
	}
	return s.L
}

func (s *UsuarioService) Get(id *string) (*model.Usuario, error) {
	return Get[model.Usuario](id, s.Loader().Load)
}

func (s *UsuarioService) GetBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	return GetBatch(keys.Keys(), &domain.UsuarioFilter{}, s.R.GetAll, s.ConvertOut)
}

func (s *UsuarioService) GetAll(i *model.UsuarioInput) ([]*model.Usuario, error) {
	return GetAllArray(i, &paging.PageArgs{}, s.ConvertFilter, s.R.GetAll, s.ConvertOut)
}

func (s *UsuarioService) Save(ctx context.Context, i *model.RegistrarUsuarioInput) (*model.Usuario, error) {
	return Save(ctx, i, s.ConvertIn, s.R.Save, s.R.Get, s.ConvertOut)
}

func (s *UsuarioService) Update(i *model.UpdateUsuarioInput) (*model.Usuario, error) {
	return Update(i, s.ConvertUpdate, s.R.Update, s.R.Get, s.ConvertOut)
}

func (s *UsuarioService) Delete(id string) (string, error) {
	return Delete(s.R.Delete, id)
}

func (s *UsuarioService) ConvertIn(ctx context.Context, i *model.RegistrarUsuarioInput) (*domain.Usuario, error) {
	if i == nil {
		return nil, nil
	}

	o := domain.Usuario{
		Status:        true,
		Nombres:       i.Nombres,
		Apellidos:     i.Apellidos,
		NombreUsuario: i.NombreUsuario,
		Email:         i.Email,
		Password:      HashPassword(i.Password),
		// TODO: Grupo Invitado ID
		GrupoID: 3,
	}

	return &o, nil
}

func (s *UsuarioService) ConvertUpdate(i *model.UpdateUsuarioInput, u *domain.Usuario) (*domain.Usuario, error) {
	if i == nil || u == nil {
		return nil, nil
	}

	id, err := DecryptAndConvert(i.ID)
	if err != nil {
		return nil, err
	}
	u.ID = id
	if i.Status != nil {
		u.Status = *i.Status
	}
	if i.Nombres != nil {
		u.Nombres = *i.Nombres
	}
	if i.Apellidos != nil {
		u.Apellidos = *i.Apellidos
	}
	if i.NombreUsuario != nil {
		u.NombreUsuario = *i.NombreUsuario
	}
	if i.Email != nil {
		u.Email = *i.Email
	}
	if i.GrupoID != nil {
		id, err := DecryptAndConvert(*i.GrupoID)
		if err != nil {
			return nil, err
		}
		u.GrupoID = id
	}
	if i.PermisosIDs != nil {
		ids, err := DecryptAndConvertArray(ConvertPointersToStrings(i.PermisosIDs))
		if err != nil {
			return nil, err
		}

		f := domain.PermisoFilter{}
		f.IDs = &ids
		pp, err := s.PR.GetAll(&f)
		if err != nil {
			return nil, err
		}
		u.Permisos = pp
	}

	return u, nil
}

func (s *UsuarioService) ConvertFilter(i *model.UsuarioInput, p *paging.PageArgs) (*domain.UsuarioFilter, error) {
	f := domain.UsuarioFilter{}
	if i != nil {
		if i.Ids != nil {
			ids, err := DecryptAndConvertArray(ConvertPointersToStrings(i.Ids))
			if err != nil {
				return nil, err
			}
			f.IDs = &ids
		}
		f.Status = i.Status
		f.Verificado = i.Verificado
		f.Nombres = i.Nombres
		f.Apellidos = i.Apellidos
		f.NombreUsuario = i.NombreUsuario
		f.Email = i.Email
		if i.GrupoID != nil {
			id, err := DecryptAndConvert(*i.GrupoID)
			if err != nil {
				return nil, err
			}
			f.GrupoID = &id
		}
		if len(i.PermisosIDs) > 0 {
			ids, err := DecryptAndConvertArray(ConvertPointersToStrings(i.PermisosIDs))
			if err != nil {
				return nil, err
			}
			f.PermisosIDs = &ids
		}
		
	}

	if p != nil {
		if p.First != nil {
			f.Limit = p.First
		}
		if p.After != nil {
			decoded := paging.DecodeOffsetCursor(p.After)
			f.Offset = &decoded
		}
	}

	return &f, nil
}

func (s *UsuarioService) ConvertOut(i *domain.Usuario) (*model.Usuario, error) {
	if i == nil {
		return nil, nil
	}

	o := model.Usuario{
		Status:        i.Status,
		Verificado:    i.Verificado,
		Nombres:       i.Nombres,
		Apellidos:     i.Apellidos,
		NombreUsuario: i.NombreUsuario,
		Email:         i.Email,
	}

	id, err := ConvertAndEncrypt(i.ID)
	if err != nil {
		return nil, err
	}
	o.ID = id

	id, err = ConvertAndEncrypt(i.GrupoID)
	if err != nil {
		return nil, err
	}
	o.GrupoID = id

	return &o, nil
}
