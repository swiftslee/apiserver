package mysql

import (
	"fmt"
	"log"
	"time"

	"github.com/go-ozzo/ozzo-dbx"
	_ "github.com/go-sql-driver/mysql"

	"github.com/yuswift/apiserver/pkg/lib/conf"
)

var dbPool *DBPool

type DBPool struct {
	dsn string
	db  *dbx.DB
}

// Init is init mysql
func Init() error {
	dbPool = &DBPool{
		dsn: fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", conf.RuntimeConf.Username, conf.RuntimeConf.Password,
			conf.RuntimeConf.Host, conf.RuntimeConf.Port, conf.RuntimeConf.Db),
	}

	if dbx, err := dbx.Open("mysql", dbPool.dsn); err != nil {
		return err
	} else {
		dbPool.db = dbx
	}

	dbPool.db.DB().SetMaxOpenConns(conf.RuntimeConf.MaxOpenConn)
	dbPool.db.DB().SetMaxIdleConns(conf.RuntimeConf.MaxIdleConn)

	if err := dbPool.Ping(); err != nil {
		return err
	}

	go func() {
		taskConnect := time.NewTicker(3 * time.Second)
		for {
			<-taskConnect.C
			go dbPool.Ping()
		}
	}()

	return nil
}

func GetDB() *dbx.DB {
	return dbPool.db
}

// todo ping hangs if connection can't connet, maybe add timeout
func (dbPool *DBPool) Ping() error {
	if err := dbPool.db.DB().Ping(); err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Close() error {
	if dbPool == nil || dbPool.db == nil {
		return nil
	}
	log.Println("Close mysql client pool.")
	return dbPool.db.DB().Close()
}
