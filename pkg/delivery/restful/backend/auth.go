package backend

import (
	internalGin "boyi/pkg/Infra/gin"
	"boyi/pkg/model/vo"

	"github.com/gin-gonic/gin"
)

// @Summary     登入
// @Description  loginUser.
// @Tags         login
// @Produce      json
// @Param        book  body     vo.LoginReq  true  "Book JSON"
// @Success      200   {object}  claims.Claims
// @Router       /b/apis/v1/auth [post]
func (h *handler) Login(c *gin.Context) {
	var (
		in vo.LoginReq
	)

	if err := c.ShouldBindJSON(&in); err != nil {
		internalGin.RespError(c, err)
		return
	}

	result, err := h.authSvc.Login(c.Request.Context(), in)
	if err != nil {
		internalGin.RespError(c, err)
		return
	}
	internalGin.Ok(c, result)
}
