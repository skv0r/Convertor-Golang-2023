package main

import (
	"fmt"
)

func main() {
	groupCity := map[int][]string{
		10:   []string{"Село", "Деревня", "ПГТ"},  // города с населением 10-99 тыс. человек
		100:  []string{"Город", "Станица"},        // города с населением 100-999 тыс. человек
		1000: []string{"Мегаполис", "Миллионник"}, // города с населением 1000 тыс. человек и более
	 }
	 
	 cityPopulation := map[string]int{
		"Город":     101,
		"Станица":   102,
		"Село":      103,
		"Мегаполис": 104,
	 }
	


	for x :=10; x<1001; x*=10{
		for i :=0; i < len(groupCity[x]); i++ {
			_, good := cityPopulation[groupCity[x][i]]
			if good {
				if x != 100 {
					delete(cityPopulation, groupCity[x][i])
				} 
			}
		}
	}
}
