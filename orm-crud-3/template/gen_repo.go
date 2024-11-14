package main

//go:generate go run gen_repo.go -entName User

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path"
	"strings"
	"text/template"
)

var entNameParam = flag.String("entName", "", "entName")

func main() {

	flag.Parse()

	entName := *entNameParam
	if len(entName) == 0 {
		panic("entName参数")
	}

	data := map[string]any{
		"entName": entName,
	}
	tmplFile := "./repo.tmpl"
	targetFile := fmt.Sprintf("../repository/%s.go", strings.ToLower(entName)+"_crud")
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
