package helper

import (
	"testing"
)

func TestGetLastDomain(t *testing.T) {
	lastDomain, err := GetLastDomain("test1.test2.com")
	if err != nil {
		t.Error(err)
	}
	if lastDomain != "test1" {
		t.Error("wrong reponse")
	}
	_, err = GetLastDomain("test2.com")
	if err == nil {
		t.Error("should has error, is top-level domain")
	}
}
func TestGetDomainParamKey(t *testing.T) {
	_, err := GetDomainParamKey("test1.test2.com")
	if err == nil {
		t.Error("shoild has error")
	}
	paramKey, err := GetDomainParamKey("test1.paramKey.test2.com")
	if err != nil {
		t.Error("shoild no error")
	}
	if paramKey != "paramKey" {
		t.Error("wrong paramKey")
	}
}
