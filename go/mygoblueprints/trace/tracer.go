package trace

import (
	"fmt"
	"io"
)

type Tracer interface {
	Trace(...any)
}

func New(w io.Writer) Tracer {
	return &tracer{out: w}
}

type tracer struct {
	out io.Writer
}

func (t *tracer) Trace(a ...any) {
	fmt.Fprint(t.out, a...)
	fmt.Fprintln(t.out)
}

type nilTracer struct{}

func (t *nilTracer) Trace(a ...any) {}

func Off() Tracer {
	return &nilTracer{}
}
