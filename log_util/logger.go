package log_util

import (
	"fmt"
	"runtime"
)

func Tracef(format string, args ...any) {
	Logger.Tracef(format, args...)
}

func Debugf(format string, args ...any) {
	Logger.Debugf(format, args...)
}

func Infof(format string, args ...any) {
	Logger.Infof(format, args...)
}

func Warnf(format string, args ...any) {
	Logger.Warnf(format, args...)
}

func Errorf(format string, args ...any) {
	Logger.Errorf(format, args...)
}

func Fatalf(format string, args ...any) {
	Logger.Fatalf(format, args...)
}

func Panicf(format string, args ...any) {
	Logger.Panicf(format, args...)
}

const skip = 9

func callerPrettifier(f *runtime.Frame) (string, string) {
	if pc, file, line, ok := runtime.Caller(skip); ok {
		return runtime.FuncForPC(pc).Name(), fmt.Sprintf("%s:%d", file, line)
	}
	return f.Func.Name(), f.File
}
