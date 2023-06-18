package logger

import (
	"strconv"
)

// CompatibilityLogger follows a common interface that other packages use, so it can be used in places that expect it.
type CompatibilityLogger struct {
	defaultMessage string
	log            Logger
	prefix         string
}

func NewCompatibilityLogger(log Logger, prefix, defaultMessage string) CompatibilityLogger {
	return CompatibilityLogger{
		defaultMessage: defaultMessage,
		log:            log,
		prefix:         prefix,
	}
}

func (l CompatibilityLogger) Debug(args ...interface{}) {
	msg, newArgs := l.extractMessage(args...)
	l.log.Debug(msg, l.argsToData(newArgs...))
}

func (l CompatibilityLogger) Info(args ...interface{}) {
	msg, newArgs := l.extractMessage(args...)
	l.log.Info(msg, l.argsToData(newArgs...))
}

func (l CompatibilityLogger) Warn(args ...interface{}) {
	msg, newArgs := l.extractMessage(args...)
	l.log.Warn(msg, l.argsToData(newArgs...))
}

func (l CompatibilityLogger) Error(args ...interface{}) {
	msg, newArgs := l.extractMessage(args...)
	l.log.Error(msg, l.argsToData(newArgs...))
}

func (l CompatibilityLogger) Fatal(args ...interface{}) {
	msg, newArgs := l.extractMessage(args...)
	l.log.Fatal(msg, l.argsToData(newArgs...))
}

func (l CompatibilityLogger) extractMessage(args ...interface{}) (string, []interface{}) {
	if len(args) == 0 {
		return l.defaultMessage, args
	}
	newArgs := args[1:]
	message, ok := args[0].(string)
	if !ok {
		message = l.defaultMessage
		newArgs = args
	}
	return l.prefix + message, newArgs
}

func (l CompatibilityLogger) argsToData(args ...interface{}) Data {
	data := Data{}
	for i, arg := range args {
		data[strconv.Itoa(i)] = arg
	}
	return data
}
