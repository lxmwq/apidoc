// Copyright 2016 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"path/filepath"
	"strings"
	"testing"

	"github.com/issue9/assert"

	"github.com/caixw/apidoc/input"
	"github.com/caixw/apidoc/internal/vars"
	"github.com/caixw/apidoc/output"
)

func TestConfig_generateConfig_loadConfig(t *testing.T) {
	a := assert.New(t)

	path, err := filepath.Abs("./.apidoc.yaml")
	a.NotError(err).NotEmpty(path)

	a.NotError(generateConfig("./..", path))
	cfg, err := loadConfig(path)
	a.NotError(err).NotNil(cfg)

	a.Equal(cfg.Version, vars.Version())
}

func TestConfig_sanitize(t *testing.T) {
	a := assert.New(t)

	conf := &config{}
	err := conf.sanitize()
	a.Error(err)
	a.Equal(err.(*configError).Field, "version")

	// 版本号错误
	conf.Version = "5.0"
	err = conf.sanitize()
	a.Error(err)
	a.Equal(err.(*configError).Field, "version")

	// 未声明 inputs
	conf.Version = "5.0.1"
	err = conf.sanitize()
	a.Error(err)
	a.Equal(err.(*configError).Field, "inputs")

	// 未声明 output
	conf.Inputs = []*input.Options{{}}
	err = conf.sanitize()
	a.Error(err)
	a.Equal(err.(*configError).Field, "output")

	// 查看错误提示格式是否正确
	conf.Output = &output.Options{}
	conf.Inputs = append(conf.Inputs, &input.Options{
		Lang: "123",
	})
	err = conf.sanitize()
	a.Error(err)
	a.True(strings.HasPrefix(err.(*configError).Field, "inputs[0]"))
}
