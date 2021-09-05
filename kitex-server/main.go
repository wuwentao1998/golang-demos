package main

import (
	"log"

	"github.com/wuwentao1998/golang-demos/kitex-server/kitex_gen/demo/demoservice"
)

func main() {
	svr := demoservice.NewServer(new(DemoServiceImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
