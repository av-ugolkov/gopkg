package safe

import "github.com/av-ugolkov/gopkg/logger"

func SafeGo(fn func()) {
	go func() {
		defer func() {
			if err := recover(); err != nil {
				logger.Errorf("safe.SafeGo: %v", err)
			}
		}()

		fn()
	}()
}
