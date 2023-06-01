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

package ricdms

import (
	"archive/tar"
	"bytes"
	"compress/gzip"
	"io"
	"math/rand"
	"mime/multipart"
	"net/http"
	netUrl "net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/url"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/errors"
	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

const (
	pathLength    = 20
	letterBytes   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	letterIdxBits = 6
	letterIdxMask = 1<<letterIdxBits - 1
	letterIdxMax  = 63 / letterIdxBits
)

var randSrc rand.Source

func init() {
	randSrc = rand.NewSource(time.Now().UnixNano())
}

func buildURLforChartOnboard() string {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL := netUrl.URL{
		Scheme: "http",
		Host:   ricdmsIP + ":" + ricdmsPort,
		Path:   url.Onboard(),
	}

	return targetURL.String()
}

func buildURLforChartFetch(name string, version string) (string, error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL := netUrl.URL{
		Scheme: "http",
		Host:   ricdmsIP + ":" + ricdmsPort,
		Path:   url.Download(),
	}

	if name != "" {
		targetURL.Path = strings.Replace(targetURL.Path, "{xApp_name}", name, -1)
	} else {
		return "", errors.InternalServerError{Message: "ifsv name is not given"}
	}

	if version != "" {
		targetURL.Path = strings.Replace(targetURL.Path, "{version}", version, -1)
	} else {
		return "", errors.InternalServerError{Message: "ifsv version is not given"}
	}

	return targetURL.String(), nil
}

func getChartFromRicdms(url string) (*http.Response, error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	resp, err := http.Get(url)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return nil, errors.InternalServerError{Message: err.Error()}
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.InternalServerError{
			Message: "ricdms return error, status code : " + strconv.Itoa(int(resp.StatusCode)),
		}
		logger.Logging(logger.ERROR, err.Error())
		return nil, err
	}
	return resp, err
}

func Untar(dst string, r io.Reader) error {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	gzr, err := gzip.NewReader(r)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		return err
	}
	defer gzr.Close()

	tr := tar.NewReader(gzr)

	for {
		header, err := tr.Next()

		switch {
		case err == io.EOF:
			logger.Logging(logger.ERROR, err.Error())
			return nil
		case err != nil:
			logger.Logging(logger.ERROR, err.Error())
			return err
		case header == nil:
			continue
		}

		target := filepath.Join(dst, header.Name)

		switch header.Typeflag {
		case tar.TypeDir:
			err := mkdirectory(target)
			if err != nil {
				logger.Logging(logger.ERROR, err.Error())
				return err
			}
		case tar.TypeReg:
			paths := strings.Split(target, "/")
			if len(paths) != 0 {
				var path string
				for idx := range paths {
					if !strings.Contains(paths[idx], ".") {
						path = filepath.Join(path, paths[idx])
					}
				}
				err = mkdirectory(path)
				if err != nil {
					logger.Logging(logger.ERROR, err.Error())
					return err
				}
			}
			f, err := os.OpenFile(target, os.O_CREATE|os.O_RDWR, os.FileMode(header.Mode))
			if err != nil {
				logger.Logging(logger.ERROR, err.Error())
				return err
			}
			_, err = io.Copy(f, tr)
			if err != nil {
				logger.Logging(logger.ERROR, err.Error())
				return err
			}
			f.Close()
		}
	}
}

func mkdirectory(name string) (err error) {
	_, err = os.Stat(name)
	if err != nil {
		err = os.MkdirAll(name, 0755)
		if err != nil {
			logger.Logging(logger.ERROR, err.Error())
			return
		}
	}
	return
}

func generateRandPath(n int) string {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	sb := strings.Builder{}
	sb.Grow(n)
	for i, cache, remain := n-1, randSrc.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = randSrc.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			sb.WriteByte(letterBytes[idx])
			i--
		}
		cache >>= letterIdxBits
		remain--
	}

	return sb.String()
}

func onboardChartToRicdms(targetURL string, chartPath string) (*http.Response, error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	chartFile, err := os.Open(chartPath)
	if err != nil {
		logger.Logging(logger.DEBUG, err.Error())
	}
	defer chartFile.Close()

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, _ := writer.CreateFormFile("helmpkg", chartFile.Name())
	io.Copy(fw, chartFile)
	writer.Close()
	req, err := http.NewRequest("POST", targetURL, body)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return resp, errors.InternalServerError{Message: err.Error()}
	}

	if resp.StatusCode != http.StatusOK {
		err = errors.InternalServerError{
			Message: "ricdms return error, status code : " + strconv.Itoa(int(resp.StatusCode)),
		}
		logger.Logging(logger.ERROR, err.Error())
		return resp, err
	}
	return resp, err
}
