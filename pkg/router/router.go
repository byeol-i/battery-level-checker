package router

import (
	"net/http"
	"regexp"

	"go.uber.org/zap"

	"github.com/byeol-i/battery-level-checker/pkg/logger"
)

type routeRule struct {
	name    string
	method  string
	pattern *regexp.Regexp
	handler http.Handler
}

type Router struct {
	rules          []*routeRule
	defaultHandler http.Handler
	middleware     func(http.Handler, http.ResponseWriter, *http.Request) http.Handler
	version        string
}

func NewRouter(defaultHandler http.Handler, version string) *Router {
	logger.Info("New Router base url : api/" + version)
	return &Router{
		rules:          make([]*routeRule, 0),
		defaultHandler: defaultHandler,
		version:        version,
	}
}

func (rtr *Router) Use(middleware func(http.Handler, http.ResponseWriter, *http.Request) http.Handler) {
	rtr.middleware = middleware
}

func (rtr *Router) AddRule(name string, method, pattern string, handler http.HandlerFunc) {
	newRule := &routeRule{
		name:    name,
		method:  method,
		pattern: regexp.MustCompile("^/api/" + rtr.version + pattern),
		handler: handler,
	}
	rtr.rules = append(rtr.rules, newRule)
	logger.Info("add router rule :", zap.String("name", name), zap.String("pattern", pattern))
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

		// logger.Info("found handler", zap.String("rule", rule.name), zap.String("path", path))

		if rtr.middleware != nil {
			handler := rtr.middleware(rule.handler, resp, req)
			if handler != nil {
				handler.ServeHTTP(resp, req)
			}
		} else {
			handler := rule.handler
			handler.ServeHTTP(resp, req)
		}

		return
	}

	rtr.defaultHandler.ServeHTTP(resp, req)
}

