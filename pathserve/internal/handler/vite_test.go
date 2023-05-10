package handler

import (
	"fmt"
	"pathserve/internal/config"
	"testing"

	"go.uber.org/zap"
)

func TestXxx(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	cfg := config.Config{
		HandlerConfigs: make(map[string]config.HandlerConfig),
	}
	cfg.HandlerConfigs = map[string]config.HandlerConfig{
		"vite": {
			Params: make(map[string]string),
		},
	}
	cfg.HandlerConfigs["vite"].Params["html_template"] = "template.html"
	cfg.HandlerConfigs["vite"].Params["script_template"] = "script.ts"
	cfg.HandlerConfigs["vite"].Params["auto_gen_dir"] = "autogen"
	cfg.HandlerConfigs["vite"].Params["dir_public"] = ""

	vite := Vite{
		Logger: logger,
		Config: &cfg,
	}
	filePath, err := vite.GenerateScript("test.vue", "vite", "dd", &cfg)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(filePath)
}
