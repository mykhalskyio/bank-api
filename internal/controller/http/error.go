package http

import "github.com/gin-gonic/gin"

type response struct {
	Error string `json:"error"`
}

func errorResponse(ctx *gin.Context, code int, err error) {
	ctx.AbortWithStatusJSON(code, response{Error: err.Error()})
}
