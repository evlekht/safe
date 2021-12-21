package safe

import (
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
