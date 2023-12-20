package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client struct {
	mock.Mock
}

// Client adalah tipe mock yang dihasilkan secara otomatis untuk tipe Client
func (m *Client) Connect(ctx context.Context) error {
	ret := m.Called(ctx)

	var returnErr error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		returnErr = rf(ctx)
	} else {
		returnErr = ret.Error(0)
	}

	return returnErr
}

// Connect menyediakan sebuah fungsi tiruan dengan bidang yang diberikan: ctx
func (m *Client) Database(dbName string) mongo.Database {
	ret := m.Called(dbName)

	var returnErr mongo.Database
	if databaseFunc, ok := ret.Get(0).(func(string) mongo.Database); ok {
		returnErr = databaseFunc(dbName)
	} else {
		if ret.Get(0) != nil {
			returnErr = ret.Get(0).(mongo.Database)
		}
	}

	return returnErr
}

func (m *Client) Disconnect(ctx context.Context) error {
	ret := m.Called(ctx)

	var returnErr error
	if disconnectFunc, ok := ret.Get(0).(func(context.Context) error); ok {
		returnErr = disconnectFunc(ctx)
	} else {
		returnErr = ret.Error(0)
	}

	return returnErr
}

func (m *Client) Ping(ctx context.Context) error {
	ret := m.Called(ctx)

	var returnErr error
	if pingFunc, ok := ret.Get(0).(func(context.Context) error); ok {
		returnErr = pingFunc(ctx)
	} else {
		returnErr = ret.Error(0)
	}

	return returnErr
}

func (m *Client) StartSession() (mongo.Session, error) {
	ret := m.Called()

	var returnErr mongo.Session
	if sessionFunc, ok := ret.Get(0).(func() mongo.Session); ok {
		returnErr = sessionFunc()
	} else {
		if ret.Get(0) != nil {
			returnErr = ret.Get(0).(mongo.Session)
		}
	}

	var err error
	if errFunc, ok := ret.Get(1).(func() error); ok {
		err = errFunc()
	} else {
		err = ret.Error(1)
	}

	return returnErr, err
}

func (m *Client) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	ret := m.Called(ctx, fn)

	var returnErr error
	if useSessionFunc, ok := ret.Get(0).(func(context.Context, func(mongo.SessionContext) error) error); ok {
		returnErr = useSessionFunc(ctx, fn)
	} else {
		returnErr = ret.Error(0)
	}

	return returnErr
}

type mockConstructorTestingTNewClient interface {
	mock.TestingT
	Cleanup(func())
}

func NewClient(t mockConstructorTestingTNewClient) *Client {
	mock := &Client{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
