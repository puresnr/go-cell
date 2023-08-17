package csync

import (
	"github.com/puresnr/go-cell/gosafe"
	"sync"
)

// GoWait 用于并发的执行一组函数, 并等待所有函数执行完毕
func GoWait(funcs ...func()) {
	var wg sync.WaitGroup

	for _, f := range funcs {
		wg.Add(1)

		gosafe.GoP(func(ef func()) {
			defer wg.Done()

			ef()
		}, f)
	}

	wg.Wait()
}
