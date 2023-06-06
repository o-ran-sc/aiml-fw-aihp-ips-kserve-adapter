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

package preparation

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/utils"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/helm"
)

type Command interface {
	Post(c *gin.Context)
}

type Executor struct {
	Command
}

func init() {
}

func (Executor) Post(c *gin.Context) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	configFile := c.Query("configfile")
	schemaFile := c.Query("schemafile")

	chartBuilder := helm.NewChartBuilder(configFile, schemaFile)
	chartPath, err := chartBuilder.PackageChart()

	if err != nil {
		utils.WriteError(c.Writer, errors.InternalServerError{Message: err.Error()})
		return
	}

	data, err := json.Marshal(chartPath)
	if err != nil {
		utils.WriteError(c.Writer, err)
		return
	}
	utils.WriteSuccess(c.Writer, http.StatusAccepted, data)
}