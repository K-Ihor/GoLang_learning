// package main

// import (
// 	"fmt"
// 	"log"

// 	"golangtest/datafile"
// )

// func main() {

// 	lines, err := datafile.GetStrings("data1.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	counts := make(map[string]int)

// 	for _, line := range lines {
// 		counts[line]++
// 	}

// 	fmt.Print(counts)

// }