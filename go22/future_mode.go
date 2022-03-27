package main

import (
	"fmt"
	"time"
)

// 洗菜
func washVegetables() <-chan string {
	vegetables := make(chan string)

	go func() {
		time.Sleep(5 * time.Second)
		vegetables <- "washed vegetable"
	}()
	return vegetables
}

// 烧水
func boilWater() <-chan string {
	water := make(chan string)
	go func() {
		time.Sleep(5 * time.Second)
		water <- "boiled water"
	}()
	return water
}

func main() {
	vegetablesCh := washVegetables()
	waterCh := boilWater()
	fmt.Println("Wash and boil starts..")
	fmt.Println("rest for a while...")
	time.Sleep(2 * time.Second)
	fmt.Println("rest complete, have the water and vegetable prepared?")
	vegetables := <-vegetablesCh
	water := <-waterCh
	fmt.Println(vegetables)
	fmt.Println(water)
	fmt.Println("preparation complete")
}
