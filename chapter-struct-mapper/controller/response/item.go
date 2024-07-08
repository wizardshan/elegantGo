package response

import (
	"elegantGo/chapter-struct-mapper/domain"
	"time"
)

type Items []*Item

type Item struct {
	ID         int        `json:"id"`
	Title      string     `json:"title"`
	CreateTime time.Time  `json:"createTime"`
	DeleteTime *time.Time `json:"deleteTime"`
}

func (respItem *Item) Mapper(domItem *domain.Item) *Item {
	if domItem == nil {
		return nil
	}
	respItem.ID = domItem.ID
	respItem.Title = domItem.Title
	respItem.CreateTime = domItem.CreateTime
	respItem.DeleteTime = domItem.DeleteTime
	return respItem
}

func (respItems Items) Mapper(domItems domain.Items) Items {

	size := len(domItems)
	respItems = make(Items, size)
	for i := 0; i < size; i++ {
		var respItem Item
		respItems[i] = respItem.Mapper(domItems[i])
	}

	return respItems
}
