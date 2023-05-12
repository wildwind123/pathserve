package paramslist

import (
	"fmt"
	"os"
	"path/filepath"
	"pathserve/internal/config"
	"pathserve/internal/helper"
	"pathserve/internal/watcher"
	"strings"
	"sync"

	"go.uber.org/zap"
)

type ParamsList struct {
	Logger *zap.Logger
	list   map[string]config.HostParam
	mu     sync.Mutex
	Config *config.Config
}

func Get(cfg *config.Config, logger *zap.Logger) *ParamsList {
	return &ParamsList{
		Config: cfg,
		Logger: logger,
		list:   make(map[string]config.HostParam),
	}
}

const IgnoreFilePattern = ".ignore."

func (pL *ParamsList) GetFilePaths(path string, patterns []string) (map[string][]string, error) {
	if patterns == nil {
		return nil, nil
	}
	files := make(map[string][]string)
	filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {

			for i := range patterns {
				if strings.Contains(info.Name(), patterns[i]) {
					if !filepath.IsAbs(path) {
						path, err = filepath.Abs(path)
						if err != nil {
							return fmt.Errorf("unable to filepath.Abs %v", err)
						}
					}
					files[patterns[i]] = append(files[patterns[i]], path)
				}
			}

		}
		return nil
	})

	return files, nil
}

func (pL *ParamsList) SetParamsList(path string, patterns []string) error {
	filePaths, err := pL.GetFilePaths(path, patterns)
	if err != nil {
		return err
	}

	for key := range filePaths {
		err = pL.AddPathToParam(key, filePaths[key]...)
		if err != nil {
			return err
		}
	}
	return nil
}

func (pL *ParamsList) GetParamList() map[string]config.HostParam {
	pL.mu.Lock()
	defer pL.mu.Unlock()
	return pL.list
}

func (pL *ParamsList) GetParam(key string) *config.HostParam {
	pL.mu.Lock()
	defer pL.mu.Unlock()
	val, ok := pL.list[key]
	if !ok {
		return nil
	}
	return &val
}

func (pL *ParamsList) WatchDirAndAddParam(path string, patterns ...string) {
	w := watcher.Get(pL.Logger)
	go w.WatchDir(path, watcher.Hook{
		CreatedDir: func(p string) {},
		CreatedFile: func(p string) {
			for i := range patterns {
				if !strings.Contains(filepath.Base(p), patterns[i]) {
					continue
				}
				err := pL.AddPathToParam(patterns[i], p)
				if err != nil {
					pL.Logger.Error("unable to pL.AddPathToParam", zap.Error(err))
				}
			}
		},
		RemovedOrRenamed: func(p string) {
			err := pL.DeletePath(p)
			if err != nil {
				pL.Logger.Error("unable to pL.AddPathToParam", zap.Error(err))
			}
		},
	})
}

func (pL *ParamsList) AddParam(key string, param config.HostParam) {
	pL.mu.Lock()
	defer pL.mu.Unlock()
	pL.list[key] = param
}

func (pL *ParamsList) AddPathToParam(pattern string, patchs ...string) error {
	if patchs == nil {
		return nil
	}
	pL.mu.Lock()
	defer pL.mu.Unlock()
	if pL.list == nil {
		pL.list = map[string]config.HostParam{}
	}
	for i := range patchs {
		p := filepath.Clean(patchs[i])
		if strings.Contains(p, IgnoreFilePattern) {
			continue
		}

		if !strings.Contains(filepath.Base(p), pattern) {
			continue
		}

		if !filepath.IsAbs(p) {
			absolutePath, err := filepath.Abs(patchs[i])
			if err != nil {
				return fmt.Errorf("unable to filepath.Abs %v", err)
			}
			p = filepath.Clean(absolutePath)
		}

		pL.list[pL.ParamKey(p, pattern)] = config.HostParam{
			Path:          p,
			Key:           pL.ParamKey(p, pattern),
			HandlerConfig: pattern,
		}
		pL.Logger.Debug("added AddPathToParam", zap.String("pattern", pattern), zap.Any("path", p), zap.Any("list", pL.list))

	}
	return nil
}

func (pL *ParamsList) ParamKey(path, pattern string) string {
	return helper.GetSha1(path + "_" + pattern)
}

func (pL *ParamsList) DeletePath(path string) error {
	pL.mu.Lock()
	defer pL.mu.Unlock()
	if !filepath.IsAbs(path) {
		absPath, err := filepath.Abs(path)
		if err != nil {
			return err
		}
		path = absPath
	}
	path = filepath.Clean(path)
	for i := range pL.list {
		if pL.list[i].Path != path {
			continue
		}

		handlerConfig, ok := pL.Config.HandlerConfigs[pL.list[i].HandlerConfig]
		if !ok {
			continue
		}
		if handlerConfig.Handler == "vite" {
			fileForDelete := helper.GetAutoGenFileName(pL.list[i].Path, pL.list[i].Key, IgnoreFilePattern, handlerConfig.Params.GetViteParams().ScriptTemplate)

			autoGenPath := handlerConfig.Params.GetViteParams().AutoGenDir
			if !filepath.IsAbs(autoGenPath) {
				autoGenPath = filepath.Join(filepath.Dir(pL.Config.ConfigPath), autoGenPath)
			}
			err := os.Remove(filepath.Join(autoGenPath, fileForDelete))
			if err != nil {
				pL.Logger.Error("unable to os.Remove", zap.Error(err))
			}
		}
		delete(pL.list, i)
		// os.Remove(filepath.)
	}
	return nil
}
