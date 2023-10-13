package routes

import (
	"main/api"
	"main/oauth"
	"net/http"

	"github.com/gorilla/mux"
)

var routes []Route

func init() {
	register("POST", "/api/common/{table}", api.Insert, nil)
	register("GET", "/api/common/{table}", api.Query, nil)
	register("DELETE", "/api/common/{table}", api.Delete, oauth.CorsHandler)
	register("PUT", "/api/common/{table}", api.Update, oauth.CorsHandler)
	register("GET", "/api/common/{table}/{column}", api.Scalar, nil)
	register("POST", "/api/UpdateTable", api.UpdateTable, oauth.CorsHandler)
	register("post", "/api/UpdateView", api.UpdateView, nil)
	register("POST", "/api/UpdateStoredProcedure", api.UpdateStoredProcedure, nil)
	register("POST", "/api/UpdateSchema", api.UpdateSchema, nil)
	register("GET", "/{page}", api.PageGetter, nil)
	register("GET", "/", api.PageGetter, nil)
	register("POST", "/api/DataInsert", api.RssInsert, nil)

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
	staticrouter(r)
	return r
}
func staticrouter(r *mux.Router) {
	fs := http.FileServer(http.Dir("templatesite/assets/"))
	//ds := http.FileServer(http.Dir("templatesite/"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))

}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
	routes = append(routes, Route{method, pattern, handler, middleware})
}
