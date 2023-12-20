package mocks

import "github.com/stretchr/testify/mock"

type SingleResult struct {
	mock.Mock
}

func (m *SingleResult) Decode(result interface{}) error {
	ret := m.Called(result)

	var err error
	if decodeFunc, ok := ret.Get(0).(func(interface{}) error); ok {
		err = decodeFunc(result)
	} else {
		err = ret.Error(0)
	}

	return err
}

type mockConstructorTestingTNewSingleResult interface {
	mock.TestingT
	Cleanup(func())
}

func NewSingleResult(t mockConstructorTestingTNewSingleResult) *SingleResult {
	mock := &SingleResult{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
