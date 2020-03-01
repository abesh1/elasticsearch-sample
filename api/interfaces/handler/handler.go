package handler

import (
	"net/http"

	"github.com/jiro94/elasticsearch-sample/api/interfaces/handler/product"

	"github.com/jiro94/elasticsearch-sample/api/registry"
	"github.com/julienschmidt/httprouter"
)

func Register(server *http.Server, services *registry.Services) {
	r := httprouter.New()
	r.NotFound = http.NotFoundHandler()

	product.Register(r, services.Product)

	server.Handler = NoCacheHandler(r)
}

func NoCacheHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "private, no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
