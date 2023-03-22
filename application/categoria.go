package application

import (
	"context"

	"github.com/graph-gophers/dataloader"
	"github.com/nrfta/go-paging"
	"github.com/sonderkevin/governance/domain"
	"github.com/sonderkevin/governance/domain/repository"
	"github.com/sonderkevin/governance/graph/model"
)

type CategoriaService struct {
	R repository.CategoriaRepository
	L *dataloader.Loader
}

func (s *CategoriaService) Loader() *dataloader.Loader {
	if s.L == nil {
		s.L = dataloader.NewBatchedLoader(s.GetBatch)
	}
	return s.L
}

func (s *CategoriaService) Get(id *string) (*model.Categoria, error) {
	return Get[model.Categoria](id, s.Loader().Load)
}

func (s *CategoriaService) GetBatch(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	return GetBatch(keys.Keys(), &domain.CategoriaFilter{}, s.R.GetAll, s.ConvertOut)
}

func (s *CategoriaService) GetAll(i *model.CategoriaInput, page *paging.PageArgs) (*model.CategoriaNodeConnection, error) {
	return GetAll(i, page, s.ConvertFilter, s.R.GetAll, s.ConvertOut, s.BuildNodeConnection)
}

func (s *CategoriaService) Save(ctx context.Context, i *model.SaveCategoriaInput) (*model.Categoria, error) {
	return Save(ctx, i, s.ConvertIn, s.R.Save, s.R.Get, s.ConvertOut)
}

func (s *CategoriaService) Update(i *model.UpdateCategoriaInput) (*model.Categoria, error) {
	return Update(i, s.ConvertUpdate, s.R.Update, s.R.Get, s.ConvertOut)
}

func (s *CategoriaService) Delete(id string) (string, error) {
	return Delete(s.R.Delete, id)
}

func (s *CategoriaService) ConvertIn(ctx context.Context, i *model.SaveCategoriaInput) (*domain.Categoria, error) {
	if i == nil {
		return nil, nil
	}

	o := domain.Categoria{
		Nombre: i.Nombre,
		Tipo:   i.Tipo,
	}

	if i.PadreID != nil {
		id, err := DecryptAndConvert(*i.PadreID)
		if err != nil {
			return nil, err
		}
		o.PadreID = id
	} else {
		// TODO: This is Categoria Padre
		o.PadreID = 1
	}

	return &o, nil
}

func (s *CategoriaService) ConvertUpdate(i *model.UpdateCategoriaInput, u *domain.Categoria) (*domain.Categoria, error) {
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
	if i.Tipo != nil {
		u.Tipo = *i.Tipo
	}
	if i.PadreID != nil {
		id, err := DecryptAndConvert(*i.PadreID)
		if err != nil {
			return nil, err
		}
		u.PadreID = id
	}

	return u, nil
}

func (s *CategoriaService) ConvertFilter(i *model.CategoriaInput, p *paging.PageArgs) (*domain.CategoriaFilter, error) {
	f := domain.CategoriaFilter{}
	if i != nil {
		if i.Ids != nil {
			uintids, err := DecryptAndConvertArray(ConvertPointersToStrings(i.Ids))
			if err != nil {
				return nil, err
			}
			f.IDs = &uintids
		}
		f.Nombre = i.Nombre
		f.Tipo = i.Tipo
		if i.PadreID != nil {
			id, err := DecryptAndConvert(*i.PadreID)
			if err != nil {
				return nil, err
			}
			f.PadreID = &id
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

func (s *CategoriaService) ConvertOut(i *domain.Categoria) (*model.Categoria, error) {
	if i == nil {
		return nil, nil
	}

	o := model.Categoria{
		Nombre: i.Nombre,
		Tipo:   i.Tipo,
	}

	id, err := ConvertAndEncrypt(i.ID)
	if err != nil {
		return nil, err
	}
	o.ID = id

	id, err = ConvertAndEncrypt(i.PadreID)
	if err != nil {
		return nil, err
	}
	o.PadreID = id

	return &o, nil
}

func (s *CategoriaService) BuildNodeConnection(cc []*model.Categoria, p *paging.PageArgs) (*model.CategoriaNodeConnection, error) {
	totalCount := len(cc)

	paginator := paging.NewOffsetPaginator(p, int64(totalCount))

	result := &model.CategoriaNodeConnection{
		PageInfo: &paginator.PageInfo,
	}

	for i, c := range cc {
		result.Edges = append(result.Edges, &model.CategoriaNodeEdge{
			Cursor: *paging.EncodeOffsetCursor(paginator.Offset + i),
			Node:   c,
		})
	}

	return result, nil
}
