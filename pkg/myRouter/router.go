package myRouter

import (
	"fmt"
	"net/http"
	"strings"
)

type Router struct {
	mux     *http.ServeMux
	Handler http.Handler
	Routes  map[string]Route
}

func NewRouter() *Router {
	r := Router{}
	r.mux = http.NewServeMux()
	r.Handler = NewMiddleware(r.mux, &r)
	r.Routes = map[string]Route{}
	return &r
}

type Route struct {
	HandlerFunction http.HandlerFunc
	Method          string
}

func (r *Router) POST(relativePath string, h http.HandlerFunc) {
	r.Routes[relativePath] = Route{
		HandlerFunction: h,
		Method:          http.MethodPost,
	}
	r.mux.HandleFunc(relativePath, h)
	return
}
func (r *Router) GET(relativePath string, h http.HandlerFunc) {
	r.Routes[relativePath] = Route{
		HandlerFunction: h,
		Method:          http.MethodGet,
	}
	r.mux.HandleFunc(relativePath, h)
	return
}

type MethodCheckingMiddleware struct {
	Handler *http.ServeMux
	Router  *Router
}

func NewMiddleware(handlerToWrap *http.ServeMux, Router *Router) *MethodCheckingMiddleware {
	m := MethodCheckingMiddleware{Handler: handlerToWrap}
	m.Router = Router
	return &m
}
func (m *MethodCheckingMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {

	requestPattern := req.URL.Path
	requestMethod := req.Method

	response := fmt.Sprintf("cannot %s%s \n", strings.ToUpper(requestMethod), requestPattern)

	hit := 0
	for k, _ := range m.Router.Routes {
		if k == requestPattern {
			hit += 1
		}
	}
	if hit == 0 {
		w.WriteHeader(404)
		w.Write([]byte(response))
		return
	}

	servedMethod := m.Router.Routes[requestPattern].Method

	if requestMethod != servedMethod {
		w.WriteHeader(404)
		w.Write([]byte(response))
		return
	}
	m.Handler.ServeHTTP(w, req)
	return
}
