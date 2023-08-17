package trace

import (
	"fmt"
	"os"
	"runtime"

	"github.com/davecgh/go-spew/spew"
)

// 产生panic时的调用栈打印
func PrintPanicStack(extras ...interface{}) {
	i := 0
	funcName, file, line, ok := runtime.Caller(i)
	for ok {
		fmt.Fprintf(os.Stderr, "frame %v:[func:%v,file:%v,line:%v]\n", i, runtime.FuncForPC(funcName).Name(), file, line)
		i++
		funcName, file, line, ok = runtime.Caller(i)
	}
	for k := range extras {
		fmt.Fprintf(os.Stderr, "EXTRAS#%v DATA:%v\n", k, spew.Sdump(extras[k]))
	}
}
