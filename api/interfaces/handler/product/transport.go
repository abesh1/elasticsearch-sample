package product

import (
	"context"
	"encoding/json"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/jiro94/elasticsearch-sample/api/domain/service"
	"github.com/julienschmidt/httprouter"
)

func Register(r *httprouter.Router, s service.Product) {
	r.Handler(http.MethodGet, "/:index/search",
		httptransport.NewServer(
			getSearchEndpoint(s),
			decodeGetSearchRequest,
			encodeResponse,
		))
	r.Handler(http.MethodGet, "/:index/search/suggestion",
		httptransport.NewServer(
			getSearchSuggestionEndpoint(s),
			decodeGetSearchRequest,
			encodeResponse,
		))
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}
