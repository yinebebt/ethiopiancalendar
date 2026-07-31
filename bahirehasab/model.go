package bahirehasab

import "fmt"

// Traditional ትንተ constants used in መጥቅ / አበቅቴ arithmetic.
const (
	TinteMetiq   = 19
	TinteAbektie = 11
)

// Day-of-week ተውሳክ (algorithm-internal).
const (
	tewsakSaturday  = 8
	tewsakSunday    = 7
	tewsakMonday    = 6
	tewsakTuesday   = 5
	tewsakWednesday = 4
	tewsakThursday  = 3
	tewsakFriday    = 2
)

// weekdayTewsak maps ትንተ ቀመር weekday index (0=Mon … 6=Sun) to ተውሳክ.
var weekdayTewsak = [7]int{
	tewsakMonday, tewsakTuesday, tewsakWednesday, tewsakThursday,
	tewsakFriday, tewsakSaturday, tewsakSunday,
}

// Month is a 1-based Ethiopian month (መስከረም=1 … ጳጉሜ=13).
type Month int

const (
	Meskerem Month = iota + 1
	Tikimt
	Hidar
	Tahsas
	Tir
	Yekatit
	Megabit
	Miazia
	Ginbot
	Sene
	Hamle
	Nehase
	Pagumen
)

func (m Month) String() string {
	names := [...]string{
		"", "መስከረም", "ጥቅምት", "ኅዳር", "ታኅሣሥ", "ጥር", "የካቲት",
		"መጋቢት", "ሚያዝያ", "ግንቦት", "ሰኔ", "ሐምሌ", "ነሐሴ", "ጳጉሜ",
	}
	if m < Meskerem || m > Pagumen {
		return fmt.Sprintf("Month(%d)", int(m))
	}
	return names[m]
}

// Evangelist is አመተ ዓለም % 4 (0=John … 3=Luke).
type Evangelist int

const (
	EvangelistJohn Evangelist = iota
	EvangelistMatthew
	EvangelistMark
	EvangelistLuke
)

func (e Evangelist) String() string {
	switch e {
	case EvangelistJohn:
		return "ዮሐንስ(John)"
	case EvangelistMatthew:
		return "ማቲዎስ(Matthew)"
	case EvangelistMark:
		return "ማርቆስ(Mark)"
	case EvangelistLuke:
		return "ሉቃስ(Luke)"
	default:
		return fmt.Sprintf("Evangelist(%d)", int(e))
	}
}

// Weekday uses this package's ትንተ ቀመር numbering: 0=Monday … 6=Sunday.
type Weekday int

const (
	Monday Weekday = iota
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func (w Weekday) String() string {
	names := [...]string{
		"ሰኞ(Monday)", "ማክሰኞ(Tuesday)", "እሮብ(Wednesday)", "ሃሙስ(Thursday)",
		"አርብ(Friday)", "ቅዳሜ(Saturday)", "እሁድ(Sunday)",
	}
	if w < Monday || w > Sunday {
		return fmt.Sprintf("Weekday(%d)", int(w))
	}
	return names[w]
}

// BealeMetiq is the month in which በዓለ-መጥቅ falls.
type BealeMetiq int

const (
	BealeMetiqMeskerem BealeMetiq = 1
	BealeMetiqTikimt   BealeMetiq = 2
)

// Festival is the full ባሕረ ሐሳብ result for one Ethiopian year.
type Festival struct {
	Year    Year    `json:"year"`
	Basic   Basic   `json:"basic"`
	Fasting Fasting `json:"fasting"`
}

// Year holds New Year metadata (evangelist and weekday).
type Year struct {
	Year           int        `json:"year"`
	Evangelist     Evangelist `json:"evangelist"`
	NewYearWeekday Weekday    `json:"new_year_weekday"`
}

// Basic holds መደብ through ነነዌ, the seed for all later feasts.
type Basic struct {
	Medeb       int        `json:"medeb"`
	Wenber      int        `json:"wenber"`
	Abektie     int        `json:"abektie"`
	Metiq       int        `json:"metiq"`
	BealeMetiq  BealeMetiq `json:"beale_metiq"`
	MebajaHamer int        `json:"mebaja_hamer"`
	Nenewie     Date       `json:"nenewie"`
}

// Date is a day within an Ethiopian calendar year (year lives on Festival).
type Date struct {
	Day   int   `json:"day"`
	Month Month `json:"month"`
}

// IsValid reports whether d is an Ethiopian calendar date.
func (d Date) IsValid() bool {
	if d.Month < Meskerem || d.Month > Pagumen {
		return false
	}
	m := 30
	if d.Month == Pagumen {
		m = 6
	}
	return d.Day >= 1 && d.Day <= m
}

// String returns a zero-padded MM-DD representation.
func (d Date) String() string {
	return fmt.Sprintf("%02d-%02d", int(d.Month), d.Day)
}

// AddDays advances d by n days using 30-day months.
func (d Date) AddDays(n int) Date {
	total := (int(d.Month)-1)*30 + (d.Day - 1) + n
	monthIndex := total / 30
	day := total%30 + 1
	return Date{Day: day, Month: Month(monthIndex + 1)}
}

// Fasting holds አጽዋማትና በዓላት.
type Fasting struct {
	Abiy       Date `json:"abiy"`
	DebreZeit  Date `json:"debre_zeit"`
	Hosanna    Date `json:"hosanna"`
	Siklet     Date `json:"siklet"`
	Tinsaye    Date `json:"tinsaye"`
	RkbeKahnat Date `json:"rkbe_kahnat"`
	Dihnet     Date `json:"dihnet"`
	Hawariyat  Date `json:"hawariyat"`
	Erget      Date `json:"erget"`
	Peraklitos Date `json:"peraklitos"`

	// Fixed-entry fasts.
	Nebiyat Date `json:"nebiyat"`
	Filseta Date `json:"filseta"`
	Gehad   Date `json:"gehad"`
}
