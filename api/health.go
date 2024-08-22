package api

import "github.com/gin-gonic/gin"

func (server *Server) healthCheck(ctx *gin.Context) {
	ctx.String(200, "OK")
}
