// Copyright 2015-2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/engine (interfaces: TaskEngine,DockerClient,ImageManager)

// Package engine is a generated GoMock package.
package engine

import (
	context "context"
	io "io"
	reflect "reflect"
	time "time"

	api "github.com/aws/amazon-ecs-agent/agent/api"
	"github.com/aws/amazon-ecs-agent/agent/ecr"
	dockerclient "github.com/aws/amazon-ecs-agent/agent/engine/dockerclient"
	image "github.com/aws/amazon-ecs-agent/agent/engine/image"
	statechange "github.com/aws/amazon-ecs-agent/agent/statechange"
	statemanager "github.com/aws/amazon-ecs-agent/agent/statemanager"
	go_dockerclient "github.com/fsouza/go-dockerclient"
	gomock "github.com/golang/mock/gomock"
)

// MockTaskEngine is a mock of TaskEngine interface
type MockTaskEngine struct {
	ctrl     *gomock.Controller
	recorder *MockTaskEngineMockRecorder
}

// MockTaskEngineMockRecorder is the mock recorder for MockTaskEngine
type MockTaskEngineMockRecorder struct {
	mock *MockTaskEngine
}

// NewMockTaskEngine creates a new mock instance
func NewMockTaskEngine(ctrl *gomock.Controller) *MockTaskEngine {
	mock := &MockTaskEngine{ctrl: ctrl}
	mock.recorder = &MockTaskEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTaskEngine) EXPECT() *MockTaskEngineMockRecorder {
	return m.recorder
}

