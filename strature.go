// Структуры включают в себя разные типы данных можно определять свои типы на базе структуры
// var subscriber struct {     // структура
// 	name string
// 	rate float64
// 	active bool
// }

// subscriber.name = "Alex"
// subscriber.rate = 11.00
// subscriber.active = true
// package main

// import "fmt"

// type Subscriber struct { // определили свой тип данных subscriber
// 	Name       string
// 	Rate       float64
// 	Active     bool
// 	HomaAdress Adress
// }

// type Adress struct {
// 	Street   string
// 	City     string
// 	Building string
// }

// func printInfo(s Subscriber) {
// 	fmt.Println("Name: ", s.Name)
// 	fmt.Println("Rate: ", s.Rate)
// 	fmt.Println("Active: ", s.Active)
// 	fmt.Println("Adress: ", s.HomaAdress)
// }

// func defaultSubscriber(name string) Subscriber {
// 	var s Subscriber
// 	s.Name = name
// 	s.Rate = 11.99
// 	s.Active = true
// 	s.HomaAdress.City = "London"
// 	return s
// }

// func applyDiscount(s *Subscriber) { // * - указатель на обект если бередадим без * а подписчика то мы передадим его копию
// 	s.Rate = 4.99
// }

// func main() {

// 	adress := Adress{
// 		Street:   "Pushina",
// 		City:     "Odessa",
// 		Building: "1",
// 	}

// 	subscriber1 := defaultSubscriber("Alex")
// 	subscriber1.HomaAdress = adress

// 	applyDiscount(&subscriber1) //  & - когда передаем указатель !!!!
// 	printInfo(subscriber1)

// 	// var subscriber1 subscriber
// 	// subscriber1.name = "Alex"

// 	// var subscriber2 subscriber
// 	// subscriber2.name = "Alena"

// }