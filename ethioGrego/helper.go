package ethioGrego

// returns first day of that Ethiopian year
func startDayOfEthiopian(year int) int {
	//magic formula gives start of year
	newYearDay := (year / 100) - (year / 400) - 4

	//if the prev ethiopian year is a leap year, new-year occurs on 12th
	if (year-1)%4 == 3 {
		newYearDay += 1
	}
	return newYearDay
}

// prevent incorrect input
func isValid(year, month, date int) bool {
	inputs := []int{year, month, date}
	for i := range inputs {
		if inputs[i] <= 0 {
			return false
		}
	}
	return !(len(inputs) != 3 || inputs[2] > 31 || inputs[1] > 12)
}
