package main

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"os/exec"
	"strings"
)

func main() {
	termState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Failed to set raw mode on STDIN: %v\n",
			err)
		return
	}
	n := term.NewTerminal(os.Stdin, "tdp> ")
	_ = n.SetSize(int(^uint(0)>>1), 0)
	var ln, input string

	for {
		ln, err = n.ReadLine()
		_ = term.Restore(int(os.Stdin.Fd()), termState)
		if err != nil {
			break
		}
		input = strings.Trim(ln, "\n ")
		switch {
		case input == "exit":
			os.Exit(0)
		case input == "pwd":
			cmd := exec.Command("pwd")
			//cmd := exec.Command("gradle", "build")
			cmd.Stdout = os.Stdout
			cmd.Env = append([]string{"UI=path_to_tdp_ui"}, os.Environ()...)
			// fmt.Println(cmd.Env)
			if err := cmd.Run(); err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "Failed: %v\n",
					err)
			}
			// Add more
		default:
			fmt.Println(input)
		}
		termState, err = term.MakeRaw(int(os.Stdin.Fd()))
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "Failed to set raw mode on STDIN: %v\n",
				err)
			break
		}
	}
}
