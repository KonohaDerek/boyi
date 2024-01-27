package backend

import (
	"github.com/gin-gonic/gin"
)

// Register is register rest api router bind handle
func (h *handler) Register(g *gin.RouterGroup) {
	withVerify := g.Use(h.authSvc.SetClaims())
	withVerify.GET("/files/:xid", h.GetFile)

	withVerify.GET("/menu/tree", h.GetMenuTree)

	h.RegisterAuth(g)
}

func (h *handler) RegisterAuth(g *gin.RouterGroup) {
	g.POST("/auth/Login", h.Login)

}
