package main

import (
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm"
	"github.com/tetratelabs/proxy-wasm-go-sdk/proxywasm/types"

	"github.com/jptosso/coraza-waf/v2"
)

type vmContext struct {
	types.DefaultVMContext
}

// https://github.com/apache/apisix/blob/master/t/wasm/fault-injection/main.go
func (*vmContext) NewPluginContext(contextID uint32) types.PluginContext {
	return &pluginContext{}
}

type pluginContext struct {
	types.DefaultPluginContext
}

func (ctx *pluginContext) OnPluginStart(pluginConfigurationSize int) types.OnPluginStartStatus {
	coraza.NewWaf()
	return types.OnPluginStartStatusOK
}

func (ctx *pluginContext) NewHttpContext(contextID uint32) types.HttpContext {
	return &httpLifecycle{parent: ctx}
}

type httpLifecycle struct {
	types.DefaultHttpContext
	parent *pluginContext
}

func (ctx *httpLifecycle) OnHttpRequestHeaders(numHeaders int, endOfStream bool) types.Action {
	err := proxywasm.SendHttpResponse(200, nil, []byte{}, -1)
	if err != nil {
		proxywasm.LogErrorf("waf.sendHttpResponse error: %v", err)
		return types.ActionPause
	}
	return types.ActionPause
}

func main() {
	proxywasm.SetVMContext(&vmContext{})
}
