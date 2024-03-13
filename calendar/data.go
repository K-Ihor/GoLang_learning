package calendar

import (
	"errors"
)

type Date struct {
	year  int // чтоб эти поля были закрытими нужно определить Date  в отдельном пакете
	month int
	day   int
}

// определяем сеттер для нашего типа данных
func (d *Date) SetYear(year int) error { // СЕТТЕР
	if year < 1 {
		return errors.New("invalid year")
	}
	d.year = year
	return nil
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid month")
	}
	d.month = month
	return nil
}

func (d *Date) SetDay(day int) {
	d.day = day
}

func (d *Date) Day() int { // ГЕТЕР
	return d.day
}
