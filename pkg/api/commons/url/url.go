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

package url

// V1 returns the v1 url as a type of string.
func V1() string { return "/v1" }

// IPS returns the ips url as a type of string.
func IPS() string { return "/ips" }

// Healthcheck returns the healthcheck url as a type of string.
func Healthcheck() string { return "/healthcheck" }

// Revision returns the revision url as a type of string.
func Revision() string { return "/revision" }

// Status returns the status url as a type of string.
func Status() string { return "/status" }

// Info returns the status url as a type of string.
func Info() string { return "/info" }

// IPSName returns the name url as a type of string.
func IPSName() string { return "/{name}" }

// Version returns the version url as a type of string.
func Version() string { return "/ver" }

// Preparation returns the package url as s type of string.
func Preparation() string { return "/preparation" }

// Download returns the chart download url as a type of string.
func Download() string { return "/api/v1/charts/xApp/download/{xApp_name}/ver/{version}" }

// Onboard returns the chart Onboard url as a type of string.
func Onboard() string { return "/api/v1/custom-onboard" }
