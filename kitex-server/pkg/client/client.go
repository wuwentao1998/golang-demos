package main

import (
	"context"
	"log"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/wuwentao1998/golang-demos/kitex-server/kitex_gen/demo"
	"github.com/wuwentao1998/golang-demos/kitex-server/kitex_gen/demo/demoservice"
)

func main() {
	cli, err := demoservice.NewClient("p.s.m", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		panic(err)
	}

	for {
		req := demo.NewRequest()
		req.Msg = "hello"

		resp, err := cli.Echo(context.Background(), req)
		if err != nil {
			panic(err)
		}
		log.Println(resp)

		err = cli.Send(context.Background(), req)
		if err != nil {
			panic(err)
		}

		time.Sleep(time.Second)
	}
}
