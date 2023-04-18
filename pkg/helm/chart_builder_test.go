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
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseConfigFile(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")
	config, err := chartBuilder.parseConfigFile("data/sample_config.json")

	assert.Nil(t, err)
	assert.NotEmpty(t, config)
}

func TestParseSchemaFile(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")
	schema, err := chartBuilder.parseSchemaFile("data/sample_schema.json")

	assert.Nil(t, err)
	assert.NotEmpty(t, schema)
}

func TestHelmLint(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")
	err := chartBuilder.helmLint()

	assert.Nil(t, err)
}

func TestAppendConfigToValuesYaml(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")

	err := chartBuilder.appendConfigToValuesYaml()

	assert.Nil(t, err)
}

func TestChangeChartNameVersion(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")
	err := chartBuilder.changeChartNameVersion()

	assert.Nil(t, err)
}

func TestPackageChart(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")
	err := chartBuilder.PackageChart()

	assert.Nil(t, err)
}

func TestCopyFile(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")
	SRC_FILE := "data/sample_config.json"
	DEST_FILE := "data/sample_config_copy.json"
	err := chartBuilder.copyFile(SRC_FILE, DEST_FILE)
	defer os.RemoveAll(DEST_FILE)

	assert.Nil(t, err)

	srcBytes, _ := ioutil.ReadFile(SRC_FILE)
	destBytes, _ := ioutil.ReadFile(DEST_FILE)

	assert.Equal(t, srcBytes, destBytes)
}

func TestCopyDirectory(t *testing.T) {
	chartBuilder := NewChartBuilder("data/sample_config.json", "data/sample_schema.json")

	SRC_DIR := "data/resources/std"
	DEST_DIR := "data/resources/std-copy"

	err := chartBuilder.copyDirectory(SRC_DIR, DEST_DIR)
	defer os.RemoveAll(DEST_DIR)
	assert.Nil(t, err)

	srcFiles, _ := ioutil.ReadDir(SRC_DIR)
	destFiles, _ := ioutil.ReadDir(DEST_DIR)

	assert.Equal(t, len(srcFiles), len(destFiles))

	for i, srcFile := range srcFiles {
		for j, destFile := range destFiles {
			if i == j {
				srcFileInfo, _ := ioutil.ReadFile(srcFile.Name())
				destFileInfo, _ := ioutil.ReadFile(destFile.Name())
				assert.Equal(t, srcFileInfo, destFileInfo)
			}
		}
	}
}
