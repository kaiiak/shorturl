package main

import (
	"flag"
	"fmt"

	"github.com/kaiiak/shorturl/config"
	"github.com/kaiiak/shorturl/controller"
	"github.com/kaiiak/shorturl/data"
	"github.com/kaiiak/shorturl/route"
)

func main() {
	var configpath string
	flag.StringVar(&configpath, "config", "", "open config file without empty")
	flag.Parse()
	cnf, err := config.New(configpath)
	if err != nil {
		panic(fmt.Sprintf("new config error:%s", err))
	}
	d, err := data.New(cnf)
	if err != nil {
		panic(fmt.Sprintf("new data error:%s", err))
	}
	c := controller.New(d)
	r := route.New(c, cnf)

	r.Init()
	if err = r.Run(); err != nil {
		panic(fmt.Sprintf("run router error:%s", err))
	}
}
