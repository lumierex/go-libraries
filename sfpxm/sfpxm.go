package sfpxm

import (
	"fmt"
	"log"
	"net/http"
)

// Request Func
type HandlerFunc func(http.ResponseWriter, *http.Request)

// engine implement interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

//  New sfpxm constructor
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//  ServeHTTP implements http ListenAnServe interface
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := engine.router[key]; ok {
		handler(w, req)
	} else {
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %q\n", req.URL)
		// err server error log
		if err != nil {
			log.Println(err.Error())
		}
	}
}

//
func (engine *Engine) Run(addr string) (err error, ) {
	return http.ListenAndServe(addr, engine)
}
