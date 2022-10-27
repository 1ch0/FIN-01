#  原文地址 [colobu.com](https://colobu.com/2016/12/21/how-to-dump-goroutine-stack-traces/)

> Stack trace 是指堆栈回溯信息，在当前时间，以当前方法的执行点开始，回溯调用它的方法的方法的执行点，然后继续回溯，这样就可以跟踪整个方法的调用, 大家比较熟悉的是 JDK 所带的 jstack 工具，可以把 Java 的所有线程的 stack trace 都打印出来。

Stack trace 是指堆栈回溯信息，在当前时间，以当前方法的执行点开始，回溯调用它的方法的方法的执行点，然后继续回溯，这样就可以跟踪整个方法的调用, 大家比较熟悉的是 JDK 所带的`jstack`工具，可以把 Java 的所有线程的 stack trace 都打印出来。

它有什么用呢？用处非常的大，当应用出现一些状况的时候，比如某个模块不执行， 锁竞争、CPU 占用非常高等问题， 又没有足够的 log 信息可以分析，那么可以查看 stack trace 信息，看看线程都被阻塞或者运行在那些代码上，然后定位问题所在。

对于 Go 开发的程序，有没有类似`jstack`这样的利器呢，目前我还没有看到，但是我们可以通过其它途径也很方便的输出 goroutine 的 stack trace 信息。

本文介绍了几种方法，尤其是最后介绍的方法比较有用。

异常退出情况下输出 stacktrace
--------------------

### 通过 panic

如果应用中有没 recover 的 panic, 或者应用在运行的时候出现运行时的异常，那么程序自动会将当前的 goroutine 的 stack trace 打印出来。

比如下面的程序，如果你运行会抛出一个 panic。

```
package main

import (

"time"

)

func main() {

go a()

	m1()

}

func m1() {

	m2()

}

func m2() {

	m3()

}

func m3() {

panic("panic from m3")

}

func a() {

	time.Sleep(time.Hour)

}
```

输出下面的 stack trace:

```
dump go run p.go

panic: panic from m3

goroutine 1 [running]:

panic(0x596a0, 0xc42000a1a0)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/panic.go:500 +0x1a1

main.m3()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:21 +0x6d

main.m2()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:17 +0x14

main.m1()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:13 +0x14

main.main()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:9 +0x3a

exit status 2
```

从这个信息中我们可以看到 p.go 的第 9 行是 main 方法，它在这一行调用 m1 方法，m1 方法在第 13 行调用 m2 方法，m2 方法在第 17 行调用 m3 方法，m3 方法在第 21 出现 panic， 它们运行在 goroutine 1 中，当前 goroutine 1 的状态是 running 状态。

如果想让它把所有的 goroutine 信息都输出出来，可以设置 `GOTRACEBACK=1`:

```
GOTRACEBACK=1 go run p.go

panic: panic from m3

goroutine 1 [running]:

panic(0x596a0, 0xc42000a1b0)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/panic.go:500 +0x1a1

main.m3()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:21 +0x6d

main.m2()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:17 +0x14

main.m1()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:13 +0x14

main.main()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:9 +0x3a

goroutine 4 [sleep]:

time.Sleep(0x34630b8a000)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/time.go:59 +0xe1

main.a()

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:25 +0x30

created by main.main

	/Users/yuepan/go/src/github.com/smallnest/dump/p.go:8 +0x35

exit status 2
```

同样你也可以分析这个 stack trace 的信息，得到方法调用点的情况，同时这个信息将两个 goroutine 的 stack trace 都打印出来了，而且 goroutine 4 的状态是 sleep 状态。

