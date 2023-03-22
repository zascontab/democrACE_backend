package application

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"github.com/sonderkevin/governance/graph/model"
)

type PermisoService struct {
	R repository.PermisoRepository
	L *dataloader.Loader
}

func (s *PermisoService) Loader() *dataloader.Loader {
	if s.L == nil {
		s.L = dataloader.NewBatchedLoader(s.GetBatch)
	}
	return s.L
}

func (s *PermisoService) Get(id *string) (*model.Permiso, error) {
	return Get[model.Permiso](id, s.Loader().Load)
}

func (s *PermisoService) GetBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	return GetBatch(keys.Keys(), &domain.PermisoFilter{}, s.R.GetAll, s.ConvertOut)
}

func (s *PermisoService) GetAll(i *model.PermisoInput) ([]*model.Permiso, error) {
	return GetAllArray(i, &paging.PageArgs{}, s.ConvertFilter, s.R.GetAll, s.ConvertOut)
}

func (s *PermisoService) ConvertFilter(i *model.PermisoInput, p *paging.PageArgs) (*domain.PermisoFilter, error) {
	f := domain.PermisoFilter{}
	if i != nil {
		if i.Ids != nil {
			uintids, err := DecryptAndConvertArray(ConvertPointersToStrings(i.Ids))
			if err != nil {
				return nil, err
			}
			f.IDs = &uintids
		}
		f.Nombre = i.Nombre
		if i.GrupoID != nil {
			id, err := DecryptAndConvert(*i.GrupoID)
			if err != nil {
				return nil, err
			}
			f.GruposIDs = &[]uint{id}
		}
		if i.UsuarioID != nil {
			id, err := DecryptAndConvert(*i.UsuarioID)
			if err != nil {
				return nil, err
			}
			f.UsuariosIDs = &[]uint{id}
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

func (s *PermisoService) ConvertOut(i *domain.Permiso) (*model.Permiso, error) {
	if i == nil {
		return nil, nil
	}

	o := model.Permiso{
		Nombre: i.Nombre,
	}

	id, err := ConvertAndEncrypt(i.ID)
	if err != nil {
		return nil, err
	}
	o.ID = id

	return &o, nil
}
