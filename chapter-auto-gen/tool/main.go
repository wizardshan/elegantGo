package main

import (
	"elegantGo/chapter-auto-gen/repository/ent"
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"reflect"
	"strings"
	"text/template"
)

type flags []string

func (fs *flags) Set(s string) error {
	*fs = append(*fs, s)
	return nil
}

func (fs *flags) String() string {
	return strings.Join(*fs, " ")
}

type Layer struct {
	target string
	tmpl   string
}

var entityMapper = map[string]any{
	"User": ent.User{},
}

var layerMapping = map[string]Layer{
	"req": {
		target: "../controller/request/%s.go",
		tmpl:   "./tmpl/req.tmpl",
	},
	"repo": {
		target: "../repository/%s_crud.go",
		tmpl:   "./tmpl/repo.tmpl",
	},
	"ent": {
		target: "../repository/ent/%s_mapper.go",
		tmpl:   "./tmpl/ent_mapper.tmpl",
	},
	"dom": {
		target: "../domain/%s.go",
		tmpl:   "./tmpl/dom.tmpl",
	},
	"resp": {
		target: "../controller/response/%s.go",
		tmpl:   "./tmpl/resp.tmpl",
	},
	"ctr": {
		target: "../controller/%s.go",
		tmpl:   "./tmpl/ctr.tmpl",
	},
}

type Field struct {
	Name     string
	TypeName string
}

func main() {
	var name string
	var layers flags
	var c bool
	flag.StringVar(&name, "name", "", "名参数")
	flag.Var(&layers, "layers", "层参数")
	flag.BoolVar(&c, "c", false, "清除文件参数")
	flag.Parse()

	if len(name) == 0 {
		panic("name参数异常")
	}

	entity, ok := entityMapper[name]
	if !ok {
		panic("entity不存在")
	}

	fields := parseFields(entity)

	data := map[string]any{
		"name":   strings.ToLower(name),
		"Name":   name,
		"fields": fields,
	}

	if c {
		for _, k := range layers {
			layer, ok := layerMapping[k]
			if !ok {
				continue
			}
			os.Remove(fmt.Sprintf(layer.target, data["name"]))
		}
		return
	}

	for _, k := range layers {
		layer, ok := layerMapping[k]
		if !ok {
			continue
		}

		err := parse(layer.tmpl, fmt.Sprintf(layer.target, data["name"]), data)
		if err != nil {
			fmt.Printf("os.Create %v %v", err, layer)
			fmt.Println()
		}
	}
}

func parse(tmplFile string, targetFile string, data map[string]any) error {

	funcMap := template.FuncMap{
		//"lower": strings.ToLower,
	}

	tmpl, err := template.New(path.Base(tmplFile)).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	if fileExist(targetFile) {
		return errors.New("文件已存在")
	}

	target, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer target.Close()

	return tmpl.Execute(target, data)
}

func parseFields(entity any) []*Field {

	t := reflect.TypeOf(entity)
	//fmt.Printf("结构体名: %v\n", t.Name())
	var fields []*Field
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fieldType := field.Type
		//fmt.Printf("属性名: %v, 属性类型：%v 字段是否可导出: %v\n", field.Name, fieldType.String(), field.IsExported())
		if !field.IsExported() {
			continue
		}

		if field.Name == "Edges" {
			continue
		}

		fields = append(fields, &Field{Name: field.Name, TypeName: fieldType.String()})
	}
	return fields
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
