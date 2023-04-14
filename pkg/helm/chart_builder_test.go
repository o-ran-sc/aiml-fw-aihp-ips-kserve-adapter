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

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigFile(t *testing.T) {
	chartBuilder := NewChartBuilder()
	config, err := chartBuilder.parseConfigFile("data/sample_config.json")

	assert.Nil(t, err)
	assert.NotEmpty(t, config)
}

func TestParseSchemaFile(t *testing.T) {
	chartBuilder := NewChartBuilder()
	schema, err := chartBuilder.parseSchemaFile("data/sample_schema.json")

	assert.Nil(t, err)
	assert.NotEmpty(t, schema)
}

func TestHelmLint(t *testing.T) {
	chartBuilder := NewChartBuilder()
	err := chartBuilder.helmLint()

	assert.Nil(t, err)
}

func TestAppendConfigToValuesYaml(t *testing.T) {
	chartBuilder := NewChartBuilder()
	err := chartBuilder.appendConfigToValuesYaml()

	assert.Nil(t, err)
}

func TestChangeChartNameVersion(t *testing.T) {
	chartBuilder := NewChartBuilder()
	err := chartBuilder.changeChartNameVersion()

	assert.Nil(t, err)
}

func TestPackageChart(t *testing.T) {
	chartBuilder := NewChartBuilder()
	err := chartBuilder.PackageChart()

	assert.Nil(t, err)
}
