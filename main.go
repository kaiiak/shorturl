package main

import (
	"flag"

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
		panic(err)
	}
	d, err := data.New(cnf)
	if err != nil {
		panic(err)
	}
	c := controller.New(d)
	r := route.New(c, cnf)

	r.Init()
	if err = r.Run(); err != nil {
		panic(err)
	}
}
