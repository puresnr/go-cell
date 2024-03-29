package cerror

import (
	"fmt"
	"runtime"
)

func Wrap(err error) error {
	if err == nil {
		return nil
	}

	pc, f, l, _ := runtime.Caller(1)

	return fmt.Errorf("<file: %s><line: %d><func: %s> | %w", f, l, runtime.FuncForPC(pc).Name(), err)
}

func WrapWithSvcName(svc string, err error) error {
	if err == nil {
		return nil
	}

	pc, f, l, _ := runtime.Caller(1)

	return fmt.Errorf("<svc: %s><file: %s><line: %d><func: %s> | %w", svc, f, l, runtime.FuncForPC(pc).Name(), err)
}
