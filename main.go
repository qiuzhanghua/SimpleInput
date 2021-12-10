package main

import (
	"fmt"
	"github.com/magiconair/properties"
	"github.com/qiuzhanghua/i10n"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func main() {
	filename := "data" + string(filepath.Separator) + "app_zh.properties"
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	p, _ := properties.Load(data, properties.UTF8)
	tag := i10n.ParseTagWithDefault(filename)
	i10n.SetDefaultLang("zh-CN")
	i10n.AddTagMap(tag, p.Map())
	fmt.Println(i10n.TT("hello", tag))
	fmt.Println(i10n.T("hello", "Daniel"))
}
