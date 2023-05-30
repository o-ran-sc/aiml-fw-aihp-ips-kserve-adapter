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

package errors

type Unknown struct {
	Message string `json:"message" example:"unknown error: {reason}" format:"string"`
}

func (e Unknown) Error() string {
	return "unknown error: " + e.Message
}

type NotFoundURL struct {
	Message string `json:"message" example:"unsupported url: {reason}" format:"string"`
}

func (e NotFoundURL) Error() string {
	return "unsupported url: " + e.Message
}

type InvalidMethod struct {
	Message string `json:"message" example:"invalid method: {reason}" format:"string"`
}

func (e InvalidMethod) Error() string {
	return "invalid method: " + e.Message
}

type InvalidIPSName struct {
	Message string `json:"message" example:"invalid ips name: {reason}" format:"string"`
}

func (e InvalidIPSName) Error() string {
	return "invalid IPS name: " + e.Message
}

type NotFoundIPS struct {
	Message string `json:"message" example:"not found ips: {reason}" format:"string"`
}

func (e NotFoundIPS) Error() string {
	return "not found ips: " + e.Message
}

type IOError struct {
	Message string `json:"message" example:"io error: {reason}" format:"string"`
}

func (e IOError) Error() string {
	return "io error: " + e.Message
}

type TimeoutError struct {
	Message string `json:"message" example:"time out error: {reason}" format:"string"`
}

func (e TimeoutError) Error() string {
	return "time out error: " + e.Message
}

type InternalServerError struct {
	Message string `json:"message" example:"internal server error: {reason}" format:"string"`
}

func (e InternalServerError) Error() string {
	return "internal server error: " + e.Message
}

type InvalidConfigFile struct {
	Message string `json:"message" example:"invalid config file: {reason}" format:"string"`
}

func (e InvalidConfigFile) Error() string {
	return "invalid Config file: " + e.Message
}

type InvalidSchemaFile struct {
	Message string `json:"message" example:"invalid schema file: {reason}" format:"string"`
}

func (e InvalidSchemaFile) Error() string {
	return "invalid Schema file: " + e.Message
}
