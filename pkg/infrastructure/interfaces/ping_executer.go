package interfaces

import "context"

type wrapPingContextExecuter struct {
	pinger PingExecutor
}

func (w *wrapPingContextExecuter) PingContext(ctx context.Context) error {
	var err error
	done := make(chan struct{})
	go func() {
		defer close(done)
		err = w.pinger.Ping()
	}()
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-done:
		return err
	}
}
