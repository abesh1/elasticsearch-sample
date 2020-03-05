package product

import (
	"context"
	"encoding/json"
	"strings"
	"unicode"

	"github.com/jiro94/elasticsearch-sample/internal/mecab"
	"github.com/jiro94/elasticsearch-sample/internal/util"

	"github.com/spf13/cast"

	"github.com/olivere/elastic/v7"

	"github.com/jiro94/elasticsearch-sample/api/domain/entity"
	"github.com/jiro94/elasticsearch-sample/api/domain/entityreq"
)

func (r repo) GetWebESSearch(ctx context.Context, req entityreq.ProductSearch) (list entity.ESProductList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	bq := elastic.NewBoolQuery()
	nq := normalizeQuery(req.Keyword)
	bq.Should(
		elastic.NewPrefixQuery("name", nq).Boost(100),
		elastic.NewWildcardQuery("name", "*"+nq+"*").Boost(50),
		elastic.NewNestedQuery("authors", elastic.NewPrefixQuery("authors.name", nq).Boost(100)),
		elastic.NewNestedQuery("authors", elastic.NewWildcardQuery("authors.name", "*"+nq+"*").Boost(50)),
	)
	for _, v := range queryToSlice(req.Keyword) {
		if nq := normalizeQuery(v); nq != "" {
			bq.Should(
				elastic.NewPrefixQuery("name", nq).Boost(100),
				elastic.NewWildcardQuery("name", "*"+nq+"*").Boost(50),
				elastic.NewNestedQuery("authors", elastic.NewPrefixQuery("authors.name", nq).Boost(100)),
				elastic.NewNestedQuery("authors", elastic.NewWildcardQuery("authors.name", "*"+nq+"*").Boost(50)),
			)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			bq.Should(
				elastic.NewPrefixQuery("name_kana", nkq).Boost(80),
				elastic.NewWildcardQuery("name_kana", "*"+nkq+"*").Boost(30),
				elastic.NewNestedQuery("authors", elastic.NewPrefixQuery("authors.name_kana", nkq).Boost(80)),
				elastic.NewNestedQuery("authors", elastic.NewWildcardQuery("authors.name_kana", "*"+nkq+"*").Boost(30)),
			)
		}
	}

	res, err := r.ES.Session("store").
		Search().
		Index("web").
		Query(bq).
		From(0).
		Size(int(req.Limit)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return
	}
	list2 := make(entity.ESProductList, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		var d entity.ESProduct
		if err = json.Unmarshal(hit.Source, &d); err != nil {
			return
		}
		list2 = append(list2, d)
	}
	list = list2

	return
}

func (r repo) GetV5PrefixESSearch(ctx context.Context, req entityreq.ProductSearch) (list entity.ESProductList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	bq := elastic.NewBoolQuery()
	for _, v := range queryToSlice(req.Keyword) {
		kq := elastic.NewBoolQuery()
		if nq := normalizeQuery(v); nq != "" {
			kq.Should(
				elastic.NewTermQuery("name", nq).Boost(100),
				elastic.NewNestedQuery("authors", elastic.NewTermQuery("authors.name", nq)).Boost(100),
			)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			kq.Should(
				elastic.NewTermQuery("name_kana", nkq).Boost(80),
				elastic.NewNestedQuery("authors", elastic.NewTermQuery("authors.name_kana", nkq)).Boost(80),
			)
		}
		bq.Must(kq)
	}

	res, err := r.ES.Session("store").
		Search().
		Index("v5_prefix").
		Query(bq).
		From(0).
		Size(int(req.Limit)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return
	}
	list2 := make(entity.ESProductList, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		var d entity.ESProduct
		if err = json.Unmarshal(hit.Source, &d); err != nil {
			return
		}
		list2 = append(list2, d)
	}
	list = list2

	return
}

func (r repo) GetV5PartialESSearch(ctx context.Context, req entityreq.ProductSearch) (list entity.ESProductList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	bq := elastic.NewBoolQuery()
	for _, v := range queryToSlice(req.Keyword) {
		kq := elastic.NewBoolQuery()
		if nq := normalizeQuery(v); nq != "" {
			kq.Should(
				elastic.NewMatchPhraseQuery("name", nq).Boost(100),
				elastic.NewNestedQuery("authors", elastic.NewMatchPhraseQuery("authors.name", nq)).Boost(100),
			)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			kq.Should(
				elastic.NewMatchPhraseQuery("name_kana", nkq).Boost(80),
				elastic.NewNestedQuery("authors", elastic.NewMatchPhraseQuery("authors.name_kana", nkq)).Boost(80),
			)
		}
		bq.Must(kq)
	}

	res, err := r.ES.Session("store").
		Search().
		Index("v5_partial").
		Query(bq).
		From(0).
		Size(int(req.Limit)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return
	}
	list2 := make(entity.ESProductList, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		var d entity.ESProduct
		if err = json.Unmarshal(hit.Source, &d); err != nil {
			return
		}
		list2 = append(list2, d)
	}
	list = list2

	return
}

