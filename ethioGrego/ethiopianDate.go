/*
Ethiopian Calendar tool for Golang
Copyright (c) 2022 Yinebeb Tariku <yintar5@gmail.com>
*/

package ethioGrego

import (
	"fmt"
	"strconv"
	"time"
)

type response struct {
	msg string
}

// returns first day of that Ethiopian year
func start_day_of_ethiopian(year int) int {
	//magic formula gives start of year
	new_year_day := (year / 100) - (year / 400) - 4

	//if the prev ethiopian year is a leap year, new-year occrus on 12th
	if (year-1)%4 == 3 {
		new_year_day += 1
	}
	return new_year_day
}

// Gregorian date object representation of provided Ethiopian date
func To_gregorian(year, month, date int) time.Time {
	var gregorian_date int
	var dateResult string
	// prevent incorect input
	inputs := []int{year, month, date}
	for i := range inputs {
		if inputs[i] == 0 || len(inputs) != 3 {
			fmt.Print("Malformed input can't be converted.")
		}
	}

	// Ethiopian new year in Gregorian calendar
	new_year_day := start_day_of_ethiopian(year)

	// September (Ethiopian) sees 7y difference
	gregorian_year := year + 7

	// Number of days in gregorian months starting with September (index 1)
	// Index 0 is reserved for leap years switches.
	gregorian_months := []int{0, 30, 31, 30, 31, 31, 28,
		31, 30, 31, 30, 31, 31, 30}

	//if next gregorian year is leap year, February has 29 days.
	next_year := gregorian_year + 1
	if (next_year%4 == 0 && next_year%100 != 0) || next_year%400 == 0 {
		gregorian_months[6] = 29
	}
	// calculate number of days up to that date
	until := ((month - 1) * 30) + date
	if until <= 37 && year <= 1575 { //mysterious rule
		until += 28
		gregorian_months[0] = 31
	} else {
		until += new_year_day - 1
	}

	// if ethiopian year is leap year, paguemain has six days
	if year-1%4 == 3 {
		until += 1
	}

	//calculate month and date incremently
	m := 0
	for i := range gregorian_months {
		if until <= gregorian_months[i] {
			m = i
			gregorian_date = until
			break
		} else {
			m = i
			until -= gregorian_months[i]
		}
	}
	//  if m > 4, we're already on next Gregorian year
	if m > 4 {
		gregorian_year += 1
	}
	//  Gregorian months ordered according to Ethiopian
	order := []int{8, 9, 10, 11, 12, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	gregorian_month := order[m]

	da := strconv.Itoa(gregorian_date)
	mon := strconv.Itoa(gregorian_month)
	if len(da) == 1 {
		da = "0" + da
	}
	if len(mon) == 1 {
		mon = "0" + mon
	}

	dateResult = "" + strconv.Itoa(gregorian_year) + "-" + mon + "-" + da
	res, err := time.Parse("2006-01-02", dateResult)
	if err != nil {
		fmt.Print("unabe to parse dateResult.", err)
	}
	return res
}

// Ethiopian date string representation of provided Gregorian date
func To_ethiopian(year, month, date int) time.Time {
	var tahissas int
	var ethiopian_date int
	var dateResult string

	// prevent incorect input
	inputs := []int{year, month, date}
	for i := range inputs {
		if inputs[i] == 0 || len(inputs) != 3 {
			fmt.Printf("Malformed input can't be converted.")
		}
	}

	//  date between 5 and 14 of May 1582 are invalid
	if month == 10 && date >= 5 && date <= 14 && year == 1582 {
		fmt.Printf("Invalid Date between 5-14 May 1582.")
	}

	// Number of days in gregorian months starting with January (index 1)
	//
	//	Index 0 is reserved for leap years switches.
	gregorian_months := []int{0, 31, 28, 31, 30, 31, 30,
		31, 31, 30, 31, 30, 31}

	ethiopian_months := []int{0, 30, 30, 30, 30, 30, 30, 30,
		30, 30, 5, 30, 30, 30, 30}

	//  if gregorian leap year, February has 29 days.
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		gregorian_months[2] = 29
	}
	//  September sees 8y difference
	ethiopian_year := year - 8

	//  if ethiopian leap year pagumain has 6 days
	if ethiopian_year%4 == 3 {
		ethiopian_months[10] = 6
	} else {
		ethiopian_months[10] = 5
	}
	//  Ethiopian new year in Gregorian calendar
	new_year_day := start_day_of_ethiopian(year - 8)

	// calculate number of days up to that date
	until := 0
	for i := 1; i < month; i++ {
		until += gregorian_months[i]
	}
	until += date

	// # update tahissas (december) to match january 1st
	if ethiopian_year%4 == 0 {
		tahissas = 26
	} else {
		tahissas = 25
	}

	// take into account the 1582 change
	if year < 1582 {
		ethiopian_months[1] = 0
		ethiopian_months[2] = tahissas
	} else if until <= 277 && year == 1582 {
		ethiopian_months[1] = 0
		ethiopian_months[2] = tahissas
	} else {
		tahissas = new_year_day - 3
		ethiopian_months[1] = tahissas
	}
	// calculate month and date incremently
	m := 0
	for m = range ethiopian_months {
		if until <= ethiopian_months[m] {
			if m == 1 || ethiopian_months[m] == 0 {
				ethiopian_date = until + (30 - tahissas)
			} else {
				ethiopian_date = until
			}
			break
		} else {
			until -= ethiopian_months[m]
		}
	}
	//  if m > 4, we're already on next Ethiopian year
	if m > 10 {
		ethiopian_year += 1
	}
	// Ethiopian months ordered according to Gregorian
	order := []int{0, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 1, 2, 3, 4}
	ethiopian_month := order[m]

	da := strconv.Itoa(ethiopian_date)
	mon := strconv.Itoa(ethiopian_month)
	if len(da) == 1 {
		da = "0" + da
	}
	if len(mon) == 1 {
		mon = "0" + mon
	}

	dateResult = strconv.Itoa(ethiopian_year) + "-" + mon + "-" + da
	res, err := time.Parse("2006-01-02", dateResult)
	if err != nil {
		fmt.Print("unabe to parse dateResult.", err)
	}
	return res
}
