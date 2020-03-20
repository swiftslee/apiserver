package models

import (
	"testing"

	"github.com/yuswift/apiserver/pkg/lib/conf"
	"github.com/yuswift/apiserver/pkg/lib/mysql"
)

// suppose we have tested the Load function and Init function.
func init() {
	_ = conf.Load("../../conf/conf.yaml")
	_ = mysql.Init()
}

func TestUserCreat(t *testing.T) {
	user := User{Name: "yuswift", Age: 23}
	_, err := UserModel.Create(&user)
	if err != nil {
		t.Error(err)
	} else {
		t.Log(user)
	}
}

func TestUserList(t *testing.T) {
	user, err := UserModel.List("")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(user)
	}

}

func TestUserGet(t *testing.T) {
	user, err := UserModel.Get("yuswift")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(user)
	}

	user, err = UserModel.Get("swift")
	if err != nil {
		t.Log(err)
	} else {
		t.Fatal()
	}

}

func TestUserDelete(t *testing.T) {
	err := UserModel.Delete("yuswift")
	if err != nil {
		t.Error(err)
	}
}

func TestUserUpdate(t *testing.T) {
	user := User{Name: "swift", Age: 23}
	_, err := UserModel.Create(&user)
	if err != nil {
		t.Error(err)
	}

	user.Name = "swift-"
	err = UserModel.Update("swift", &user)
	if err != nil {
		t.Error(err)
	}

	user, err = UserModel.Get("swift-")
	if err != nil {
		t.Error(err)
	} else {
		t.Log(user)
	}

	user, err = UserModel.Get("swift")
	if err == nil {
		t.Fatal()
	} else {
		t.Log(err)
	}

}
