package router

import (
	"net/http"
	"strings"
)

type Router struct {
	root *trieRoute
}

type trieRoute struct {
	next            map[string]*trieRoute
	handlerByMethod map[string]http.Handler
}

// Register a handler to a path.
//
// `handler` handles incoming requests that matches `path`.
// If `path` is not a valid URL path, `handler` won't be registered.
func (r *Router) Route(method string, path string, handler http.Handler) {
	route := r.root.traverse(path, true)
	route.handlerByMethod[method] = handler
}

// Traverse the trie tree and find a route.
//
// `path` is split into multiple `pattern` strings. For example, a `path` "/foo/bar"
// is split into `pattern` ["foo", "bar"].
// `pattern` identifies a route.
//
// In the traversal, `pattern` is matched with the current route to find the next route.
// Traversal ends until all `pattern` are matched or one of them is not matched.
// However, if `newOnTraversal` is true, a new route is created when a `pattern` is not matched.
func (t *trieRoute) traverse(path string, newOnTraversal bool) *trieRoute {
	patterns := pathToPatterns(path)
	for _, pattern := range patterns {
		if t.next[pattern] == nil && newOnTraversal {
			t.next[pattern] = newRoute()
		}
		t = t.next[pattern]
		if t == nil {
			return t
		}
	}

	return t
}

func newRoute() *trieRoute {
	return &trieRoute{
		next:            make(map[string]*trieRoute),
		handlerByMethod: make(map[string]http.Handler),
	}
}

func NewRouter() Router {
	root := newRoute()
	return Router{root}
}

func isPath(s string) bool {
	if len(s) != 0 && s[:1] == "/" {
		return true
	}

	return false
}

func pathToPatterns(path string) []string {
	if !isPath(path) {
		return []string{}
	}

	return append([]string{"/"}, strings.Split(path[1:], "/")...)
}
