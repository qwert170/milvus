// Code generated by mockery v2.14.0. DO NOT EDIT.

package meta

import (
	mock "github.com/stretchr/testify/mock"
	clientv3 "go.etcd.io/etcd/client/v3"

	querypb "github.com/milvus-io/milvus/internal/proto/querypb"
)

// MockStore is an autogenerated mock type for the Store type
type MockStore struct {
	mock.Mock
}

type MockStore_Expecter struct {
	mock *mock.Mock
}

func (_m *MockStore) EXPECT() *MockStore_Expecter {
	return &MockStore_Expecter{mock: &_m.Mock}
}

// GetCollections provides a mock function with given fields:
func (_m *MockStore) GetCollections() ([]*querypb.CollectionLoadInfo, error) {
	ret := _m.Called()

	var r0 []*querypb.CollectionLoadInfo
	if rf, ok := ret.Get(0).(func() []*querypb.CollectionLoadInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.CollectionLoadInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetCollections_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCollections'
type MockStore_GetCollections_Call struct {
	*mock.Call
}

// GetCollections is a helper method to define mock.On call
func (_e *MockStore_Expecter) GetCollections() *MockStore_GetCollections_Call {
	return &MockStore_GetCollections_Call{Call: _e.mock.On("GetCollections")}
}

func (_c *MockStore_GetCollections_Call) Run(run func()) *MockStore_GetCollections_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStore_GetCollections_Call) Return(_a0 []*querypb.CollectionLoadInfo, _a1 error) *MockStore_GetCollections_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetPartitions provides a mock function with given fields:
func (_m *MockStore) GetPartitions() (map[int64][]*querypb.PartitionLoadInfo, error) {
	ret := _m.Called()

	var r0 map[int64][]*querypb.PartitionLoadInfo
	if rf, ok := ret.Get(0).(func() map[int64][]*querypb.PartitionLoadInfo); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[int64][]*querypb.PartitionLoadInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetPartitions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPartitions'
type MockStore_GetPartitions_Call struct {
	*mock.Call
}

// GetPartitions is a helper method to define mock.On call
func (_e *MockStore_Expecter) GetPartitions() *MockStore_GetPartitions_Call {
	return &MockStore_GetPartitions_Call{Call: _e.mock.On("GetPartitions")}
}

func (_c *MockStore_GetPartitions_Call) Run(run func()) *MockStore_GetPartitions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStore_GetPartitions_Call) Return(_a0 map[int64][]*querypb.PartitionLoadInfo, _a1 error) *MockStore_GetPartitions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetReplicas provides a mock function with given fields:
func (_m *MockStore) GetReplicas() ([]*querypb.Replica, error) {
	ret := _m.Called()

	var r0 []*querypb.Replica
	if rf, ok := ret.Get(0).(func() []*querypb.Replica); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*querypb.Replica)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockStore_GetReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetReplicas'
type MockStore_GetReplicas_Call struct {
	*mock.Call
}

// GetReplicas is a helper method to define mock.On call
func (_e *MockStore_Expecter) GetReplicas() *MockStore_GetReplicas_Call {
	return &MockStore_GetReplicas_Call{Call: _e.mock.On("GetReplicas")}
}

func (_c *MockStore_GetReplicas_Call) Run(run func()) *MockStore_GetReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStore_GetReplicas_Call) Return(_a0 []*querypb.Replica, _a1 error) *MockStore_GetReplicas_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// LoadHandoffWithRevision provides a mock function with given fields:
func (_m *MockStore) LoadHandoffWithRevision() ([]string, []string, int64, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 []string
	if rf, ok := ret.Get(1).(func() []string); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).([]string)
		}
	}

	var r2 int64
	if rf, ok := ret.Get(2).(func() int64); ok {
		r2 = rf()
	} else {
		r2 = ret.Get(2).(int64)
	}

	var r3 error
	if rf, ok := ret.Get(3).(func() error); ok {
		r3 = rf()
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockStore_LoadHandoffWithRevision_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'LoadHandoffWithRevision'
type MockStore_LoadHandoffWithRevision_Call struct {
	*mock.Call
}

// LoadHandoffWithRevision is a helper method to define mock.On call
func (_e *MockStore_Expecter) LoadHandoffWithRevision() *MockStore_LoadHandoffWithRevision_Call {
	return &MockStore_LoadHandoffWithRevision_Call{Call: _e.mock.On("LoadHandoffWithRevision")}
}

func (_c *MockStore_LoadHandoffWithRevision_Call) Run(run func()) *MockStore_LoadHandoffWithRevision_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockStore_LoadHandoffWithRevision_Call) Return(_a0 []string, _a1 []string, _a2 int64, _a3 error) *MockStore_LoadHandoffWithRevision_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

