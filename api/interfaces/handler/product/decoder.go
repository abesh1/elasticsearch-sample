package product

import (
	"context"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/jiro94/elasticsearch-sample/api/domain/entityreq"
	"github.com/spf13/cast"
)

func decodeGetSearchRequest(_ context.Context, r *http.Request) (interface{}, error) {
	q := r.URL.Query()
	params := httprouter.ParamsFromContext(r.Context())
	req := entityreq.ProductSearch{
		Index:   params.ByName("index"),
		Keyword: q.Get("keyword"),
		Limit:   cast.ToUint32(q.Get("limit")),
	}

	return req, nil
}
