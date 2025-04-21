package safe

import (
	"context"

	"github.com/av-ugolkov/gopkg/logger"
)

func Go(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("safe.SafeGo: %v", err)
			}
		}()

		fn()
	}()
}

func GoCtx(ctx context.Context, fn func(ctx context.Context)) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("safe.SafeGo: %v", err)
			}
		}()

		fn(ctx)
	}()
}

func GoCh[T any](ctx context.Context, fn func(ctx context.Context) (T, error)) (chan T, chan error) {
	dataCh := make(chan T)
	errCh := make(chan error)

	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("safe.SafeGo: %v", err)
			}
		}()

		data, err := fn(ctx)
		dataCh <- data
		errCh <- err
	}()

	return dataCh, errCh
}
