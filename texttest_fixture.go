package main

/*

This code was part of the original clone of the repo and served as an example of the output

*/

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Sulfuras, Hand of Ragnaros", -1, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("name, sellIn, quality")
		for _, item := range items {
			fmt.Println(item)
		}
		fmt.Println("")
		UpdateItems(items)
	}
}
