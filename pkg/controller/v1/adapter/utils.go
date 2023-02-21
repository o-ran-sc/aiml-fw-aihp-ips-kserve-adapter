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
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v1"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/types"
)

func valueParse(path string) (values types.Values, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	filePath, err := getValuesFilePath(path)
	if err != nil {
		return
	}

	b, err := os.ReadFile(filePath)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		err = errors.IOError{Message: err.Error()}
		return
	}

	err = yaml.Unmarshal(b, &values)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		err = errors.IOError{Message: err.Error()}
		return
	}
	values.CanaryTrafficPercent = -1

	return
}

func getValuesFilePath(root string) (path string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	files, err := os.ReadDir(root)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return "", errors.IOError{Message: err.Error()}
	}

	for _, file := range files {
		if file.IsDir() {
			path = filepath.Join(root, file.Name())
		}
	}
	path = filepath.Join(path, "values.yaml")
	return
}

func setCanaryTrafficRatio(values *types.Values, ratio string) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	trafficRatio, err := strconv.Atoi(ratio)
	if err != nil {
		trafficRatio = -1
	}
	values.CanaryTrafficPercent = int64(trafficRatio)
}
