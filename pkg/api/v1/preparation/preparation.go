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
	"net/http"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/utils"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/controller/v1/adapter"
	"github.com/gin-gonic/gin"
)

type Command interface {
	Preperation(c *gin.Context)
}

type Executor struct {
	Command
}

var ipsAdapter adapter.Command

func init() {
	ipsAdapter = adapter.Executor{}
}

func (Executor) Preperation(c *gin.Context) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	configFile := c.Query("configfile")
	if configFile == "" {
		utils.WriteError(c.Writer, errors.InvalidConfigFile{Message: "Empty Query"})
		return
	}

	schemaFile := c.Query("schemafile")
	if schemaFile == "" {
		utils.WriteError(c.Writer, errors.InvalidSchemaFile{Message: "Empty Query"})
		return
	}

	chartPath, err := ipsAdapter.Preperation(configFile, schemaFile)
	if err != nil {
		utils.WriteError(c.Writer, err)
		return
	}
	utils.WriteSuccess(c.Writer, http.StatusCreated, []byte(chartPath))
}
