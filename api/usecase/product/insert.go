package product

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"path"

	"github.com/jiro94/elasticsearch-sample/api/domain/entity"
)

func (s serv) InsertSearchSeed(ctx context.Context) error {
	titles, err := getSeedObjects(path.Join(seedBasePath, productsSeedFileName))
	if err != nil {
		return err
	}

	authors, err := getSeedObjects(path.Join(seedBasePath, authorsSeedFileName))
	if err != nil {
		return err
	}
	idAuthorsMap := authors.ToIDAuthorsMap()

	relations, err := getSeedRelations(path.Join(seedBasePath, productsAuthorsSeedFileName))
	if err != nil {
		return err
	}
	titleAuthorsMap := relations.ToTitleAuthorMap()

	products := make(entity.ESProductList, len(titles))
	for i, v := range titles {
		var p entity.ESProduct

		p.ID = v.ID
		p.Name = v.Name
		p.NameKana = v.NameKana

		p.Authors = idAuthorsMap.GetAuthorsByIDs(titleAuthorsMap[p.ID])

		products[i] = p
	}

	var (
		start int
		end   int
		limit = 1000
		count = 1
	)
	//for {
	//	end += start + (count * limit)
	//	if len(products) > end {
	//		if err := s.repo.InsertSearchSeed(ctx, products[start:end]); err != nil {
	//			return err
	//		}
	//	} else {
	//		if err := s.repo.InsertSearchSeed(ctx, products[start:]); err != nil {
	//			return err
	//		}
	//		break
	//	}
	//	start = end
	//	count++
	//}

	for {
		end += start + (count * limit)
		if len(authors) > end {
			if err := s.repo.InsertSearchAuthorSeed(ctx, authors[start:end]); err != nil {
				return err
			}
		} else {
			if err := s.repo.InsertSearchAuthorSeed(ctx, authors[start:]); err != nil {
				return err
			}
			break
		}
		start = end
		count++
	}

	return nil
}

func getSeedRelations(seedFilePath string) (relations entity.RelationList, err error) {
	buf, err := readFile(seedFilePath)
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf, &relations); err != nil {
		return
	}

	return
}

func getSeedObjects(seedFilePath string) (objects entity.ESObjectList, err error) {
	buf, err := readFile(seedFilePath)
	if err != nil {
		return
	}

	if err = json.Unmarshal(buf, &objects); err != nil {
		return
	}

	return
}

func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return ioutil.ReadAll(f)
}
