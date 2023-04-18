package ppts

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"monster-base-backend/internal/service/base"
	"monster-base-backend/pkg"
)

type PromptsService interface {
	base.IBaseRegister
}

type promptService struct {
	redisCli redis.UniversalClient
}

func (p *promptService) Register(e *gin.Engine) {
	prompt := e.Group("/prompt")
	prompt.GET("/list", p.list)
}

func (p promptService) list(ctx *gin.Context) {
	p.redisCli.Set(context.Background(), "hello", "world", -1)
	get := p.redisCli.Get(context.TODO(), "hello")
	ctx.JSON(200, get.Val())
}

func NewPromptsService() PromptsService {
	return &promptService{
		redisCli: pkg.GetRedisClient(),
	}
}
