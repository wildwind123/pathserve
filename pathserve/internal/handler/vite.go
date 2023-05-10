package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"pathserve/internal/config"
	"pathserve/internal/helper"
	"pathserve/internal/paramslist"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Vite struct {
	Config    *config.Config
	Logger    *zap.Logger
	ParamList *paramslist.ParamsList
}

func (v *Vite) Handle(c *gin.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		cfg := v.Config

		lastDomain, err := helper.GetLastDomain(c.Request.Host)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		viteParams := cfg.HandlerConfigs[lastDomain].Params.GetViteParams()

		segment := c.Param(SegmentKey)
		if segment != "" && segment != "/" {
			filePath, exist, err := GetFilePath(cfg, c.Request.Host, segment, viteParams.DirPublic)
			if err != nil {
				v.Logger.Error("unable to GetFilePath", zap.Error(err))
				c.JSON(http.StatusBadRequest, gin.H{
					"error": fmt.Sprintf("unable to GetFilePath, %v", err),
				})
				return
			}
			if exist {
				c.File(filePath)
				return
			}
		}

		paramKey, err := helper.GetDomainParamKey(c.Request.Host)
		if err != nil {
			v.Logger.Error("unable to helper.GetDomainParamKey", zap.Error(err))
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("unable to helper.GetDomainParamKey, %v", err),
			})
			return
		}
		param := v.ParamList.GetParam(paramKey)

		if param == nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": fmt.Sprintf("param key '%s' does not exist", paramKey),
			})
			return
		}

		newScriptPath, err := v.GenerateScript(param.Path, lastDomain, paramKey, cfg)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		templateHtmlPath := viteParams.HtmlTemplate
		if !filepath.IsAbs(templateHtmlPath) {
			templateHtmlPath = filepath.Join(filepath.Dir(cfg.ConfigPath), templateHtmlPath)
		}
		if !filepath.IsAbs(newScriptPath) {
			newScriptPath, err = filepath.Abs(newScriptPath)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": err.Error(),
				})
				return
			}
		}
		newHtml, err := v.GetViteHtml(viteParams.Host, templateHtmlPath, newScriptPath)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.Writer.Write(newHtml)
	}
}

func (v *Vite) GenerateScript(filePath string, handlerConfigName string, key string, config *config.Config) (string, error) {
	handlerConfig, ok := config.HandlerConfigs[handlerConfigName]
	if !ok {
		return "", fmt.Errorf("config.HandlerConfigs does not exist, %s", handlerConfigName)
	}

	autoGenDir := handlerConfig.Params.GetViteParams().AutoGenDir
	if !filepath.IsAbs(autoGenDir) {
		autoGenDir = filepath.Join(filepath.Dir(config.ConfigPath), autoGenDir)
	}

	if _, err := os.Stat(autoGenDir); os.IsNotExist(err) {
		err = os.MkdirAll(autoGenDir, os.ModePerm)
		if err != nil {
			return "", err
		}
	}
	viteParams := handlerConfig.Params.GetViteParams()
	if !filepath.IsAbs(filePath) {
		filePath = filepath.Join(filepath.Dir(config.ConfigPath), filePath)
	}
	scriptTemplatePath := viteParams.ScriptTemplate
	if !filepath.IsAbs(viteParams.ScriptTemplate) {
		scriptTemplatePath = filepath.Join(filepath.Dir(config.ConfigPath), scriptTemplatePath)
	}

	newFileName := helper.GetAutoGenFileName(filePath, key, paramslist.IgnoreFilePattern, filepath.Ext(scriptTemplatePath))

	newFilePath := filepath.Join(autoGenDir, newFileName)

	newScript, err := v.GetViteScript(scriptTemplatePath, filePath)
	if err != nil {
		return "", err
	}
	newFileByte, err := ioutil.ReadFile(newFilePath)
	if err != nil && !os.IsNotExist(err) {
		return "", err
	}

	if newScript != nil && string(newFileByte) == string(newScript) {
		return newFilePath, nil
	}

	f, err := os.OpenFile(newFilePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer f.Close()

	_, err = f.Write(newScript)
	if err != nil {
		return "", err
	}
	return newFilePath, nil
}

func (h *Vite) GetViteScript(scriptTemplatePath, componentPath string) ([]byte, error) {

	if !filepath.IsAbs(componentPath) {
		var err error
		componentPath, err = filepath.Abs(componentPath)
		if err != nil {
			return nil, err
		}
	}

	b, err := ioutil.ReadFile(scriptTemplatePath)
	if err != nil {
		return nil, err
	}
	newMainScript := bytes.Replace(b, []byte("{{component_path}}"), []byte(componentPath), -1)
	newMainScript = bytes.Replace(newMainScript, []byte("///@ts-nocheck"), []byte(""), -1)
	return newMainScript, nil
}

func (h *Vite) GetViteHtml(viteHost, viteTempalteHtml, viteScriptPath string) ([]byte, error) {

	viteClient := fmt.Sprintf("%s/@vite/client", viteHost)
	viteScript := fmt.Sprintf("%s/%s", viteHost, viteScriptPath)

	htmlByte, err := ioutil.ReadFile(viteTempalteHtml)
	if err != nil {
		return nil, err
	}
	newHtmlByte := bytes.Replace(htmlByte, []byte("{{vite_client}}"), []byte(viteClient), -1)
	newHtmlByte = bytes.Replace(newHtmlByte, []byte("{{vite_script}}"), []byte(viteScript), -1)
	return newHtmlByte, nil
}
