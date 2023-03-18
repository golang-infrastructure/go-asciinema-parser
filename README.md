# asciinema parser 

# 一、这是什么？解决了什么问题？

这个网站[https://asciinema.org/](https://asciinema.org/)是一个命令行录屏分享网站，它定义了一套`Ascii Cast`格式的文件来存储录屏内容，这个库就是用来解析`Ascii Cast`文件的，支持[v1](https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v1.md)和[v2](https://github.com/asciinema/asciinema/blob/develop/doc/asciicast-v2.md)两个版本。

# 二、 API代码示例

## 2.1 检查录屏文件是哪个版本

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

	// 识别版本
	version, err := asciinema_parser.DetectVersion(context.Background(), asciiCastV2String)
	if err != nil {
		panic(err)
	}
	fmt.Println(version) // Output: 2

}
```

## 2.2 解析V1格式的录屏软件

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

## 2.3 解析V2格式的录屏软件

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



