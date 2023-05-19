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
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/xeipuuv/gojsonschema"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

const (
	VALUES_YAML = "values.yaml"
	CHART_YAML  = "Chart.yaml"

	ENV_CHART_WORKSPACE_PATH = "CHART_WORKSPACE_PATH"

	SERVING_VERSION_ALPHA = "serving.kubeflow.org/v1alpha2"
	SERVING_VERSION_BETA  = "serving.kubeflow.org/v1beta1"
	RESOURCE_ALPHA        = "resources/inf-alpha"
	RESOURCE_BETA         = "resources/inf-beta"
)

type HelmChartBuilder interface {
	PackageChart() (err error)
}

type ChartBuilder struct {
	config             Config
	schema             Schema
	chartWorkspacePath string
	chartName          string
	chartVersion       string
}

func NewChartBuilder(configFile string, schemaFile string) *ChartBuilder {
	chartBuilder := &ChartBuilder{}
	var err error

	chartBuilder.config, err = chartBuilder.parseConfigFile(configFile)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return nil
	}

	if chartBuilder.config.XappName == "" || chartBuilder.config.Version == "" {
		logger.Logging(logger.ERROR, fmt.Sprintf("some value is empty : xAppName = %s, Version = %s", chartBuilder.config.XappName, chartBuilder.config.Version))
		return nil
	}

	chartBuilder.schema, err = chartBuilder.parseSchemaFile(schemaFile)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return nil
	}

	chartBuilder.chartName = chartBuilder.config.XappName
	chartBuilder.chartVersion = chartBuilder.config.Version
	chartBuilder.chartWorkspacePath = os.Getenv(ENV_CHART_WORKSPACE_PATH) + "/" + chartBuilder.chartName + "-" + chartBuilder.chartVersion

	_, err = os.Stat(chartBuilder.chartWorkspacePath)
	if err != nil {
		if !os.IsNotExist(err) {
			os.RemoveAll(chartBuilder.chartWorkspacePath)
		}
	}

	err = os.Mkdir(chartBuilder.chartWorkspacePath, os.FileMode(0744))
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return nil
	}

	resourceVersionMap := map[string]string{
		SERVING_VERSION_ALPHA: RESOURCE_ALPHA,
		SERVING_VERSION_BETA:  RESOURCE_BETA,
	}

	apiVersion := chartBuilder.config.InferenceService.ApiVersion
	resource := resourceVersionMap[apiVersion]
	err = chartBuilder.copyDirectory("data/"+resource, chartBuilder.chartWorkspacePath+"/"+chartBuilder.chartName)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return nil
	}

	return chartBuilder

}

func (c *ChartBuilder) copyFile(src string, dest string) (err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return
	}
	defer func() {
		if e := destFile.Close(); e != nil {
			err = e
		}
	}()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return
	}
	err = destFile.Sync()
	if err != nil {
		return
	}
	srcInfo, err := os.Stat(src)
	if err != nil {
		return
	}

	err = os.Chmod(dest, srcInfo.Mode())
	if err != nil {
		return
	}
	return
}

func (c *ChartBuilder) copyDirectory(src string, dest string) (err error) {
	src = filepath.Clean(src)
	dest = filepath.Clean(dest)

	srcInfo, err := os.Stat(src)
	if err != nil {
		return
	}
	if !srcInfo.IsDir() {
		return fmt.Errorf("src(%s) is not a directory", src)
	}

	_, err = os.Stat(dest)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("destination(%s) is already exists", dest)
	}

	err = os.Mkdir(dest, os.FileMode(0744))
	if err != nil && !os.IsExist(err) {
		return
	}
	entries, err := ioutil.ReadDir(src)
	if err != nil {
		return
	}

	for _, entry := range entries {
		srcPath := filepath.Join(src, entry.Name())
		destPath := filepath.Join(dest, entry.Name())

		if entry.IsDir() {
			err = c.copyDirectory(srcPath, destPath)
			if err != nil {
				return
			}
		} else {
			if entry.Mode()&os.ModeSymlink != 0 {
				continue
			}
			err = c.copyFile(srcPath, destPath)
			if err != nil {
				return
			}
		}
	}
	return
}

func (c *ChartBuilder) PackageChart() (err error) {
	return errors.New("not yet implemented")
}

func (c *ChartBuilder) parseConfigFile(configFile string) (config Config, err error) {
	data, err := os.Open(configFile)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}
	byteValue, _ := ioutil.ReadAll(data)

	err = json.Unmarshal(byteValue, &config)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}
	config.configFile = configFile
	return
}

func (c *ChartBuilder) parseSchemaFile(schemaFile string) (schema Schema, err error) {
	data, err := os.Open(schemaFile)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}
	byteValue, _ := ioutil.ReadAll(data)

	err = json.Unmarshal(byteValue, &schema)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return
	}

	schema.schemaFile = schemaFile
	return
}

func (c *ChartBuilder) helmLint() (err error) {
	return errors.New("not yet implemented")
}

func (c *ChartBuilder) appendConfigToValuesYaml() (err error) {
	return errors.New("not yet implemented")
}

func (c *ChartBuilder) changeChartNameVersion() (err error) {
	return errors.New("not yet implemented")
}

func (c *ChartBuilder) ValidateChartMaterials() (err error) {

	schemaLoader := gojsonschema.NewReferenceLoader("file://" + c.schema.schemaFile)
	configLoader := gojsonschema.NewReferenceLoader("file://" + c.config.configFile)

	result, err := gojsonschema.Validate(schemaLoader, configLoader)

	if err != nil {
		return
	}

	if !result.Valid() {
		for _, desc := range result.Errors() {
			logger.Logging(logger.ERROR, fmt.Sprintf("[Invalid Config] - %s\n", desc))
		}
		return errors.New("Invalid Config")
	}
	return
}
