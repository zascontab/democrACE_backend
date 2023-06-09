// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/sonderkevin/governance/graph/model"

	paging "github.com/nrfta/go-paging"
)

// QueryResolver is an autogenerated mock type for the QueryResolver type
type QueryResolver struct {
	mock.Mock
}

// Categoria provides a mock function with given fields: ctx, id
func (_m *QueryResolver) Categoria(ctx context.Context, id string) (*model.Categoria, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Categoria
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Categoria, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Categoria); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Categoria)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Categorias provides a mock function with given fields: ctx, input, page
func (_m *QueryResolver) Categorias(ctx context.Context, input *model.CategoriaInput, page *paging.PageArgs) (*model.CategoriaNodeConnection, error) {
	ret := _m.Called(ctx, input, page)

	var r0 *model.CategoriaNodeConnection
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.CategoriaInput, *paging.PageArgs) (*model.CategoriaNodeConnection, error)); ok {
		return rf(ctx, input, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.CategoriaInput, *paging.PageArgs) *model.CategoriaNodeConnection); ok {
		r0 = rf(ctx, input, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.CategoriaNodeConnection)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.CategoriaInput, *paging.PageArgs) error); ok {
		r1 = rf(ctx, input, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Grupo provides a mock function with given fields: ctx, id
func (_m *QueryResolver) Grupo(ctx context.Context, id string) (*model.Grupo, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Grupo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Grupo, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Grupo); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Grupo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Grupos provides a mock function with given fields: ctx, input, page
func (_m *QueryResolver) Grupos(ctx context.Context, input *model.GrupoInput, page *paging.PageArgs) ([]*model.Grupo, error) {
	ret := _m.Called(ctx, input, page)

	var r0 []*model.Grupo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.GrupoInput, *paging.PageArgs) ([]*model.Grupo, error)); ok {
		return rf(ctx, input, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.GrupoInput, *paging.PageArgs) []*model.Grupo); ok {
		r0 = rf(ctx, input, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Grupo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.GrupoInput, *paging.PageArgs) error); ok {
		r1 = rf(ctx, input, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Me provides a mock function with given fields: ctx
func (_m *QueryResolver) Me(ctx context.Context) (*model.Usuario, error) {
	ret := _m.Called(ctx)

	var r0 *model.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*model.Usuario, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *model.Usuario); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Permiso provides a mock function with given fields: ctx, id
func (_m *QueryResolver) Permiso(ctx context.Context, id string) (*model.Permiso, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Permiso
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Permiso, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Permiso); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Permiso)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Permisos provides a mock function with given fields: ctx, input, page
func (_m *QueryResolver) Permisos(ctx context.Context, input *model.PermisoInput, page *paging.PageArgs) ([]*model.Permiso, error) {
	ret := _m.Called(ctx, input, page)

	var r0 []*model.Permiso
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.PermisoInput, *paging.PageArgs) ([]*model.Permiso, error)); ok {
		return rf(ctx, input, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.PermisoInput, *paging.PageArgs) []*model.Permiso); ok {
		r0 = rf(ctx, input, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Permiso)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.PermisoInput, *paging.PageArgs) error); ok {
		r1 = rf(ctx, input, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
// Usuario provides a mock function with given fields: ctx, id
func (_m *QueryResolver) Usuario(ctx context.Context, id string) (*model.Usuario, error) {
	ret := _m.Called(ctx, id)

	var r0 *model.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*model.Usuario, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.Usuario); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usuarios provides a mock function with given fields: ctx, input
func (_m *QueryResolver) Usuarios(ctx context.Context, input *model.UsuarioInput) ([]*model.Usuario, error) {
	ret := _m.Called(ctx, input)

	var r0 []*model.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.UsuarioInput) ([]*model.Usuario, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.UsuarioInput) []*model.Usuario); ok {
		r0 = rf(ctx, input)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.UsuarioInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewQueryResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewQueryResolver creates a new instance of QueryResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewQueryResolver(t mockConstructorTestingTNewQueryResolver) *QueryResolver {
	mock := &QueryResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
