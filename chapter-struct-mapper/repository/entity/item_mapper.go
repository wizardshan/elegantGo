package entity

import (
	"app/chapter-struct-mapper/domain"
)

func (entItem *Item) Mapper() *domain.Item {
	if entItem == nil {
		return nil
	}

	domItem := new(domain.Item)
	domItem.ID = entItem.ID
	domItem.Title = entItem.Title
	domItem.CreateTime = entItem.CreateTime
	domItem.DeleteTime = entItem.DeleteTime

	return domItem
}

func (entItems Items) Mapper() domain.Items {

	size := len(entItems)
	domItems := make(domain.Items, size)

	for i := 0; i < size; i++ {
		domItems[i] = entItems[i].Mapper()
	}

	return domItems
}
