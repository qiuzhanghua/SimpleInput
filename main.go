package main

import (
	"fmt"
	"github.com/qiuzhanghua/go-input"
	"os"
)

func main() {

	ui := &input.UI{
		Writer: os.Stdout,
		Reader: os.Stdin,
	}

	query := "What is your name?"
	name, err := ui.Ask(query, &input.Options{
		//Default:  "Daniel",
		Required: true,
		Loop:     true,
	})

	fmt.Println(name, err)

	query = "Which language do you prefer to use?"
	lang, err := ui.Select(query, []string{"go", "Go", "golang"}, &input.Options{
		Default: "Go",
		Loop:    true,
	})

	fmt.Println(lang, err)

}
