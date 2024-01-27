package http

import "github.com/gin-gonic/gin"

type Handler interface {
	Register(g *gin.RouterGroup)
	Version() string
	BeforeRun() error
}