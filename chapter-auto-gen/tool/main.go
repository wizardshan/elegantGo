package main

import (
	"elegantGo/chapter-orm-entgo/repository/ent/schema"
	entgo "entgo.io/ent"
	"errors"
	"flag"
	"fmt"
	"github.com/elliotchance/pie/v2"
	"github.com/gobeam/stringy"
	"os"
	"path"
	"strings"
	"text/template"
)

type FieldsGetter interface {
	Fields() []entgo.Field
}

type EdgesGetter interface {
	Edges() []entgo.Edge
}

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
	"User":    schema.User{},
	"Post":    schema.Post{},
	"Comment": schema.Comment{},
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
	"entMapper": {
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
	"respMapper": {
		target: "../controller/response/%s_mapper.go",
		tmpl:   "./tmpl/resp_mapper.tmpl",
	},
	"ctr": {
		target: "../controller/%s.go",
		tmpl:   "./tmpl/ctr.tmpl",
	},
}

type Field struct {
	Name     string
	TypeName string
	IsSlice  bool
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
	edgesFields := parseEdgesFields(entity)

	data := map[string]any{
		"name":        strings.ToLower(name),
		"Name":        name,
		"fields":      fields,
		"edgesFields": edgesFields,
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

func parseEdgesFields(entity any) []*Field {
	var fields []*Field
	edgesGetter, ok := entity.(EdgesGetter)
	if !ok {
		return fields
	}
	pie.All(edgesGetter.Edges(), func(f entgo.Edge) bool {
		fields = append(fields, &Field{Name: pascalCase(f.Descriptor().Name), TypeName: f.Descriptor().Type, IsSlice: !f.Descriptor().Inverse})
		return true
	})

	return fields
}

func parseFields(entity any) []*Field {
	var fields []*Field
	fieldsGetter, ok := entity.(FieldsGetter)
	if !ok {
		return fields
	}

	pie.All(fieldsGetter.Fields(), func(f entgo.Field) bool {
		fields = append(fields, &Field{Name: pascalCase(f.Descriptor().Name), TypeName: f.Descriptor().Info.String()})
		return true
	})

	idField := &Field{
		Name:     "ID",
		TypeName: "int",
	}

	createTimeField := &Field{
		Name:     "CreateTime",
		TypeName: "time.Time",
	}

	updateTimeField := &Field{
		Name:     "UpdateTime",
		TypeName: "time.Time",
	}

	fields = append(fields, idField, createTimeField, updateTimeField)
	return fields
}

func pascalCase(s string) string {
	s = strings.ReplaceAll(s, "id", "ID")
	return stringy.New(s).PascalCase().Get()
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
