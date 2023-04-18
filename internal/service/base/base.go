package base

import "github.com/gin-gonic/gin"

type IBaseRegister interface {
	Register(e *gin.Engine)
}
