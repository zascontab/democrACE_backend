package application

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"github.com/sonderkevin/governance/graph/model"
)

type GrupoService struct {
	R  repository.GrupoRepository
	PR repository.PermisoRepository
	L  *dataloader.Loader
}

func (s *GrupoService) Loader() *dataloader.Loader {
	if s.L == nil {
		s.L = dataloader.NewBatchedLoader(s.GetBatch)
	}
	return s.L
}

func (s *GrupoService) Get(id *string) (*model.Grupo, error) {
	return Get[model.Grupo](id, s.Loader().Load)
}

func (s *GrupoService) GetBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	return GetBatch(keys.Keys(), &domain.GrupoFilter{}, s.R.GetAll, s.ConvertOut)
}

func (s *GrupoService) GetAll(i *model.GrupoInput) ([]*model.Grupo, error) {
	return GetAllArray(i, &paging.PageArgs{}, s.ConvertFilter, s.R.GetAll, s.ConvertOut)
}

func (s *GrupoService) Save(ctx context.Context, i *model.SaveGrupoInput) (*model.Grupo, error) {
	return Save(ctx, i, s.ConvertIn, s.R.Save, s.R.Get, s.ConvertOut)
}

func (s *GrupoService) Update(i *model.UpdateGrupoInput) (*model.Grupo, error) {
	return Update(i, s.ConvertUpdate, s.R.Update, s.R.Get, s.ConvertOut)
}

func (s *GrupoService) Delete(id string) (string, error) {
	return Delete(s.R.Delete, id)
}

func (s *GrupoService) ConvertIn(ctx context.Context, i *model.SaveGrupoInput) (*domain.Grupo, error) {
	if i == nil {
		return nil, nil
	}

	o := domain.Grupo{
		Nombre:      i.Nombre,
		Descripcion: i.Descripcion,
	}

	if len(i.PermisosIDs) > 0 {
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

		o.Permisos = pp
	}

	return &o, nil
}

func (s *GrupoService) ConvertUpdate(i *model.UpdateGrupoInput, u *domain.Grupo) (*domain.Grupo, error) {
	if i == nil || u == nil {
		return nil, nil
	}

	id, err := DecryptAndConvert(i.ID)
	if err != nil {
		return nil, err
	}
	u.ID = id
	if i.Nombre != nil {
		u.Nombre = *i.Nombre
	}
	if i.Descripcion != nil {
		u.Descripcion = i.Descripcion
	}

	if len(i.PermisosIDs) > 0 {
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

func (s *GrupoService) ConvertFilter(i *model.GrupoInput, p *paging.PageArgs) (*domain.GrupoFilter, error) {
	f := domain.GrupoFilter{}
	if i != nil {
		if i.Ids != nil {
			uintids, err := DecryptAndConvertArray(ConvertPointersToStrings(i.Ids))
			if err != nil {
				return nil, err
			}
			f.IDs = &uintids
		}
		f.Nombre = i.Nombre
		f.Descripcion = i.Descripcion
		if i.PermisoID != nil {
			id, err := DecryptAndConvert(*i.PermisoID)
			if err != nil {
				return nil, err
			}
			f.PermisoID = &id
		}
		if i.UsuarioID != nil {
			id, err := DecryptAndConvert(*i.UsuarioID)
			if err != nil {
				return nil, err
			}
			f.UsuarioID = &id
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

func (s *GrupoService) ConvertOut(i *domain.Grupo) (*model.Grupo, error) {
	if i == nil {
		return nil, nil
	}

	o := model.Grupo{
		Nombre:      i.Nombre,
		Descripcion: i.Descripcion,
	}

	id, err := ConvertAndEncrypt(i.ID)
	if err != nil {
		return nil, err
	}
	o.ID = id

	return &o, nil
}
