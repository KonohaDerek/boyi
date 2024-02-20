package gin

import (
	"fmt"
	"net/url"
	"runtime"
	"strings"
	"time"

	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
	"github.com/rs/zerolog/log"
)

// NewRecoverMiddleware handles panic error
func NewRecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// unknown error.  handler status code is 500 series.
				trace := make([]byte, 4096)
				runtime.Stack(trace, true)
				traceID := ctxutil.GetTraceIDFromContext(c.Request.Context())
				logger := log.With().
					Str("url", c.Request.RequestURI).
					Str("stack_error", string(trace)).
					Str("trace_id", traceID).Logger()

				var msg string
				for i := 2; ; i++ {
					_, file, line, ok := runtime.Caller(i)
					if !ok {
						break
					}
					msg += fmt.Sprintf("%s:%d\n", file, line)
				}

				logger.Error().Msgf("http: unknown error: %s", msg)

				httpError := errors.HttpError{
					Code:    errors.ErrInternalError.Code,
					Message: "unknow error",
				}
				c.AbortWithStatusJSON(500, httpError)
			}
		}()
		c.Next()
	}
}

// NewRequestIDMiddleware Default returns the location middleware with default configuration.
func NewRequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.Request.Header.Get(ctxutil.XTraceID.String())
		if traceID == "" {
			traceID = xid.New().String()
		}

		logger := log.With().Str("trace_id", traceID).Logger()
		realIP := c.GetHeader(string(ctxutil.XRealIP))
		realIPSplit := strings.Split(realIP, ",")

		deviceID := c.GetHeader(DeviceID)

		ctx := ctxutil.ContextWithXTraceID(c.Request.Context(), traceID)
		ctx = ctxutil.ContextWithXUserAgent(ctx, c.Request.UserAgent())
		ctx = ctxutil.ContextWithXDeviceID(ctx, deviceID)
		if len(realIPSplit) != 0 {
			ctx = ctxutil.ContextWithXRealIP(ctx, realIPSplit[0])
		}

		url, err := url.Parse(c.Request.Header.Get("Origin"))
		if err != nil {
			ctx = ctxutil.ContextWithXOrigin(ctx, c.Request.Header.Get("Origin"))
		} else {
			ctx = ctxutil.ContextWithXOrigin(ctx, url.Host)
		}

		ctx = logger.WithContext(ctx)

		c.Request = c.Request.WithContext(ctx)

		c.Request.Header.Set(ctxutil.XTraceID.String(), traceID)
		c.Writer.Header().Set(ctxutil.XTraceID.String(), traceID)
		c.Next()
	}
}

func CORS() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATH", "OPTIONS"},
		AllowHeaders: []string{
			HeaderAccept,
			HeaderAcceptEncoding,
			HeaderAuthorization,
			HeaderContentDisposition,
			HeaderContentEncoding,
			HeaderContentLength,
			HeaderContentType,
			HeaderCookie,
			HeaderSetCookie,
			HeaderIfModifiedSince,
			HeaderLastModified,
			HeaderLocation,
			HeaderUpgrade,
			HeaderVary,
			HeaderWWWAuthenticate,
			HeaderXForwardedFor,
			HeaderXForwardedProto,
			HeaderXForwardedProtocol,
			HeaderXForwardedSsl,
			HeaderXUrlScheme,
			HeaderXHTTPMethodOverride,
			HeaderXRealIP,
			HeaderXRequestID,
			HeaderXRequestedWith,
			HeaderServer,
			HeaderOrigin,
			string(ctxutil.XTraceID),
			DeviceID,
		},
		AllowWebSockets: true,
		AllowOriginFunc: func(_ string) bool {
			return true
		},
		AllowBrowserExtensions: true,
		AllowCredentials:       true,
		MaxAge:                 24 * time.Hour,
	})
}

func Swagger() gin.HandlerFunc {
	return func(c *gin.Context) {
		// if strings.Contains(c.Request.URL.Path, "swagger") {
		// 	c.Header("Content-Type", "application/json")
		// 	c.File("./swagger.json")
		// 	c.Abort()
		// }
		c.Next()
	}
}

const (
	HeaderAccept              = "Accept"
	HeaderAcceptEncoding      = "Accept-Encoding"
	HeaderAllow               = "Allow"
	HeaderAuthorization       = "Authorization"
	HeaderContentDisposition  = "Content-Disposition"
	HeaderContentEncoding     = "Content-Encoding"
	HeaderContentLength       = "Content-Length"
	HeaderContentType         = "Content-Type"
	HeaderCookie              = "Cookie"
	HeaderSetCookie           = "Set-Cookie"
	HeaderIfModifiedSince     = "If-Modified-Since"
	HeaderLastModified        = "Last-Modified"
	HeaderLocation            = "Location"
	HeaderUpgrade             = "Upgrade"
	HeaderVary                = "Vary"
	HeaderWWWAuthenticate     = "WWW-Authenticate"
	HeaderXForwardedFor       = "X-Forwarded-For"
	HeaderXForwardedProto     = "X-Forwarded-Protocol"
	HeaderXForwardedProtocol  = "X-Forwarded-Protocol"
	HeaderXForwardedSsl       = "X-Forwarded-Ssl"
	HeaderXUrlScheme          = "X-Url-Scheme"
	HeaderXHTTPMethodOverride = "X-HTTP-Method-Override"
	HeaderXRealIP             = "X-Real-Ip"
	HeaderXRequestID          = "X-Request-ID"
	HeaderXRequestedWith      = "X-Requested-With"
	HeaderServer              = "Server"
	HeaderOrigin              = "Origin"

	DeviceID = "X-Device-ID"
)
