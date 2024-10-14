package tracer

import (
	"errors"
	"fmt"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"runtime"
)

func Error(err error) error {
	pc := make([]uintptr, 15)
	n := runtime.Callers(2, pc)
	frames := runtime.CallersFrames(pc[:n])
	frame, _ := frames.Next()
	return fmt.Errorf("%s:%d: %w", frame.Function, frame.Line, err)
}

func GetMessageFromStackTrace(err error) string {
	beforeOneError := errors.New("")
	for {
		err = errors.Unwrap(err)
		if err == nil {
			return beforeOneError.Error()
		}
		beforeOneError = err
	}
}

func RecordErrorOtel(span trace.Span, err error) {
	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	}
}
