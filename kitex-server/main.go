package main

import (
	"log"

	"github.com/wuwentao1998/golang-demos/kitex-server/kitex_gen/demo/demo"
)

func main() {
	svr := demo.NewServer(new(DemoImpl))

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
