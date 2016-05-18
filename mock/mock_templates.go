// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/docker/libmachete (interfaces: TemplateLoader)

package mock

import (
	gomock "github.com/golang/mock/gomock"
)

// Mock of TemplateLoader interface
type MockTemplateLoader struct {
	ctrl     *gomock.Controller
	recorder *_MockTemplateLoaderRecorder
}

// Recorder for MockTemplateLoader (not exported)
type _MockTemplateLoaderRecorder struct {
	mock *MockTemplateLoader
}

func NewMockTemplateLoader(ctrl *gomock.Controller) *MockTemplateLoader {
	mock := &MockTemplateLoader{ctrl: ctrl}
	mock.recorder = &_MockTemplateLoaderRecorder{mock}
	return mock
}

func (_m *MockTemplateLoader) EXPECT() *_MockTemplateLoaderRecorder {
	return _m.recorder
}

func (_m *MockTemplateLoader) Read(_param0 string, _param1 string) ([]byte, error) {
	ret := _m.ctrl.Call(_m, "Read", _param0, _param1)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockTemplateLoaderRecorder) Read(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Read", arg0, arg1)
}
