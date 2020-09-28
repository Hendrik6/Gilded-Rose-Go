package gildedrose

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

func Test_UpdateBackStagePasses(t *testing.T) {
	var inputs = []*Item{
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 15},
		{"Backstage passes to a TAFKAL80ETC concert", 12, 13},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 8},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 1},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 15},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 48},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 3},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
	}

	var expected = []*Item{
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 18},
		{"Backstage passes to a TAFKAL80ETC concert", 11, 14},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 10},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 4},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 2, 6},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
	}

	for i, input := range inputs {
		input.UpdateBackstagePasses()
		assert.Equal(t, input, expected[i], "Test_UpdateBackStagePasses failure : item %d, expected %d", input, expected[i])
	}
}

func Test_UpdateAgedBrie(t *testing.T) {
	var inputs = []*Item{
		{"Aged Brie", 15, 20},
		{"Aged Brie", 4, 49},
		{"Aged Brie", 3, 15},
		{"Aged Brie", 2, 13},
		{"Aged Brie", 1, 8},
		{"Aged Brie", 0, 1},
		{"Aged Brie", -1, 15},
		{"Aged Brie", -2, 49},
		{"Aged Brie", -3, 50},
		{"Aged Brie", -4, 47},
	}

	var expected = []*Item{
		{"Aged Brie", 14, 21},
		{"Aged Brie", 3, 50},
		{"Aged Brie", 2, 16},
		{"Aged Brie", 1, 14},
		{"Aged Brie", 0, 9},
		{"Aged Brie", -1, 3},
		{"Aged Brie", -2, 17},
		{"Aged Brie", -3, 50},
		{"Aged Brie", -4, 50},
		{"Aged Brie", -5, 49},
	}

	for i, input := range inputs {
		input.UpdateAgedBrie()
		assert.Equal(t, input, expected[i], "Test_UpdateAgedBrie failure : item %d, expected %d", input, expected[i])
	}
}

func Test_RegularItems(t *testing.T) {
	var inputs = []*Item{
		{"+5 Dexterity Vest", 10, 20},
		{"+5 Dexterity Vest", 2, 5},
		{"+5 Dexterity Vest", -1, 3},
		{"Elixir of the Mongoose", 5, 7},
		{"Elixir of the Mongoose", 3, 6},
		{"Elixir of the Mongoose", 1, 3},
		{"Conjured Mana Cake", 4, 12},
		{"Conjured Mana Cake", 2, 48},
		{"Conjured Mana Cake", 1, 2},
	}

	var expected = []*Item{
		{"+5 Dexterity Vest", 9, 19},
		{"+5 Dexterity Vest", 1, 4},
		{"+5 Dexterity Vest", -2, 1},
		{"Elixir of the Mongoose", 4, 6},
		{"Elixir of the Mongoose", 2, 5},
		{"Elixir of the Mongoose", 0, 2},
		{"Conjured Mana Cake", 3, 10},
		{"Conjured Mana Cake", 1, 46},
		{"Conjured Mana Cake", 0, 0},
	}

	for i, input := range inputs {
		input.RegularDegrade()
		assert.Equal(t, input, expected[i], "Test_RegularItems failure : item %d, expected %d", input, expected[i])
	}
}

func Test_SulfurasHandOfRagnaros(t *testing.T) {
	var inputs = []*Item{
		{"Sulfuras, Hand of Ragnaros", 2, 80},
		{"Sulfuras, Hand of Ragnaros", 1, 80},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Sulfuras, Hand of Ragnaros", 4, 80},
	}

	var expected = []*Item{
		{"Sulfuras, Hand of Ragnaros", 2, 80},
		{"Sulfuras, Hand of Ragnaros", 1, 80},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Sulfuras, Hand of Ragnaros", 4, 80},
	}

	UpdateQuality(inputs)

	for i, input := range inputs {
		assert.Equal(t, input, expected[i], "Test_SulfurasHandOfRagnaros failure : item %d, expected %d", input, expected[i])
	}
}

//These 11 testcases cover 11 consecutive days in which the item quality is being updated.
var items = [][]*Item{
	{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6},
	},
	{
		{"+5 Dexterity Vest", 9, 19},
		{"Aged Brie", 1, 1},
		{"Elixir of the Mongoose", 4, 6},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 14, 21},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Conjured Mana Cake", 2, 4},
	}, {
		{"+5 Dexterity Vest", 8, 18},
		{"Aged Brie", 0, 2},
		{"Elixir of the Mongoose", 3, 5},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 13, 22},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		{"Conjured Mana Cake", 1, 2},
	}, {
		{"+5 Dexterity Vest", 7, 17},
		{"Aged Brie", -1, 4},
		{"Elixir of the Mongoose", 2, 4},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 12, 23},
		{"Backstage passes to a TAFKAL80ETC concert", 7, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 2, 50},
		{"Conjured Mana Cake", 0, 0},
	}, {
		{"+5 Dexterity Vest", 6, 16},
		{"Aged Brie", -2, 6},
		{"Elixir of the Mongoose", 1, 3},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 11, 24},
		{"Backstage passes to a TAFKAL80ETC concert", 6, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 50},
		{"Conjured Mana Cake", -1, 0},
	}, {
		{"+5 Dexterity Vest", 5, 15},
		{"Aged Brie", -3, 8},
		{"Elixir of the Mongoose", 0, 2},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 25},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 50},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Conjured Mana Cake", -2, 0},
	},
	{
		{"+5 Dexterity Vest", 4, 14},
		{"Aged Brie", -4, 10},
		{"Elixir of the Mongoose", -1, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 9, 27},
		{"Backstage passes to a TAFKAL80ETC concert", 4, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -1, 0},
		{"Conjured Mana Cake", -3, 0},
	},
	{
		{"+5 Dexterity Vest", 3, 13},
		{"Aged Brie", -5, 12},
		{"Elixir of the Mongoose", -2, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 8, 29},
		{"Backstage passes to a TAFKAL80ETC concert", 3, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -2, 0},
		{"Conjured Mana Cake", -4, 0},
	},
	{
		{"+5 Dexterity Vest", 2, 12},
		{"Aged Brie", -6, 14},
		{"Elixir of the Mongoose", -3, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 7, 31},
		{"Backstage passes to a TAFKAL80ETC concert", 2, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -3, 0},
		{"Conjured Mana Cake", -5, 0},
	},
	{
		{"+5 Dexterity Vest", 1, 11},
		{"Aged Brie", -7, 16},
		{"Elixir of the Mongoose", -4, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 6, 33},
		{"Backstage passes to a TAFKAL80ETC concert", 1, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -4, 0},
		{"Conjured Mana Cake", -6, 0},
	},
	{
		{"+5 Dexterity Vest", 0, 10},
		{"Aged Brie", -8, 18},
		{"Elixir of the Mongoose", -5, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 35},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
		{"Conjured Mana Cake", -7, 0},
	},
}
