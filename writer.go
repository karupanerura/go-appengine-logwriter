package aelogwriter

import (
	"context"
)

type Logger func(ctx context.Context, format string, args ...interface{})

type Writer struct {
	ctx context.Context
	l   Logger
}

func NewWriter(ctx context.Context, logger func(ctx context.Context, format string, args ...interface{})) *Writer {
	return &Writer{ctx: ctx, l: Logger(logger)}
}

func (w *Writer) Write(p []byte) (n int, err error) {
	w.l(w.ctx, string(p))
	return len(p), nil
}
