package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Server struct {
		Host string `yaml:"host"`
	} `yaml:"server"`
	HandlerConfigs map[string]HandlerConfig `yaml:"handler_configs"`
	HostParams     Params                   `yaml:"host_params"`
	ConfigPath     string                   `yaml:"-"`
}

type HandlerConfig struct {
	Handler  string `yaml:"handler"`
	WorkDir  string `yaml:"work_dir"`
	WatchDir string `yaml:"watch_dir"`
	Params   Params `yaml:"params"`
}

type ViteParams struct {
	HtmlTemplate   string
	ScriptTemplate string
	AutoGenDir     string
	DirPublic      string
	Host           string
}

type Params map[string]string

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

func (p Params) GetViteParams() *ViteParams {
	viteParams := ViteParams{}
	viteParams.HtmlTemplate = p.getConfigByKey("html_template")
	viteParams.ScriptTemplate = p.getConfigByKey("script_template")
	viteParams.AutoGenDir = p.getConfigByKey("auto_gen_dir")
	viteParams.DirPublic = p.getConfigByKey("dir_public")
	viteParams.Host = p.getConfigByKey("host")
	return &viteParams
}

func (p *Params) getConfigByKey(key string) string {
	paramsClone := *p
	value, ok := paramsClone[key]
	if ok {
		return value
	}
	return ""
}
