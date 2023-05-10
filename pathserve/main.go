package main

import (
	"flag"
	"net/http"
	"path/filepath"
	"pathserve/internal/config"
	"pathserve/internal/logger"
	"pathserve/internal/paramslist"

	"pathserve/internal/handler"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func main() {
	logger := logger.GetLogger()
	defer logger.Sync()
	configFilePath := flag.String("config", "pathserve.config.yaml", "path of config")
	flag.Parse()
	if configFilePath == nil || *configFilePath == "" {
		logger.Fatal("config is required")
	}
	cfg, err := config.GetConfig(*configFilePath)
	if err != nil {
		logger.Fatal("cant get config", zap.Error(err))
	}

	paramsList := paramslist.Get(cfg, logger)
	for key, val := range cfg.HandlerConfigs {
		if val.WatchDir == "" {
			continue
		}
		watchDirPath := val.WorkDir
		if !filepath.IsAbs(val.WatchDir) {
			watchDirPath = filepath.Join(filepath.Dir(cfg.ConfigPath), val.WatchDir)
		}
		err := paramsList.SetParamsList(watchDirPath, []string{key})
		if err != nil {
			logger.Fatal("unbable to paramsList.SetParamsList", zap.Error(err))
		}
		paramsList.WatchDirAndAddParam(watchDirPath, []string{key}...)

	}
	for key := range cfg.HostParams {
		paramsList.AddParam(key, paramslist.Param{
			Path:          cfg.HostParams[key],
			Key:           key,
			HandlerConfig: "from_config",
		})
	}

	logger.Info("config loaded", zap.Any("config", cfg))

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	router.GET("/*"+handler.SegmentKey, func(c *gin.Context) {

		segment := c.Param(handler.SegmentKey)
		if segment == "/favicon.ico" {
			return
		}
		reqHandler, err := handler.GetHandler(c, paramsList, cfg, logger)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		reqHandler(c)
	})
	err = router.Run(cfg.Server.Host)
	if err != nil {
		logger.Error("unable to start server", zap.Error(err))
	}
}
