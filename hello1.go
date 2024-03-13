// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"log"
// 	"math/rand"
// 	"os"
// 	"strconv"
// 	"strings"
// 	//"time"
// )

// func main() {

// 	// second := time.Now().Unix()

// 	target := rand.Intn(100) + 1 // +1 Для того чтоб небыло случайно нуля

// 	fmt.Println("Я выбираю число от 1 до 100")
// 	fmt.Println("Число выбрано")
// 	// fmt.Println(target)

// 	reader := bufio.NewReader(os.Stdin)

// 	success := false

// 	for guesses := 0; guesses < 10; guesses++ {

// 		fmt.Println("У вас осталось ", 10-guesses, " попыток")

// 		fmt.Println("Напишите ваше число: ")
// 		input, err := reader.ReadString('\n') // Читаем наш ввод до '\n'
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		input = strings.TrimSpace(input) // перезапись чтоб избавиться от табуляции и переноса строки

// 		guess, err := strconv.Atoi(input) // конвертация в int
// 		if err != nil {
// 			log.Fatal(err)
// 		}

// 		if guess > target {
// 			fmt.Println("Ваше число больше")
// 		} else if guess < target {
// 			fmt.Println("Ваше число меньше")
// 		} else {
// 			success = true
// 			fmt.Println("Поздравляю вы угадали")
// 			break
// 		}
// 	}

// 	if !success {
// 		fmt.Println("Попытки кончелись. Число было ", target)
// 	}

// }

// // go build hello.go - скомпилить
// // ./hello  - запустить скомпиленный вайл
// // "strconv" - просто запуск кода без компиляции
