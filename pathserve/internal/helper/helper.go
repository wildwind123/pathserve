package helper

import (
	"crypto/sha1"
	"fmt"
	"path/filepath"
	"strings"
)

func GetLastDomain(host string) (string, error) {
	items := strings.Split(host, ".")
	if len(items) < 3 {
		return "", fmt.Errorf("is it top-level domain %s", host)
	}
	return items[0], nil
}

func GetDomainParamKey(host string) (string, error) {
	items := strings.Split(host, ".")
	if len(items) < 3 {
		return "", fmt.Errorf("is it top-level domain %s", host)
	}

	if len(items) < 4 {
		return "", fmt.Errorf("key doesn't exist %s. host shuld be like test.[key].top-level-domain.com or test.[key].*.top-level-domain.com", host)
	}

	return items[1], nil
}

func GetSha1(txt string) string {
	hash := sha1.Sum([]byte(txt))
	return fmt.Sprintf("%x", hash)
}

func GetAutoGenFileName(sourceFilePath, key, ignorePattern string, templatePath string) string {
	return fmt.Sprintf("%s.%s%s%s", filepath.Base(sourceFilePath), key, ignorePattern, filepath.Ext(templatePath))
}
