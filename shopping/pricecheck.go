package shopping

import (
	"shopping/db"
)

func PriceCheck(itemId int) (float64, bool) {
	item := db.LoadItem(itemId)
	//item is a pointer
	if item == nil {
		return 0, false
	}
	return item.Price, true
}
