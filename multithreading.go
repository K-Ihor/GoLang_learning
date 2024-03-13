// package main

// import (
// 	"fmt"
// 	"io"
// 	"log"
// 	"net/http"
// )

// type Page struct {
// 	URL  string
// 	Size int
// }

// func main() {

// 	// // создаем переменную для канала и cоздаем сам канал
// 	// myChannel := make(chan float64)
// 	// myChannel <- 3.14 // передаем запись в канал
// 	// qwe := <-myChannel // чтоб из канала то запись в обратную строну;

// 	pages := make(chan Page) // создаем канал
// 	urls := []string{
// 		"https://golang.org",
// 		"https://golang.org/doc",
// 		"https://exemple.com",
// 	}

// 	for _, url := range urls {
// 		go responseSize(url, pages)
// 	}

// 	// time.Sleep(3 * time.Second)

// 	for i := 0; i < len(urls); i++ {
// 		page := <-pages
// 		fmt.Println(page.URL, " ", page.Size)
// 		// fmt.Println(<-sizes)
// 	}
// }

// func responseSize(url string, chanal chan Page) {
// 	fmt.Println("Getting", url)
// 	response, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	defer response.Body.Close()
// 	body, err := io.ReadAll(response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// fmt.Println(len(body))
// 	chanal <- Page{URL: url, Size: len(body)} // в сьруктуре скобочки фигурные

// }
