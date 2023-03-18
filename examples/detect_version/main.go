package main

import (
	"context"
	"fmt"
	asciinema_parser "github.com/golang-infrastructure/go-asciinema-parser"
)

func main() {

	asciiCastV2String := `{"version": 2, "width": 80, "height": 24, "timestamp": 1504467315, "title": "Demo", "env": {"TERM": "xterm-256color", "SHELL": "/bin/zsh"}}
[0.248848, "o", "\u001b[1;31mHello \u001b[32mWorld!\u001b[0m\n"]
[1.001376, "o", "That was ok\rThis is better."]
[2.143733, "o", " "]
[6.541828, "o", "Bye!"]`

	// 识别版本
	version, err := asciinema_parser.DetectVersion(context.Background(), asciiCastV2String)
	if err != nil {
		panic(err)
	}
	fmt.Println(version) // Output: 2

}
