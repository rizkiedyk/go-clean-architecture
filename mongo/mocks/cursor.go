package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type Cursor struct {
	mock.Mock
}

func (m *Cursor) All(ctx context.Context, result interface{}) error {
	ret := m.Called(ctx, result)

	var err error
	if allFuc, ok := ret.Get(0).(func(context.Context, interface{}) error); ok {
		err = allFuc(ctx, result)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (m *Cursor) Close(ctx context.Context) error {
	ret := m.Called(ctx)

	var err error
	if closeFunc, ok := ret.Get(0).(func(context.Context) error); ok {
		err = closeFunc(ctx)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (m *Cursor) Decode(result interface{}) error {
	ret := m.Called(result)

	var err error
	if decodeFunc, ok := ret.Get(0).(func(interface{}) error); ok {
		err = decodeFunc(result)
	} else {
		err = ret.Error(0)
	}

	return err
}

func (m *Cursor) Next(ctx context.Context) bool {
	ret := m.Called(ctx)

	var next bool
	if nextFunc, ok := ret.Get(0).(func(context.Context) bool); ok {
		next = nextFunc(ctx)
	} else {
		next = ret.Get(0).(bool)
	}

	return next
}

type mockConstructorTestingTNewCursor interface {
	mock.TestingT
	Cleanup(func())
}

func NewCursor(t mockConstructorTestingTNewCursor) *Cursor {
	mock := &Cursor{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