func (r repo) GetV5PrefixAndPartialESSearch(ctx context.Context, req entityreq.ProductSearch) (list entity.ESProductList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	bq := elastic.NewBoolQuery()
	for _, v := range queryToSlice(req.Keyword) {
		kq := elastic.NewBoolQuery()
		if nq := normalizeQuery(v); nq != "" {
			kq.Should(
				elastic.NewTermQuery("name", nq).Boost(100),
				elastic.NewWildcardQuery("name", "*"+nq+"*").Boost(50),
				elastic.NewNestedQuery("authors", elastic.NewTermQuery("authors.name", nq).Boost(100)),
				elastic.NewNestedQuery("authors", elastic.NewWildcardQuery("authors.name", "*"+nq+"*").Boost(50)),
			)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			kq.Should(
				elastic.NewTermQuery("name_kana", nkq).Boost(80),
				elastic.NewWildcardQuery("name_kana", "*"+nkq+"*").Boost(30),
				elastic.NewNestedQuery("authors", elastic.NewTermQuery("authors.name_kana", nkq).Boost(80)),
				elastic.NewNestedQuery("authors", elastic.NewWildcardQuery("authors.name_kana", "*"+nkq+"*").Boost(30)),
			)
		}
		bq.Must(kq)
	}

	res, err := r.ES.Session("store").
		Search().
		Index("v5_prefix").
		Query(bq).
		From(0).
		Size(int(req.Limit)).
		Pretty(true).
		Do(ctx)
	if err != nil {
		return
	}
	list2 := make(entity.ESProductList, 0, len(res.Hits.Hits))
	for _, hit := range res.Hits.Hits {
		var d entity.ESProduct
		if err = json.Unmarshal(hit.Source, &d); err != nil {
			return
		}
		list2 = append(list2, d)
	}
	list = list2

	return
}

func (r repo) GetWebESSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (list entity.MultiESProductList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	productBQ := elastic.NewBoolQuery()
	authorBQ := elastic.NewBoolQuery()

	for _, v := range queryToSlice(req.Keyword) {
		if nq := normalizeQuery(v); nq != "" {
			productBQ.Should(
				elastic.NewPrefixQuery("name", nq).Boost(100),
				elastic.NewWildcardQuery("name", "*"+nq+"*").Boost(50),
			)
			authorBQ.Should(
				elastic.NewNestedQuery("authors", elastic.NewPrefixQuery("authors.name", nq).Boost(100)),
				elastic.NewNestedQuery("authors", elastic.NewWildcardQuery("authors.name", "*"+nq+"*").Boost(50)),
			)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			productBQ.Should(
				elastic.NewPrefixQuery("name_kana", nkq).Boost(80),
				elastic.NewWildcardQuery("name_kana", "*"+nkq+"*").Boost(30),
			)
			authorBQ.Should(
				elastic.NewNestedQuery("authors", elastic.NewPrefixQuery("authors.name_kana", nkq).Boost(80)),
				elastic.NewNestedQuery("authors", elastic.NewWildcardQuery("authors.name_kana", "*"+nkq+"*").Boost(30)),
			)
		}
	}

	ms := r.ES.Session("store").MultiSearch()
	ms.Add(
		elastic.NewSearchRequest().
			Index("web").
			Size(int(req.Limit)).
			Query(productBQ).
			FetchSourceIncludeExclude([]string{"id", "name"}, nil),
		elastic.NewSearchRequest().
			Index("web").
			Size(int(req.Limit)).
			Query(authorBQ).
			FetchSourceIncludeExclude([]string{"authors.id", "authors.name"}, nil),
	)

	multiRes, err := ms.Do(ctx)
	if err != nil {
		return
	}
	for idx, res := range multiRes.Responses {
		switch idx {
		case 0:
			// product
			list.Product = make(entity.ESProductList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Product[i]); err != nil {
					return
				}
			}
		case 1:
			// author
			list.Author = make(entity.ESProductList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Author[i]); err != nil {
					return
				}
			}
		}
	}
	return
}

