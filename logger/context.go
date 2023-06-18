package logger

import (
	"context"
)

type key int

const ctxKey key = 0

// FromContext returns a Logger from the given context.Context. If there is no
// attached logger, then this will just return a new Logger instance.
func FromContext(ctx context.Context) Logger {
	var log Logger
	log, ok := ctx.Value(ctxKey).(Logger)
	if !ok {
		log = New()
	}
	return log
}

func (log Logger) WithContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, ctxKey, log)
}
