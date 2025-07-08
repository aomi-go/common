package logger

type traceId struct{}

var (
	TraceIdCtxKey = traceId{}
	EmptyLogger   = emptyLogger{}
)
