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

package adapter

import (
	"os"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/client/kserve"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/client/onboard"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

type Command interface {
	Deploy(name string, version string) error
}

type Executor struct {
	Command
}

var kserveClient kserve.Command
var onboardClient onboard.Command

func init() {
	kserveClient = &kserve.Client{}
	onboardClient = onboard.Executor{}

	kubeconfigPath := os.Getenv("KUBECONFIG")
	err := kserveClient.Init(kubeconfigPath)
	if err != nil {
		os.Exit(8)
	}
}

func (Executor) Deploy(name string, version string) error {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	// TODO: Get object from onboard & Deploy using kserveClient

	return nil
}
