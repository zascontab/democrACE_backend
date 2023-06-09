// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	model "github.com/sonderkevin/governance/graph/model"

	paging "github.com/nrfta/go-paging"
)

// GrupoResolver is an autogenerated mock type for the GrupoResolver type
type GrupoResolver struct {
	mock.Mock
}

// Permisos provides a mock function with given fields: ctx, obj, input, page
func (_m *GrupoResolver) Permisos(ctx context.Context, obj *model.Grupo, input *model.PermisoInput, page *paging.PageArgs) ([]*model.Permiso, error) {
	ret := _m.Called(ctx, obj, input, page)

	var r0 []*model.Permiso
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Grupo, *model.PermisoInput, *paging.PageArgs) ([]*model.Permiso, error)); ok {
		return rf(ctx, obj, input, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Grupo, *model.PermisoInput, *paging.PageArgs) []*model.Permiso); ok {
		r0 = rf(ctx, obj, input, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Permiso)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Grupo, *model.PermisoInput, *paging.PageArgs) error); ok {
		r1 = rf(ctx, obj, input, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Usuarios provides a mock function with given fields: ctx, obj, input, page
func (_m *GrupoResolver) Usuarios(ctx context.Context, obj *model.Grupo, input *model.UsuarioInput, page *paging.PageArgs) ([]*model.Usuario, error) {
	ret := _m.Called(ctx, obj, input, page)

	var r0 []*model.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, *model.Grupo, *model.UsuarioInput, *paging.PageArgs) ([]*model.Usuario, error)); ok {
		return rf(ctx, obj, input, page)
	}
	if rf, ok := ret.Get(0).(func(context.Context, *model.Grupo, *model.UsuarioInput, *paging.PageArgs) []*model.Usuario); ok {
		r0 = rf(ctx, obj, input, page)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, *model.Grupo, *model.UsuarioInput, *paging.PageArgs) error); ok {
		r1 = rf(ctx, obj, input, page)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewGrupoResolver interface {
	mock.TestingT
	Cleanup(func())
}

// NewGrupoResolver creates a new instance of GrupoResolver. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewGrupoResolver(t mockConstructorTestingTNewGrupoResolver) *GrupoResolver {
	mock := &GrupoResolver{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
