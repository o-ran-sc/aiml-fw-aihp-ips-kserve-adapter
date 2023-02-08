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

package types

type Values struct {
	APIVersion            string `yaml:"api_version"`
	Engine                string `yaml:"engine"`
	Name                  string `yaml:"name"`
	FullName              string `yaml:"fullname"`
	ImagePullPolicy       string `yaml:"image_pull_policy"`
	MAXReplicas           string `yaml:"max_replicas"`
	MINReplicas           string `yaml:"min_replicas"`
	Resources             string `yaml:"resources"`
	RICServiceAccountName string `yaml:"ric_serviceaccount_name"`
	StorageURI            string `yaml:"storageUri"`
	Image                 string `yaml:"image"`
	ResourceVersion       string
	CanaryTrafficPercent  int64
}
