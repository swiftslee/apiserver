package conf

import (
	"fmt"
	"testing"
)

func TestLoad(t *testing.T) {
	path := "../../../conf/conf.yaml"
	err := Load(path)
	if err != nil {
		t.Error(fmt.Sprintf("Failed to load conf %s", err))
	} else {
		t.Log(*RuntimeConf)
	}
}
