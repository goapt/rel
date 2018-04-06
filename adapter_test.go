package grimoire

import (
	"github.com/stretchr/testify/mock"
)

type TestAdapter struct {
	mock.Mock
}

var _ Adapter = (*TestAdapter)(nil)

func (adapter TestAdapter) Open(dsn string) error {
	args := adapter.Called(dsn)
	return args.Error(0)
}

func (adapter TestAdapter) Close() error {
	args := adapter.Called()
	return args.Error(0)
}
func (adapter TestAdapter) All(query Query, doc interface{}) (int, error) {
	args := adapter.Called(query, doc)
	return args.Int(0), args.Error(1)
}

func (adapter TestAdapter) Insert(query Query, ch map[string]interface{}) (int, error) {
	args := adapter.Called(query, ch)
	return args.Int(0), args.Error(1)
}

func (adapter TestAdapter) Update(query Query, ch map[string]interface{}) error {
	args := adapter.Called(query, ch)
	return args.Error(0)
}

func (adapter TestAdapter) Delete(query Query) error {
	args := adapter.Called(query)
	return args.Error(0)
}

func (adapter TestAdapter) Begin() (Adapter, error) {
	args := adapter.Called()
	return adapter, args.Error(0)
}

func (adapter TestAdapter) Commit() error {
	args := adapter.Called()
	return args.Error(0)
}

func (adapter TestAdapter) Rollback() error {
	args := adapter.Called()
	return args.Error(0)
}

func (adapter TestAdapter) Query(out interface{}, qs string, qargs []interface{}) (int64, error) {
	args := adapter.Called(out, qs, qargs)
	return args.Get(0).(int64), args.Error(1)
}

func (adapter TestAdapter) Exec(qs string, qargs []interface{}) (int64, int64, error) {
	args := adapter.Called(qs, qargs)
	return args.Get(0).(int64), args.Get(1).(int64), args.Error(2)
}
