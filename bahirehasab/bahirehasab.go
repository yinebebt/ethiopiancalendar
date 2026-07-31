package bahirehasab

import "fmt"

// የበዓላትና የአጽዋማት ተውሳክ
// https://eotcmk.org/a/የ፳፻፲፩-ዓ-ም-የአጽዋማት-ባሕረ-ሐሳባዊ/
const (
	offsetAbiy       = 14
	offsetDebreZeit  = 41
	offsetHosanna    = 62
	offsetSiklet     = 67
	offsetTinsaye    = 69
	offsetRkbeKahnat = 93
	offsetErget      = 108
	offsetPeraklitos = 118
	offsetHawariyat  = 119
	offsetDihnet     = 121
)

// NewFestival computes ባሕረ ሐሳብ for the given Ethiopian year.
func NewFestival(year int) (Festival, error) {
	if year < 0 {
		return Festival{}, fmt.Errorf("invalid Ethiopian year %d", year)
	}
	basic := getBasic(year)
	return Festival{
		Year:    getYear(year),
		Basic:   basic,
		Fasting: getFasting(basic),
	}, nil
}

// BahireHasab computes the Ethiopian religious calendar (fasting and festival
// dates) for the given Ethiopian year.
//
// Deprecated: kept for compatibility. Use NewFestival instead.
func BahireHasab(etYear int) (Festival, error) {
	return NewFestival(etYear)
}

func getYear(year int) Year {
	ameteAlem := year + 5500
	wng := Evangelist(ameteAlem % 4)
	meteneRabit := (ameteAlem - int(wng)) / 4
	tinteQemer := Weekday((ameteAlem + meteneRabit) % 7)
	return Year{
		Year:           year,
		Evangelist:     wng,
		NewYearWeekday: tinteQemer,
	}
}

func getBasic(year int) Basic {
	amtAlem := year + 5500
	medeb := amtAlem % 19
	wenber := medeb - 1
	if medeb == 0 {
		wenber = 18
	}

	abektie := (wenber * TinteAbektie) % 30
	metiq := (wenber * TinteMetiq) % 30
	// (wenber*19)%30 never yields 14; a 0 remainder is day 30 (መስከረም ፴).
	if metiq == 0 {
		metiq = 30
	}
	wngNum := amtAlem % 4
	meteneRabit := (amtAlem - wngNum) / 4
	tinteQemer := (amtAlem + meteneRabit) % 7

	// መጥቅ > 14 → መስከረም; መጥቅ < 14 → ጥቅምት. Equality with 14 is unreachable.
	var beale BealeMetiq
	var dayOfWeekForMtq int
	if metiq > 14 {
		beale = BealeMetiqMeskerem
		dayOfWeekForMtq = ((metiq-1)%7 + tinteQemer) % 7
	} else {
		beale = BealeMetiqTikimt
		dayOfWeekForMtq = ((metiq+29)%7 + tinteQemer) % 7
	}

	twsakOfDay := weekdayTewsak[dayOfWeekForMtq]
	mebajaHamer := metiq + twsakOfDay

	var nenewie Date
	if beale == BealeMetiqMeskerem && mebajaHamer <= 30 {
		nenewie = Date{Day: mebajaHamer, Month: Tir}
	} else {
		day := mebajaHamer % 30
		if day == 0 {
			day = 30
		}
		nenewie = Date{Day: day, Month: Yekatit}
	}

	return Basic{
		Medeb:       medeb,
		Wenber:      wenber,
		Abektie:     abektie,
		Metiq:       metiq,
		BealeMetiq:  beale,
		MebajaHamer: mebajaHamer,
		Nenewie:     nenewie,
	}
}

func getFasting(b Basic) Fasting {
	n := b.Nenewie
	return Fasting{
		Abiy:       n.AddDays(offsetAbiy),
		DebreZeit:  n.AddDays(offsetDebreZeit),
		Hosanna:    n.AddDays(offsetHosanna),
		Siklet:     n.AddDays(offsetSiklet),
		Tinsaye:    n.AddDays(offsetTinsaye),
		RkbeKahnat: n.AddDays(offsetRkbeKahnat),
		Erget:      n.AddDays(offsetErget),
		Peraklitos: n.AddDays(offsetPeraklitos),
		Hawariyat:  n.AddDays(offsetHawariyat),
		Dihnet:     n.AddDays(offsetDihnet),
		Nebiyat:    Date{Day: 15, Month: Hidar},
		Filseta:    Date{Day: 1, Month: Nehase},
		Gehad:      gehad(timket),
	}
}

// timket is ጥምቀት (Tir 11); its eve is ጾመ ገሀድ.
var timket = Date{Day: 11, Month: Tir}

// gehad is the day before a timket.
func gehad(feast Date) Date {
	return feast.AddDays(-1)
}