func (r repo) GetV5PrefixESSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (list entity.MultiESObjectList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	productBQ := elastic.NewBoolQuery()
	authorBQ := elastic.NewBoolQuery()
	for _, v := range queryToSlice(req.Keyword) {
		productBQ2 := elastic.NewBoolQuery()
		authorBQ2 := elastic.NewBoolQuery()
		if nq := normalizeQuery(v); nq != "" {
			nameQ := elastic.NewTermQuery("name", nq).Boost(100)
			productBQ2.Should(nameQ)
			authorBQ2.Should(nameQ)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			nameKanaQ := elastic.NewTermQuery("name_kana", nkq).Boost(80)
			productBQ2.Should(nameKanaQ)
			authorBQ2.Should(nameKanaQ)
		}
		productBQ.Must(productBQ2)
		authorBQ.Must(authorBQ2)
	}

	ms := r.ES.Session("store").MultiSearch()
	ms.Add(
		elastic.NewSearchRequest().
			Index("v5_prefix").
			Size(int(req.Limit)).
			Query(productBQ).
			FetchSourceIncludeExclude([]string{"id", "name"}, nil),
		elastic.NewSearchRequest().
			Index("authors_prefix").
			Size(int(req.Limit)).
			Query(authorBQ).
			FetchSourceIncludeExclude([]string{"id", "name"}, nil),
	)

	multiRes, err := ms.Do(ctx)
	if err != nil {
		return
	}
	for idx, res := range multiRes.Responses {
		switch idx {
		case 0:
			// product
			list.Product = make(entity.ESObjectList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Product[i]); err != nil {
					return
				}
			}
		case 1:
			// author
			list.Author = make(entity.ESObjectList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Author[i]); err != nil {
					return
				}
			}
		}
	}
	return
}

func (r repo) GetV5PartialESSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (list entity.MultiESObjectList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	productBQ := elastic.NewBoolQuery()
	authorBQ := elastic.NewBoolQuery()
	for _, v := range queryToSlice(req.Keyword) {
		productBQ2 := elastic.NewBoolQuery()
		authorBQ2 := elastic.NewBoolQuery()
		if nq := normalizeQuery(v); nq != "" {
			nameQ := elastic.NewMatchPhraseQuery("name", nq).Boost(100)
			productBQ2.Should(nameQ)
			authorBQ2.Should(nameQ)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			nameKanaQ := elastic.NewMatchPhraseQuery("name_kana", nkq).Boost(80)
			productBQ2.Should(nameKanaQ)
			authorBQ2.Should(nameKanaQ)
		}
		productBQ.Must(productBQ2)
		authorBQ.Must(authorBQ2)
	}

	ms := r.ES.Session("store").MultiSearch()
	ms.Add(
		elastic.NewSearchRequest().
			Index("v5_partial").
			Size(int(req.Limit)).
			Query(productBQ).
			FetchSourceIncludeExclude([]string{"id", "name"}, nil),
		elastic.NewSearchRequest().
			Index("authors_partial").
			Size(int(req.Limit)).
			Query(authorBQ).
			FetchSourceIncludeExclude([]string{"id", "name"}, nil),
	)

	multiRes, err := ms.Do(ctx)
	if err != nil {
		return
	}
	for idx, res := range multiRes.Responses {
		switch idx {
		case 0:
			// product
			list.Product = make(entity.ESObjectList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Product[i]); err != nil {
					return
				}
			}
		case 1:
			// author
			list.Author = make(entity.ESObjectList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Author[i]); err != nil {
					return
				}
			}
		}
	}
	return
}

