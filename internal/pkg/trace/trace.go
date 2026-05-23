package trace

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"
)

// Header 请求/响应头中的 traceId 字段名
const Header = "X-Trace-Id"

type ctxKey struct{}

// NewID 生成 traceId：毫秒时间戳 + 4 字节随机十六进制
func NewID() string {
	var b [4]byte
	_, _ = rand.Read(b[:])
	return fmt.Sprintf("%d%s", time.Now().UnixMilli(), hex.EncodeToString(b[:]))
}

// With 将 traceId 写入 context
func With(ctx context.Context, id string) context.Context {
	if ctx == nil || id == "" {
		return ctx
	}
	return context.WithValue(ctx, ctxKey{}, id)
}

// FromContext 从 context 读取 traceId
func FromContext(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	if v, ok := ctx.Value(ctxKey{}).(string); ok {
		return v
	}
	return ""
}
