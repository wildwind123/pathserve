package watcher

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
	"go.uber.org/zap"
	"golang.org/x/exp/slices"
)

type Watcher struct {
	Logger *zap.Logger
}

func Get(logger *zap.Logger) *Watcher {
	return &Watcher{
		Logger: logger,
	}
}

func (w *Watcher) WatchDir(dir string, hook Hook) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	err = watcher.Add(dir)
	if err != nil {
		w.Logger.Fatal("unable to watcher.Add", zap.Error(err))
	}
	err = filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				w.Logger.Warn("unable to watcher.Add(path)", zap.Error(err))
			}
		}
		return nil
	})
	if err != nil {
		w.Logger.Fatal("filepath.Walk", zap.Error(err))
	}

	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			logger := w.Logger.With(zap.String("event", event.Op.String()), zap.String("path", event.Name))
			if event.Has(fsnotify.Create) {
				fileInfo, err := os.Stat(event.Name)
				if err != nil {
					logger.Error("unable to os.Stat", zap.Error(err))
					continue
				}
				if fileInfo.IsDir() {
					hook.CreatedDir(event.Name)
					logger.Debug("wath new dir")
					err = watcher.Add(event.Name)
					if err != nil {
						logger.Error("unable to watcher.Add", zap.Error(err))
						continue
					}
				} else {
					hook.CreatedFile(event.Name)
					logger.Debug("new file created")
				}
			} else if event.Has(fsnotify.Remove) || event.Has(fsnotify.Rename) {
				logger.Debug("renamed or deleted")
				hook.RemovedOrRenamed(event.Name)
				if slices.Contains(watcher.WatchList(), event.Name) {

					err := watcher.Remove(event.Name)
					if err != nil {
						logger.Warn("unable to watcher.Remove", zap.Error(err))
					}
				}
			}

		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			w.Logger.Warn("watcher.Errors", zap.Error(err))
		}
	}
}
