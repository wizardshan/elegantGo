package main

import (
	"app/chapter-struct-mapper/repository/entity"
	"os"
	"reflect"
	"strings"
	"text/template"
)

var FuncMap = template.FuncMap{}

func init() {
	FuncMap["lowerFirst"] = func(s string) string {
		if s == "" {
			return ""
		}

		if s == "ID" {
			return strings.ToLower(s)
		}

		return strings.ToLower(s[:1]) + s[1:]
	}
}

var entityPath = "./repository/entity/"
var domainPath = "./domain/"
var responsePath = "./controller/response/"

var entities []interface{}

type Field struct {
	Name     string
	TypeName string
}

type ModelData struct {
	StructName string
	Fields     []*Field
}

func init() {
	entities = []interface{}{
		entity.Item{},
	}
}

func main() {
	for _, model := range entities {
		gen(model)
	}
}

func gen(model interface{}) {
	modelData := new(ModelData)

	t := reflect.TypeOf(model)
	modelData.StructName = t.Name()
	//fmt.Printf("结构体名: %v\n", t.Name())
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		//fmt.Printf("属性名: %v, 属性类型：%v 字段是否可导出: %v\n", field.Name, fieldType.String(), field.IsExported())
		if !field.IsExported() {
			continue
		}
		modelData.Fields = append(modelData.Fields, &Field{Name: field.Name, TypeName: fieldType.String()})
	}

	entityTpl :=
		`package entity

func (ent{{.StructName}} *{{.StructName}}) Mapper() *domain.{{.StructName}} {
	if ent{{.StructName}} == nil {
		return nil
	}

	dom{{.StructName}} := new(domain.{{.StructName}})
{{- range .Fields}}
	dom{{$.StructName}}.{{.Name}} = ent{{$.StructName}}.{{.Name}}
{{- end}}

	return dom{{.StructName}}
}

func (ent{{.StructName}}s {{.StructName}}s) Mapper() domain.{{.StructName}}s {

	size := len(ent{{.StructName}}s)
	dom{{.StructName}}s := make(domain.{{.StructName}}s, size)

	for i := 0; i < size; i++ {
		dom{{.StructName}}s[i] = ent{{.StructName}}s[i].Mapper()
	}

	return dom{{.StructName}}s
}`

	fileName := strings.ToLower(modelData.StructName) + "_mapper.go"
	filePath := entityPath + fileName
	genFile("entity", filePath, entityTpl, modelData)

	domainTpl :=
		`package domain

import "time"

type {{.StructName}}s []*{{.StructName}}

type {{.StructName}} struct {
{{- range .Fields}}
	{{.Name}} {{.TypeName}}
{{- end}}
}`
	fileName = strings.ToLower(modelData.StructName) + ".go"
	filePath = domainPath + fileName
	genFile("domain", filePath, domainTpl, modelData)

	responseTpl :=
		`package response

import "time"

type {{.StructName}}s []*{{.StructName}}

type {{.StructName}} struct {
{{- range .Fields}}
	{{.Name}} {{.TypeName}}` + " `json:\"{{lowerFirst .Name}}\"`" +
			`{{- end}}
}

func (resp{{.StructName}} *{{.StructName}}) Mapper(dom{{.StructName}} *domain.{{.StructName}}) *{{.StructName}} {
	if dom{{.StructName}} == nil {
		return nil
	}
{{- range .Fields}}
	resp{{$.StructName}}.{{.Name}} = dom{{$.StructName}}.{{.Name}}
{{- end}}
	return resp{{.StructName}}
}

func (resp{{.StructName}}s {{.StructName}}s) Mapper(dom{{.StructName}}s domain.{{.StructName}}s) {{.StructName}}s {

	size := len(dom{{.StructName}}s)
	resp{{.StructName}}s = make({{.StructName}}s, size)
	for i := 0; i < size; i++ {
		var resp{{.StructName}} {{.StructName}}
		resp{{.StructName}}s[i] = resp{{.StructName}}.Mapper(dom{{.StructName}}s[i])
	}

	return resp{{.StructName}}s
}
`

	filePath = responsePath + fileName
	genFile("response", filePath, responseTpl, modelData)
}

func genFile(templateName string, filePath string, text string, data any) {
	tmpl, err := template.New(templateName).Funcs(FuncMap).Parse(text)
	if err != nil {
		panic("template.New" + err.Error())
	}

	if fileExist(filePath) {
		panic("file exist: " + filePath)
	}

	f, err := os.Create(filePath)
	if err != nil {
		panic("os.Create" + err.Error())
	}

	err = tmpl.Execute(f, data)
	if err != nil {
		panic("template.Execute" + err.Error())
	}
}

func fileExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
