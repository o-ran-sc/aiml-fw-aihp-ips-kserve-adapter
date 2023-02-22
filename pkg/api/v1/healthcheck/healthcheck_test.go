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

package healthcheck

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/url"

	"github.com/gin-gonic/gin"
)

func TestReceivedPingRequest_ExpectSuccess(t *testing.T) {
	w := httptest.NewRecorder()

	r, err := http.NewRequest("GET", url.V1()+url.Healthcheck(), nil)
	if err != nil {
		t.Errorf("http.NewRequest return Error : %s", err.Error())
	}
	r.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = r

	status := Executor{}
	status.Ping(c)
}
