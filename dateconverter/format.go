package dateconverter

import (
	"fmt"
	"time"
)

// Amharic month names (1-based index).
var ethMonthNames = [13]string{
	"መስከረም", "ጥቅምት", "ኅዳር", "ታኅሣሥ", "ጥር", "የካቲት",
	"መጋቢት", "ሚያዝያ", "ግንቦት", "ሰኔ", "ሐምሌ", "ነሐሴ", "ጳጉሜ",
}

// Amharic weekday names indexed by time.Weekday (Sunday = 0).
var ethWeekdayNames = [7]string{
	"እሁድ", "ሰኞ", "ማክሰኞ", "ረቡዕ", "ሐሙስ", "አርብ", "ቅዳሜ",
}

// FormatGregorian returns a long English Gregorian date with weekday.
// e.g. "Monday, July 2, 2012".
func FormatGregorian(t time.Time) string {
	return t.Format("Monday, January 2, 2006")
}

// FormatEthiopian returns an Amharic long form with weekday.
// e.g. "መስከረም 5፣ 2017 (እሁድ)".
func FormatEthiopian(year, month, day int) (string, error) {
	if !isValidEthiopian(year, month, day) {
		return "", fmt.Errorf("not a valid Ethiopian date")
	}
	g, err := Gregorian(year, month, day)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s %d፣ %d (%s)", ethMonthNames[month-1], day, year, ethWeekdayNames[g.Weekday()]), nil
}
