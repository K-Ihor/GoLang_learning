// package main

// import (
// 	"fmt"
// 	"log"

// 	"golangtest/datafile"
// )

// func main() {
// 	numbers, err := datafile.GetFloats("data.txt")

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	var sum float64 = 0

// 	for _, nnumber := range numbers {
// 		sum += nnumber
// 	}

// 	count := float64(len(numbers))
// 	fmt.Print(sum / count)

// }

// -----------------------------------------------

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// )

// func main() {

// 	file, err := os.Open("data.txt")
// 	if err != nil {
// 		log.Fatal()
// 	}

// 	scanner := bufio.NewScanner(file)

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}

// 	err = file.Close()

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	if scanner.Err() != nil {
// 		log.Fatal(scanner.Err())
// 	}

// }

// // go build hello.go - скомпилить
// // ./hello  - запустить скомпиленный вайл
// // "strconv" - просто запуск кода без компиляции
