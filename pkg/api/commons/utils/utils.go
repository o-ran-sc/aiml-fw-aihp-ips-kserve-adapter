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

package utils

import (
	"encoding/json"
	"net/http"
)

func WriteSuccess(w http.ResponseWriter, code int, data []byte) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}

func WriteError(w http.ResponseWriter, err error) {
	data := make(map[string]interface{})
	data["message"] = err.Error()

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	w.Write(ChangeToJson(data))
}

func ChangeToJson(src map[string]interface{}) []byte {
	dst, err := json.Marshal(src)
	if err != nil {
		return nil
	}
	return dst
}
