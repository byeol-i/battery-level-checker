package router

import (
	"log"
	"net/http"
	"regexp"

	"go.uber.org/zap"

	"github.com/aglide100/battery-level-checker/pkg/logger"
)


type routeRule struct {
	name    string
	method  string
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	rules           []*routeRule
	notFoundHandler http.Handler
	version 		string
}

func NewRouter(notFoundHandler http.Handler, version string) *Router {
	return &Router{
		rules:           make([]*routeRule, 0),
		notFoundHandler:  notFoundHandler,
		version: version,
	}
}

func (rtr *Router) AddRule(name string, method, pattern string, handler http.HandlerFunc) {
	newRule := &routeRule{
		name:    name,
		method:  method,
		pattern: regexp.MustCompile("^/api/" + rtr.version + pattern),
		handler: handler,
	}
	rtr.rules = append(rtr.rules, newRule)
	log.Println("add router rule :", name, pattern)
}

func (rtr *Router) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	method := req.Method
	path := req.URL.Path
	for _, rule := range rtr.rules {
		if rule.method != method {
			continue
		}
		
		if !rule.pattern.MatchString(path) {
			continue
		}
		logger.Info("found handler", zap.String("rule", rule.name), zap.String("path", path))
		handler := rule.handler
		handler.ServeHTTP(resp, req)
		return
	}

	rtr.notFoundHandler.ServeHTTP(resp, req)
}
