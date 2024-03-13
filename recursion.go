// package main

// import (
// 	"fmt"
// 	"os"
// 	"path/filepath"
// )

// func reportPanic() {
// 	p := recover() // для того чтоб останавливать панику, для того чтоб востановить управление програмой
// 	if p == nil {
// 		return
// 	}
// 	err, ok := p.(error)
// 	if ok {
// 		fmt.Println(err)
// 	} else {
// 		panic(p)
// 	}
// }

// func scanDirectory(path string) {
// 	fmt.Println(path)
// 	files, err := os.ReadDir(path)
// 	if err != nil {
// 		panic(err)
// 	}

// 	for _, file := range files {
// 		filePath := filepath.Join(path, file.Name())

// 		if file.IsDir() {
// 			scanDirectory(filePath)
// 		} else {
// 			fmt.Println(filePath)
// 		}
// 	}
// }

// func main() {
// 	defer reportPanic() // defer - отложенный вызов ф-ции
// 	scanDirectory("my_directory")

// 	// ------------------------------------------------
// 	// files, err := os.ReadDir("my_directory")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }

// 	// for _, file := range files {
// 	// 	if file.IsDir() {
// 	// 		fmt.Println("Dir: ", file.Name())
// 	// 	} else {
// 	// 		fmt.Println("File: ", file.Name())
// 	// 	}
// 	// }
// 	// ----------------------------------------------------

// }
