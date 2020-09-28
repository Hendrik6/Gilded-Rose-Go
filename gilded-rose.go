package gildedrose

import "strings"

type Item struct {
	name       string
	daysToSell int
	quality    int
}

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

		if item.name == "Sulfuras, Hand of Ragnaros" {
			continue
		}

		item.NormalDegrade()
		continue
	}
}

func (item *Item) UpdateAgedBrie() {
	item.increaseQuality()
	item.daysToSell--

	//This logic is solely for the Aged Brie since Aged Brie increases in Quality twice as fast if it's expired (daysToSell < 0)
	if item.daysToSell < 0 {
		item.increaseQuality()
	}
}

func (item *Item) NormalDegrade() {
	item.decreaseQuality()

	if strings.Contains(item.name, "Conjured") {
		item.decreaseQuality()
	}

	item.daysToSell--
	item.checkExpiration()
}

func (item *Item) checkExpiration() {
	if item.daysToSell < 0 && item.quality > 0 {
		item.decreaseQuality()
	}
}

func (item *Item) UpdateBackstagePasses() {
	item.increaseQuality()

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

func (item *Item) increaseQuality() {
	if item.quality < 50 {
		item.quality++
	}
}

func (item *Item) decreaseQuality() {
	if item.quality > 0 {
		item.quality--
	}
}
