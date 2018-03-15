package gc_util

import (
	"fmt"
	"runtime/debug"
	"runtime"
	"os"
)

// Report prints caller file and number and prints all current stack
// Report gives off a warning requesting the user to submit an issue to the github tracker.
func Report(extra ...interface{}) {
	fmt.Fprintln(os.Stderr, "You've encountered a sought after, hard to reproduce bug. Please report this to the developers <3 https://github.com/linxinzhe/go-codebase/issues")
	fmt.Fprintln(os.Stderr, extra...)

	_, file, line, _ := runtime.Caller(1)         //lxz 跳过Report自身，打印调用Report的函数的堆栈的行和文件
	fmt.Fprintf(os.Stderr, "%v:%v\n", file, line) //lxz F表示自行决定打印的输出流，printf表示按参数打印

	debug.PrintStack() //lxz 打印所有的内存堆栈

	fmt.Fprintln(os.Stderr, "#### BUG! PLEASE REPORT ####")
}
