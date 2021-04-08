package sfpxm_web_core

import (
	"github.com/gin-gonic/gin"
	"sfpxm-web-core/pkg/plugin"
)

type ServerOptions struct {
	port string // listening port

	selectorSvrAddr string   // service discovery server address
	tracingSvrAddr  string   // tracing plugin server address
	tracingSpanName string   //
	pluginNames     []string // plugin name
}

type ServerOption func(options *ServerOptions)

func WithPort(port string) ServerOption {
	return func(o *ServerOptions) {
		o.port = port
	}
}

func WithSelectorSvrAddr(selectorSvrAddr string) ServerOption {
	return func(o *ServerOptions) {
		o.selectorSvrAddr = selectorSvrAddr
	}
}

func WithTracingSvrAddr(tracingSvrAddr string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSvrAddr = tracingSvrAddr
	}
}

func WithTracingSpanName(tracingSpanName string) ServerOption {
	return func(o *ServerOptions) {
		o.tracingSpanName = tracingSpanName
	}
}

func WithPlugin(pluginName ...string) ServerOption {
	return func(o *ServerOptions) {
		o.pluginNames = append(o.pluginNames, pluginName...)
	}
}

type Server struct {
	*gin.Engine
	Adder   string
	plugins []plugin.Plugin
}
