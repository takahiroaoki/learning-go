package trace

import (
	"bytes"
	"testing"
)

func TestNew(t *testing.T) {
	var buf bytes.Buffer
	tracer := New(&buf)
	if tracer == nil {
		t.Error("nil from New")
	} else {
		tracer.Trace("Hello, package 'trace'")
		if buf.String() != "Hello, package 'trace'\n" {
			t.Errorf("The wrong string '%s' was output", buf.String())
		}
	}
}

func TestOff(t *testing.T) {
	silentTracer := Off()
	silentTracer.Trace("Hello")
}
