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

package onboard

import (
	"archive/tar"
	"compress/gzip"
	"io"
	"net/http"
	netUrl "net/url"
	"os"
	"path/filepath"
	"strings"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/api/commons/url"
	"github.com/pkg/errors"

	"gerrit.o-ran-sc.org/r/aiml-fw/aihp/ips/kserve-adapter/pkg/commons/logger"
)

func fetchHelmPackageFromOnboard(url string) (path string, err error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	resp, err := requestToOnboard(url)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	path = "/tmp"

	err = Untar(path, resp.Body)
	if err != nil {
		return
	}
	return
}

func buildURL(name string, version string) string {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	targetURL := netUrl.URL{
		Scheme: "http",
		Host:   onboarderIP + ":" + onboarderPort,
		Path:   url.Onboard(),
	}

	if name != "" {
		targetURL.Path += url.IPSName() + "/" + name

		if version != "" {
			targetURL.Path += url.Version() + "/" + version

		}
	}

	return targetURL.String()
}

func requestToOnboard(url string) (*http.Response, error) {
	logger.Logging(logger.DEBUG, "IN")
	defer logger.Logging(logger.DEBUG, "OUT")

	resp, err := http.Get(url)
	if err != nil {
		logger.Logging(logger.ERROR, err.Error())
		// TODO : define errors
		return nil, errors.Wrap(err, "internal server error")
	}

	if resp.StatusCode != http.StatusOK {
		// TODO : define errors
		err = errors.New("internal server error")
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
