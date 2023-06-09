package ppts

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	dsl "monster-base-backend/internal/generated/dsl"
	"monster-base-backend/internal/service/base"
	"monster-base-backend/internal/service/common"
	"monster-base-backend/pkg"
	"net/http"
	"time"
)

type PromptsService interface {
	base.IBaseRegister
}

type promptService struct {
	redisCli  redis.UniversalClient
	repo      PromptsRepository
	commonSvc common.CommonsService
}

type (
	ListSimpleRequest struct {
		base.PageBaseRequest
		Type int `json:"type"`
	}

	SimpleModel struct {
		Id       int64    `json:"id"`
		Title    string   `json:"title"`
		Images   []string `json:"images"`
		Model    string   `json:"model"`
		Value    int      `json:"value"`
		ValueNum int      `json:"valueNum"`
		LikeNum  int      `json:"likeNum"`
		NickName string   `json:"nickName"`
	}

	ListSimpleResponse struct {
		base.PageBaseResponse
		Models []SimpleModel `json:"data"`
	}

	AddModel struct {
		Type         int8               `json:"type"`
		ArtModel     ArtPromptModel     `json:"artModel"`
		ChatgptModel ChatgptPromptModel `json:"chatgptModel"`
	}
)

func (p *promptService) Register(e *gin.Engine) {
	prompt := e.Group("/prompts")
	//推荐
	prompt.POST("/listRecommends", p.listRecommends)
	//增加 art
	prompt.POST("/addPrompt", p.add)
}

var ArtModelRef = map[string]int32{
	"Stable Diffusion": 1,
}

type ArtContentJsonModel struct {
	Title            string `json:"title"`
	Content          string `json:"content"`
	Seed             string `json:"seed"`
	GuidanceScale    string `json:"guidanceScale"`
	Steps            string `json:"steps"`
	Sampler          string `json:"sampler"`
	ExtraInstruction string `json:"extraInstruction"`
}

var ArtVersionRef = map[string]int32{}

func (p *promptService) addArt(model AddModel) error {
	id, err := pkg.GenerateInt64()
	if err != nil {
		return err
	}
	timestamp := time.Now().UnixMilli()
	jsonModel := &ArtContentJsonModel{
		Title:            model.ArtModel.Title,
		Content:          model.ArtModel.Content,
		Seed:             model.ArtModel.Seed,
		GuidanceScale:    model.ArtModel.GuidanceScale,
		Steps:            model.ArtModel.Steps,
		Sampler:          model.ArtModel.Sampler,
		ExtraInstruction: model.ArtModel.ExtraInstruction,
	}

	marshal, err := json.Marshal(jsonModel)
	if err != nil {
		return err
	}
	prompt := &dsl.PromptsModel{
		Id:          id,
		UserId:      model.ArtModel.UserId,
		PType:       common.ArtType,
		Title:       model.ArtModel.Title,
		Author:      "",
		Model:       ArtModelRef[model.ArtModel.Model],
		Version:     ArtVersionRef[model.ArtModel.Version],
		ContentJson: string(marshal),
		CreatedTime: timestamp,
		UpdatedTime: timestamp,
	}

	categories := make([]*dsl.PromptsIndexModel, len(model.ArtModel.Category))
	for idx, category := range model.ArtModel.Category {
		c := &dsl.PromptsIndexModel{
			PromptsId:   id,
			Model:       ArtModelRef[model.ArtModel.Model],
			Category:    int32(p.commonSvc.GetDict()[common.ArtType][common.CategorySubType]["en"][category]),
			Version:     ArtVersionRef[model.ArtModel.Version],
			Value:       0,
			ValueNum:    0,
			LikeNum:     0,
			CreatedTime: timestamp,
			UpdatedTime: timestamp,
		}
		categories[idx] = c
	}

	medias := make([]*dsl.MediasModel, 0)
	for _, image := range model.ArtModel.Medias {
		mediaId, err := pkg.GenerateInt64()
		if err != nil {
			return err
		}
		medias = append(medias, &dsl.MediasModel{
			Id:          mediaId,
			TargetId:    id,
			MediaId:     image,
			CreatedTime: timestamp,
			UpdatedTime: timestamp,
		})
	}
	return p.repo.Add(prompt, categories, medias)
}

func (p *promptService) add(ctx *gin.Context) {
	var model AddModel
	err := ctx.Bind(&model)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error request")
		return
	}

	if model.Type == common.ArtType {
		err = p.addArt(model)
	} else if model.Type == common.ChatGptType {

	}
	if err != nil {
		ctx.JSON(http.StatusBadRequest, map[string]string{
			"errorInfo": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, "ok")
}

func (p promptService) listRecommends(ctx *gin.Context) {
	p.redisCli.Set(context.Background(), "hello", "world", -1)
	var listRequest ListSimpleRequest
	err := ctx.Bind(&listRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error request")
		return
	}

	models := make([]SimpleModel, 0)
	for i := 0; i < 2; i++ {
		models = append(models, SimpleModel{
			Id:    int64(i) + 1,
			Title: fmt.Sprintf("title %d", i),
			Images: []string{
				fmt.Sprintf("image %d", i*1),
				fmt.Sprintf("image %d", i*2),
				fmt.Sprintf("image %d", i*3),
			},
			Model:    "midjourney",
			Value:    10 * i,
			ValueNum: 199 * i,
			LikeNum:  99 * i,
			NickName: fmt.Sprintf("liyz %d", i),
		})
	}
	response := &ListSimpleResponse{
		PageBaseResponse: base.PageBaseResponse{
			TotalNum:  1,
			TotalPage: 2,
		},
		Models: models,
	}
	ctx.JSON(200, response)
}

func NewPromptsService(commonSvc common.CommonsService) PromptsService {
	svc := &promptService{
		redisCli:  pkg.GetRedisClient(),
		repo:      NewPromptsRepository(),
		commonSvc: commonSvc,
	}
	return svc
}
