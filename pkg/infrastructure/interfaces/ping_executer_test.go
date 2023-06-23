package interfaces

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type testPinger struct {
	t   time.Duration
	err error
}

func (t *testPinger) Ping() error {
	time.Sleep(t.t)
	return t.err
}

func TestPingExecutor(t *testing.T) {
	var (
		tt     *testPinger
		ctx    context.Context
		err    error
		cancel context.CancelFunc
		p      *wrapPingContextExecuter
	)
	// 1msスリープ
	tt = &testPinger{t: 1 * time.Millisecond}
	// 10msで timeout する context
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	p = &wrapPingContextExecuter{pinger: tt}
	err = p.PingContext(ctx)
	assert.NoError(t, err)
	assert.Nil(t, err)

	// 1msスリープしてエラーを返す
	testErr := errors.New("")
	tt = &testPinger{t: 1 * time.Millisecond, err: testErr}
	// 10msで timeout する context
	ctx, cancel = context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	p = &wrapPingContextExecuter{pinger: tt}
	err = p.PingContext(ctx)
	assert.Error(t, err)
	assert.ErrorIs(t, err, testErr)

	// 10msスリープ
	tt = &testPinger{t: 10 * time.Millisecond}
	// 1msで timeout する context
	ctx, cancel = context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()
	p = &wrapPingContextExecuter{pinger: tt}
	err = p.PingContext(ctx)
	assert.Error(t, err)
	assert.ErrorIs(t, err, context.DeadlineExceeded)
}
