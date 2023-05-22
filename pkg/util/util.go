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

package util

import (
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
	"time"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

func Exec(args string) (out []byte, err error) {
	cmd := exec.Command("/bin/sh", "-c", args)

	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	logger.Logging(logger.INFO, fmt.Sprintf("Running command: %s", cmd.Args))
	for i := 0; i < 3; i++ {
		if err = cmd.Run(); err != nil {
			logger.Logging(logger.ERROR, fmt.Sprintf("Command failed : %v - %s, retrying", err.Error(), stderr.String()))
			time.Sleep(time.Duration(2) * time.Second)
			continue
		}
		break
	}

	if err == nil {
		logger.Logging(logger.INFO, fmt.Sprintf("command success: %s", stdout.String()))
		return stdout.Bytes(), nil
	}

	return stdout.Bytes(), errors.New(stderr.String())
}

var HelmExec = func(args string) (out []byte, err error) {
	return Exec(strings.Join([]string{"helm", args}, " "))
}
