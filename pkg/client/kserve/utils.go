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
	"encoding/json"
	"strconv"

	api_v1beta1 "github.com/kserve/kserve/pkg/apis/serving/v1beta1"
	core_v1 "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/types"
)

func convertValuesToInferenceService(values types.Values) (ifsv api_v1beta1.InferenceService) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	maxReplicas, err := strconv.Atoi(values.MAXReplicas)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		maxReplicas = 1
	}

	minReplicas, err := strconv.Atoi(values.MINReplicas)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		minReplicas = 1
	}

	ifsv = api_v1beta1.InferenceService{
		TypeMeta: v1.TypeMeta{
			Kind:       "InferenceService",
			APIVersion: "serving.kserve.io/v1beta1",
		},
		ObjectMeta: v1.ObjectMeta{
			Name:      values.FullName,
			Namespace: ips_namespace,
			Labels: map[string]string{
				"controller-tools.k8s.io": "1.0",
				"app":                     values.FullName,
			},
		},
		Spec: api_v1beta1.InferenceServiceSpec{
			Predictor: api_v1beta1.PredictorSpec{
				ComponentExtensionSpec: api_v1beta1.ComponentExtensionSpec{
					MaxReplicas: maxReplicas,
					MinReplicas: &minReplicas,
				},
				PodSpec: api_v1beta1.PodSpec{
					ServiceAccountName: values.RICServiceAccountName,
				},
			},
		},
		Status: api_v1beta1.InferenceServiceStatus{},
	}

	switch values.Engine {
	case "tensorflow":
		ifsv.Spec.Predictor.Tensorflow = &api_v1beta1.TFServingSpec{
			PredictorExtensionSpec: api_v1beta1.PredictorExtensionSpec{
				StorageURI: &values.StorageURI,
				Container: core_v1.Container{
					Image: values.Image,
					Ports: []core_v1.ContainerPort{
						{
							Name:          "h2c",
							ContainerPort: 9000,
							Protocol:      "TCP",
						},
					},
				},
			},
		}
	case "sklearn":
		ifsv.Spec.Predictor.SKLearn = &api_v1beta1.SKLearnSpec{
			PredictorExtensionSpec: api_v1beta1.PredictorExtensionSpec{
				StorageURI: &values.StorageURI,
				Container: core_v1.Container{
					Image: values.Image,
					Ports: []core_v1.ContainerPort{
						{
							Name:          "h2c",
							ContainerPort: 9000,
							Protocol:      "TCP",
						},
					},
				},
			},
		}
	}

	if values.CanaryTrafficPercent >= 0 {
		ifsv.Spec.Predictor.CanaryTrafficPercent = &values.CanaryTrafficPercent
	}

	if values.ResourceVersion != "" {
		ifsv.ResourceVersion = values.ResourceVersion
	}

	return
}

func makeRevision(ifsv api_v1beta1.InferenceService) (revision types.RevisionItem, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	revisionInfo, exist := ifsv.Status.Components[api_v1beta1.PredictorComponent]
	if !exist {
		err = errors.InternalServerError{
			Message: "PredictorComponent does not have revision information.",
		}
		return
	}

	data, _ := json.Marshal(revisionInfo)
	json.Unmarshal(data, &revision)
	return
}

func makeStatus(ifsv api_v1beta1.InferenceService) (status []types.StatusItem, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	if len(ifsv.Status.Conditions) == 0 {
		logger.Logging(logger.ERROR, "condition is empty")
		err = errors.InternalServerError{
			Message: "condition is empty",
		}
		return
	}

	for _, condition := range ifsv.Status.Conditions {
		status = append(status, types.StatusItem{
			Type:               string(condition.Type),
			Status:             string(condition.Status),
			Severity:           string(condition.Severity),
			Reason:             condition.Reason,
			Message:            condition.Message,
			LastTransitionTime: condition.LastTransitionTime.Inner.String(),
		})
	}
	return
}
