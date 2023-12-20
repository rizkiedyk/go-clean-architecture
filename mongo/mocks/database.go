package mocks

import (
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type Database struct {
	mock.Mock
}

func (m *Database) Client() mongo.Client {
	ret := m.Called()

	var returnErr mongo.Client
	if databaseFunc, ok := ret.Get(0).(func() mongo.Client); ok {
		returnErr = databaseFunc()
	} else {
		if ret.Get(0) != nil {
			returnErr = ret.Get(0).(mongo.Client)
		}
	}

	return returnErr
}

func (m *Database) Collection(collectionName string) mongo.Collection {
	ret := m.Called(collectionName)

	var returnErr mongo.Collection
	if collectionFunc, ok := ret.Get(0).(func(string) mongo.Collection); ok {
		returnErr = collectionFunc(collectionName)
	} else {
		if ret.Get(0) != nil {
			returnErr = ret.Get(0).(mongo.Collection)
		}
	}

	return returnErr
}

type mockConstructorTestingTNewDatabase interface {
	mock.TestingT
	Cleanup(func())
}

func NewDatabase(t mockConstructorTestingTNewDatabase) *Database {
	mock := &Database{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
