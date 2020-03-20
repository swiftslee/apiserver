package cmd

import (
	"flag"
	"log"
	"net/http"
	"os"

	"github.com/yuswift/apiserver/pkg/lib/conf"
	"github.com/yuswift/apiserver/pkg/lib/mysql"
	"github.com/yuswift/apiserver/pkg/routers"
)

var (
	confPath = flag.String("conf", "./conf/conf.yaml", "The conf file path.")
)

func Run() {
	if err := conf.Load(*confPath); err != nil {
		exit(err)
	}

	if err := mysql.Init(); err != nil{
		exit(err)
	}

	routers.Run()

	log.Printf("start listening on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func exit(err error) {
	if err != nil {
		log.Fatal("err happens, exit", err)
	}

	if err := mysql.Close(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
