package middleware

import (
	"gin-basic/internal/pkg/trace"

	"github.com/gin-gonic/gin"
)

const (
	ContextKeyTraceID = "trace_id"
	HeaderTraceID     = trace.Header
)

// Trace 为每个请求生成/透传 traceId，写入 gin.Context 与 request.Context
func Trace() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(HeaderTraceID)
		if traceID == "" {
			traceID = trace.NewID()
		}
		c.Set(ContextKeyTraceID, traceID)
		c.Header(HeaderTraceID, traceID)

		c.Request = c.Request.WithContext(trace.With(c.Request.Context(), traceID))
		c.Next()
	}
}

// TraceIDFromGin 从 gin 取 traceId（AccessLog、Handler 用）
func TraceIDFromGin(c *gin.Context) string {
	if v, ok := c.Get(ContextKeyTraceID); ok {
		if s, ok := v.(string); ok && s != "" {
			return s
		}
	}
	return trace.FromContext(c.Request.Context())
}
