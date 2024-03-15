package router

import (
	"context"
	"net/http"
)

type Router struct {
	routes map[string][]Handler
	paths map[string]bool
	ctx context.Context
}

func (r* Router) Get(path string, handle Handle) {
	handler := Handler {
		Handle: handle,
		Method: "GET",
	}
	r.routes[path] = append(r.routes[path], handler)
	r.paths[path] = true
}

func (r* Router) Post(path string, handle Handle) {
	handler := Handler {
		Handle: handle,
		Method: "POST",
	}
	r.routes[path] = append(r.routes[path], handler)
	r.paths[path] = true
}

func (r* Router) Put(path string, handle Handle) {
	handler := Handler {
		Handle: handle,
		Method: "PUT",
	}
	r.routes[path] = append(r.routes[path], handler)
	r.paths[path] = true
}

func (r* Router) Delete(path string, handle Handle) {
	handler := Handler {
		Handle: handle,
		Method: "DELETE",
	}
	r.routes[path] = append(r.routes[path], handler)
	r.paths[path] = true
}

func (r* Router) Ctx(ctx context.Context) {
	r.ctx = ctx
}

func (router* Router) Build() {
	for k := range router.paths {
		path := k
		routesInPath := router.routes[path]

		http.HandleFunc(path, func(w http.ResponseWriter, r* http.Request) {
			for i := 0; i < len(routesInPath); i++ {
				route := routesInPath[i]

				if r.Method == route.Method {
					route.Handle(router.ctx, w, r)
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

type Handle func(ctx context.Context, w http.ResponseWriter, r* http.Request)
