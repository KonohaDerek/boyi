package ctxutil

import (
	"context"
	"fmt"
	"strings"

	"github.com/mileusna/useragent"
	"github.com/rs/xid"
)

type ctxKey string

const (
	// XTraceID request id
	XTraceID   ctxKey = "x-trace-id"
	XOrigin    ctxKey = "x-origin"
	XRealIP    ctxKey = "x-real-ip"
	XDeviceID  ctxKey = "x-device-id"
	XUserAgent ctxKey = "x-user-agent"
)

func (ctxKey) String() string {
	return "x-trace-id"
}

// GetTraceIDFromContext get trace-id from context
func GetTraceIDFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XTraceID).(string)
	if !ok {
		return ""
	}
	return v
}

func TraceIDWithContext(ctx context.Context) context.Context {
	traceID := GetTraceIDFromContext(ctx)
	if traceID == "" {
		traceID = NewTraceID()
	}

	return ContextWithXTraceID(ctx, traceID)
}

// ContextWithXTraceID returns a context.Context with given trace-id value.
func ContextWithXTraceID(ctx context.Context, traceID string) context.Context {
	return context.WithValue(ctx, XTraceID, traceID)
}

func NewTraceID() string {
	return xid.New().String()
}

// GetRealIPFromContext get trace-id from context
func GetRealIPFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XRealIP).(string)
	if !ok {
		return ""
	}
	return v
}

// ContextWithXRealIP returns a context.Context with given trace-id value.
func ContextWithXRealIP(ctx context.Context, readIP string) context.Context {
	return context.WithValue(ctx, XRealIP, readIP)
}

// GetDeviceUIDFromContext get device-id from context
func GetDeviceUIDFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XDeviceID).(string)
	if !ok || v == "" {
		return xid.New().String()
	}
	return v
}

// ContextWithXDeviceID returns a context.Context with given device-id value.
func ContextWithXDeviceID(ctx context.Context, deviceID string) context.Context {
	return context.WithValue(ctx, XDeviceID, deviceID)
}

// GetUserAgentFromContext get user-agent from context
func GetUserAgentFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XUserAgent).(string)
	if !ok {
		return ""
	}
	return v
}

// ContextWithXUserAgent returns a context.Context with given user-agent value.
func ContextWithXUserAgent(ctx context.Context, userAgent string) context.Context {
	return context.WithValue(ctx, XUserAgent, userAgent)
}

// GetOriginFromContext get host from context
func GetOriginFromContext(ctx context.Context) string {
	v, ok := ctx.Value(XOrigin).(string)
	if !ok {
		return ""
	}
	return v
}

// ContextWithXOrigin returns a context.Context with given host value.
func ContextWithXOrigin(ctx context.Context, origin string) context.Context {
	return context.WithValue(ctx, XOrigin, origin)
}

// ContextWithXRealIP returns a context.Context with given trace-id value.
func ContextWithToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, "Authorization", token)
}

func GetTokenFromContext(ctx context.Context) string {
	v, ok := ctx.Value("Authorization").(string)
	if !ok {
		return ""
	}
	return v
}

func GetDeviceOSFronContext(ctx context.Context) string {
	userAgent := GetUserAgentFromContext(ctx)
	if userAgent == "" {
		return ""
	}

	ua := useragent.Parse(userAgent)
	fmt.Println()
	fmt.Println(ua.String)
	fmt.Println(strings.Repeat("=", len(ua.String)))
	fmt.Println("Name:", ua.Name, "v", ua.Version)
	fmt.Println("OS:", ua.OS, "v", ua.OSVersion)
	fmt.Println("Device:", ua.Device)
	return fmt.Sprintf("%s %s, %s", ua.OS, ua.OSVersion, ua.Device)
}

func ContextWithMerchantID(ctx context.Context, merchantID uint64) context.Context {
	return context.WithValue(ctx, "MerchantID", merchantID)
}

func GetMerchantIDFromContext(ctx context.Context) uint64 {
	v, ok := ctx.Value("MerchantID").(uint64)
	if !ok {
		return 0
	}
	return v
}
