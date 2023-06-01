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

package ricdms

import (
	"net/http"
	"os"
	"strconv"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

//go:generate mockgen -source=client.go -destination=./mock/mock_client.go -package=mock
type Command interface {
	OnboardHelmChart(chartPath string) error
	FetchHelmChart(name string) error
	FetchHelmChartAndUntar(name string, version string) (string, error)
}

type Executor struct {
	Command
}

var (
	ricdmsIP   string
	ricdmsPort string
)

func init() {
	ricdmsIP = os.Getenv("RIC_DMS_IP")
	ricdmsPort = os.Getenv("RIC_DMS_PORT")
}

func (Executor) OnboardHelmChart(chartPath string) error {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL := buildURLforChartOnboard()

	resp, err := onboardChartToRicdms(targetURL, chartPath)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.InternalServerError{
			Message: "ricdms return error, status code : " + strconv.Itoa(int(resp.StatusCode)),
		}
		logger.Logging(logger.ERROR, err.Error())
		return err
	}
	return err
}

func (Executor) FetchHelmChart(name string) (err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL, err := buildURLforChartFetch(name, "")
	if err != nil {
		return
	}

	_, err = getChartFromRicdms(targetURL)
	if err != nil {
		return
	}
	return
}

func (Executor) FetchHelmChartAndUntar(name string, version string) (path string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL, err := buildURLforChartFetch(name, version)
	if err != nil {
		return
	}

	resp, err := getChartFromRicdms(targetURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	path = generateRandPath(pathLength)

	err = Untar(path, resp.Body)

	if err != nil {
		return
	}
	return
}
