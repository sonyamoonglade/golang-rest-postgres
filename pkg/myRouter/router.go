package myRouter

import (
	"fmt"
	"github.com/sonyamoonglade/golang-rest-postgres/pkg/myLogger"
	"net/http"
	"strings"
	"time"
)

type Route struct {
	HandlerFunction http.HandlerFunc
	Method          string
}

type Router struct {
	logger  *myLogger.Logger
	mux     *http.ServeMux
	Handler http.Handler
	Routes  map[string]Route
}

func (r *Router) LogAllRoutes() {
	go func() {
		time.Sleep(300 * time.Millisecond)
		for path, route := range r.Routes {
			r.logger.PrintWithRoute(fmt.Sprintf("-- %s ( %s )", path, route.Method))
		}
	}()
}

func NewRouter(logger *myLogger.Logger) *Router {
	r := Router{}
	r.mux = http.NewServeMux()
	r.Handler = newMiddleware(r.mux, &r)
	r.Routes = map[string]Route{}
	r.logger = logger
	return &r
}

func (r *Router) Post(relativePath string, h http.HandlerFunc) {
	r.Routes[relativePath] = Route{
		HandlerFunction: h,
		Method:          http.MethodPost,
	}

	r.mux.HandleFunc(relativePath, h)

	return
}
func (r *Router) Put(relativePath string, h http.HandlerFunc) {
	r.Routes[relativePath] = Route{
		HandlerFunction: h,
		Method:          http.MethodPut,
	}

	r.mux.HandleFunc(relativePath, h)

	return
}
func (r *Router) Get(relativePath string, h http.HandlerFunc) {
	r.Routes[relativePath] = Route{
		HandlerFunction: h,
		Method:          http.MethodGet,
	}
	r.mux.HandleFunc(relativePath, h)

	return
}
func (r *Router) Delete(relativePath string, h http.HandlerFunc) {
	r.Routes[relativePath] = Route{
		HandlerFunction: h,
		Method:          http.MethodDelete,
	}
	r.mux.HandleFunc(relativePath, h)

	return
}

type methodCheckingMiddleware struct {
	Handler *http.ServeMux
	Router  *Router
}

func newMiddleware(handlerToWrap *http.ServeMux, Router *Router) *methodCheckingMiddleware {
	m := methodCheckingMiddleware{Handler: handlerToWrap}
	m.Router = Router
	return &m
}
func (m *methodCheckingMiddleware) ServeHTTP(w http.ResponseWriter, req *http.Request) {

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
	w.Header().Set("Content-Type", "application/json")
	m.Handler.ServeHTTP(w, req)
	return
}
