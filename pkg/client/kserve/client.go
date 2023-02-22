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
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/tools/clientcmd"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/types"
)

const (
	ips_namespace = "ricips"
)

var ifsvGetter func(string) (client_v1beta1.InferenceServiceInterface, error)

//go:generate mockgen -source=client.go -destination=./mock/mock_client.go -package=mock
type Command interface {
	Init(kubeconfigPath string) error
	Create(values types.Values) (string, error)
	Delete(name string) error
	Get(name string) (*api_v1beta1.InferenceService, error)
	Update(values types.Values) (string, error)
	Revision(name string) (revisionList types.Revision, err error)
	Status(name string) (statusList types.Status, err error)
	Info(name string) (infoList types.Info, err error)
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

func (c *Client) Create(values types.Values) (revision string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	info := convertValuesToInferenceService(values)
	if err != nil {
		return
	}

	_, err = c.api.Create(&info)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		err = errors.InternalServerError{Message: err.Error()}
		return
	}
	return
}

func (c *Client) Delete(name string) (err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	err = c.api.Delete(name, &v1.DeleteOptions{})
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		err = errors.NotFoundIPS{
			Message: err.Error(),
		}
		return
	}
	return
}

func (c *Client) Get(name string) (ifsv *api_v1beta1.InferenceService, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	ifsv, err = c.api.Get(name, v1.GetOptions{})
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		err = errors.NotFoundIPS{Message: err.Error()}
		return
	}
	return
}

func (c *Client) Update(values types.Values) (revision string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	info := convertValuesToInferenceService(values)
	if err != nil {
		return
	}

	_, err = c.api.Update(&info)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		err = errors.NotFoundIPS{Message: err.Error()}
	}

	return
}

func (c *Client) Revision(name string) (revisionList types.Revision, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	revisionList.Revision = make(map[string]types.RevisionItem)

	if name == "" {
		ifsvList, e := c.api.List(v1.ListOptions{})
		if e != nil {
			err = errors.InternalServerError{Message: e.Error()}
			return
		}

		for _, ifsv := range ifsvList.Items {
			revision, e := makeRevision(ifsv)
			if e != nil {
				err = e
				return
			}
			revisionList.Revision[ifsv.Name] = revision
		}

	} else {
		ifsv, e := c.Get(name)
		if e != nil {
			err = e
			return
		}

		revision, e := makeRevision(*ifsv)
		if e != nil {
			err = e
			return
		}
		revisionList.Revision[ifsv.Name] = revision
	}

	return
}

func (c *Client) Status(name string) (statusList types.Status, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	statusList.Status = make(map[string][]types.StatusItem)

	if name == "" {
		ifsvList, e := c.api.List(v1.ListOptions{})
		if e != nil {
			err = errors.InternalServerError{Message: e.Error()}
			return
		}

		for _, ifsv := range ifsvList.Items {
			status, e := makeStatus(ifsv)
			if e != nil {
				err = errors.InternalServerError{Message: e.Error()}
				return
			}
			statusList.Status[ifsv.Name] = status
		}

	} else {
		ifsv, e := c.Get(name)
		if e != nil {
			err = e
			return
		}

		status, e := makeStatus(*ifsv)
		if e != nil {
			err = e
			return
		}
		statusList.Status[ifsv.Name] = status
	}

	return
}

func (c *Client) Info(name string) (infoList types.Info, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	infoList.Info = make(map[string]types.InfoItem)

	if name == "" {
		ifsvList, e := c.api.List(v1.ListOptions{})
		if e != nil {
			logger.Logging(logger.ERROR, e.Error())
			err = errors.InternalServerError{
				Message: e.Error(),
			}
			return
		}

		for _, ifsv := range ifsvList.Items {
			info, e := makeInfo(ifsv)
			if e != nil {
				err = e
				return
			}
			infoList.Info[ifsv.Name] = info
		}

	} else {
		ifsv, e := c.Get(name)
		if e != nil {
			err = e
			return
		}

		info, e := makeInfo(*ifsv)
		if e != nil {
			err = e
			return
		}
		infoList.Info[ifsv.Name] = info
	}

	return
}
