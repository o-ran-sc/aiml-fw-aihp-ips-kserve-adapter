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

package kserve

import (
	api_v1beta1 "github.com/kserve/kserve/pkg/apis/serving/v1beta1"
	"github.com/kserve/kserve/pkg/client/clientset/versioned"
	client_v1beta1 "github.com/kserve/kserve/pkg/client/clientset/versioned/typed/serving/v1beta1"
	"k8s.io/client-go/tools/clientcmd"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

const (
	ips_namespace = "ricips"
)

var ifsvGetter func(string) (client_v1beta1.InferenceServiceInterface, error)

type Command interface {
	Init(kubeconfigPath string) error
}

type Client struct {
	Command

	api client_v1beta1.InferenceServiceInterface
}

type InferenceServiceInfo struct {
	api_v1beta1.InferenceServiceSpec
}

func init() {
	ifsvGetter = inferenceServiceGetter
}

func inferenceServiceGetter(path string) (api client_v1beta1.InferenceServiceInterface, err error) {
	config, err := clientcmd.BuildConfigFromFlags("", path)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}

	client, err := versioned.NewForConfig(config)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}

	api = client.ServingV1beta1().InferenceServices(ips_namespace)

	return
}

func (c *Client) Init(kubeconfigPath string) (err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	c.api, err = ifsvGetter(kubeconfigPath)
	if err != nil {
		err = errors.InternalServerError{Message: err.Error()}
		return
	}
	return
}
