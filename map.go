// // Карты подобие словарей

// package main

// import (
// 	"fmt"
// 	"sort"
// )

// func main() {

// 	langs := map[string]float64{
// 		"java":   100,
// 		"python": 70,
// 		"go":     50,
// 	}

// 	var names []string
// 	for name := range langs {
// 		names = append(names, name)
// 	}

// 	sort.Strings(names)

// 	for _, value := range names {
// 		fmt.Println(value, langs[value])
// 	}

// 	// delete(myMap, "a") // удаление по ключу a

// }
