package service

import (
	"github.com/gin-gonic/gin"
	"monster-base-backend/internal/service/common"
	"monster-base-backend/internal/service/ppts"
)

type StartService struct {
	ctx *gin.Engine
}

func (s *StartService) Start() {
	commonSvc := common.NewCommonsService()
	ppts.NewPromptsService(commonSvc).Register(s.ctx)
	commonSvc.Register(s.ctx)
}

func NewStartService(ctx *gin.Engine) *StartService {
	return &StartService{
		ctx: ctx,
	}
}
