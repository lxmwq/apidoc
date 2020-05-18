// SPDX-License-Identifier: MIT

package token

import (
	"testing"

	"github.com/issue9/assert"
	"golang.org/x/text/language"

	"github.com/caixw/apidoc/v7/internal/locale"
)

func TestNewTypes(t *testing.T) {
	a := assert.New(t)
	id := locale.DefaultLocaleID

	ts, err := NewTypes(&objectTag{}, language.MustParse(id))
	a.NotError(err).NotNil(ts)
	ts2 := &Types{Types: []*Type{
		{
			Name:  "apidoc",
			Usage: InnerXML{locale.Translate(id, "usage-root")},
			Items: []*Item{
				{
					Name:     "@id",
					Usage:    locale.Translate(id, "usage"),
					Type:     "number",
					Array:    false,
					Required: true,
				},
				{
					Name:     "name",
					Usage:    locale.Translate(id, "usage"),
					Type:     "string",
					Array:    false,
					Required: true,
				},
			},
		},
		{
			Name:  "number",
			Usage: InnerXML{Text: locale.Translate(id, "usage-number")},
			Items: []*Item{},
		},
		{
			Name:  "string",
			Usage: InnerXML{Text: locale.Translate(id, "usage-string")},
			Items: []*Item{},
		},
	}}
	a.Equal(len(ts.Types), len(ts2.Types))
	a.Equal(ts, ts2, "not equal\nv1=%#v\nv2=%#v\n", ts, ts2)
}