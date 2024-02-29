package ethioGrego

func startDayOfEthiopian(year int) int {
	//magic formula gives start of year
	newYearDay := (year / 100) - (year / 400) - 4

	// if the prev ethiopian year is a leap year, new-year occurs on 12th
	if (year-1)%4 == 3 {
		newYearDay += 1
	}
	return newYearDay
}

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
