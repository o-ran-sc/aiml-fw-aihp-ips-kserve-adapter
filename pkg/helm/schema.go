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
}

type Definitions struct {
}

type Properties struct {
	ServicePort ServicePort `json:"service_ports"`
	Rmr         Rmr         `json:"rmr"`
	Envs        Envs        `json:"envs"`
}

type Envs struct {
	Id         string        `json:"$id"`
	Type       string        `json:"type"`
	Title      string        `json:"title"`
	Required   []string      `json:"required"`
	Properties EnvProperties `json:"properties"`
}

type EnvProperties struct {
	GNodeB                 EnvPropertyTemplate `json:"gNodeB"`
	Threads                EnvPropertyTemplate `json:"THREADS"`
	A1SchemaFile           EnvPropertyTemplate `json:"A1_SCHEMA_FILE"`
	VesSchemaFile          EnvPropertyTemplate `json:"VES_SCHEMA_FILE"`
	SampleFile             EnvPropertyTemplate `json:"SAMPLE_FILE"`
	VesCollectorURL        EnvPropertyTemplate `json:"VES_COLLECTOR_URL"`
	VesMeasurementInterval EnvPropertyTemplate `json:"VES_MEASUREMENT_INTERVAL"`
}

type EnvPropertyTemplate struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  string   `json:"default"`
	Examples []string `json:"examples"`
	Pattern  string   `json:"pattern"`
}

type Rmr struct {
	Id            string        `json:"$id"`
	Type          string        `json:"type"`
	Title         string        `json:"title"`
	Required      []string      `json:"required"`
	RmrProperties RmrProperties `json:"properties"`
}

type RmrProperties struct {
	ProtPort   ProtPort       `json:"protPort"`
	MaxSize    RmrSubProperty `json:"maxSize"`
	NumWorkers RmrSubProperty `json:"numWorkers"`
	TxMessages RmrMessage     `json:"txMessages"`
	RxMessages RmrMessage     `json:"rxMessages"`
	FilePath   FilePath       `json:"file_path"`
	Contents   Contents       `json:"contents"`
}

type FilePath struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  string   `json:"default"`
	Examples []string `json:"examples"`
	Pattern  string   `json:"pattern"`
}

type Contents struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  string   `json:"default"`
	Examples []string `json:"examples"`
	Pattern  string   `json:"pattern"`
}

type RmrMessage struct {
	Id    string          `json:"$id"`
	Type  string          `json:"type"`
	Title string          `json:"title"`
	Items RmrMessageItems `json:"items"`
}

type RmrMessageItems struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  string   `json:"default"`
	Examples []string `json:"examples"`
	Pattern  string   `json:"pattern"`
}

type RmrSubProperty struct {
	Id       string  `json:"$id"`
	Type     string  `json:"type"`
	Title    string  `json:"title"`
	Default  int32   `json:"default"`
	Examples []int32 `json:"examples"`
}

type ProtPort struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  string   `json:"default"`
	Examples []string `json:"examples"`
	Pattern  string   `json:"pattern"`
}

type ServicePort struct {
	Id                    string                `json:"$id"`
	Type                  string                `json:"type"`
	Title                 string                `json:"title"`
	Required              []string              `json:"required"`
	ServicePortProperties ServicePortProperties `json:"properties"`
}

type ServicePortProperties struct {
	XappPort XappPort `json:"xapp_port"`
	RmrPort  RmrPort  `json:"rmr_port"`
}
type RmrPort struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  uint32   `json:"default"`
	Examples []uint32 `json:"examples"`
}

type XappPort struct {
	Id       string   `json:"$id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Default  uint32   `json:"default"`
	Examples []uint32 `json:"examples"`
}
