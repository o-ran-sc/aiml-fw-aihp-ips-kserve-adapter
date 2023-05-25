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

type Schema struct {
	Definitions Definitions `json:"definitions"`
	Schema      string      `json:"$schema"`
	Id          string      `json:"$id"`
	Type        string      `json:"type"`
	Title       string      `json:"title"`
	Required    []string    `json:"required"`
	Properties  Properties  `json:"properties"`
	schemaFile string
}

type Definitions struct {
}

type Properties struct {
	InferenceService InferenceServiceSchema `json:"inferenceservice"`
}

type InferenceServiceSchema struct {
	Id         string                   `json:"$id"`
	Type       string                   `json:"type"`
	Title      string                   `json:"title"`
	Required   []string                 `json:"required"`
	Properties InferenceServiceProperty `json:"properties"`
}

type InferenceServiceProperty struct {
	Engine      StringProperty  `json:"engine"`
	StorageURI  StringProperty  `json:"storage_uri"`
	ApiVersion  StringProperty  `json:"api_version"`
	MinReplicas IntegerProperty `json:"min_replicas"`
	MaxReplicas IntegerProperty `json:"max_replicas"`
}

type StringProperty struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  string   `json:"default"`
	Examples []string `json:"examples"`
}

type IntegerProperty struct {
	Id       string  `json:"$id"`
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Default  int32   `json:"default"`
	Examples []int32 `json:"examples"`
}
