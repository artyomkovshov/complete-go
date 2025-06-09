package exercise

import "fmt"

// Starting code for the exercise

// Item struct
type Item struct {
	Name string
	Type string
}

// Player struct
type Player struct {
	Name      string
	Inventory []Item
}

// Pick up an item (modifies the player's inventory)
func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s\n", p.Name, item.Name)
}

// Drop an item (removes from inventory)
func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			// Remove item from inventory
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s\n", p.Name, item.Name)
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory\n", p.Name, itemName)
}

// Use an item (if potion, remove it after use)
func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "potion" {
				// Remove potion from inventory
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
				fmt.Printf("%s used potion %s and it has been removed from inventory\n", p.Name, item.Name)
			} else {
				fmt.Printf("%s used item %s\n", p.Name, item.Name)
			}
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory\n", p.Name, itemName)
}
