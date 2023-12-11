package http

import (
	"fmt"
	"time"
)

type HandlerFunc func(ctx *Context)

type FilterBuilder func(next Filter) Filter

type Filter func(ctx *Context)

// 断定它是不是某类型的常见写法
var _ FilterBuilder = MetricsFilterBuilder

func MetricsFilterBuilder(next Filter) Filter {
	return func(ctx *Context) {
		start := time.Now().Nanosecond()
		next(ctx)
		end := time.Now().Nanosecond()
		fmt.Printf("用了 %d 秒", end-start)
	}
}
