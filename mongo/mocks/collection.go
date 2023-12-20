package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Collection struct {
	mock.Mock
}

func (m *Collection) Aggregate(ctx context.Context, pipeline interface{}) (mongo.Cursor, error) {
	ret := m.Called(ctx, pipeline)

	var cursorResult mongo.Cursor
	if cursorFunc, ok := ret.Get(0).(func(context.Context, interface{}) mongo.Cursor); ok {
		cursorResult = cursorFunc(ctx, pipeline)
	} else {
		if ret.Get(0) != nil {
			cursorResult = ret.Get(0).(mongo.Cursor)
		}
	}

	var errResult error
	if errFunc, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		errResult = errFunc(ctx, pipeline)
	} else {
		errResult = ret.Error(1)
	}

	return cursorResult, errResult
}

func (m *Collection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	// Menyiapkan variabel slice untuk menyimpan parameter yang akan diberikan ke metode Called()
	var parameters []interface{}
	// Memasukkan parameter ctx dan filter ke dalam slice parameters
	parameters = append(parameters, ctx, filter)
	// Mengonversi opsional CountOptions ke dalam bentuk slice interface{}
	var optionalArgs []interface{}
	for _, opt := range opts {
		optionalArgs = append(optionalArgs, opt)
	}
	// Menambahkan elemen opsional CountOptions ke dalam slice parameters
	parameters = append(parameters, optionalArgs...)

	// Memanggil metode Called() dengan parameter yang telah disiapkan sebelumnya
	ret := m.Called(parameters...)

	var resultCount int64
	// Memeriksa jika nilai yang direkam adalah fungsi dengan signature yang tepat, jika ya, panggil fungsi tersebut
	if countFunc, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.CountOptions) int64); ok {
		resultCount = countFunc(ctx, filter, opts...)
	} else {
		// Jika nilai yang direkam bukan fungsi, maka itu adalah nilai int64, langsung asumsikan sebagai nilai kembalian
		resultCount = ret.Get(0).(int64)
	}

	var resultError error
	// Sama seperti sebelumnya, memeriksa dan menentukan apakah nilai yang direkam adalah fungsi, jika ya, panggil fungsi tersebut
	if errorFunc, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.CountOptions) error); ok {
		resultError = errorFunc(ctx, filter, opts...)
	} else {
		// Jika nilai yang direkam bukan fungsi, itu adalah nilai error, langsung asumsikan sebagai nilai kembalian
		resultError = ret.Error(1)
	}

	return resultCount, resultError
}

func (m *Collection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	// Memanggil metode Called() dengan parameter yang diberikan
	ret := m.Called(ctx, filter)

	var deletedCount int64
	// Memeriksa apakah nilai yang direkam adalah fungsi dengan signature yang tepat, jika ya, panggil fungsi tersebut
	if countFunc, ok := ret.Get(0).(func(context.Context, interface{}) int64); ok {
		deletedCount = countFunc(ctx, filter)
	} else {
		// Jika nilai yang direkam bukan fungsi, maka itu adalah nilai int64, langsung asumsikan sebagai nilai kembalian
		deletedCount = ret.Get(0).(int64)
	}

	var deleteError error
	// Memeriksa dan menentukan apakah nilai yang direkam adalah fungsi, jika ya, panggil fungsi tersebut
	if errorFunc, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		deleteError = errorFunc(ctx, filter)
	} else {
		// Jika nilai yang direkam bukan fungsi, itu adalah nilai error, langsung asumsikan sebagai nilai kembalian
		deleteError = ret.Error(1)
	}

	return deletedCount, deleteError
}

