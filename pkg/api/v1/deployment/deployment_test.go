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

package deployment

import (
	"net/http"
	"net/http/httptest"
	. "net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/url"

	controllermock "gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/controller/v1/adapter/mock"
)

const (
	testName    = "test_name"
	testVesrion = "test_version"
)

func TestReceivedDeploymentRequest_ExpectCalledDeployController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockobj := controllermock.NewMockCommand(ctrl)
	gomock.InOrder(
		mockobj.EXPECT().Deploy(testName, testVesrion),
	)

	ipsAdapter = mockobj

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("name", testName)
	param.Set("version", testVesrion)

	r, err := http.NewRequest("POST", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Deploy(c)
}

func TestNegativeReceivedDeploymentRequestWithoutVersionQuery_ExpectErrorReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("name", testName)

	r, err := http.NewRequest("POST", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Deploy(c)

	if w.Code == 201 {
		t.Errorf("Unexpected Response Code : %d", w.Code)
	}
}

func TestNegativeReceivedDeploymentRequestWithoutNameQuery_ExpectErrorREturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("version", testVesrion)

	r, err := http.NewRequest("POST", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Deploy(c)

	if w.Code == http.StatusCreated {
		t.Errorf("Unexpected Response Code : %d", w.Code)
	}
}

func TestReceivedDeleteRequest_ExpectCalledDeleteController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockobj := controllermock.NewMockCommand(ctrl)
	gomock.InOrder(
		mockobj.EXPECT().Delete(testName),
	)

	ipsAdapter = mockobj

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("name", testName)

	r, err := http.NewRequest("DELETE", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Delete(c)
}

func TestNegativeReceivedDeleteRequestWithoutNameQuery_ExpectErrorReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()

	param := Values{}

	r, err := http.NewRequest("DELETE", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Delete(c)

	if w.Code == 201 {
		t.Errorf("Unexpected Response Code : %d", w.Code)
	}
}

func TestReceivedUpdateRequest_ExpectCalledUpdateController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockobj := controllermock.NewMockCommand(ctrl)
	gomock.InOrder(
		mockobj.EXPECT().Update(testName, testVesrion, gomock.Any()),
	)

	ipsAdapter = mockobj

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("name", testName)
	param.Set("version", testVesrion)

	r, err := http.NewRequest("PUT", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Update(c)
}

func TestNegativeReceivedUpdateRequestWithoutVersionQuery_ExpectErrorReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("name", testName)

	r, err := http.NewRequest("PUT", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Update(c)

	if w.Code == 201 {
		t.Errorf("Unexpected Response Code : %d", w.Code)
	}
}

func TestNegativeReceivedUpdateRequestWithoutNameQuery_ExpectErrorReturn(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("version", testVesrion)

	r, err := http.NewRequest("PUT", url.V1()+url.IPS(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	deployment := Executor{}
	deployment.Update(c)

	if w.Code == 201 {
		t.Errorf("Unexpected Response Code : %d", w.Code)
	}
}
