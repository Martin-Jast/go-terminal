package main

type Item struct {
	Name   string
	Type   string
	Weight int
	Power  int
	Slot   int
}

func (i *Item) AsInventoryItem(equipped bool) *InventoryItem {
	return &InventoryItem{
		Item:       i,
		IsEquipped: equipped,
	}
}
