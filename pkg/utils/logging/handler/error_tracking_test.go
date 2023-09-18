package handler

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestErrorTracking(t *testing.T) {
	assert := assert.New(t)
	buf := bytes.Buffer{}
	handler := slog.NewTextHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	h := NewErrorTracking(handler)
	log := slog.New(h)
	log.Error("aaa")
	assert.Contains(buf.String(), "msg=aaa")
}

func TestIgnoreErrorTracking(t *testing.T) {
	assert := assert.New(t)
	buf := bytes.Buffer{}
	handler := slog.NewTextHandler(&buf, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})
	h := NewErrorTracking(handler)
	log := slog.New(h)
	log.With(slog.Any("aaa", "bbb")).With(IgnoreTracing).Error("aaa")
	assert.Equal(buf.String(), "")
}
