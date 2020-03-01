package usecase

import (
	"github.com/jiro94/elasticsearch-sample/api/domain/entityres"
	"github.com/jiro94/elasticsearch-sample/internal/util"
)

// NewItems
func NewItems(list interface{}) entityres.Items {
	return entityres.Items{
		Items: util.SliceInterface(list, 0),
	}
}
