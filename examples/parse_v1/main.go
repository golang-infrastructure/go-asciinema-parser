package main

import (
	"context"
	"fmt"
	asciinema_parser "github.com/golang-infrastructure/go-asciinema-parser"
)

func main() {

	asciiCastV1String := `{
  "version": 1,
  "width": 80,
  "height": 24,
  "duration": 1.515658,
  "command": "/bin/zsh",
  "title": "",
  "env": {
    "TERM": "xterm-256color",
    "SHELL": "/bin/zsh"
  },
  "stdout": [
    [
      0.248848,
      "\u001b[1;31mHello \u001b[32mWorld!\u001b[0m\n"
    ],
    [
      1.001376,
      "I am \rThis is on the next line."
    ]
  ]
}`

	v1, err := asciinema_parser.ParseV1(context.Background(), []byte(asciiCastV1String))
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%v", v1))

}
