package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

var entityNameParam = flag.String("entityName", "", "entityName参数")

func main() {

	flag.Parse()

	entityName := *entityNameParam
	if len(entityName) == 0 {
		panic("entityName参数")
	}

	data := map[string]any{
		"entityName": entityName,
	}
	tmplFile := "./repo.tmpl"
	targetFile := fmt.Sprintf("../repository/%s.go", strings.ToLower(entityName)+"-crud")
	err := parse(tmplFile, targetFile, data)
	if err != nil {
		fmt.Println("os.Create", targetFile, err)
	}
}

func parse(tmplFile string, targetFile string, data map[string]any) error {

	funcMap := template.FuncMap{
		"ToLower": strings.ToLower,
	}

	tmpl, err := template.New(path.Base(tmplFile)).Funcs(funcMap).ParseFiles(tmplFile)
	if err != nil {
		return err
	}

	_, err = os.Stat(targetFile)
	if err == nil {
		return errors.New("文件已存在")
	}

	target, err := os.Create(targetFile)
	if err != nil {
		return err
	}
	defer target.Close()

	return tmpl.Execute(target, data)
}
