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
	"github.com/gin-gonic/gin"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/url"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/v1/deployment"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/v1/healthcheck"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/v1/info"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/v1/revision"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/v1/status"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

var (
	deploymentExecutor  deployment.Command
	healthcheckExecutor healthcheck.Command
	revisionExecutor    revision.Command
	statusExecutor      status.Command
	infoExecutor        info.Command
)

func init() {
	deploymentExecutor = deployment.Executor{}
	healthcheckExecutor = healthcheck.Executor{}
	revisionExecutor = revision.Executor{}
	statusExecutor = status.Executor{}
	infoExecutor = info.Executor{}
}

func setupRouter() (router *gin.Engine) {
	router = gin.Default()

	v1 := router.Group(url.V1())
	{
		deployment := v1.Group(url.IPS())
		{
			deployment.POST("", deploymentExecutor.Deploy)
			deployment.DELETE("", deploymentExecutor.Delete)
			deployment.PUT("", deploymentExecutor.Update)
		}

		healthcheck := v1.Group(url.Healthcheck())
		{
			healthcheck.GET("", healthcheckExecutor.Ping)
		}

		revision := v1.Group(url.IPS() + url.Revision())
		{
			revision.GET("", revisionExecutor.Get)
		}

		status := v1.Group(url.IPS() + url.Status())
		{
			status.GET("", statusExecutor.Get)
		}

		info := v1.Group(url.IPS() + url.Info())
		{
			info.GET("", infoExecutor.Get)
		}

		_, _, _, _ = healthcheck, revision, status, info
	}

	return
}

func RunWebServer(port string) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	r := setupRouter()

	r.Run(":" + port)
}
