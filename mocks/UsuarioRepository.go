// Code generated by mockery v2.20.2. DO NOT EDIT.

package mocks

import (
	domain "github.com/sonderkevin/governance/domain"
	mock "github.com/stretchr/testify/mock"
)

// UsuarioRepository is an autogenerated mock type for the UsuarioRepository type
type UsuarioRepository struct {
	mock.Mock
}

// Delete provides a mock function with given fields: _a0
func (_m *UsuarioRepository) Delete(_a0 uint) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: _a0
func (_m *UsuarioRepository) Get(_a0 uint) (*domain.Usuario, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*domain.Usuario, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(uint) *domain.Usuario); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetAll provides a mock function with given fields: _a0
func (_m *UsuarioRepository) GetAll(_a0 *domain.UsuarioFilter) ([]*domain.Usuario, error) {
	ret := _m.Called(_a0)

	var r0 []*domain.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.UsuarioFilter) ([]*domain.Usuario, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*domain.UsuarioFilter) []*domain.Usuario); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.UsuarioFilter) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByEmail provides a mock function with given fields: _a0
func (_m *UsuarioRepository) GetByEmail(_a0 string) (*domain.Usuario, error) {
	ret := _m.Called(_a0)

	var r0 *domain.Usuario
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.Usuario, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.Usuario); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Usuario)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: _a0
func (_m *UsuarioRepository) Save(_a0 *domain.Usuario) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Usuario) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Update provides a mock function with given fields: _a0
func (_m *UsuarioRepository) Update(_a0 *domain.Usuario) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.Usuario) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

type mockConstructorTestingTNewUsuarioRepository interface {
	mock.TestingT
	Cleanup(func())
}

// NewUsuarioRepository creates a new instance of UsuarioRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewUsuarioRepository(t mockConstructorTestingTNewUsuarioRepository) *UsuarioRepository {
	mock := &UsuarioRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}