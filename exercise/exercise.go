package main

import(
	"fmt"
)

type Player struct {
	Name string
	Inventory []Item
}
type Item struct{
	Name string
	Type string
}

func (p *Player) PickUpItem(item Item) {
	p.Inventory = append(p.Inventory, item)
	fmt.Printf("%s picked up %s!\n", p.Name, item.Name)
}

func (p *Player) DropItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]...)
			fmt.Printf("%s dropped %s!\n", p.Name, item.Name)
			return
		}
	}
	fmt.Printf("%s does not have %s in inventory. \n", p.Name, itemName)
}

func (p *Player) UseItem(itemName string) {
	for i, item := range p.Inventory {
		if item.Name == itemName {
			if item.Type == "Consumable" {
				fmt.Printf("%s used %s!\n", p.Name, item.Name)
				p.Inventory = append(p.Inventory[:i], p.Inventory[i+1:]... )
			} else {
				fmt.Printf("%s used %s.\n", p.Name, item.Name)
			}
			return
		}
	}
	fmt.Printf("%s don't have %s in inventory.\n", p.Name, itemName)
}


func main() {
	// Create a player
	player := Player{Name: "Hero"}

	// Create some items
	sword := Item{Name: "Sword", Type: "Weapon"}
	potion := Item{Name: "Health Potion", Type: "Consumable"}
	shield := Item{Name: "Shield", Type: "Armor"}

	// Pick up some items
	player.PickUpItem(sword)
	player.PickUpItem(potion)
	player.PickUpItem(shield)

	// Use some items
	player.UseItem("Health Potion")
	player.UseItem("Sword")

	// Drop an item
	player.DropItem("Shield")

	// Try using an item not in the inventory
	player.UseItem("Bow") // This item isn't in the inventory

	// Try dropping an item that doesn't exist
	player.DropItem("Bow") // This item isn't in the inventory
}

