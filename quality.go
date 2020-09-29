package main

import "strings"

//If the item is expired (past it's days to sell), it needs to decrease faster in quality
func (item *Item) checkExpiration() {
	if item.daysToSell < 0 && item.quality > 0 {
		item.decreaseQuality()
	}
}

//The Quality of an item is never more than 50
func (item *Item) increaseQuality() {
	if item.quality < 50 {
		item.quality++
	}
}

//Quality can never be lower than 0
func (item *Item) decreaseQuality() {
	if item.quality > 0 {
		item.quality--
	}

	//"Conjured" items degrade in Quality twice as fast as normal items so the quality value needs to be subtracted by 1 again if this is true
	//At first I had placed this piece of code in the RegularDegrade() function. But if a conjured item is added that isn't of normal quality this won't trigger
	//This will always check for Conjured items so if it's a Conjured item of other rarity, the quality will still degrade faster .
	if item.isConjured() {
		if item.quality > 0 {
			item.quality--
		}
	}
}

func (item *Item) isConjured() bool {
	return strings.Contains(item.name, "Conjured")
}
