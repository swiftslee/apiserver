package mysql

import (
	"log"
)

// One is select for one struct
func One(sql string, value interface{}) error {
	log.Printf("execute select sql for one: %s", sql)
	err := dbPool.db.NewQuery(sql).One(value)
	if err != nil {
		log.Printf("error to get result for sql: %s. err: %s", sql, err)
	}
	return err
}

// All is select for slice
func All(sql string, value interface{}) error {
	log.Printf("execute select sql for all: %s", sql)
	err := dbPool.db.NewQuery(sql).All(value)
	if err != nil {
		log.Printf("error to get result for sql: %s. err: %s", sql, err)
	}
	return err
}

// Insert is insert data to mysql
// return int64: last insert id, -1 if error happens
// return int64: affected rows, 0 if error happens
// return error: error info
func Insert(sql string) (int64, int64, error) {
	log.Printf("execute insert sql: %s", sql)
	result, err := dbPool.db.NewQuery(sql).Execute()
	if err != nil {
		log.Printf("error to execute sql: %s. err: %s", sql, err)
		return -1, -1, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		log.Printf("error to get rows affected id for sql: %s. err: %s", sql, err)
		return -1, 0, err
	}
	lastInsertId, err := result.LastInsertId()
	if err != nil {
		log.Printf("error to get last insert id for sql: %s. err: %s", sql, err)
		return -1, 0, err
	}
	return lastInsertId, rowsAffected, nil
}

// Update is insert data to mysql
// return int64: affected rows, 0 if error happens
// return error: error info
func Update(sql string) (int64, error) {
	log.Printf("execute update sql: %s", sql)
	result, err := dbPool.db.NewQuery(sql).Execute()
	if err != nil {
		log.Printf("error to execute sql: %s. err: %s", sql, err)
		return 0, err
	}
	return result.RowsAffected()
}

// Delete is delete from mysql
// return int64: affected rows, 0 if error happens
// return error: error info
func Delete(sql string) (int64, error) {
	log.Printf("execute delete sql: %s", sql)
	result, err := dbPool.db.NewQuery(sql).Execute()
	if err != nil {
		log.Printf("error to execute sql: %s. err: %s", sql, err)
		return 0, err
	}
	return result.RowsAffected()
}
