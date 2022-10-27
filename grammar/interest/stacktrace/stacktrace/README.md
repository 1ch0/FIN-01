在 go 的安装目录修改 Gosrcsyscall ypes_windows.go，增加如下代码：

```go
var signals = [...]string{
// 这里省略N行。。。。

    /** 兼容windows start */
    16: "SIGUSR1",
    17: "SIGUSR2",
    18: "SIGTSTP",
    /** 兼容windows end */
}

/** 兼容windows start */
func Kill(...interface{}) {
return;
}
const (
SIGUSR1 = Signal(0x10)
SIGUSR2 = Signal(0x11)
SIGTSTP = Signal(0x12)
)
/** 兼容windows end */

```
