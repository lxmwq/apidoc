// SPDX-License-Identifier: MIT

package mock

import (
	"testing"

	"github.com/issue9/assert"

	"github.com/caixw/apidoc/v7/internal/ast"
	"github.com/caixw/apidoc/v7/internal/token"
)

func TestJSONValidator_find(t *testing.T) {
	a := assert.New(t)

	v := &jsonValidator{
		param: (&ast.Request{
			Name: &ast.Attribute{Value: token.String{Value: "root"}},
			Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeObject}},
			Items: []*ast.Param{
				{
					Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeString}},
					Name: &ast.Attribute{Value: token.String{Value: "name"}},
				},
				{
					Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeNumber}},
					Name: &ast.Attribute{Value: token.String{Value: "id"}},
				},
				{
					Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeObject}},
					Name: &ast.Attribute{Value: token.String{Value: "group"}},
					Items: []*ast.Param{
						{
							Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeString}},
							Name: &ast.Attribute{Value: token.String{Value: "name"}},
						},
						{
							Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeNumber}},
							Name: &ast.Attribute{Value: token.String{Value: "id"}},
						},
						{
							Name: &ast.Attribute{Value: token.String{Value: "tags"}},
							Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeObject}},
							Items: []*ast.Param{
								{
									Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeString}},
									Name: &ast.Attribute{Value: token.String{Value: "name"}},
								},
								{
									Type: &ast.TypeAttribute{Value: token.String{Value: ast.TypeNumber}},
									Name: &ast.Attribute{Value: token.String{Value: "id"}},
								},
							},
						}, // end tags
					},
				}, // end group
			},
		}).Param(),
	}

	v.names = []string{}
	p := v.find()
	a.Equal(p, v.param)

	v.names = nil
	p = v.find()
	a.Equal(p, v.param)

	v.names = []string{""}
	p = v.find()
	a.Nil(p)

	v.names = []string{"name"}
	p = v.find()
	a.NotNil(p).Equal(p.Type.V(), ast.TypeString)

	v.names = []string{"not-exists"}
	p = v.find()
	a.Nil(p)

	v.names = []string{"group", "id"}
	p = v.find()
	a.NotNil(p).Equal(p.Type.V(), ast.TypeNumber)

	v.names = []string{"group", "tags", "id"}
	p = v.find()
	a.NotNil(p).Equal(p.Type.V(), ast.TypeNumber)
}

func TestValidJSON(t *testing.T) {
	a := assert.New(t)

	for _, item := range data {
		err := validJSON(item.Type, []byte(item.JSON))
		a.NotError(err, "测试 %s 时返回错误值 %s", item.Title, err)
	}
}

func TestBuildJSON(t *testing.T) {
	a := assert.New(t)

	for _, item := range data {
		data, err := buildJSON(item.Type, indent, testOptions)

		a.NotError(err, "测试 %s 返回了错误值 %s", item.Title, err).
			Equal(string(data), item.JSON, "测试 %s 失败 v1:%s,v2:%s", item.Title, string(data), item.JSON)
	}
}
