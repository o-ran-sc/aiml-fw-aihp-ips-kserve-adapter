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

package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/url"
	deploymentmock "gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/v1/deployment/mock"

	"github.com/golang/mock/gomock"
)

func TestRecivedDeploymentRequest_ExpectCalledDeploymentHandle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockobj := deploymentmock.NewMockCommand(ctrl)
	gomock.InOrder(
		mockobj.EXPECT().Deploy(gomock.Any()),
	)

	// pass mockObj to a real object.
	deploymentExecutor = mockobj

	ts := httptest.NewServer(setupRouter())
	defer ts.Close()
	http.Post(ts.URL+url.V1()+url.IPS(),
		"application/json", nil)
}