func (m *Collection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (mongo.Cursor, error) {
	// Mengonversi opsional FindOptions ke dalam bentuk slice interface{}
	var optionalArgs []interface{}
	for _, opt := range opts {
		optionalArgs = append(optionalArgs, opt)
	}
	// Menyiapkan slice _ca untuk menyimpan parameter yang akan diberikan ke metode Called()
	var parameters []interface{}
	// Memasukkan parameter ctx dan filter ke dalam slice parameters
	parameters = append(parameters, ctx, filter)
	// Menambahkan elemen opsional FindOptions ke dalam slice parameters
	parameters = append(parameters, optionalArgs...)

	// Memanggil metode Called() dengan parameter yang telah disiapkan sebelumnya
	ret := m.Called(parameters...)

	var foundCursor mongo.Cursor
	// Memeriksa jika nilai yang direkam adalah fungsi dengan signature yang tepat, jika ya, panggil fungsi tersebut
	if cursorFunc, ok := ret.Get(0).(func(context.Context, interface{}, ...*options.FindOptions) mongo.Cursor); ok {
		foundCursor = cursorFunc(ctx, filter, opts...)
	} else {
		// Jika nilai yang direkam bukan fungsi, maka itu adalah nilai Cursor, langsung asumsikan sebagai nilai kembalian
		foundCursor = ret.Get(0).(mongo.Cursor)
	}

	var findError error
	// Memeriksa dan menentukan apakah nilai yang direkam adalah fungsi, jika ya, panggil fungsi tersebut
	if errorFunc, ok := ret.Get(1).(func(context.Context, interface{}, ...*options.FindOptions) error); ok {
		findError = errorFunc(ctx, filter, opts...)
	} else {
		// Jika nilai yang direkam bukan fungsi, itu adalah nilai error, langsung asumsikan sebagai nilai kembalian
		findError = ret.Error(1)
	}

	return foundCursor, findError
}

func (m *Collection) FindOne(ctx context.Context, filter interface{}) mongo.SingleResult {
	// Memanggil metode Called() dengan parameter yang diberikan
	ret := m.Called(ctx, filter)

	var foundResult mongo.SingleResult
	// Memeriksa apakah nilai yang direkam adalah fungsi dengan signature yang tepat, jika ya, panggil fungsi tersebut
	if resultFunc, ok := ret.Get(0).(func(context.Context, interface{}) mongo.SingleResult); ok {
		foundResult = resultFunc(ctx, filter)
	} else {
		// Jika nilai yang direkam bukan fungsi, maka itu adalah nilai SingleResult, langsung asumsikan sebagai nilai kembalian
		foundResult = ret.Get(0).(mongo.SingleResult)
	}

	return foundResult
}

func (m *Collection) InsertMany(ctx context.Context, documents []interface{}) ([]interface{}, error) {
	ret := m.Called(ctx, documents)

	var result []interface{}
	if rf, ok := ret.Get(0).(func(context.Context, []interface{}) []interface{}); ok {
		result = rf(ctx, documents)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).([]interface{})
		}
	}

	var err error
	if rf, ok := ret.Get(1).(func(context.Context, []interface{}) error); ok {
		err = rf(ctx, documents)
	} else {
		err = ret.Error(1)
	}

	return result, err
}

func (m *Collection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	ret := m.Called(ctx, document)

	var result interface{}
	if rf, ok := ret.Get(0).(func(context.Context, interface{}) interface{}); ok {
		result = rf(ctx, document)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).(interface{})
		}
	}

	var err error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}) error); ok {
		err = rf(ctx, document)
	} else {
		err = ret.Error(1)
	}

	return result, err
}

func (m *Collection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := m.Called(_ca...)

	var result *mongo.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo.UpdateResult); ok {
		result = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	var err error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		err = rf(ctx, filter, update, opts...)
	} else {
		err = ret.Error(1)
	}

	return result, err
}

func (m *Collection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	_va := make([]interface{}, len(opts))
	for _i := range opts {
		_va[_i] = opts[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, filter, update)
	_ca = append(_ca, _va...)
	ret := m.Called(_ca...)

	var result *mongo.UpdateResult
	if rf, ok := ret.Get(0).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) *mongo.UpdateResult); ok {
		result = rf(ctx, filter, update, opts...)
	} else {
		if ret.Get(0) != nil {
			result = ret.Get(0).(*mongo.UpdateResult)
		}
	}

	var err error
	if rf, ok := ret.Get(1).(func(context.Context, interface{}, interface{}, ...*options.UpdateOptions) error); ok {
		err = rf(ctx, filter, update, opts...)
	} else {
		err = ret.Error(1)
	}

	return result, err
}

type mockConstructorTestingTNewCollection interface {
	mock.TestingT
	Cleanup(func())
}

// NewCollection creates a new instance of Collection. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewCollection(t mockConstructorTestingTNewCollection) *Collection {
	mock := &Collection{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
