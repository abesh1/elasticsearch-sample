package product

import (
	"context"

	"github.com/jiro94/elasticsearch-sample/api/usecase"

	mapset "github.com/deckarep/golang-set"
	"github.com/jiro94/elasticsearch-sample/api/domain/entityres"

	"github.com/jiro94/elasticsearch-sample/api/domain/entityreq"
)

func (s serv) GetSearch(ctx context.Context, req entityreq.ProductSearch) (res interface{}, err error) {
	switch req.Index {
	case "web":
		res, err = s.getWebESSearch(ctx, req)
	case "v5_prefix":
		res, err = s.getV5PrefixESSearch(ctx, req)
	case "v5_partial":
		res, err = s.getV5PartialESSearch(ctx, req)
	default:
		return
	}
	return
}

func (s serv) getWebESSearch(ctx context.Context, req entityreq.ProductSearch) (interface{}, error) {
	list, err := s.repo.GetWebESSearch(ctx, req)
	if err != nil {
		return nil, err
	}

	return usecase.NewItems(list), nil
}

func (s serv) getV5PrefixESSearch(ctx context.Context, req entityreq.ProductSearch) (interface{}, error) {
	list, err := s.repo.GetV5PrefixESSearch(ctx, req)
	if err != nil {
		return nil, err
	}

	return usecase.NewItems(list), nil
}

func (s serv) getV5PartialESSearch(ctx context.Context, req entityreq.ProductSearch) (interface{}, error) {
	list, err := s.repo.GetV5PartialESSearch(ctx, req)
	if err != nil {
		return nil, err
	}

	return usecase.NewItems(list), nil
}

func (s serv) GetSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (res interface{}, err error) {
	switch req.Index {
	case "web":
		res, err = s.getWebSearchSuggestion(ctx, req)
	case "v5_prefix":
		res, err = s.getV5PrefixSearchSuggestion(ctx, req)
	case "v5_partial":
		res, err = s.getV5PartialSearchSuggestion(ctx, req)
	}
	return
}

func (s serv) getWebSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (interface{}, error) {
	multiList, err := s.repo.GetWebESSearchSuggestion(ctx, req)
	if err != nil {
		return nil, err
	}

	set := mapset.NewSet()
	authors := make(entityres.SuggestionAuthorList, 0, len(multiList.Author))
	for _, v := range multiList.Author {
		for _, v2 := range v.Authors {
			if set.Contains(v2.ID) {
				continue
			}
			set.Add(v2.ID)
			authors = append(authors, entityres.SuggestionAuthor{
				ID:   v2.ID,
				Name: v2.Name,
			})
		}
	}

	products := make(entityres.SuggestionProductList, 0, len(multiList.Product))
	for _, v := range multiList.Product {
		products = append(products, entityres.SuggestionProduct{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	res := entityres.Suggestion{
		Query:   req.Keyword,
		Author:  usecase.NewItems(authors),
		Product: usecase.NewItems(products),
	}

	return res, nil
}

func (s serv) getV5PrefixSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (interface{}, error) {
	multiList, err := s.repo.GetV5PrefixESSearchSuggestion(ctx, req)
	if err != nil {
		return nil, err
	}

	authors := make(entityres.SuggestionAuthorList, 0, len(multiList.Author))
	for _, v := range multiList.Author {
		authors = append(authors, entityres.SuggestionAuthor{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	products := make(entityres.SuggestionProductList, 0, len(multiList.Product))
	for _, v := range multiList.Product {
		products = append(products, entityres.SuggestionProduct{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	res := entityres.Suggestion{
		Query:   req.Keyword,
		Author:  usecase.NewItems(authors),
		Product: usecase.NewItems(products),
	}

	return res, nil
}

func (s serv) getV5PartialSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (interface{}, error) {
	multiList, err := s.repo.GetV5PartialESSearchSuggestion(ctx, req)
	if err != nil {
		return nil, err
	}

	authors := make(entityres.SuggestionAuthorList, 0, len(multiList.Author))
	for _, v := range multiList.Author {
		authors = append(authors, entityres.SuggestionAuthor{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	products := make(entityres.SuggestionProductList, 0, len(multiList.Product))
	for _, v := range multiList.Product {
		products = append(products, entityres.SuggestionProduct{
			ID:   v.ID,
			Name: v.Name,
		})
	}

	res := entityres.Suggestion{
		Query:   req.Keyword,
		Author:  usecase.NewItems(authors),
		Product: usecase.NewItems(products),
	}

	return res, nil
}
