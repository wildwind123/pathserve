package handler

import (
	"fmt"
	"net/http"
	"pathserve/internal/config"
	"pathserve/internal/helper"

	"pathserve/internal/paramslist"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Index struct {
	Config    *config.Config
	Logger    *zap.Logger
	ParamList *paramslist.ParamsList
}

func (i *Index) Handle(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := i.Config

		filePath, exist, err := GetFilePath(cfg, c.Request.Host, c.Param(SegmentKey), "")
		if err != nil {
			i.Logger.Error("unable to GetFilePath", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("unable to GetFilePath, %v", err),
			})
			return
		}
		if exist {
			c.File(filePath)
			return
		}

		paramKey, err := helper.GetDomainParamKey(c.Request.Host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		param := i.ParamList.GetParam(paramKey)

		if param == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("param key '%s' does not exist", paramKey),
			})
			return
		}
		c.File(param.Path)
	}
}
