// Code generated by mockery v2.1.0. DO NOT EDIT.

package flow

import (
	context "context"

	mock "github.com/stretchr/testify/mock"
	git "gopkg.in/src-d/go-git.v4"

	plumbing "gopkg.in/src-d/go-git.v4/plumbing"
)

// MockGitService is an autogenerated mock type for the GitService type
type MockGitService struct {
	mock.Mock
}

// Checkout provides a mock function with given fields: ctx, rootPath, hash
func (_m *MockGitService) Checkout(ctx context.Context, rootPath string, hash plumbing.Hash) error {
	ret := _m.Called(ctx, rootPath, hash)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, plumbing.Hash) error); ok {
		r0 = rf(ctx, rootPath, hash)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Clone provides a mock function with given fields: _a0, _a1
func (_m *MockGitService) Clone(_a0 context.Context, _a1 string) (*git.Repository, error) {
	ret := _m.Called(_a0, _a1)

	var r0 *git.Repository
	if rf, ok := ret.Get(0).(func(context.Context, string) *git.Repository); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*git.Repository)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: ctx, rootPath, changesPath, msg
func (_m *MockGitService) Commit(ctx context.Context, rootPath string, changesPath string, msg string) error {
	ret := _m.Called(ctx, rootPath, changesPath, msg)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) error); ok {
		r0 = rf(ctx, rootPath, changesPath, msg)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// LocateServiceReleaseRollbackSkip provides a mock function with given fields: ctx, r, env, service, n
func (_m *MockGitService) LocateServiceReleaseRollbackSkip(ctx context.Context, r *git.Repository, env string, service string, n uint) (plumbing.Hash, error) {
	ret := _m.Called(ctx, r, env, service, n)

	var r0 plumbing.Hash
	if rf, ok := ret.Get(0).(func(context.Context, *git.Repository, string, string, uint) plumbing.Hash); ok {
		r0 = rf(ctx, r, env, service, n)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(plumbing.Hash)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *git.Repository, string, string, uint) error); ok {
		r1 = rf(ctx, r, env, service, n)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MasterPath provides a mock function with given fields:
func (_m *MockGitService) MasterPath() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// SyncMaster provides a mock function with given fields: _a0
func (_m *MockGitService) SyncMaster(_a0 context.Context) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