// AddTask mocks base method
func (m *MockTaskEngine) AddTask(arg0 *api.Task) error {
	ret := m.ctrl.Call(m, "AddTask", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddTask indicates an expected call of AddTask
func (mr *MockTaskEngineMockRecorder) AddTask(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTask", reflect.TypeOf((*MockTaskEngine)(nil).AddTask), arg0)
}

// Disable mocks base method
func (m *MockTaskEngine) Disable() {
	m.ctrl.Call(m, "Disable")
}

// Disable indicates an expected call of Disable
func (mr *MockTaskEngineMockRecorder) Disable() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Disable", reflect.TypeOf((*MockTaskEngine)(nil).Disable))
}

// GetTaskByArn mocks base method
func (m *MockTaskEngine) GetTaskByArn(arg0 string) (*api.Task, bool) {
	ret := m.ctrl.Call(m, "GetTaskByArn", arg0)
	ret0, _ := ret[0].(*api.Task)
	ret1, _ := ret[1].(bool)
	return ret0, ret1
}

// GetTaskByArn indicates an expected call of GetTaskByArn
func (mr *MockTaskEngineMockRecorder) GetTaskByArn(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskByArn", reflect.TypeOf((*MockTaskEngine)(nil).GetTaskByArn), arg0)
}

// Init mocks base method
func (m *MockTaskEngine) Init(arg0 context.Context) error {
	ret := m.ctrl.Call(m, "Init", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Init indicates an expected call of Init
func (mr *MockTaskEngineMockRecorder) Init(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Init", reflect.TypeOf((*MockTaskEngine)(nil).Init), arg0)
}

// ListTasks mocks base method
func (m *MockTaskEngine) ListTasks() ([]*api.Task, error) {
	ret := m.ctrl.Call(m, "ListTasks")
	ret0, _ := ret[0].([]*api.Task)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTasks indicates an expected call of ListTasks
func (mr *MockTaskEngineMockRecorder) ListTasks() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTasks", reflect.TypeOf((*MockTaskEngine)(nil).ListTasks))
}

// MarshalJSON mocks base method
func (m *MockTaskEngine) MarshalJSON() ([]byte, error) {
	ret := m.ctrl.Call(m, "MarshalJSON")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// MarshalJSON indicates an expected call of MarshalJSON
func (mr *MockTaskEngineMockRecorder) MarshalJSON() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarshalJSON", reflect.TypeOf((*MockTaskEngine)(nil).MarshalJSON))
}

// MustInit mocks base method
func (m *MockTaskEngine) MustInit(arg0 context.Context) {
	m.ctrl.Call(m, "MustInit", arg0)
}

// MustInit indicates an expected call of MustInit
func (mr *MockTaskEngineMockRecorder) MustInit(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MustInit", reflect.TypeOf((*MockTaskEngine)(nil).MustInit), arg0)
}

// SetSaver mocks base method
func (m *MockTaskEngine) SetSaver(arg0 statemanager.Saver) {
	m.ctrl.Call(m, "SetSaver", arg0)
}

// SetSaver indicates an expected call of SetSaver
func (mr *MockTaskEngineMockRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSaver", reflect.TypeOf((*MockTaskEngine)(nil).SetSaver), arg0)
}

// StateChangeEvents mocks base method
func (m *MockTaskEngine) StateChangeEvents() chan statechange.Event {
	ret := m.ctrl.Call(m, "StateChangeEvents")
	ret0, _ := ret[0].(chan statechange.Event)
	return ret0
}

// StateChangeEvents indicates an expected call of StateChangeEvents
func (mr *MockTaskEngineMockRecorder) StateChangeEvents() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StateChangeEvents", reflect.TypeOf((*MockTaskEngine)(nil).StateChangeEvents))
}

// UnmarshalJSON mocks base method
func (m *MockTaskEngine) UnmarshalJSON(arg0 []byte) error {
	ret := m.ctrl.Call(m, "UnmarshalJSON", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// UnmarshalJSON indicates an expected call of UnmarshalJSON
func (mr *MockTaskEngineMockRecorder) UnmarshalJSON(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UnmarshalJSON", reflect.TypeOf((*MockTaskEngine)(nil).UnmarshalJSON), arg0)
}

// Version mocks base method
func (m *MockTaskEngine) Version() (string, error) {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockTaskEngineMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockTaskEngine)(nil).Version))
}

// MockDockerClient is a mock of DockerClient interface
type MockDockerClient struct {
	ctrl     *gomock.Controller
	recorder *MockDockerClientMockRecorder
}

// MockDockerClientMockRecorder is the mock recorder for MockDockerClient
type MockDockerClientMockRecorder struct {
	mock *MockDockerClient
}

// NewMockDockerClient creates a new mock instance
func NewMockDockerClient(ctrl *gomock.Controller) *MockDockerClient {
	mock := &MockDockerClient{ctrl: ctrl}
	mock.recorder = &MockDockerClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDockerClient) EXPECT() *MockDockerClientMockRecorder {
	return m.recorder
}

// APIVersion mocks base method
func (m *MockDockerClient) APIVersion() (dockerclient.DockerVersion, error) {
	ret := m.ctrl.Call(m, "APIVersion")
	ret0, _ := ret[0].(dockerclient.DockerVersion)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// APIVersion indicates an expected call of APIVersion
func (mr *MockDockerClientMockRecorder) APIVersion() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "APIVersion", reflect.TypeOf((*MockDockerClient)(nil).APIVersion))
}

