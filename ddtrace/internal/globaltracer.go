package internal // import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace/internal"

import "gopkg.in/DataDog/dd-trace-go.v1/ddtrace"

// GlobalTracer holds the currently active tracer. It's "zero value" should
// always be the NoopTracer.
var GlobalTracer ddtrace.Tracer = &NoopTracer{}

// Testing is set to true when the mock tracer is active. It usually signifies that we are in a test
// environment. This value is used by tracer.Start to prevent overriding the GlobalTracer in tests.
var Testing = false

var _ ddtrace.Tracer = (*NoopTracer)(nil)

// NoopTracer is an implementation of ddtrace.Tracer that is a no-op.
type NoopTracer struct{}

// StartSpan implements ddtrace.Tracer.
func (NoopTracer) StartSpan(operationName string, opts ...ddtrace.StartSpanOption) ddtrace.Span {
	return NoopSpan{}
}

// SetServiceInfo implements ddtrace.Tracer.
func (NoopTracer) SetServiceInfo(name, app, appType string) {}

// Extract implements ddtrace.Tracer.
func (NoopTracer) Extract(carrier interface{}) (ddtrace.SpanContext, error) {
	return NoopSpanContext{}, nil
}

// Inject implements ddtrace.Tracer.
func (NoopTracer) Inject(context ddtrace.SpanContext, carrier interface{}) error { return nil }

// Stop implements ddtrace.Tracer.
func (NoopTracer) Stop() {}

var _ ddtrace.Span = (*NoopSpan)(nil)

// NoopSpan is an implementation of ddtrace.Span that is a no-op.
type NoopSpan struct{}

// SetTag implements ddtrace.Span.
func (NoopSpan) SetTag(key string, value interface{}) {}

// SetOperationName implements ddtrace.Span.
func (NoopSpan) SetOperationName(operationName string) {}

// BaggageItem implements ddtrace.Span.
func (NoopSpan) BaggageItem(key string) string { return "" }

// SetBaggageItem implements ddtrace.Span.
func (NoopSpan) SetBaggageItem(key, val string) {}

// Finish implements ddtrace.Span.
func (NoopSpan) Finish(opts ...ddtrace.FinishOption) {}

// Tracer implements ddtrace.Span.
func (NoopSpan) Tracer() ddtrace.Tracer { return NoopTracer{} }

// Context implements ddtrace.Span.
func (NoopSpan) Context() ddtrace.SpanContext { return NoopSpanContext{} }

var _ ddtrace.SpanContext = (*NoopSpanContext)(nil)

// NoopSpanContext is an implementation of ddtrace.SpanContext that is a no-op.
type NoopSpanContext struct{}

// ForeachBaggageItem implements ddtrace.SpanContext.
func (NoopSpanContext) ForeachBaggageItem(handler func(k, v string) bool) {}