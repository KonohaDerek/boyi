package platform

import (
	"boyi/pkg/infra/errors"
	internalGin "boyi/pkg/infra/gin"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetFile(c *gin.Context) {
	result, err := h.s3Svc.GetFile(c.Request.Context(), c.Param("xid"))
	if err != nil {
		internalGin.RespError(c, err)
		return
	}
	_, err = c.Writer.WriteString(result)
	if err != nil {
		internalGin.RespError(c, errors.ErrInternalError)
		return
	}
}
