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

package helm

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

type ChartBuilder struct {
}

func NewChartBuilder() ChartBuilder {
	return ChartBuilder{}
}

func (c *ChartBuilder) parseConfigFile(configFile string) (config Config, err error) {
	data, err := os.Open(configFile)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}
	byteValue, _ := ioutil.ReadAll(data)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}
	return
}

func (c *ChartBuilder) parseSchemaFile(schemaFile string) (schema Schema, err error) {
	data, err := os.Open(schemaFile)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}
	byteValue, _ := ioutil.ReadAll(data)

	err = json.Unmarshal(byteValue, &schema)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}
	return
}
