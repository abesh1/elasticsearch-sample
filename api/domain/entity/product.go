package entity

type ESProduct struct {
	ID       uint64       `json:"id"`
	Name     string       `json:"name"`
	NameKana string       `json:"name_kana"`
	Authors  ESObjectList `json:"authors"`
}

type ESProductList []ESProduct

type ESObject struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	NameKana string `json:"name_kana"`
}

type ESObjectList []ESObject

type MultiESObjectList struct {
	Product ESObjectList
	Author  ESObjectList
}

type MultiESProductList struct {
	Product ESProductList
	Author  ESProductList
}

type IDAuthorsMap map[uint64]ESObject

func (l ESObjectList) ToIDAuthorsMap() IDAuthorsMap {
	m := make(IDAuthorsMap, len(l))
	for _, v := range l {
		m[v.ID] = v
	}
	return m
}

func (m IDAuthorsMap) GetAuthorsByIDs(ids []uint64) ESObjectList {
	l := make(ESObjectList, len(ids))
	for i, v := range ids {
		l[i] = m[v]
	}
	return l
}

type Relation struct {
	TitleID  uint64 `json:"title_id"`
	AuthorID uint64 `json:"author_id"`
}

type RelationList []Relation

type TitleAuthorsMap map[uint64][]uint64

func (l RelationList) ToTitleAuthorMap() TitleAuthorsMap {
	m := make(TitleAuthorsMap)
	for _, v := range l {
		m[v.TitleID] = append(m[v.TitleID], v.AuthorID)
	}
	return m
}
