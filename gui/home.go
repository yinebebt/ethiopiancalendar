package gui

import (
	"fmt"
	"sort"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/yinebebt/ethiocal/bahirehasab"
	"github.com/yinebebt/ethiocal/dateconverter"
)

const welcomeLine = "እንኳን ደህና መጡ"

type namedFest struct {
	name string
	date bahirehasab.Date
}

func majorFestivals(f bahirehasab.Festival) []namedFest {
	return []namedFest{
		{"ነነዌ ጾም", f.Basic.Nenewie},
		{"አብይ ጾም", f.Fasting.Abiy},
		{"ደብረ ዘይት", f.Fasting.DebreZeit},
		{"ሆሳህና", f.Fasting.Hosanna},
		{"ስቅለት", f.Fasting.Siklet},
		{"ፋሲካ", f.Fasting.Tinsaye},
		{"ርክበ ካህናት", f.Fasting.RkbeKahnat},
		{"እርገት", f.Fasting.Erget},
		{"ጰራቅሊጦስ", f.Fasting.Peraklitos},
		{"ጾመ ሐዋሪያት", f.Fasting.Hawariyat},
		{"ጾመ ነቢያት", f.Fasting.Nebiyat},
		{"ጾመ ፍልሰታ", f.Fasting.Filseta},
		{"ጾመ ገሀድ", f.Fasting.Gehad},
	}
}

func ethOrdinal(d bahirehasab.Date) int {
	return (int(d.Month)-1)*30 + d.Day
}

// nearbyFestivals returns up to limit upcoming major feasts within the next
// 45 Ethiopian tabular days.
func nearbyFestivals(f bahirehasab.Festival, month, day, limit int) []namedFest {
	today := (month-1)*30 + day
	var upcoming []namedFest
	for _, e := range majorFestivals(f) {
		if !e.date.IsValid() {
			continue
		}
		o := ethOrdinal(e.date)
		if o >= today && o <= today+45 {
			upcoming = append(upcoming, e)
		}
	}
	sort.Slice(upcoming, func(i, j int) bool {
		return ethOrdinal(upcoming[i].date) < ethOrdinal(upcoming[j].date)
	})
	if len(upcoming) > limit {
		upcoming = upcoming[:limit]
	}
	return upcoming
}

func fmtFestLine(e namedFest, year int) string {
	long, err := dateconverter.FormatEthiopian(year, int(e.date.Month), e.date.Day)
	if err != nil {
		return fmt.Sprintf("%s — %s", e.name, fmtDateNamed(e.date))
	}
	return fmt.Sprintf("%s — %s", e.name, long)
}

func newHomeTab() fyne.CanvasObject {
	now := time.Now()
	gregLong := dateconverter.FormatGregorian(now)

	etY, etM, etD := 1, 1, 1
	etLong := ""
	if et, err := dateconverter.Ethiopian(now.Year(), int(now.Month()), now.Day()); err == nil {
		etY, etM, etD = et.Year(), int(et.Month()), et.Day()
		etLong, _ = dateconverter.FormatEthiopian(et.Year(), int(et.Month()), et.Day())
	}

	brand := accentHeading("Ethiocal")
	welcome := widget.NewLabel(welcomeLine)
	welcome.Wrapping = fyne.TextWrapWord
	welcome.Alignment = fyne.TextAlignCenter

	todayHead := widget.NewLabel("ዛሬ")
	todayHead.TextStyle = fyne.TextStyle{Bold: true}
	todayHead.Alignment = fyne.TextAlignCenter

	gregLbl := widget.NewLabel(gregLong)
	gregLbl.Wrapping = fyne.TextWrapWord
	gregLbl.Alignment = fyne.TextAlignCenter
	etLbl := widget.NewLabel(etLong)
	etLbl.Wrapping = fyne.TextWrapWord
	etLbl.Alignment = fyne.TextAlignCenter
	if etLong == "" {
		etLbl.SetText("—")
	}

	nearHead := widget.NewLabel("በአቅራቢያ")
	nearHead.TextStyle = fyne.TextStyle{Bold: true}
	nearHead.Alignment = fyne.TextAlignCenter

	nearBox := container.NewVBox()
	fest, err := bahirehasab.NewFestival(etY)
	if err != nil {
		msg := widget.NewLabel("የበዓላት መረጃ አልተገኘም።")
		msg.Alignment = fyne.TextAlignCenter
		nearBox.Add(msg)
	} else {
		events := nearbyFestivals(fest, etM, etD, 4)
		if len(events) == 0 {
			msg := widget.NewLabel("በሚቀጥሉት ቀናት ዋና በዓል የለም።")
			msg.Alignment = fyne.TextAlignCenter
			nearBox.Add(msg)
		} else {
			for _, e := range events {
				row := widget.NewLabel(fmtFestLine(e, etY))
				row.Wrapping = fyne.TextWrapWord
				row.Alignment = fyne.TextAlignCenter
				nearBox.Add(row)
			}
		}
	}

	body := container.NewVBox(
		container.NewCenter(brand),
		container.NewCenter(welcome),
		widget.NewSeparator(),
		todayHead,
		gregLbl,
		etLbl,
		widget.NewSeparator(),
		nearHead,
		nearBox,
	)

	return centered(container.NewScroll(body))
}
