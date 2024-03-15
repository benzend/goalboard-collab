package router

import (
	"context"
	"net/http"
)

type router struct {
	routes map[string]RouteMethod
	paths map[string]bool
	ctx context.Context
}

func NewRouter() router {
	return router {
		routes: make(map[string]RouteMethod),
		paths: make(map[string]bool),
	}
}

func (r* router) Get(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Get = handle
	} else {
		r.routes[path] = RouteMethod{ Get: handle }
	}
	r.paths[path] = true
}

func (r* router) Post(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Post = handle
	} else {
		r.routes[path] = RouteMethod{ Post: handle }
	}
	r.paths[path] = true
}

func (r* router) Put(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Put = handle
	} else {
		r.routes[path] = RouteMethod{ Put: handle }
	}
	r.paths[path] = true
}

func (r* router) Delete(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Delete = handle
	} else {
		r.routes[path] = RouteMethod{ Delete: handle }
	}
	r.paths[path] = true
}

func (r* router) Ctx(ctx context.Context) {
	r.ctx = ctx
}

func (router* router) Build() {
	for k := range router.paths {
		path := k
		routesInPath, ok := router.routes[path]

		if !ok {
			panic("idk what the hell you did, but it aint working")
		}

		http.HandleFunc(path, func(w http.ResponseWriter, r* http.Request) {
			if routesInPath.Get != nil {
				if r.Method == http.MethodGet {
					routesInPath.Get(router.ctx, w, r)
					return
				}
			}
			if routesInPath.Post != nil {
				if r.Method == http.MethodPost {
					routesInPath.Post(router.ctx, w, r)
					return
				}
			}
			if routesInPath.Put != nil {
				if r.Method == http.MethodPut {
					routesInPath.Put(router.ctx, w, r)
					return
				}
			}
			if routesInPath.Delete != nil {
				if r.Method == http.MethodDelete {
					routesInPath.Delete(router.ctx, w, r)
					return
				}
			}

			http.Error(w, http.ErrNotSupported.ErrorString, http.StatusMethodNotAllowed)
		})
	}
}

type Route struct {
	Path string
	Handlers []Handler
}

type Handler struct {
	Method string
	Handle Handle
}

type RouteMethod struct {
	Get Handle
	Post Handle
	Put Handle
	Delete Handle
}

type Handle func(ctx context.Context, w http.ResponseWriter, r* http.Request)
