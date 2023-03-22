package application

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"github.com/sonderkevin/governance/graph/model"
)

type ArchivoService struct {
	R repository.ArchivoRepository
	L *dataloader.Loader
}

func (s *ArchivoService) Loader() *dataloader.Loader {
	if s.L == nil {
		s.L = dataloader.NewBatchedLoader(s.GetBatch)
	}
	return s.L
}

func (s *ArchivoService) Get(id *string) (*model.Archivo, error) {
	return Get[model.Archivo](id, s.Loader().Load)
}

func (s *ArchivoService) GetBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	return GetBatch(keys.Keys(), &domain.ArchivoFilter{}, s.R.GetAll, s.ConvertOut)
}

func (s *ArchivoService) GetAll(i *model.ArchivoInput) ([]*model.Archivo, error) {
	return GetAllArray(i, &paging.PageArgs{}, s.ConvertFilter, s.R.GetAll, s.ConvertOut)
}

func (s *ArchivoService) Save(ctx context.Context, i *model.Archivo) (*model.Archivo, error) {
	return Save(ctx, i, s.ConvertIn, s.R.Save, s.R.Get, s.ConvertOut)
}

func (s *ArchivoService) Delete(id string) (string, error) {
	return Delete(s.R.Delete, id)
}

func (s *ArchivoService) ConvertIn(ctx context.Context, i *model.Archivo) (*domain.Archivo, error) {
	if i == nil {
		return nil, nil
	}

	o := domain.Archivo{
		Nombre: i.Nombre,
	}

	return &o, nil
}

func (s *ArchivoService) ConvertFilter(i *model.ArchivoInput, p *paging.PageArgs) (*domain.ArchivoFilter, error) {
	f := domain.ArchivoFilter{}
	if i != nil {
		if i.Ids != nil {
			uintids, err := DecryptAndConvertArray(ConvertPointersToStrings(i.Ids))
			if err != nil {
				return nil, err
			}
			f.IDs = &uintids
		}
		f.Nombre = i.Nombre
		if i.ProgramaID != nil {
			id, err := DecryptAndConvert(*i.ProgramaID)
			if err != nil {
				return nil, err
			}
			f.ProgramaID = &id
		}
		if i.ProyectoID != nil {
			id, err := DecryptAndConvert(*i.ProyectoID)
			if err != nil {
				return nil, err
			}
			f.ProyectoID = &id
		}
		if i.ActividadID != nil {
			id, err := DecryptAndConvert(*i.ActividadID)
			if err != nil {
				return nil, err
			}
			f.ActividadID = &id
		}
		if i.EjecucionID != nil {
			id, err := DecryptAndConvert(*i.EjecucionID)
			if err != nil {
				return nil, err
			}
			f.EjecucionID = &id
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

func (s *ArchivoService) ConvertOut(i *domain.Archivo) (*model.Archivo, error) {
	if i == nil {
		return nil, nil
	}

	o := model.Archivo{
		Nombre: i.Nombre,
	}

	id, err := ConvertAndEncrypt(i.ID)
	if err != nil {
		return nil, err
	}
	o.ID = id

	return &o, nil
}
