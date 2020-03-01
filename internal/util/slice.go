package util

import "reflect"

func SliceInterface(i interface{}, limit int) []interface{} {
	itemsVal := reflect.ValueOf(i)
	if itemsVal.Kind() == reflect.Slice {
		if limit > 0 && limit < itemsVal.Len() {
			itemsVal = itemsVal.Slice(0, limit)
		}

		itemList := make([]interface{}, itemsVal.Len())
		for k := 0; k < itemsVal.Len(); k++ {
			itemList[k] = itemsVal.Index(k).Interface()
		}
		return itemList
	}
	return nil
}
