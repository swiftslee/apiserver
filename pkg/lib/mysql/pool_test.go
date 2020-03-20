package mysql

import (
	"github.com/yuswift/apiserver/pkg/lib/conf"
	"testing"
)

func TestInit(t *testing.T) {
	_ = conf.Load("../../../conf/conf.yaml")
	err := Init()
	if err != nil {
		t.Error(err)
	}
}
