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

type RevisionItem struct {
	LatestReadyRevision       string `json:"latestReadyRevision,omitempty" example:"{revision}" format:"string"`
	LatestCreatedRevision     string `json:"latestCreatedRevision,omitempty" example:"{revision}" format:"string"`
	PreviousRolledoutRevision string `json:"previousRolledoutRevision,omitempty" example:"{revision}" format:"string"`
	LatestRolledoutRevision   string `json:"latestRolledoutRevision,omitempty" example:"{revision}" format:"string"`
}

type StatusItem struct {
	Type               string `json:"type" example:"{status type}" format:"string"`
	Status             string `json:"status" example:"{True / False}" format:"string"`
	Severity           string `json:"severity,omitempty" example:"{severity}" format:"string"`
	LastTransitionTime string `json:"lastTransitionTime,omitempty" example:"{time stamp}" format:"string"`
	Reason             string `json:"reason,omitempty" example:"{reason}" format:"string"`
	Message            string `json:"message,omitempty" example:"{message}" format:"string"`
}

type Traffic struct {
	Tag            string `json:"tag,omitempty" example:"prev" format:"string"`
	RevisionName   string `json:"revisionName" example:"{revision}" format:"string"`
	LatestRevision string `json:"latestRevision,omitempty" example:"{True / False}" format:"string"`
	Percent        string `json:"percent,omitempty" example:"90" format:"string"`
	URL            string `json:"url,omitempty" example:"url" format:"string"`
}

type Resources struct {
	Limits   Limits   `json:"limits"`
	Requests Requests `json:"requests"`
}

type Limits struct {
	CPU    string `json:"cpu" example:"{1}" format:"string"`
	Memory string `json:"memory" example:"{2Gi}" format:"string"`
}

type Requests struct {
	CPU    string `json:"cpu" example:"{1}" format:"string"`
	Memory string `json:"memory" example:"{2Gi}" format:"string"`
}

type InfoItem struct {
	Traffic   []Traffic `json:"traffic"`
	Resources Resources `json:"resources"`
}
type Revision struct {
	Revision map[string]RevisionItem `json:"revision"`
}

type Status struct {
	Status map[string][]StatusItem `json:"status"`
}

type Info struct {
	Info map[string]InfoItem `json:"info"`
}
