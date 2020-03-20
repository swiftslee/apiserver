package models

import (
	"github.com/yuswift/apiserver/pkg/lib/mysql"
)

type baseModel struct{}

func (m *baseModel) add(sql string) (int64, error) {
	lastInsertId, _, err := mysql.Insert(sql)
	return lastInsertId, err
}

func (m *baseModel) delete(sql string) error {
	_, err := mysql.Delete(sql)
	return err
}

func (m *baseModel) get(sql string, v interface{}) error {
	return mysql.One(sql, v)
}

func (m *baseModel) list(sql string, v interface{}) error {
	return mysql.All(sql, v)

}

func (m *baseModel) update(sql string) error {
	_, err := mysql.Update(sql)
	return err
}
