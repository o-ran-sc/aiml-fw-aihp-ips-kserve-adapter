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

	"github.com/gin-gonic/gin"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/utils"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

type Command interface {
	Deploy(c *gin.Context)
}

type Executor struct {
	Command
}

func (Executor) Deploy(c *gin.Context) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	name := c.Query("name")
	if name == "" {
		utils.WriteError(c.Writer, errors.InvalidIPSName{Message: "Empty Query"})
		return
	}

	version := c.Query("version")
	if version == "" {
		utils.WriteError(c.Writer, errors.InvalidIPSName{Message: "Empty Query"})
		return
	}

	// kserve client deploy logic

	utils.WriteSuccess(c.Writer, http.StatusCreated, nil)
}
