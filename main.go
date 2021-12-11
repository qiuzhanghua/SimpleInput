package main

import (
	"fmt"
	"github.com/magiconair/properties"
	"github.com/qiuzhanghua/go-input"
	"github.com/qiuzhanghua/i10n"
	"log"
	"strings"
)

func init() {
	_ = i10n.SetDefaultLang("zh-CN")
	for _, name := range AssetNames() {
		if strings.HasPrefix(name, "locales") && strings.HasSuffix(name, ".properties") {
			buffer, err := Asset(name)
			if err != nil {
				log.Fatal(err)
			}
			p, err := properties.Load(buffer, properties.UTF8)
			if err != nil {
				log.Fatal(err)
			}
			tag := i10n.ParseTagWithDefault(name)
			i10n.AddTagMap(tag, p.Map())
		}
	}
	// 使用i10的方法
	input.T = i10n.T
}

func main() {
	ui := input.DefaultUI()

	query := "What is your name?"
	name, err := ui.Ask(query, &input.Options{
		//		Default:  "qiuzhanghua",
		Required: true,
		Loop:     true,
	})
	fmt.Println(name, err)

	query = "What is your password?"
	name, err = ui.Ask(query, &input.Options{
		Required:    true,
		Mask:        false,
		MaskDefault: true,
		Loop:        true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer is %s\n", name)

	query = "Which language do you prefer to use?"
	lang, err := ui.Select(query, []string{"go", "Go", "golang"}, &input.Options{
		Default: "Go",
		Loop:    true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer is %s\n", lang)

	query = "Do you love golang [Y/n]"
	name, err = ui.Ask(query, &input.Options{
		Required: true,
		// Validate input
		ValidateFunc: func(s string) error {
			if s != "Y" && s != "n" {
				return fmt.Errorf("input must be Y or n")
			}
			return nil
		},
		Loop: true,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Answer is %s\n", name)

}
