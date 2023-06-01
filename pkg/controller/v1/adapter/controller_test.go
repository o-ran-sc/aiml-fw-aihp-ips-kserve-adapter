/*
==================================================================================

Copyright (c) 2023 Samsung Electronics Co., Ltd. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
==================================================================================
*/

package adapter

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kserve/kserve/pkg/apis/serving/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	kservemock "gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/client/kserve/mock"
	ricdmsmock "gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/client/ricdms/mock"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
)

const (
	testName        = "test_name"
	testVersion     = "test_version"
	samplePath      = "/../../../../sample/validServing"
	invalidPath     = "/../../../../invalid"
	invalidYAMLPath = "/../../../../sample/invalidServing"
)

var (
	sampleIFSV = v1beta1.InferenceService{
		TypeMeta:   v1.TypeMeta{},
		ObjectMeta: v1.ObjectMeta{},
		Spec: v1beta1.InferenceServiceSpec{
			Predictor: v1beta1.PredictorSpec{
				ComponentExtensionSpec: v1beta1.ComponentExtensionSpec{},
				Tensorflow: &v1beta1.TFServingSpec{
					PredictorExtensionSpec: v1beta1.PredictorExtensionSpec{},
				},
			},
		},
		Status: v1beta1.InferenceServiceStatus{},
	}
)

func fakeRemoveFunc(path string) (err error) {
	return nil
}

func TestNegativeCalledDepolyWithInvalidYAMLPath_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	workspace, _ := os.Getwd()

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return(workspace+invalidYAMLPath, nil),
	)

	// pass mockObj to a real object.
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	_, err := exec.Deploy(testName, testVersion)
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}

func TestNegativeCalledDepolyWithInvalidPath_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	workspace, _ := os.Getwd()

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return(workspace+invalidPath, nil),
	)

	// pass mockObj to a real object.
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	_, err := exec.Deploy(testName, testVersion)
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}

func TestCalledDepoly_ExpectReturnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kserveMockobj := kservemock.NewMockCommand(ctrl)
	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	workspace, _ := os.Getwd()

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return(workspace+samplePath, nil),
		kserveMockobj.EXPECT().Create(gomock.Any()),
	)

	// pass mockObj to a real object.
	kserveClient = kserveMockobj
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	exec.Deploy(testName, testVersion)
}

func TestNegativeCalledDepoly_WhenOnboardClientReturnError_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return("", errors.InternalServerError{}),
	)

	// pass mockObj to a real object.
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	_, err := exec.Deploy(testName, testVersion)
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}

func TestNegativeCalledDepoly_WhenKServeClientReturnError_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kserveMockobj := kservemock.NewMockCommand(ctrl)
	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	workspace, _ := os.Getwd()

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return(workspace+samplePath, nil),
		kserveMockobj.EXPECT().Create(gomock.Any()).Return("", errors.InternalServerError{}),
	)

	// pass mockObj to a real object.
	kserveClient = kserveMockobj
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	_, err := exec.Deploy(testName, testVersion)
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}

func TestCalledDelete_ExpectReturnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kserveMockobj := kservemock.NewMockCommand(ctrl)
	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChart(testName).Return(nil),
		kserveMockobj.EXPECT().Delete(testName).Return(nil),
	)

	// pass mockObj to a real object.
	kserveClient = kserveMockobj
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	err := exec.Delete(testName)
	if err != nil {
		t.Errorf("Expect error is nil, but error return, %s", err.Error())
	}
}

func TestNegativeCalledDelete_WhenOnboardClientReturnError_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChart(testName).Return(errors.IOError{}),
	)

	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	err := exec.Delete(testName)
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}

func TestNegativeCalledDelete_WhenKServeClientReturnError_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kserveMockobj := kservemock.NewMockCommand(ctrl)
	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChart(testName).Return(nil),
		kserveMockobj.EXPECT().Delete(testName).Return(errors.InternalServerError{}),
	)

	kserveClient = kserveMockobj
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	err := exec.Delete(testName)
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}

func TestCalledUpdate_ExpectReturnSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kserveMockobj := kservemock.NewMockCommand(ctrl)
	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	workspace, _ := os.Getwd()

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return(workspace+samplePath, nil),
		kserveMockobj.EXPECT().Get(testName).Return(&sampleIFSV, nil),
		kserveMockobj.EXPECT().Update(gomock.Any()),
	)

	// pass mockObj to a real object.
	kserveClient = kserveMockobj
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	exec.Update(testName, testVersion, "0")
}

func TestNegativeCalledUpdate_WhenOnboardClientReturnError_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return("", errors.InternalServerError{}),
	)

	// pass mockObj to a real object.
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	_, err := exec.Update(testName, testVersion, "0")
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}

func TestNegativeCalledUpdate_WhenKServeClientReturnError_ExpectReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	kserveMockobj := kservemock.NewMockCommand(ctrl)
	ricdmsmockobj := ricdmsmock.NewMockCommand(ctrl)

	workspace, _ := os.Getwd()

	gomock.InOrder(
		ricdmsmockobj.EXPECT().FetchHelmChartAndUntar(testName, testVersion).Return(workspace+samplePath, nil),
		kserveMockobj.EXPECT().Get(testName).Return(&sampleIFSV, nil),
		kserveMockobj.EXPECT().Update(gomock.Any()).Return("", errors.InternalServerError{}),
	)

	// pass mockObj to a real object.
	kserveClient = kserveMockobj
	ricdmsClient = ricdmsmockobj
	removeFunc = fakeRemoveFunc

	exec := Executor{}

	_, err := exec.Update(testName, testVersion, "0")
	if err == nil {
		t.Errorf("Expect error return, but error is nil")
	}
}