func (r repo) GetV5PrefixAndPartialESSearchSuggestion(ctx context.Context, req entityreq.ProductSearch) (list entity.MultiESObjectList, err error) {
	if req.Keyword == "" {
		return
	}
	if req.Index == "" {
		return
	}

	productBQ := elastic.NewBoolQuery()
	authorBQ := elastic.NewBoolQuery()
	for _, v := range queryToSlice(req.Keyword) {
		productBQ2 := elastic.NewBoolQuery()
		authorBQ2 := elastic.NewBoolQuery()
		if nq := normalizeQuery(v); nq != "" {
			nameQ := elastic.NewTermQuery("name", nq).Boost(100)
			nameWildQ := elastic.NewWildcardQuery("name", "*"+nq+"*").Boost(50)
			productBQ2.Should(nameQ, nameWildQ)
			authorBQ2.Should(nameQ, nameWildQ)
		}
		if nkq := normalizeKanaQuery(v); nkq != "" {
			nameKanaQ := elastic.NewTermQuery("name_kana", nkq).Boost(80)
			nameKanaWildQ := elastic.NewTermQuery("name_kana", "*"+nkq+"*").Boost(30)
			productBQ2.Should(nameKanaQ, nameKanaWildQ)
			authorBQ2.Should(nameKanaQ, nameKanaWildQ)
		}
		productBQ.Must(productBQ2)
		authorBQ.Must(authorBQ2)
	}

	ms := r.ES.Session("store").MultiSearch()
	ms.Add(
		elastic.NewSearchRequest().
			Index("v5_prefix").
			Size(int(req.Limit)).
			Query(productBQ).
			FetchSourceIncludeExclude([]string{"id", "name"}, nil),
		elastic.NewSearchRequest().
			Index("authors_prefix").
			Size(int(req.Limit)).
			Query(authorBQ).
			FetchSourceIncludeExclude([]string{"id", "name"}, nil),
	)

	multiRes, err := ms.Do(ctx)
	if err != nil {
		return
	}
	for idx, res := range multiRes.Responses {
		switch idx {
		case 0:
			// product
			list.Product = make(entity.ESObjectList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Product[i]); err != nil {
					return
				}
			}
		case 1:
			// author
			list.Author = make(entity.ESObjectList, len(res.Hits.Hits))
			for i, hit := range res.Hits.Hits {
				if err = json.Unmarshal(hit.Source, &list.Author[i]); err != nil {
					return
				}
			}
		}
	}
	return
}

func (r repo) InsertSearchSeed(ctx context.Context, list entity.ESProductList) error {
	sess := r.ES.Session("store")

	bulkReq := sess.Bulk()
	// v5_prefix
	for _, v := range list {
		bulkReq = bulkReq.Add(
			elastic.NewBulkIndexRequest().
				Index("v5_prefix").
				Id(cast.ToString(v.ID)).
				Doc(v),
		)
	}
	// v5_partial
	for _, v := range list {
		bulkReq = bulkReq.Add(
			elastic.NewBulkIndexRequest().
				Index("v5_partial").
				Id(cast.ToString(v.ID)).
				Doc(v),
		)
	}
	// web
	for _, v := range list {
		bulkReq = bulkReq.Add(
			elastic.NewBulkIndexRequest().
				Index("web").
				Id(cast.ToString(v.ID)).
				Doc(v),
		)
	}

	if bulkReq.NumberOfActions() > 0 {
		if _, err := bulkReq.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (r repo) InsertSearchAuthorSeed(ctx context.Context, list entity.ESObjectList) error {
	sess := r.ES.Session("store")

	bulkReq := sess.Bulk()
	// authors_prefix
	for _, v := range list {
		bulkReq = bulkReq.Add(
			elastic.NewBulkIndexRequest().
				Index("authors_prefix").
				Id(cast.ToString(v.ID)).
				Doc(v),
		)
	}
	// authors_partial
	for _, v := range list {
		bulkReq = bulkReq.Add(
			elastic.NewBulkIndexRequest().
				Index("authors_partial").
				Id(cast.ToString(v.ID)).
				Doc(v),
		)
	}

	if bulkReq.NumberOfActions() > 0 {
		if _, err := bulkReq.Do(ctx); err != nil {
			return err
		}
	}
	return nil
}

func toKatakanaWithMecab(s string) (string, error) {
	// mecab
	m, err := mecab.New(mecab.YomiOption("mecab-ipadic-neologd"))
	if err != nil {
		return "", err
	}
	defer m.Destroy()

	// mecab tag
	tg, err := m.NewTagger()
	if err != nil {
		return "", err
	}
	defer tg.Destroy()

	lt, err := m.NewLattice(s)
	if err != nil {
		return "", nil
	}

	return util.HiraganaToKatakana(strings.TrimRight(tg.Parse(lt), "\n")), nil
}

func queryToSlice(q string) []string {
	return strings.FieldsFunc(q, unicode.IsSpace)
}

func normalizeQuery(query string) string {
	return util.AlnumToHalfLowwer(query)
}

func normalizeKanaQuery(query string) string {
	k, _ := toKatakanaWithMecab(query)
	return k
}
