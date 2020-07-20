// Copyright DataStax, Inc.
// Please see the included license file for details.

// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import mock "github.com/stretchr/testify/mock"
import reconcile "sigs.k8s.io/controller-runtime/pkg/reconcile"

// Reconciler is an autogenerated mock type for the Reconciler type
type Reconciler struct {
	mock.Mock
}

// Apply provides a mock function with given fields:
func (_m *Reconciler) Apply() (reconcile.Result, error) {
	ret := _m.Called()

	var r0 reconcile.Result
	if rf, ok := ret.Get(0).(func() reconcile.Result); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(reconcile.Result)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}