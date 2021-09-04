package main

import (
	"context"

	"github.com/wuwentao1998/golang-demos/kitex-server/kitex_gen/demo"
)

// DemoImpl implements the last service interface defined in the IDL.
type DemoImpl struct{}

// Echo implements the DemoImpl interface.
func (s *DemoImpl) Echo(ctx context.Context, req *demo.Request) (resp *demo.Response, err error) {
	// TODO: Your code here...
	return
}