// ContainerEvents mocks base method
func (m *MockDockerClient) ContainerEvents(arg0 context.Context) (<-chan DockerContainerChangeEvent, error) {
	ret := m.ctrl.Call(m, "ContainerEvents", arg0)
	ret0, _ := ret[0].(<-chan DockerContainerChangeEvent)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ContainerEvents indicates an expected call of ContainerEvents
func (mr *MockDockerClientMockRecorder) ContainerEvents(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ContainerEvents", reflect.TypeOf((*MockDockerClient)(nil).ContainerEvents), arg0)
}

// CreateContainer mocks base method
func (m *MockDockerClient) CreateContainer(arg0 *go_dockerclient.Config, arg1 *go_dockerclient.HostConfig, arg2 string, arg3 time.Duration) DockerContainerMetadata {
	ret := m.ctrl.Call(m, "CreateContainer", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// CreateContainer indicates an expected call of CreateContainer
func (mr *MockDockerClientMockRecorder) CreateContainer(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateContainer", reflect.TypeOf((*MockDockerClient)(nil).CreateContainer), arg0, arg1, arg2, arg3)
}

// DescribeContainer mocks base method
func (m *MockDockerClient) DescribeContainer(arg0 string) (api.ContainerStatus, DockerContainerMetadata) {
	ret := m.ctrl.Call(m, "DescribeContainer", arg0)
	ret0, _ := ret[0].(api.ContainerStatus)
	ret1, _ := ret[1].(DockerContainerMetadata)
	return ret0, ret1
}

// DescribeContainer indicates an expected call of DescribeContainer
func (mr *MockDockerClientMockRecorder) DescribeContainer(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeContainer", reflect.TypeOf((*MockDockerClient)(nil).DescribeContainer), arg0)
}

// ImportLocalEmptyVolumeImage mocks base method
func (m *MockDockerClient) ImportLocalEmptyVolumeImage() DockerContainerMetadata {
	ret := m.ctrl.Call(m, "ImportLocalEmptyVolumeImage")
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// ImportLocalEmptyVolumeImage indicates an expected call of ImportLocalEmptyVolumeImage
func (mr *MockDockerClientMockRecorder) ImportLocalEmptyVolumeImage() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportLocalEmptyVolumeImage", reflect.TypeOf((*MockDockerClient)(nil).ImportLocalEmptyVolumeImage))
}

// InspectContainer mocks base method
func (m *MockDockerClient) InspectContainer(arg0 string, arg1 time.Duration) (*go_dockerclient.Container, error) {
	ret := m.ctrl.Call(m, "InspectContainer", arg0, arg1)
	ret0, _ := ret[0].(*go_dockerclient.Container)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectContainer indicates an expected call of InspectContainer
func (mr *MockDockerClientMockRecorder) InspectContainer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectContainer", reflect.TypeOf((*MockDockerClient)(nil).InspectContainer), arg0, arg1)
}

// InspectImage mocks base method
func (m *MockDockerClient) InspectImage(arg0 string) (*go_dockerclient.Image, error) {
	ret := m.ctrl.Call(m, "InspectImage", arg0)
	ret0, _ := ret[0].(*go_dockerclient.Image)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InspectImage indicates an expected call of InspectImage
func (mr *MockDockerClientMockRecorder) InspectImage(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InspectImage", reflect.TypeOf((*MockDockerClient)(nil).InspectImage), arg0)
}

// KnownVersions mocks base method
func (m *MockDockerClient) KnownVersions() []dockerclient.DockerVersion {
	ret := m.ctrl.Call(m, "KnownVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// KnownVersions indicates an expected call of KnownVersions
func (mr *MockDockerClientMockRecorder) KnownVersions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "KnownVersions", reflect.TypeOf((*MockDockerClient)(nil).KnownVersions))
}

// ListContainers mocks base method
func (m *MockDockerClient) ListContainers(arg0 bool, arg1 time.Duration) ListContainersResponse {
	ret := m.ctrl.Call(m, "ListContainers", arg0, arg1)
	ret0, _ := ret[0].(ListContainersResponse)
	return ret0
}

// ListContainers indicates an expected call of ListContainers
func (mr *MockDockerClientMockRecorder) ListContainers(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListContainers", reflect.TypeOf((*MockDockerClient)(nil).ListContainers), arg0, arg1)
}

// LoadImage mocks base method
func (m *MockDockerClient) LoadImage(arg0 io.Reader, arg1 time.Duration) error {
	ret := m.ctrl.Call(m, "LoadImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoadImage indicates an expected call of LoadImage
func (mr *MockDockerClientMockRecorder) LoadImage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoadImage", reflect.TypeOf((*MockDockerClient)(nil).LoadImage), arg0, arg1)
}

// PullImage mocks base method
func (m *MockDockerClient) PullImage(arg0 string, arg1 *ecr.RegistryAuthenticationData) DockerContainerMetadata {
	ret := m.ctrl.Call(m, "PullImage", arg0, arg1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// PullImage indicates an expected call of PullImage
func (mr *MockDockerClientMockRecorder) PullImage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PullImage", reflect.TypeOf((*MockDockerClient)(nil).PullImage), arg0, arg1)
}

// RemoveContainer mocks base method
func (m *MockDockerClient) RemoveContainer(arg0 string, arg1 time.Duration) error {
	ret := m.ctrl.Call(m, "RemoveContainer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainer indicates an expected call of RemoveContainer
func (mr *MockDockerClientMockRecorder) RemoveContainer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainer", reflect.TypeOf((*MockDockerClient)(nil).RemoveContainer), arg0, arg1)
}

// RemoveImage mocks base method
func (m *MockDockerClient) RemoveImage(arg0 string, arg1 time.Duration) error {
	ret := m.ctrl.Call(m, "RemoveImage", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveImage indicates an expected call of RemoveImage
func (mr *MockDockerClientMockRecorder) RemoveImage(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveImage", reflect.TypeOf((*MockDockerClient)(nil).RemoveImage), arg0, arg1)
}

// StartContainer mocks base method
func (m *MockDockerClient) StartContainer(arg0 string, arg1 time.Duration) DockerContainerMetadata {
	ret := m.ctrl.Call(m, "StartContainer", arg0, arg1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// StartContainer indicates an expected call of StartContainer
func (mr *MockDockerClientMockRecorder) StartContainer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartContainer", reflect.TypeOf((*MockDockerClient)(nil).StartContainer), arg0, arg1)
}

// Stats mocks base method
func (m *MockDockerClient) Stats(arg0 string, arg1 context.Context) (<-chan *go_dockerclient.Stats, error) {
	ret := m.ctrl.Call(m, "Stats", arg0, arg1)
	ret0, _ := ret[0].(<-chan *go_dockerclient.Stats)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Stats indicates an expected call of Stats
func (mr *MockDockerClientMockRecorder) Stats(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockDockerClient)(nil).Stats), arg0, arg1)
}

// StopContainer mocks base method
func (m *MockDockerClient) StopContainer(arg0 string, arg1 time.Duration) DockerContainerMetadata {
	ret := m.ctrl.Call(m, "StopContainer", arg0, arg1)
	ret0, _ := ret[0].(DockerContainerMetadata)
	return ret0
}

// StopContainer indicates an expected call of StopContainer
func (mr *MockDockerClientMockRecorder) StopContainer(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StopContainer", reflect.TypeOf((*MockDockerClient)(nil).StopContainer), arg0, arg1)
}

// SupportedVersions mocks base method
func (m *MockDockerClient) SupportedVersions() []dockerclient.DockerVersion {
	ret := m.ctrl.Call(m, "SupportedVersions")
	ret0, _ := ret[0].([]dockerclient.DockerVersion)
	return ret0
}

// SupportedVersions indicates an expected call of SupportedVersions
func (mr *MockDockerClientMockRecorder) SupportedVersions() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SupportedVersions", reflect.TypeOf((*MockDockerClient)(nil).SupportedVersions))
}

// Version mocks base method
func (m *MockDockerClient) Version() (string, error) {
	ret := m.ctrl.Call(m, "Version")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Version indicates an expected call of Version
func (mr *MockDockerClientMockRecorder) Version() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Version", reflect.TypeOf((*MockDockerClient)(nil).Version))
}

// WithVersion mocks base method
func (m *MockDockerClient) WithVersion(arg0 dockerclient.DockerVersion) DockerClient {
	ret := m.ctrl.Call(m, "WithVersion", arg0)
	ret0, _ := ret[0].(DockerClient)
	return ret0
}

// WithVersion indicates an expected call of WithVersion
func (mr *MockDockerClientMockRecorder) WithVersion(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithVersion", reflect.TypeOf((*MockDockerClient)(nil).WithVersion), arg0)
}

// MockImageManager is a mock of ImageManager interface
type MockImageManager struct {
	ctrl     *gomock.Controller
	recorder *MockImageManagerMockRecorder
}

// MockImageManagerMockRecorder is the mock recorder for MockImageManager
type MockImageManagerMockRecorder struct {
	mock *MockImageManager
}

// NewMockImageManager creates a new mock instance
func NewMockImageManager(ctrl *gomock.Controller) *MockImageManager {
	mock := &MockImageManager{ctrl: ctrl}
	mock.recorder = &MockImageManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageManager) EXPECT() *MockImageManagerMockRecorder {
	return m.recorder
}

// AddAllImageStates mocks base method
func (m *MockImageManager) AddAllImageStates(arg0 []*image.ImageState) {
	m.ctrl.Call(m, "AddAllImageStates", arg0)
}

// AddAllImageStates indicates an expected call of AddAllImageStates
func (mr *MockImageManagerMockRecorder) AddAllImageStates(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddAllImageStates", reflect.TypeOf((*MockImageManager)(nil).AddAllImageStates), arg0)
}

// GetImageStateFromImageName mocks base method
func (m *MockImageManager) GetImageStateFromImageName(arg0 string) *image.ImageState {
	ret := m.ctrl.Call(m, "GetImageStateFromImageName", arg0)
	ret0, _ := ret[0].(*image.ImageState)
	return ret0
}

// GetImageStateFromImageName indicates an expected call of GetImageStateFromImageName
func (mr *MockImageManagerMockRecorder) GetImageStateFromImageName(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImageStateFromImageName", reflect.TypeOf((*MockImageManager)(nil).GetImageStateFromImageName), arg0)
}

// RecordContainerReference mocks base method
func (m *MockImageManager) RecordContainerReference(arg0 *api.Container) error {
	ret := m.ctrl.Call(m, "RecordContainerReference", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RecordContainerReference indicates an expected call of RecordContainerReference
func (mr *MockImageManagerMockRecorder) RecordContainerReference(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RecordContainerReference", reflect.TypeOf((*MockImageManager)(nil).RecordContainerReference), arg0)
}

// RemoveContainerReferenceFromImageState mocks base method
func (m *MockImageManager) RemoveContainerReferenceFromImageState(arg0 *api.Container) error {
	ret := m.ctrl.Call(m, "RemoveContainerReferenceFromImageState", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveContainerReferenceFromImageState indicates an expected call of RemoveContainerReferenceFromImageState
func (mr *MockImageManagerMockRecorder) RemoveContainerReferenceFromImageState(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveContainerReferenceFromImageState", reflect.TypeOf((*MockImageManager)(nil).RemoveContainerReferenceFromImageState), arg0)
}

// SetSaver mocks base method
func (m *MockImageManager) SetSaver(arg0 statemanager.Saver) {
	m.ctrl.Call(m, "SetSaver", arg0)
}

// SetSaver indicates an expected call of SetSaver
func (mr *MockImageManagerMockRecorder) SetSaver(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSaver", reflect.TypeOf((*MockImageManager)(nil).SetSaver), arg0)
}

// StartImageCleanupProcess mocks base method
func (m *MockImageManager) StartImageCleanupProcess(arg0 context.Context) {
	m.ctrl.Call(m, "StartImageCleanupProcess", arg0)
}

// StartImageCleanupProcess indicates an expected call of StartImageCleanupProcess
func (mr *MockImageManagerMockRecorder) StartImageCleanupProcess(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartImageCleanupProcess", reflect.TypeOf((*MockImageManager)(nil).StartImageCleanupProcess), arg0)
}
