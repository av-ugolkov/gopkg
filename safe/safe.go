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

func GoErr(fn func() error) <-chan error {
	chErr := make(chan error, 1)

	go func() {
		defer func() {
			close(chErr)
			if err := recover(); err != nil {
				logger.Errorf("safe.SafeGo: %v", err)
			}
		}()

		chErr <- fn()
	}()
	return chErr
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

func GoCtxErr(ctx context.Context, fn func(ctx context.Context) error) <-chan error {
	chErr := make(chan error, 1)

	go func() {
		defer func() {
			close(chErr)
			if err := recover(); err != nil {
				logger.Errorf("safe.SafeGo: %v", err)
			}
		}()

		chErr <- fn(ctx)
	}()

	return chErr
}

func GoCh[T any](ctx context.Context, fn func(ctx context.Context) (T, error)) (chan T, chan error) {
	dataCh := make(chan T)
	errCh := make(chan error)

	go func() {
		defer func() {
			close(dataCh)
			close(errCh)

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