Go 官方文档对这个环境变量有[介绍](https://golang.org/pkg/runtime/)：

> The GOTRACEBACK variable controls the amount of output generated when a Go program fails due to an unrecovered panic or an unexpected runtime condition. By default, a failure prints a stack trace for the current goroutine, eliding functions internal to the run-time system, and then exits with exit code 2. The failure prints stack traces for all goroutines if there is no current goroutine or the failure is internal to the run-time. GOTRACEBACK=none omits the goroutine stack traces entirely. GOTRACEBACK=single (the default) behaves as described above. GOTRACEBACK=all adds stack traces for all user-created goroutines. GOTRACEBACK=system is like “all” but adds stack frames for run-time functions and shows goroutines created internally by the run-time. GOTRACEBACK=crash is like “system” but crashes in an operating system-specific manner instead of exiting. For example, on Unix systems, the crash raises SIGABRT to trigger a core dump. For historical reasons, the GOTRACEBACK settings 0, 1, and 2 are synonyms for none, all, and system, respectively. The runtime/debug package's SetTraceback function allows increasing the amount of output at run time, but it cannot reduce the amount below that specified by the environment variable. See [https://golang.org/pkg/runtime/debug/#SetTraceback](https://golang.org/pkg/runtime/debug/#SetTraceback).

你可以设置 `none`、`all`、`system`、`single`、`crash`，历史原因， 你可以可是设置数字`0`、`1`、`2`，分别代表`none`、`all`、`system`。

### 通过 SIGQUIT 信号

如果程序没有发生 panic, 但是程序有问题，" 假死 “不工作，我们想看看哪儿出现了问题，可以给程序发送`SIGQUIT`信号，也可以输出 stack trace 信息。  
比如下面的程序:

```
package main

import (

"time"

)

func main() {

go a()

	m1()

}

func m1() {

	m2()

}

func m2() {

	m3()

}

func m3() {

	time.Sleep(time.Hour)

}

func a() {

	time.Sleep(time.Hour)

}
```

你可以运行`kill -SIGQUIT <pid>` 杀死这个程序，程序在退出的时候输出 strack trace。

正常情况下输出 stacktrace
------------------

上面的情况是必须要求程序退出才能打印出 stack trace 信息，但是有时候我们只是需要跟踪分析一下程序的问题，而不希望程序中断运行。所以我们需要其它的方法来执行。

你可以暴漏一个命令、一个 API 或者监听一个信号，然后调用相应的方法把 stack trace 打印出来。

### 打印出当前 goroutine 的 stacktrace

通过`debug.PrintStack()`方法可以将当前所在的 goroutine 的 stack trace 打印出来，如下面的程序。

```
package main

import (

"runtime/debug"

"time"

)

func main() {

go a()

	m1()

}

func m1() {

	m2()

}

func m2() {

	m3()

}

func m3() {

	debug.PrintStack()

	time.Sleep(time.Hour)

}

func a() {

	time.Sleep(time.Hour)

}
```

### 打印出所有 goroutine 的 stacktrace

可以通过`pprof.Lookup("goroutine").WriteTo`将所有的 goroutine 的 stack trace 都打印出来，如下面的程序：

```
package main

import (

"os"

"runtime/pprof"

"time"

)

func main() {

go a()

	m1()

}

func m1() {

	m2()

}

func m2() {

	m3()

}

func m3() {

	pprof.Lookup("goroutine").WriteTo(os.Stdout, 1)

	time.Sleep(time.Hour)

}

func a() {

	time.Sleep(time.Hour)

}
```

### 较完美的输出 stacktrace

你可以使用`runtime.Stack`得到所有的 goroutine 的 stack trace 信息，事实上前面`debug.PrintStack()`也是通过这个方法获得的。

为了更方便的随时的得到应用所有的 goroutine 的 stack trace 信息，我们可以监听`SIGUSER1`信号，当收到这个信号的时候就将 stack trace 打印出来。发送信号也很简单，通过`kill -SIGUSER1 <pid>`就可以，不必担心`kill`会将程序杀死，它只是发了一个信号而已。

```
package main

import (

"fmt"

"os"

"os/signal"

"runtime"

"syscall"

"time"

)

func main() {

	setupSigusr1Trap()

go a()

	m1()

}

func m1() {

	m2()

}

func m2() {

	m3()

}

func m3() {

	time.Sleep(time.Hour)

}

func a() {

	time.Sleep(time.Hour)

}

func setupSigusr1Trap() {

	c := make(chan os.Signal, 1)

	signal.Notify(c, syscall.SIGUSR1)

go func() {

for range c {

			DumpStacks()

		}

	}()

}

func DumpStacks() {

	buf := make([]byte, 16384)

	buf = buf[:runtime.Stack(buf, true)]

	fmt.Printf("=== BEGIN goroutine stack dump ===\n%s\n=== END goroutine stack dump ===", buf)

}
```

输出结果很直观，方便检查。

```
=== BEGIN goroutine stack dump ===

goroutine 36 [running]:

main.DumpStacks()

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:47 +0x77

main.setupSigusr1Trap.func1(0xc420070060)

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:40 +0x73

created by main.setupSigusr1Trap

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:42 +0xec

goroutine 1 [sleep]:

time.Sleep(0x34630b8a000)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/time.go:59 +0xe1

main.m3()

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:28 +0x30

main.m2()

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:24 +0x14

main.m1()

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:20 +0x14

main.main()

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:16 +0x3f

goroutine 34 [syscall]:

os/signal.signal_recv(0xff280)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/sigqueue.go:116 +0x157

os/signal.loop()

	/usr/local/Cellar/go/1.7.4/libexec/src/os/signal/signal_unix.go:22 +0x22

created by os/signal.init.1

	/usr/local/Cellar/go/1.7.4/libexec/src/os/signal/signal_unix.go:28 +0x41

goroutine 35 [select, locked to thread]:

runtime.gopark(0xb5cc8, 0x0, 0xab3ef, 0x6, 0x18, 0x2)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/proc.go:259 +0x13a

runtime.selectgoImpl(0xc42008d730, 0x0, 0x18)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/select.go:423 +0x11d9

runtime.selectgo(0xc42008d730)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/select.go:238 +0x1c

runtime.ensureSigM.func1()

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/signal1_unix.go:304 +0x2d1

runtime.goexit()

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/asm_amd64.s:2086 +0x1

goroutine 37 [sleep]:

time.Sleep(0x34630b8a000)

	/usr/local/Cellar/go/1.7.4/libexec/src/runtime/time.go:59 +0xe1

main.a()

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:32 +0x30

created by main.main

	/Users/yuepan/go/src/github.com/smallnest/dump/d3.go:15 +0x3a

=== END goroutine stack dump ===
```

当然，这段代码也不是我原创的，这是 docker 代码库中的一段[代码](https://github.com/docker/docker/blob/95fcf76cc64a4acf95c168e8d8607e3acf405c13/pkg/signal/trap.go)，很简单，也很强大。

配置 http/pprof
-------------

如果你的代码中配置了 `http/pprof`, 你可以通过下面的地址访问所有的 groutine 的堆栈：

```
http://localhost:8888/debug/pprof/goroutine?debug=2.
```

参考文档
----

1.  [http://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces](http://stackoverflow.com/questions/19094099/how-to-dump-goroutine-stacktraces)
2.  [https://golang.org/pkg/runtime/](https://golang.org/pkg/runtime/)

[**Newer**

一个有特色的有限状态机

](https://colobu.com/2016/12/24/a-featured-fsm/)[**Older**

Github 和 gitlab 的自动连接

](https://colobu.com/2016/12/20/detect-and-link-references-in-github/)