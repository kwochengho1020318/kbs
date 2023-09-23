package routes

import (
	"main/api"
	"net/http"

	"github.com/gorilla/mux"
)

var routes []Route

func init() {
	register("POST", "/api/chat/{phase}", api.Chat, nil)
	register("POST", "/api/common/{table}", api.Insert, nil)
	register("GET", "/api/common/{table}", api.Query, nil)
	register("DELETE", "/api/common/{table}", api.Delete, api.CorsHandler)
	register("PUT", "/api/common/{table}", api.Update, api.CorsHandler)
	register("GET", "/api/common/{table}/{column}", api.Scalar, nil)
	register("POST", "/api/login", api.Login, nil)
	register("POST", "/api/logout", api.Logout, nil)
	register("GET", "/api/page/{page}", api.PageGetter, nil)
	register("GET", "/api/userinfo", api.Userinfo, api.CorsHandler)
	register("POST", "/api/register", api.Register, api.CorsHandler)
	register("POST", "/api/UpdateTable", api.UpdateTable, api.CorsHandler)
	register("post", "/api/UpdateView", api.UpdateView, nil)
}

type Route struct {
	Method     string
	Pattern    string
	Handler    http.HandlerFunc
	Middleware mux.MiddlewareFunc
}

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	for _, route := range routes {
		r.Methods(route.Method).
			Path(route.Pattern).
			Handler(route.Handler)
		if route.Middleware != nil {
			r.Use(route.Middleware)
		}
	}
	return r
}
func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
