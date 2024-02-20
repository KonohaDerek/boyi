package gin

import (
	"net/http"

	"boyi/pkg/infra/ctxutil"
	"boyi/pkg/infra/errors"

	"github.com/rs/zerolog/log"

	"github.com/gin-gonic/gin"
)

// APIResponse is api response
type APIResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
	TraceID string      `json:"traceID,omitempty"`
}

// Empty set gin result. HTTP status codes is 200
func Empty(c *gin.Context) {
	resp := APIResponse{
		Code: 0,
		Data: "",
	}

	TraceID := ctxutil.GetTraceIDFromContext(c.Request.Context())
	resp.TraceID = TraceID

	c.JSON(http.StatusOK, resp)
}

// Ok set gin result. HTTP status codes is 200
func Ok(c *gin.Context, data interface{}) {
	resp := APIResponse{
		Code:    0,
		Data:    data,
		Message: "success",
	}

	TraceID := ctxutil.GetTraceIDFromContext(c.Request.Context())
	resp.TraceID = TraceID

	c.JSON(http.StatusOK, resp)
}

// BadRequest set gin result. HTTP status codes is 400
func BadRequest(c *gin.Context, message string) {
	resp := APIResponse{
		Code:    400,
		Message: message,
	}

	TraceID := ctxutil.GetTraceIDFromContext(c.Request.Context())
	resp.TraceID = TraceID

	c.JSON(http.StatusBadRequest, resp)
}

func RespError(c *gin.Context, err error) {
	if err == nil {
		c.JSON(http.StatusOK, "")
		return
	}
	req := c.Request

	logFields := map[string]interface{}{
		"host":       req.Host,
		"uri":        req.RequestURI,
		"method":     req.Method,
		"path":       req.URL.Path,
		"referer":    req.Referer(),
		"user_agent": req.UserAgent(),
		"status":     c.Writer.Status(),
		"bytes_in":   req.ContentLength,
		"bytes_out":  c.Writer.Size(),
		"access_log": true,
	}

	httpError := errors.ConvertToHttpError(err)

	logger := log.With().Fields(logFields).Logger()
	switch {
	case httpError.Status >= http.StatusInternalServerError:
		logger.Error().Msgf("%+v", err)
	case httpError.Status >= http.StatusBadRequest:
		logger.Warn().Msgf("%+v", err)
	default:
		logger.Error().Msgf("handler error: %+v", err)
	}

	c.JSON(httpError.Status, httpError)
}
