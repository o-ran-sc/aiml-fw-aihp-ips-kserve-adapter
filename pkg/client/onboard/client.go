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

package onboard

import (
	"os"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

type Command interface {
	Get(name string) error
	Download(name string, version string) (string, error)
}

type Executor struct {
	Command
}

var (
	onboarderIP   string
	onboarderPort string
)

func init() {
	onboarderIP = os.Getenv("ONBOARDER_IP")
	onboarderPort = os.Getenv("ONBOARDER_PORT")
}

func (Executor) Get(name string) (err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL := buildURL(name, "")

	_, err = requestToOnboard(targetURL)
	if err != nil {
		return
	}
	return
}

func (Executor) Download(name string, version string) (path string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL := buildURL(name, version)

	path, err = fetchHelmPackageFromOnboard(targetURL)
	if err != nil {
		return
	}
	return
}
