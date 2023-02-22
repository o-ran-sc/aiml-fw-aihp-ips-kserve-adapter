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

package revision

import (
	"errors"
	"net/http"
	"net/http/httptest"
	. "net/url"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/url"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/types"

	controllermock "gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/controller/v1/adapter/mock"
)

const (
	testName = "test_name"
)

func TestReceivedRevisionRequest_ExpectCalledRevisionController(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockobj := controllermock.NewMockCommand(ctrl)
	gomock.InOrder(
		mockobj.EXPECT().Revision(testName),
	)

	ipsAdapter = mockobj

	w := httptest.NewRecorder()

	param := Values{}
	param.Set("name", testName)

	r, err := http.NewRequest("GET", url.V1()+url.IPS()+url.Revision(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")
	r.URL.RawQuery = param.Encode()

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	revision := Executor{}
	revision.Get(c)
}

func TestNegativeReceivedRevisionRequestWithoutParam_ExpectRevisionControllerReturnError(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockobj := controllermock.NewMockCommand(ctrl)
	gomock.InOrder(
		mockobj.EXPECT().Revision("").Return(types.Revision{}, errors.New("test error")),
	)

	ipsAdapter = mockobj

	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", url.V1()+url.IPS()+url.Revision(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	revision := Executor{}
	revision.Get(c)
}
