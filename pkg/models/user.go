package models

import "fmt"

type User struct {
	ID   int64  `json:"id" description:"identifier of the user" db:"id"`
	Name string `json:"name" description:"name of the user" default:"john" db:"name"`
	Age  int    `json:"age" description:"age of the user" db:"age"`
}

type userModel struct {
	baseModel
}

var UserModel = userModel{}

func (u *userModel) New() User {
	return User{}

}

func (u *userModel) Create(user *User) (User, error) {
	sql := fmt.Sprintf("INSERT INTO `user_test`(`name`,`age`) VALUES('%s',%d)",
		user.Name, user.Age)

	lastInsertId, err := u.baseModel.add(sql)
	if err == nil {
		user.ID = lastInsertId
	}
	return *user, err
}

func (u *userModel) Delete(name string) error {
	sql := fmt.Sprintf("DELETE FROM `user_test` WHERE `name` = '%s'",
		name)

	return u.baseModel.delete(sql)
}

func (u *userModel) Get(name string) (User, error) {
	sql := fmt.Sprintf("SELECT `id`,`name`,`age` FROM `user_test` WHERE `name` = '%s'",
		name)

	user := User{}
	err := u.baseModel.get(sql, &user)
	return user, err
}

func (u *userModel) List(name string) ([]User, error) {
	sql := fmt.Sprintf("SELECT `id`,`name`,`age` FROM `user_test`")

	var users []User
	err := u.baseModel.list(sql, &users)
	return users, err
}

func (u *userModel) Update(name string, user *User) error {
	sql := fmt.Sprintf("UPDATE `user_test` SET `name`='%s',`age`='%d' WHERE `name`='%s'",
		user.Name, user.Age, name)

	return u.baseModel.update(sql)
}
