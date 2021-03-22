package dates

import "errors"

type Date struct {
	year int
	month int
	day int 
}

func (d *Date) SetYear(year int) error {
	//d.Year = year // automagically gets the value at the pointer
	//(*d).Year = year // same as above line, but showing for clarity of implicitness above

	if year < 1 {
		return errors.New("invalid date year value- year can't be less than 0")
	}

	d.year = year
	return nil 
}

func (d *Date) SetMonth(month int) error {
	if month < 1 || month > 12 {
		return errors.New("invalid date month value- month has to be between 1 and 12")
	}

	d.month = month
	return nil
}

func (d *Date) SetDay(day int) error {
	if day < 1 || day > 31 {
		return errors.New("invalid date day value- day has to be between 1 and 31")
	}

	d.day = day
	return nil 
}

func (d *Date) Year() int {
	return d.year
}

func (d *Date) Month() int {
	return d.month
}

func (d *Date) Day() int {
	return d.day
}