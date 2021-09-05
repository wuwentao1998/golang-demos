package main

import (
	"context"
	"fmt"

	"github.com/wuwentao1998/golang-demos/kitex-server/kitex_gen/demo"
)

// DemoServiceImpl implements the last service interface defined in the IDL.
type DemoServiceImpl struct{}

// Echo implements the DemoServiceImpl interface.
func (s *DemoServiceImpl) Echo(ctx context.Context, req *demo.Request) (resp *demo.Response, err error) {
	resp = &demo.Response{
		Code: int8(0),
		Msg:  req.Msg,
	}
	return resp, nil
}

// Send implements the DemoServiceImpl interface.
func (s *DemoServiceImpl) Send(ctx context.Context, req *demo.Request) (err error) {
	fmt.Println(req.Msg)
	return nil
}
