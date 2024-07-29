## TODO
1. 自定义日志等级
2. 日志文件存储

## 使用示例
``` go
package main

import "github.com/revevide/gtools/zapx"

func main() {
	zapx.NewZapx()
	zapx.Info("hello world")
}
```

```shell
{"level":"info","ts":"2024-07-29T15:25:19.155+0800","caller":"test/main.go:7","msg":"hello world"}
```