// ReleaseCollection provides a mock function with given fields: id
func (_m *MockStore) ReleaseCollection(id int64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_ReleaseCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseCollection'
type MockStore_ReleaseCollection_Call struct {
	*mock.Call
}

// ReleaseCollection is a helper method to define mock.On call
//  - id int64
func (_e *MockStore_Expecter) ReleaseCollection(id interface{}) *MockStore_ReleaseCollection_Call {
	return &MockStore_ReleaseCollection_Call{Call: _e.mock.On("ReleaseCollection", id)}
}

func (_c *MockStore_ReleaseCollection_Call) Run(run func(id int64)) *MockStore_ReleaseCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockStore_ReleaseCollection_Call) Return(_a0 error) *MockStore_ReleaseCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

// ReleasePartition provides a mock function with given fields: collection, partitions
func (_m *MockStore) ReleasePartition(collection int64, partitions ...int64) error {
	_va := make([]interface{}, len(partitions))
	for _i := range partitions {
		_va[_i] = partitions[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, collection)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, ...int64) error); ok {
		r0 = rf(collection, partitions...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_ReleasePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleasePartition'
type MockStore_ReleasePartition_Call struct {
	*mock.Call
}

// ReleasePartition is a helper method to define mock.On call
//  - collection int64
//  - partitions ...int64
func (_e *MockStore_Expecter) ReleasePartition(collection interface{}, partitions ...interface{}) *MockStore_ReleasePartition_Call {
	return &MockStore_ReleasePartition_Call{Call: _e.mock.On("ReleasePartition",
		append([]interface{}{collection}, partitions...)...)}
}

func (_c *MockStore_ReleasePartition_Call) Run(run func(collection int64, partitions ...int64)) *MockStore_ReleasePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]int64, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(int64)
			}
		}
		run(args[0].(int64), variadicArgs...)
	})
	return _c
}

func (_c *MockStore_ReleasePartition_Call) Return(_a0 error) *MockStore_ReleasePartition_Call {
	_c.Call.Return(_a0)
	return _c
}

// ReleaseReplica provides a mock function with given fields: collection, replica
func (_m *MockStore) ReleaseReplica(collection int64, replica int64) error {
	ret := _m.Called(collection, replica)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64, int64) error); ok {
		r0 = rf(collection, replica)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_ReleaseReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseReplica'
type MockStore_ReleaseReplica_Call struct {
	*mock.Call
}

// ReleaseReplica is a helper method to define mock.On call
//  - collection int64
//  - replica int64
func (_e *MockStore_Expecter) ReleaseReplica(collection interface{}, replica interface{}) *MockStore_ReleaseReplica_Call {
	return &MockStore_ReleaseReplica_Call{Call: _e.mock.On("ReleaseReplica", collection, replica)}
}

func (_c *MockStore_ReleaseReplica_Call) Run(run func(collection int64, replica int64)) *MockStore_ReleaseReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64), args[1].(int64))
	})
	return _c
}

func (_c *MockStore_ReleaseReplica_Call) Return(_a0 error) *MockStore_ReleaseReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

// ReleaseReplicas provides a mock function with given fields: collectionID
func (_m *MockStore) ReleaseReplicas(collectionID int64) error {
	ret := _m.Called(collectionID)

	var r0 error
	if rf, ok := ret.Get(0).(func(int64) error); ok {
		r0 = rf(collectionID)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_ReleaseReplicas_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ReleaseReplicas'
type MockStore_ReleaseReplicas_Call struct {
	*mock.Call
}

// ReleaseReplicas is a helper method to define mock.On call
//  - collectionID int64
func (_e *MockStore_Expecter) ReleaseReplicas(collectionID interface{}) *MockStore_ReleaseReplicas_Call {
	return &MockStore_ReleaseReplicas_Call{Call: _e.mock.On("ReleaseReplicas", collectionID)}
}

func (_c *MockStore_ReleaseReplicas_Call) Run(run func(collectionID int64)) *MockStore_ReleaseReplicas_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockStore_ReleaseReplicas_Call) Return(_a0 error) *MockStore_ReleaseReplicas_Call {
	_c.Call.Return(_a0)
	return _c
}

// RemoveHandoffEvent provides a mock function with given fields: segmentInfo
func (_m *MockStore) RemoveHandoffEvent(segmentInfo *querypb.SegmentInfo) error {
	ret := _m.Called(segmentInfo)

	var r0 error
	if rf, ok := ret.Get(0).(func(*querypb.SegmentInfo) error); ok {
		r0 = rf(segmentInfo)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_RemoveHandoffEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RemoveHandoffEvent'
type MockStore_RemoveHandoffEvent_Call struct {
	*mock.Call
}

// RemoveHandoffEvent is a helper method to define mock.On call
//  - segmentInfo *querypb.SegmentInfo
func (_e *MockStore_Expecter) RemoveHandoffEvent(segmentInfo interface{}) *MockStore_RemoveHandoffEvent_Call {
	return &MockStore_RemoveHandoffEvent_Call{Call: _e.mock.On("RemoveHandoffEvent", segmentInfo)}
}

func (_c *MockStore_RemoveHandoffEvent_Call) Run(run func(segmentInfo *querypb.SegmentInfo)) *MockStore_RemoveHandoffEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*querypb.SegmentInfo))
	})
	return _c
}

