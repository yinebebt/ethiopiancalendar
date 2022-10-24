package ethioGrego

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

// prevent incorect input
func isValid(year, month, date int) bool {
	inputs := []int{year, month, date}
	for i := range inputs {
		if inputs[i] <= 0 {
			return false
		}
	}
	if len(inputs) != 3 || inputs[2] > 31 || inputs[1] > 12 {
		return false
	}
	return true
}
