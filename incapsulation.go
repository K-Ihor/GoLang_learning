// package main

// import (
// 	//"errors"
// 	"fmt"
// 	"golangtest/calendar"
// 	"log"
// )
// // -------------------------------------------------------------------------------------------
// // type Date struct {
// // 	Year  int // чтоб эти поля были закрытими нужно определить Date  в отдельном пакете
// // 	Month int
// // 	Day   int
// // }

// // // определяем сеттер для нашего типа данных
// // func (d *Date) SetYear(year int) error {
// // 	if year < 1 {
// // 		return errors.New("invalid year")
// // 	}
// // 	d.Year = year
// // 	return nil
// // }

// // func (d *Date) SetMonth(month int) error {
// // 	if month < 1 || month > 12 {
// // 		return errors.New("invalid month")
// // 	}
// // 	d.Month = month
// // 	return nil
// // }

// // func (d *Date) SetDay(day int) {
// // 	d.Day = day
// // }
// // -----------------------------------------------------------------------------------------------

// func main() {
// 	date := calendar.Date{}
// 	err := date.SetYear(3100)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	date.SetYear(3100)

// 	fmt.Println(date.Day())
// 	fmt.Println(date)

// }
