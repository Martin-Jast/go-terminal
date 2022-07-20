package main

type Player struct {
	Inv      Inventory
	Cash     int
	EqpSlots []int
	*Creature
}

// EqpSlots
// 1,2 hands
// 3 body
// 4 pants
// 5 boots
// 0 not equipable
func InitPlayer() *Player {
	return &Player{
		Creature: InitCreature(10, 10),
		EqpSlots: []int{1, 2, 3, 4, 5},
		Inv: Inventory{
			Items: []InventoryItem{
				{
					Item: &Item{
						Name:  "broad sword",
						Slot:  1,
						Power: 2,
					},
				},
				{
					Item: &Item{
						Name:  "wood shield",
						Slot:  2,
						Power: 1,
					},
				},
			},
		},
	}
}

func (p *Player) AddItems(items []Item) {
	for i := 0; i < len(items); i++ {
		p.Inv.Items = append(p.Inv.Items, *items[i].AsInventoryItem(false))
	}
}

func (p Player) GetTotalPower() int {
	total := 0
	for i := 0; i < len(p.Inv.Items); i++ {
		if p.Inv.Items[i].IsEquipped {
			total += p.Inv.Items[i].Power
		}
	}
	return total
}

func (p Player) GetEquipedItems() []InventoryItem {
	a := []InventoryItem{}
	for i := 0; i < len(p.Inv.Items); i++ {
		if p.Inv.Items[i].IsEquipped {
			a = append(a, p.Inv.Items[i])
		}
	}
	return a
}

func (p *Player) EquipInventoryItem(item *InventoryItem) bool {
	equipedItems := p.GetEquipedItems()
	emptySlots := []int{}
	for i := 0; i < len(p.EqpSlots); i++ {
		found := false
		for j := 0; j < len(equipedItems); j++ {
			if equipedItems[j].Slot == p.EqpSlots[i] {
				found = true
			}
		}
		if !found {
			emptySlots = append(emptySlots, p.EqpSlots[i])
		}
	}
	for t := 0; t < len(emptySlots); t++ {
		if emptySlots[t] == item.Slot {
			item.IsEquipped = true
			return true
		}
	}
	return false
}
