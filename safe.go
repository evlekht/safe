package safe

import (
	"context"
	"fmt"
)

func Invoke(f func()) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()
	f()
	return
}

func InvokeWithErr(f func() error) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()
	err = f()
	return
}

func InvokeWithLog(logger Logger, f func()) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			logger.Error(context.Background(), err)
		}
	}()
	f()
	return
}

type Logger interface {
	Error(ctx context.Context, args ...interface{})
}
