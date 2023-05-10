package paramslist

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"go.uber.org/zap"
)

func TestGetFileList(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	pl := ParamsList{
		Logger: logger,
	}
	pattern := ".dd."
	list, err := pl.GetFilePaths("test_dir", []string{pattern})
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(list)
	if len(list[pattern]) != 3 {
		t.Error("wrong result")
	}
	fileAbsolutePath, err := filepath.Abs("test_dir/test1.dd.ext")
	if err != nil {
		t.Error(err)
		return
	}
	info, err := os.Stat(fileAbsolutePath)
	if err != nil {
		t.Error(err)
		return
	}
	if info.IsDir() {
		t.Error("wrong result")
		return
	}

	if list[".dd."][0] != fileAbsolutePath {
		t.Error("wrong result")
	}

	err = pl.SetParamsList("test_dir", []string{".dd."})
	if err != nil {
		t.Error(err)
	}

	paramsList := pl.GetParamList()
	if len(paramsList) != 3 {
		t.Error("wrong result")
	}
	if paramsList[pl.ParamKey(fileAbsolutePath, ".dd.")].Path != fileAbsolutePath {
		t.Error("wrong result")
	}
	if paramsList[pl.ParamKey(fileAbsolutePath, ".dd.")].Key != pl.ParamKey(fileAbsolutePath, ".dd.") {
		t.Error("wrong result")
	}
}

func TestWatchDir(t *testing.T) {
	logger, _ := zap.NewDevelopment()
	pl := ParamsList{
		Logger: logger,
	}
	pl.WatchDirAndAddParam("test_dir2", ".dd.")
	time.Sleep(1 * time.Second)
	filePath := "test_dir2/test.dd.ext"

	_, err := os.Create(filePath)
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(filePath)
	time.Sleep(1 * time.Second)
	fullPathFull, err := filepath.Abs(filePath)
	if err != nil {
		t.Error(err)
	}
	if pl.GetParamList()[pl.ParamKey(fullPathFull, ".dd.")].Path != fullPathFull {
		t.Error("wrong result")
	}
}
