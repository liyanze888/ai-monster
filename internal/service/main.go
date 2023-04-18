package service

import (
	"github.com/gin-gonic/gin"
	"monster-base-backend/internal/service/ppts"
)

type StartService struct {
	ctx *gin.Engine
}

func (s *StartService) Start() {
	ppts.NewPromptsService().Register(s.ctx)
}

func NewStartService(ctx *gin.Engine) *StartService {
	return &StartService{
		ctx: ctx,
	}
}
