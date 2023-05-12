package config

import (
	"testing"
)

func TestConfig(t *testing.T) {
	// 12312 12
	cfg, err := GetConfig("config.yaml")
	if err != nil {
		t.Error(err)
		return
	}

	if cfg.Server.Host != "8085" {
		t.Error("cfg.Server.Host is wrong")
	}
	if cfg.HandlerConfigs["vue"].Handler != "vue" {
		t.Error(`cfg.HandlerConfigs["vue"].Handler is wrong `)
	}
	if cfg.HandlerConfigs["vue"].WorkDir != "src" {
		t.Error(`cfg.HandlerConfigs["vue"].WorkDir is wrong`)
	}
	if cfg.HandlerConfigs["vue"].Params["params1"] != "is params 1" {
		t.Error(`cfg.HandlerConfigs["vue"].Params["params1"] is wrong`)
	}
	if len(cfg.HostParams) != 2 {
		t.Error("wrong count cfg.HostParams")
	}
	if cfg.HostParams[0].Key != "is key" {
		t.Error("wrong key")
	}
	if cfg.HostParams[0].Path != "test/path" {
		t.Error("wrong path")
	}
	if cfg.HostParams[0].HandlerConfig != "vue" {
		t.Error("wrong path")
	}

	viteParams := cfg.HandlerConfigs["vite"].Params.GetViteParams()
	if viteParams.HtmlTemplate != "is html_template" {
		t.Error("wrong result")
	}
	if viteParams.ScriptTemplate != "is script_template" {
		t.Error("wrong result")
	}
	if viteParams.AutoGenDir != "is auto_gen_dir" {
		t.Error("wrong result")
	}
	if viteParams.DirPublic != "is dir_public" {
		t.Error("wrong result")
	}
	if viteParams.Host != "is host" {
		t.Error("wrong result")
	}
}
