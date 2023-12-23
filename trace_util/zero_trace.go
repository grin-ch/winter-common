package trace_util

import "github.com/zeromicro/go-zero/core/trace"

func InitTracer(c trace.Config) {
	trace.StartAgent(c)
}
