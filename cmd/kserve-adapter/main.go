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

package main

import (
	"flag"
	"os"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
	"github.com/spf13/viper"
)

const KSERVE_ADAPTER_CONFIG_FILE = "./config/kserve-adapter.yaml"

var (
	apiServerPort string
)

func init() {
	apiServerPort = os.Getenv("API_SERVER_PORT")

	var fileName *string
	fileName = flag.String("f", KSERVE_ADAPTER_CONFIG_FILE, "Specify the configuration file.")
	flag.Parse()
	viper.SetConfigFile(*fileName)
}

func main() {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	if apiServerPort == "" {
		logger.Logging(logger.ERROR, "invalid port num")
		os.Exit(1)
	}

	api.RunWebServer(apiServerPort)
}
