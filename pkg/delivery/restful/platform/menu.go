package platform

import (
	"boyi/internal/claims"

	internalGin "boyi/pkg/infra/gin"

	"github.com/gin-gonic/gin"
)

// @Summary 提供所有Menu
// @Produce  json
// @success 200 {object} []dto.Menu "desc"
// @Router /b/menu/tree [get]
func (h *handler) GetMenuTree(c *gin.Context) {
	ret := h.menuSvc.GetParsedMenuTree(claims.Claims{})
	internalGin.Ok(c, ret)
}
