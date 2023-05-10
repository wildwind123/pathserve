package handler

import (
	"net/http"
	"pathserve/internal/config"

	"pathserve/internal/paramslist"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Api struct {
	Config    *config.Config
	Logger    *zap.Logger
	ParamList *paramslist.ParamsList
}

type Info struct {
	Params []paramslist.Param `json:"params"`
	Config *config.Config
}

func (api *Api) Handle(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Param(SegmentKey) == "/info" {
			info, err := api.Info()
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
			c.JSON(http.StatusOK, info)
		}

	}
}

func (api *Api) Info() (Info, error) {
	info := Info{
		Config: api.Config,
	}
	paramList := api.ParamList.GetParamList()
	for i := range paramList {
		info.Params = append(info.Params, paramList[i])
	}

	return info, nil
}
