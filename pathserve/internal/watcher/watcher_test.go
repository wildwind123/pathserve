package watcher

import (
	"fmt"
	"testing"

	"go.uber.org/zap"
)

func TestWatchDir(t *testing.T) {
	t.Skip("manual testing")
	logger, _ := zap.NewDevelopment()
	w := Watcher{
		Logger: logger,
	}
	w.WatchDir("/home/ganbatte/Desktop/project/proxypath2/proxypath/internal/watcher/test_dir", Hook{
		CreatedDir: func(path string) {
			fmt.Println("CreatedDir dir", path)
		},
		CreatedFile: func(path string) {
			fmt.Println("CreatedFile dir", path)
		},
		RemovedOrRenamed: func(path string) {
			fmt.Println("RemovedOrRenamed", path)
		},
	})
}