func (_c *MockStore_RemoveHandoffEvent_Call) Return(_a0 error) *MockStore_RemoveHandoffEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveCollection provides a mock function with given fields: info
func (_m *MockStore) SaveCollection(info *querypb.CollectionLoadInfo) error {
	ret := _m.Called(info)

	var r0 error
	if rf, ok := ret.Get(0).(func(*querypb.CollectionLoadInfo) error); ok {
		r0 = rf(info)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_SaveCollection_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveCollection'
type MockStore_SaveCollection_Call struct {
	*mock.Call
}

// SaveCollection is a helper method to define mock.On call
//  - info *querypb.CollectionLoadInfo
func (_e *MockStore_Expecter) SaveCollection(info interface{}) *MockStore_SaveCollection_Call {
	return &MockStore_SaveCollection_Call{Call: _e.mock.On("SaveCollection", info)}
}

func (_c *MockStore_SaveCollection_Call) Run(run func(info *querypb.CollectionLoadInfo)) *MockStore_SaveCollection_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*querypb.CollectionLoadInfo))
	})
	return _c
}

func (_c *MockStore_SaveCollection_Call) Return(_a0 error) *MockStore_SaveCollection_Call {
	_c.Call.Return(_a0)
	return _c
}

// SavePartition provides a mock function with given fields: info
func (_m *MockStore) SavePartition(info ...*querypb.PartitionLoadInfo) error {
	_va := make([]interface{}, len(info))
	for _i := range info {
		_va[_i] = info[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 error
	if rf, ok := ret.Get(0).(func(...*querypb.PartitionLoadInfo) error); ok {
		r0 = rf(info...)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_SavePartition_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SavePartition'
type MockStore_SavePartition_Call struct {
	*mock.Call
}

// SavePartition is a helper method to define mock.On call
//  - info ...*querypb.PartitionLoadInfo
func (_e *MockStore_Expecter) SavePartition(info ...interface{}) *MockStore_SavePartition_Call {
	return &MockStore_SavePartition_Call{Call: _e.mock.On("SavePartition",
		append([]interface{}{}, info...)...)}
}

func (_c *MockStore_SavePartition_Call) Run(run func(info ...*querypb.PartitionLoadInfo)) *MockStore_SavePartition_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]*querypb.PartitionLoadInfo, len(args)-0)
		for i, a := range args[0:] {
			if a != nil {
				variadicArgs[i] = a.(*querypb.PartitionLoadInfo)
			}
		}
		run(variadicArgs...)
	})
	return _c
}

func (_c *MockStore_SavePartition_Call) Return(_a0 error) *MockStore_SavePartition_Call {
	_c.Call.Return(_a0)
	return _c
}

// SaveReplica provides a mock function with given fields: replica
func (_m *MockStore) SaveReplica(replica *querypb.Replica) error {
	ret := _m.Called(replica)

	var r0 error
	if rf, ok := ret.Get(0).(func(*querypb.Replica) error); ok {
		r0 = rf(replica)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockStore_SaveReplica_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SaveReplica'
type MockStore_SaveReplica_Call struct {
	*mock.Call
}

// SaveReplica is a helper method to define mock.On call
//  - replica *querypb.Replica
func (_e *MockStore_Expecter) SaveReplica(replica interface{}) *MockStore_SaveReplica_Call {
	return &MockStore_SaveReplica_Call{Call: _e.mock.On("SaveReplica", replica)}
}

func (_c *MockStore_SaveReplica_Call) Run(run func(replica *querypb.Replica)) *MockStore_SaveReplica_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(*querypb.Replica))
	})
	return _c
}

func (_c *MockStore_SaveReplica_Call) Return(_a0 error) *MockStore_SaveReplica_Call {
	_c.Call.Return(_a0)
	return _c
}

// WatchHandoffEvent provides a mock function with given fields: revision
func (_m *MockStore) WatchHandoffEvent(revision int64) clientv3.WatchChan {
	ret := _m.Called(revision)

	var r0 clientv3.WatchChan
	if rf, ok := ret.Get(0).(func(int64) clientv3.WatchChan); ok {
		r0 = rf(revision)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(clientv3.WatchChan)
		}
	}

	return r0
}

// MockStore_WatchHandoffEvent_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WatchHandoffEvent'
type MockStore_WatchHandoffEvent_Call struct {
	*mock.Call
}

// WatchHandoffEvent is a helper method to define mock.On call
//  - revision int64
func (_e *MockStore_Expecter) WatchHandoffEvent(revision interface{}) *MockStore_WatchHandoffEvent_Call {
	return &MockStore_WatchHandoffEvent_Call{Call: _e.mock.On("WatchHandoffEvent", revision)}
}

func (_c *MockStore_WatchHandoffEvent_Call) Run(run func(revision int64)) *MockStore_WatchHandoffEvent_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int64))
	})
	return _c
}

func (_c *MockStore_WatchHandoffEvent_Call) Return(_a0 clientv3.WatchChan) *MockStore_WatchHandoffEvent_Call {
	_c.Call.Return(_a0)
	return _c
}

type mockConstructorTestingTNewMockStore interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockStore creates a new instance of MockStore. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockStore(t mockConstructorTestingTNewMockStore) *MockStore {
	mock := &MockStore{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
