// Code generated by Kitex v0.0.4. DO NOT EDIT.

package demoservice

import (
	"context"
	"github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	"github.com/wuwentao1998/golang-demos/kitex-server/kitex_gen/demo"
)

func serviceInfo() *kitex.ServiceInfo {
	return demoServiceServiceInfo
}

var demoServiceServiceInfo = newServiceInfo()

func newServiceInfo() *kitex.ServiceInfo {
	serviceName := "DemoService"
	handlerType := (*demo.DemoService)(nil)
	methods := map[string]kitex.MethodInfo{
		"Echo": kitex.NewMethodInfo(echoHandler, newDemoServiceEchoArgs, newDemoServiceEchoResult, false),
		"Send": kitex.NewMethodInfo(sendHandler, newDemoServiceSendArgs, nil, true),
	}
	extra := map[string]interface{}{
		"PackageName": "demo",
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.0.4",
		Extra:           extra,
	}
	return svcInfo
}

func echoHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.DemoServiceEchoArgs)
	realResult := result.(*demo.DemoServiceEchoResult)
	success, err := handler.(demo.DemoService).Echo(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newDemoServiceEchoArgs() interface{} {
	return demo.NewDemoServiceEchoArgs()
}

func newDemoServiceEchoResult() interface{} {
	return demo.NewDemoServiceEchoResult()
}

func sendHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*demo.DemoServiceSendArgs)

	err := handler.(demo.DemoService).Send(ctx, realArg.Req)
	if err != nil {
		return err
	}

	return nil
}
func newDemoServiceSendArgs() interface{} {
	return demo.NewDemoServiceSendArgs()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Echo(ctx context.Context, req *demo.Request) (r *demo.Response, err error) {
	var _args demo.DemoServiceEchoArgs
	_args.Req = req
	var _result demo.DemoServiceEchoResult
	if err = p.c.Call(ctx, "Echo", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Send(ctx context.Context, req *demo.Request) (err error) {
	var _args demo.DemoServiceSendArgs
	_args.Req = req
	if err = p.c.Call(ctx, "Send", &_args, nil); err != nil {
		return
	}
	return nil
}
