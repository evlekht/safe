package safe

import (
	"context"
	"fmt"
)

func Invoke(f func()) {
	defer func() {
		recover()
	}()
	f()
}

func InvokeWithErr(f func() error) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
		}
	}()
	return f()
}

//
//

type LoggerWithContext interface {
	Error(ctx context.Context, args ...interface{})
}

type Logger interface {
	Error(args ...interface{})
}

func InvokeWithLogContext(ctx context.Context, logger LoggerWithContext, f func()) {
	defer func() {
		if rec := recover(); rec != nil {
			logger.Error(ctx, fmt.Errorf("%v", rec))
		}
	}()
	f()
}

func InvokeWithLog(logger Logger, f func()) {
	defer func() {
		if rec := recover(); rec != nil {
			logger.Error(fmt.Errorf("%v", rec))
		}
	}()
	f()
}

//
//

func InvokeWithErrLogContext(ctx context.Context, logger LoggerWithContext, f func() error) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			logger.Error(ctx, err)
		}
	}()
	err = f()
	if err != nil {
		logger.Error(ctx, err)
	}
	return err
}

func InvokeWithErrLog(logger Logger, f func() error) (err error) {
	defer func() {
		if rec := recover(); rec != nil {
			err = fmt.Errorf("%v", rec)
			logger.Error(err)
		}
	}()
	err = f()
	if err != nil {
		logger.Error(err)
	}
	return err
}
