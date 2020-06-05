// SPDX-License-Identifier: MIT

package cmd

import (
	"bytes"
	"testing"

	"github.com/issue9/assert"

	"github.com/caixw/apidoc/v7/internal/docs"
)

func TestCmdTest(t *testing.T) {
	a := assert.New(t)

	buf := new(bytes.Buffer)
	cmd := Init(buf)
	erro, _, succ, _ := resetPrinters()
	err := cmd.Exec([]string{"test", "-d", docs.Dir().Append("example").String()})
	a.NotError(err)
	a.Empty(buf.String()).
		Empty(erro.String()).
		NotEmpty(succ.String())
}