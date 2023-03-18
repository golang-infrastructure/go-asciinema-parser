# asciinema parser

[English Document](./README.md)
[中文文档](./README_zh.md)

# 1. What is this? What problem was solved?

This website [https://asciinema.org/](https://asciinema.org/) is record and share your terminal session website, the simple way., it defines a set of Ascii Cast format files to store screen content, this library is used to parse Ascii Cast files，support [v1](https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v1.md) and [v2](https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v2.md) two versions.

# 2. API code examples

## 2.1 Check the version of the screen recording file

```go
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

	version, err := asciinema_parser.DetectVersion(context.Background(), asciiCastV2String)
	if err != nil {
		panic(err)
	}
	fmt.Println(version) // Output: 2

}
```

## 2.2 Parse V1 format of the screen recording software

```go
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
```

## 2.3 Parse V2 format of the screen recording software

```go
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

	v2, err := asciinema_parser.ParseV2(context.Background(), []byte(asciiCastV2String))
	if err != nil {
		panic(err)
	}
	fmt.Println(fmt.Sprintf("%v", v2))

}
```



