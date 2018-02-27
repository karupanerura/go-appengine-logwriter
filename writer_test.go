package aelogwriter

import (
	"context"
	"fmt"
	"testing"
)

func TestWriter(t *testing.T) {
	// mock
	var outs []string
	ctx := context.Background()
	l := func(ctx context.Context, format string, args ...interface{}) {
		if ctx == nil {
			t.Fatal("ctx should not be nil")
		}
		outs = append(outs, format)
	}
	w := NewWriter(ctx, l)

	fmt.Fprint(w, "foo")
	fmt.Fprint(w, "bar")

	if len(outs) != 2 || outs[0] != "foo" || outs[1] != "bar" {
		t.Errorf("Unexpected outs: %#v", outs)
	}
}
