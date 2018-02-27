package aelogwriter

import (
	"context"
)

type logger func(ctx context.Context, format string, args ...interface{})

// Writer is io.Writer compatible log writer for Google App Engine.
type Writer struct {
	ctx context.Context
	l   logger
}

// NewWriter creates new Writer.
// It needs appengine's context and appengine/log's log function.
func NewWriter(ctx context.Context, l func(ctx context.Context, format string, args ...interface{})) *Writer {
	return &Writer{ctx: ctx, l: logger(l)}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.l(w.ctx, string(p))
	return len(p), nil
}
