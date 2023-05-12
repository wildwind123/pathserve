package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host" json:"host"`
	} `yaml:"server" json:"server"`
	HandlerConfigs map[string]HandlerConfig `yaml:"handler_configs" json:"handler_configs"`
	HostParams     []HostParam              `yaml:"host_params" json:"host_params"`
	ConfigPath     string                   `yaml:"-"`
}

type HandlerConfig struct {
	Handler  string        `yaml:"handler" json:"handler"`
	WorkDir  string        `yaml:"work_dir" json:"work_dir"`
	WatchDir string        `yaml:"watch_dir" json:"watch_dir"`
	Params   HandlerParams `yaml:"params" json:"params"`
}

type HostParam struct {
	Path          string `yaml:"path" json:"path"`
	Key           string `yaml:"key" json:"key"`
	HandlerConfig string `yaml:"handler_config" json:"handler_config"`
}

type ViteParams struct {
	HtmlTemplate   string
	ScriptTemplate string
	AutoGenDir     string
	DirPublic      string
	Host           string
}

type HandlerParams map[string]string

func GetConfig(path string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	cfg := Config{
		ConfigPath: path,
	}
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

func (p HandlerParams) GetViteParams() *ViteParams {
	viteParams := ViteParams{}
	viteParams.HtmlTemplate = p.getConfigByKey("html_template")
	viteParams.ScriptTemplate = p.getConfigByKey("script_template")
	viteParams.AutoGenDir = p.getConfigByKey("auto_gen_dir")
	viteParams.DirPublic = p.getConfigByKey("dir_public")
	viteParams.Host = p.getConfigByKey("host")
	return &viteParams
}

func (p *HandlerParams) getConfigByKey(key string) string {
	paramsClone := *p
	value, ok := paramsClone[key]
	if ok {
		return value
	}
	return ""
}
