// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"os"
// 	"strconv"
// 	"strings"
// )

// func main() {
// 	fmt.Print("Cколько вам лет: ")

// 	reader := bufio.NewReader(os.Stdin)

// 	input, err := reader.ReadString('\n')

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	input = strings.TrimSpace(input) // перезапись drop "\n"

// 	old, err := strconv.ParseFloat(input, 64)

// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	result := "" // обявление переменной :=

// 	if old >= 18 {
// 		result = "Достаточно" // перезапись без двух точек просто =
// 	} else {
// 		result = "НЕ достаточно"
// 	}

// 	fmt.Print("Ваш возраст : " + result)

// }

// go build hello.go - скомпилить
// ./hello  - запустить скомпиленный вайл
// "strconv" - просто запуск кода без компиляции
