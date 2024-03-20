package router

import (
	"context"
	"fmt"
	"log"
	"net/http"
)

// `router` is private to force users to use Builder methods like `NewRouter`
type router struct {

	// using maps to have the paths be the route key to make
	// things more performant and reduce amount of code

	// example routes map:
	// {
	//   "/goals": {
	//      Get: func(),
	//      Post: func(),
	//   },
	//   "/login": {
	//      Post: func(),
	//   },
	//   etc...

	// as you can see it's easier to now get the routes based on
	// a particular path:

	// goalsPath := "/goals"
	// goalMethods := router.routes[goalsPath]
	// getGoals := goalMethods.Get

	routes map[string]RouteMethod

	// paths here is used similarly to above. however, it's more of a way to quickly index
	// all paths for the `Build` method below. It also and more importantly makes the logic
	// more understandable

	paths map[string]bool

	// storing the context for injecting into all of our RouteMethods on `Build`
	ctx context.Context
}

func NewRouter() router {
	return router {
		// maps need to be initialized in order for them to be usable
		routes: make(map[string]RouteMethod),
		paths: make(map[string]bool),
	}
}

func (r* router) Get(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Get = handle
		r.routes[path] = routes
	} else {
		r.routes[path] = RouteMethod{ Get: handle }
	}
	r.paths[path] = true
}

func (r* router) Post(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Post = handle
		r.routes[path] = routes
	} else {
		r.routes[path] = RouteMethod{ Post: handle }
	}
	r.paths[path] = true
}

func (r* router) Put(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Put = handle
		r.routes[path] = routes
	} else {
		r.routes[path] = RouteMethod{ Put: handle }
	}
	r.paths[path] = true
}

func (r* router) Delete(path string, handle Handle) {
	if routes, ok := r.routes[path]; ok {
		routes.Delete = handle
		r.routes[path] = routes
	} else {
		r.routes[path] = RouteMethod{ Delete: handle }
	}
	r.paths[path] = true
}

func (r* router) Ctx(ctx context.Context) {
	r.ctx = ctx
}

// ! the final sendoff - always make sure to run this at the very end
// ! otherwise you'll wonder why your routes aren't working
func (router* router) Build() {
	// remember how i said router.paths are used for indexing? here u go
	for k := range router.paths {
		path := k
		routesInPath, ok := router.routes[path]

		if !ok {
			panic("idk what the hell you did, but it aint working")
		}

		http.HandleFunc(path, func(w http.ResponseWriter, r* http.Request) {
			log.Println(fmt.Sprintf("started - %v %v", r.Method, path))

			if r.Method == http.MethodGet && routesInPath.Get != nil {
				routesInPath.Get(router.ctx, w, r)
				return
			}

			if r.Method == http.MethodPost && routesInPath.Post != nil {
				routesInPath.Post(router.ctx, w, r)
				return
			}

			if r.Method == http.MethodPut && routesInPath.Put != nil {
				routesInPath.Put(router.ctx, w, r)
				return
			}

			if r.Method == http.MethodDelete && routesInPath.Delete != nil {
				routesInPath.Delete(router.ctx, w, r)
				return
			}

			http.Error(w, "invalid request", http.StatusForbidden)
		})
	}
}

type RouteMethod struct {
	Get Handle
	Post Handle
	Put Handle
	Delete Handle
}

type Handle func(ctx context.Context, w http.ResponseWriter, r* http.Request)
