package main

import (
	"fmt"
	"github.com/qiuzhanghua/go-input"
)

func main() {
	ui := input.DefaultUI()

	query := "What is your name?"
	name, err := ui.Ask(query, &input.Options{
		//		Default:  "qiuzhanghua",
		Required: true,
		Loop:     true,
	})
	fmt.Println(name, err)
}
