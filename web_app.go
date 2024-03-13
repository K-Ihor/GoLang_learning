package main

import (
	"fmt"
	"golangtest/datafile"
	"html/template"
	"log"
	"net/http"
	"os"
)

type Guestbook struct {
	SignatureCount int
	Signature      []string
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newHandler(writer http.ResponseWriter, request *http.Request) {
	html, err := template.ParseFiles("new.html")
	check(err)
	err = html.Execute(writer, nil)
	check(err)
}

func createHandler(writer http.ResponseWriter, request *http.Request) {
	signature := request.FormValue("signature")

	options := os.O_WRONLY | os.O_APPEND | os.O_CREATE
	file, err := os.OpenFile("signatures.txt", options, os.FileMode(0600))
	check(err)
	_, err = fmt.Fprintln(file, signature) // метод для записи
	check(err)
	err = file.Close()
	// _, err := writer.Write([]byte(signature))
	check(err)
	http.Redirect(writer, request, "/guestbook", http.StatusFound)
}

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	signature, err := datafile.GetStrings("signatures.txt")
	check(err)
	html, err := template.ParseFiles("view.html")
	check(err)

	guestbook := Guestbook{
		SignatureCount: len(signature),
		Signature:      signature,
	}

	err = html.Execute(writer, guestbook)
	check(err)
	// placeholder := []byte("Оставте вашу запись: ") // инф которую мы передаем в браузере дложна храниться в типе byte
	// _, err := writer.Write(placeholder)            // _ - кол байт и может вернуть ошибку err
	// check(err)

}

func main() {
	http.HandleFunc("/guestbook", viewHandler) // типа роут который тригерит функцию viewHandler
	http.HandleFunc("/guestbook/new", newHandler)
	http.HandleFunc("/guestbook/create", createHandler)
	// зпустил на порту 8081 так как на 8080 у иеня дженкинс
	err := http.ListenAndServe("localhost:8081", nil) // запускаем наш сервер(работает в цикле бесконечно до ошибки)
	log.Fatal(err)

}
