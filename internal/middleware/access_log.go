package middleware

import (
	"time"

	"gin-basic/utils/logger"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		rawQuery := c.Request.URL.RawQuery

		c.Next()

		fields := logrus.Fields{
			"method":         c.Request.Method,
			"path":           path,
			"query":          rawQuery,
			"status":         c.Writer.Status(),
			"latency_ms":     time.Since(start).Milliseconds(),
			"client_ip":      c.ClientIP(),
			"user_agent":     c.Request.UserAgent(),
			"content_length": c.Request.ContentLength,
			"response_size":  c.Writer.Size(),
		}
		if tid := TraceIDFromGin(c); tid != "" {
			fields["trace_id"] = tid
		}
		if rid := c.GetHeader("X-Request-Id"); rid != "" {
			fields["request_id"] = rid
		}
		// 鉴权成功后可在 TokenAuth 里 Set("mobile", ...)，这里带上（无则忽略）
		if userID, ok := c.Get(ContextKeyUserID); ok {
			fields["user_id"] = userID
		}
		if userName, ok := c.Get(ContextKeyUserName); ok {
			fields["user_name"] = userName
		}
		if userRole, ok := c.Get(ContextKeyUserRole); ok {
			fields["user_role"] = userRole
		}

		logger.C(c.Request.Context()).WithFields(fields).Info("http access")
	}
}
