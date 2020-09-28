package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_UpdateQuality(t *testing.T) {
	//The first item in the list is continously being updated
	//The expected outcome has been entered for the upcoming days
	//We keep updating the first item, and keep checking if it's equal to the upcoming day
	for day, item := range items[1:] {
		UpdateQuality(items[0])
		assert.Equal(t, item, items[0], "UpdateQuality failure : day %d", day)
	}
}

//These 11 testcases cover 11 consecutive days in which the item quality is being updated.
var items = [][]*Item{
	{
		&Item{"+5 Dexterity Vest", 10, 20},
		&Item{"Aged Brie", 2, 0},
		&Item{"Elixir of the Mongoose", 5, 7},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		&Item{"Conjured Mana Cake", 3, 6},
	},
	{
		&Item{"+5 Dexterity Vest", 9, 19},
		&Item{"Aged Brie", 1, 1},
		&Item{"Elixir of the Mongoose", 4, 6},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		&Item{"Conjured Mana Cake", 2, 5},
	}, {
		&Item{"+5 Dexterity Vest", 8, 18},
		&Item{"Aged Brie", 0, 2},
		&Item{"Elixir of the Mongoose", 3, 5},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 8, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		&Item{"Conjured Mana Cake", 1, 4},
	}, {
		&Item{"+5 Dexterity Vest", 7, 17},
		&Item{"Aged Brie", -1, 4},
		&Item{"Elixir of the Mongoose", 2, 4},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 12, 23},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 7, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 2, 50},
		&Item{"Conjured Mana Cake", 0, 3},
	}, {
		&Item{"+5 Dexterity Vest", 6, 16},
		&Item{"Aged Brie", -2, 6},
		&Item{"Elixir of the Mongoose", 1, 3},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 11, 24},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 6, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 1, 50},
		&Item{"Conjured Mana Cake", -1, 1},
	}, {
		&Item{"+5 Dexterity Vest", 5, 15},
		&Item{"Aged Brie", -3, 8},
		&Item{"Elixir of the Mongoose", 0, 2},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		&Item{"Conjured Mana Cake", -2, 0},
	},
	{
		&Item{"+5 Dexterity Vest", 4, 14},
		&Item{"Aged Brie", -4, 10},
		&Item{"Elixir of the Mongoose", -1, 0},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 9, 27},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		&Item{"Conjured Mana Cake", -3, 0},
	},
	{
		&Item{"+5 Dexterity Vest", 3, 13},
		&Item{"Aged Brie", -5, 12},
		&Item{"Elixir of the Mongoose", -2, 0},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 8, 29},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
		&Item{"Conjured Mana Cake", -4, 0},
	},
	{
		&Item{"+5 Dexterity Vest", 2, 12},
		&Item{"Aged Brie", -6, 14},
		&Item{"Elixir of the Mongoose", -3, 0},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 7, 31},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 2, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", -3, 0},
		&Item{"Conjured Mana Cake", -5, 0},
	},
	{
		&Item{"+5 Dexterity Vest", 1, 11},
		&Item{"Aged Brie", -7, 16},
		&Item{"Elixir of the Mongoose", -4, 0},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 6, 33},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 1, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", -4, 0},
		&Item{"Conjured Mana Cake", -6, 0},
	},
	{
		&Item{"+5 Dexterity Vest", 0, 10},
		&Item{"Aged Brie", -8, 18},
		&Item{"Elixir of the Mongoose", -5, 0},
		&Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
		&Item{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		&Item{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
		&Item{"Conjured Mana Cake", -7, 0},
	},
}
