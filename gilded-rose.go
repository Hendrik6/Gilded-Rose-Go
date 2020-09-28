package gildedrose

import "strings"

//An Item exists of a name, the amount of days in which it has to be sold, and the current quality of the item.
type Item struct {
	name       string
	daysToSell int
	quality    int
}

//UpdateQuality will handle all the given items by setting the quality and daysToSell
func UpdateQuality(items []*Item) {
	for _, item := range items {

		if item.name == "Aged Brie" {
			item.UpdateAgedBrie()
			continue
		}

		if item.name == "Backstage passes to a TAFKAL80ETC concert" {
			item.UpdateBackstagePasses()
			continue
		}

		// Sulfuras, Hand of Ragnaros, being a legendary item, never has to be sold or decreases in Quality
		// So we can continue with the other items in the list
		if item.name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		//If the given item isn't an edge case, we can use the Normal Degrading standards.
		item.RegularDegrade()
		continue
	}
}

//UpdateAgedBrie will update the Aged Brie in the register and checking if it's expired which causes it to rise in value faster.
func (item *Item) UpdateAgedBrie() {
	item.increaseQuality()
	item.daysToSell--

	// Aged Brie increases in Quality twice as fast if it's expired (daysToSell < 0)
	if item.daysToSell < 0 {
		item.increaseQuality()
	}
}

//UpdateBackstagePasses will always increase the Quality of the passes if it's not > 50.
//It will also increase quality by 1 if the daysToSell value is < 11 and by 2 if this value is < 6
//If the daysToSell < 0 (the concert has passed) the value drops to 0
func (item *Item) UpdateBackstagePasses() {
	item.increaseQuality()

	//Could nest these 2 if's but it's more readable like this I think..
	if item.daysToSell < 11 {
		item.increaseQuality()
	}

	if item.daysToSell < 6 {
		item.increaseQuality()
	}

	item.daysToSell--

	if item.daysToSell < 0 {
		item.quality = 0
		return
	}
}

//RegularDegrade An item quality degrades by 1 in both daysToSell and Value every day.
func (item *Item) RegularDegrade() {
	item.decreaseQuality()
	item.daysToSell--
	item.checkExpiration()
}

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

	//"Conjured" items degrade in Quality twice as fast as normal items so decreaseQuality needs to be called again if this is true
	//At first I had placed this piece of code in the RegularDegrade() function. But if a conjured item needs to be added that isn't of normal quality
	//this will always check for Conjured items so if it's a Conjured Legendary item, the quality will still degrade faster.
	if item.isConjured() {
		if item.quality > 0 {
			item.quality--
		}
	}
}

func (item *Item) isConjured() bool {
	return strings.Contains(item.name, "Conjured")
}
