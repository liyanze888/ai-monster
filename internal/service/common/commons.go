package common

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/lqs/sqlingo"
	dsl "monster-base-backend/internal/generated/dsl"
	"monster-base-backend/internal/service/base"
	"monster-base-backend/pkg"
	"net/http"
	"os"
	"strconv"
)

const (
	ArtType     = 1
	ChatGptType = 2
)

const (
	CategorySubType = 1
)

type CommonsService interface {
	base.IBaseRegister
	GetDict() map[int32]map[int32]map[string]map[string]int64
}

type commonsService struct {
	redisCli redis.UniversalClient
	db       sqlingo.Database
	// map[Type]map[SubType][Name]id
	dict map[int32]map[int32]map[string]map[string]int64
}

func (c *commonsService) GetDict() map[int32]map[int32]map[string]map[string]int64 {
	return c.dict
}

func (c *commonsService) Register(e *gin.Engine) {
	commons := e.Group("/commons")
	commons.GET("/listCategory", c.listCategory)
}

func (c *commonsService) listCategory(ctx *gin.Context) {
	TypeStr := ctx.DefaultQuery("type", "")
	SubTypeStr := ctx.DefaultQuery("subType", "")
	if TypeStr == "" {
		TypeStr = "0"
	}

	Type, err := strconv.ParseInt(TypeStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error request")
		return
	}
	if Type == 0 {
		ctx.JSON(http.StatusBadRequest, "error request")
		return
	}

	SubType, err := strconv.ParseInt(SubTypeStr, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "error request")
		return
	}
	if Type == 0 {
		ctx.JSON(http.StatusBadRequest, "error request")
		return
	}

	data := make([]string, 0)

	if _, ok := c.dict[int32(Type)]; !ok {
		ctx.JSON(http.StatusInternalServerError, "data not assist")
		return
	}

	if _, ok := c.dict[int32(Type)][int32(SubType)]; !ok {
		ctx.JSON(http.StatusInternalServerError, "data not assist")
		return
	}

	if _, ok := c.dict[int32(Type)][int32(SubType)]["en"]; !ok {
		ctx.JSON(http.StatusInternalServerError, "data not assist")
		return
	}

	for key, _ := range c.dict[int32(Type)][int32(SubType)]["en"] {
		data = append(data, key)
	}
	rlt := make(map[string]interface{})
	rlt["data"] = data
	ctx.JSON(http.StatusOK, rlt)
	return

}
func (c *commonsService) init() {
	cursor, err := c.db.SelectFrom(dsl.Dict).
		Where(dsl.Dict.TypeId.In(ArtType, ChatGptType)).
		FetchCursor()
	if err != nil {
		os.Exit(1)
	}

	dicts := make(map[int32]map[int32]map[string]map[string]int64)
	var dict *dsl.DictModel
	for cursor.Next() {
		err := cursor.Scan(&dict)
		if err != nil {
			os.Exit(1)
		}
		if _, ok := dicts[dict.TypeId]; !ok {
			dicts[dict.TypeId] = make(map[int32]map[string]map[string]int64)
		}
		if _, ok := dicts[dict.TypeId][dict.SubType]; !ok {
			dicts[dict.TypeId][dict.SubType] = make(map[string]map[string]int64)
		}
		if _, ok := dicts[dict.TypeId][dict.SubType][dict.Locale]; !ok {
			dicts[dict.TypeId][dict.SubType][dict.Locale] = make(map[string]int64)
		}
		dicts[dict.TypeId][dict.SubType][dict.Locale][dict.Content] = dict.Id
	}
	c.dict = dicts
}

func NewCommonsService() CommonsService {
	svc := &commonsService{
		redisCli: pkg.GetRedisClient(),
		db:       pkg.GetDatabase(),
	}
	svc.init()
	return svc
}
