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

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/client/kserve"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/client/ricdms"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/types"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/helm"
)

//go:generate mockgen -source=controller.go -destination=./mock/mock_controller.go -package=mock
type Command interface {
	Deploy(name string, version string) (string, error)
	Delete(name string) error
	Update(name string, version string, canaryTrafficRatio string) (string, error)
	Revision(name string) (revisionList types.Revision, err error)
	Status(name string) (types.Status, error)
	Info(name string) (types.Info, error)
	Preperation(configFile string, schemaFile string) (chartPath string, err error)
}

type Executor struct {
	Command
}

var kserveClient kserve.Command
var ricdmsClient ricdms.Command

var removeFunc func(string) error

func init() {
	kserveClient = &kserve.Client{}

	kubeUrl := os.Getenv("KUBEURL")
	kubeconfigPath := os.Getenv("KUBECONFIG")
	err := kserveClient.Init(kubeUrl, kubeconfigPath)
	if err != nil {
		os.Exit(8)
	}

	ricdmsClient = ricdms.Executor{}

	removeFunc = func(path string) (err error) {
		err = os.RemoveAll(path)
		return
	}
}

func (Executor) Deploy(name string, version string) (revision string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	path, err := ricdmsClient.FetchHelmChartAndUntar(name, version)
	if err != nil {
		return
	}
	defer removeFunc(path)

	values, err := valueParse(path)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}

	revision, err = kserveClient.Create(values)
	if err != nil {
		return
	}
	return
}

func (Executor) Delete(name string) (err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	err = ricdmsClient.FetchHelmChart(name)
	if err != nil {
		err = errors.InvalidIPSName{
			Message: err.Error(),
		}
		return
	}

	err = kserveClient.Delete(name)
	if err != nil {
		return
	}
	return
}

func (Executor) Update(name string, version string, canaryTrafficRatio string) (revision string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	path, err := ricdmsClient.FetchHelmChartAndUntar(name, version)
	if err != nil {
		return
	}
	defer removeFunc(path)

	values, err := valueParse(path)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}

	setCanaryTrafficRatio(&values, canaryTrafficRatio)

	ifsv, err := kserveClient.Get(name)
	if err != nil {
		return
	}
	values.ResourceVersion = ifsv.ResourceVersion

	revision, err = kserveClient.Update(values)
	if err != nil {
		return
	}
	return
}

func (Executor) Revision(name string) (result types.Revision, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	result, err = kserveClient.Revision(name)
	return
}

func (Executor) Status(name string) (result types.Status, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	result, err = kserveClient.Status(name)
	return
}

func (Executor) Info(name string) (result types.Info, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	result, err = kserveClient.Info(name)
	return
}

func (Executor) Preperation(configFile string, schemaFile string) (chartPath string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	chartBuilder := helm.NewChartBuilder(configFile, schemaFile)
	chartPath, err = chartBuilder.PackageChart()
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}

	ricdmsClient.OnboardHelmChart(chartPath)
	if err != nil {
		return
	}

	return
}
