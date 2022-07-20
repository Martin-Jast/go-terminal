package main

type Inventory struct {
	Items    []InventoryItem
	Equipped []InventoryItem
}

type InventoryItem struct {
	*Item
	IsEquipped bool
}

func (in Inventory) getItemsName() []string {
	names := []string{}
	for i := 0; i < len(in.Items); i++ {
		names = append(names, in.Items[i].Name)
	}
	return names
}
