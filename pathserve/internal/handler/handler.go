package handler

import (
	"fmt"
	"os"
	"path/filepath"
	"pathserve/internal/config"
	"pathserve/internal/helper"
	"pathserve/internal/paramslist"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const SegmentKey = "segment"

// import
func GetHandler(c *gin.Context, paramsList *paramslist.ParamsList, config *config.Config, logger *zap.Logger) (gin.HandlerFunc, error) {
	lastDomain, err := helper.GetLastDomain(c.Request.Host)
	if err != nil {
		return nil, err
	}
	if lastDomain == "api" {
		api := Api{
			Config:    config,
			Logger:    logger,
			ParamList: paramsList,
		}
		return api.Handle(c), nil
	}

	val, ok := config.HandlerConfigs[lastDomain]
	if !ok {
		return nil, fmt.Errorf("config with name %s, does not exist", lastDomain)
	}
	switch val.Handler {
	case "index":
		index := Index{
			Config:    config,
			Logger:    logger,
			ParamList: paramsList,
		}
		return index.Handle(c), nil
	case "vite":
		vite := Vite{
			Config:    config,
			Logger:    logger,
			ParamList: paramsList,
		}
		return vite.Handle(c), nil
	}

	return nil, fmt.Errorf("handler with name %s, does not implement", val.Handler)
}

func GetFilePath(cfg *config.Config, host string, segment string, publicDir string) (string, bool, error) {
	configName, err := helper.GetLastDomain(host)
	if err != nil {
		return "", false, err
	}
	handlerConfig := cfg.HandlerConfigs[configName]
	workDir := handlerConfig.WorkDir
	if !filepath.IsAbs(workDir) {
		workDir = filepath.Join(filepath.Dir(cfg.ConfigPath), workDir)
	}

	filePath := filepath.Join(workDir, segment)

	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return filePath, true, nil
	}

	if !filepath.IsAbs(publicDir) {
		publicDir = filepath.Join(filepath.Dir(cfg.ConfigPath), publicDir)
	}
	filePath = filepath.Join(publicDir, segment)
	if _, err := os.Stat(filePath); !os.IsNotExist(err) {
		return filePath, true, nil
	}
	return filePath, false, nil
}
