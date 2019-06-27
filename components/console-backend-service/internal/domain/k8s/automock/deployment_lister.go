// Code generated by mockery v1.0.0
package automock

import mock "github.com/stretchr/testify/mock"
import v1beta2 "k8s.io/api/apps/v1beta2"

// deploymentLister is an autogenerated mock type for the deploymentLister type
type deploymentLister struct {
	mock.Mock
}

// List provides a mock function with given fields: namespace
func (_m *deploymentLister) List(namespace string) ([]*v1beta2.Deployment, error) {
	ret := _m.Called(namespace)

	var r0 []*v1beta2.Deployment
	if rf, ok := ret.Get(0).(func(string) []*v1beta2.Deployment); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta2.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ListWithoutFunctions provides a mock function with given fields: namespace
func (_m *deploymentLister) ListWithoutFunctions(namespace string) ([]*v1beta2.Deployment, error) {
	ret := _m.Called(namespace)

	var r0 []*v1beta2.Deployment
	if rf, ok := ret.Get(0).(func(string) []*v1beta2.Deployment); ok {
		r0 = rf(namespace)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*v1beta2.Deployment)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
