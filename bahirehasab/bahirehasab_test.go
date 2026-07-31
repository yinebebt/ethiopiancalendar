package bahirehasab

import "testing"

func TestNewFestival_2016(t *testing.T) {
	f, err := NewFestival(2016)
	if err != nil {
		t.Fatalf("NewFestival: %v", err)
	}

	if f.Year.Evangelist != EvangelistJohn {
		t.Errorf("evangelist: got %v want %v", f.Year.Evangelist, EvangelistJohn)
	}
	if f.Year.NewYearWeekday != Tuesday {
		t.Errorf("new year weekday: got %v want %v", f.Year.NewYearWeekday, Tuesday)
	}

	wantBasic := Basic{
		Medeb:       11,
		Wenber:      10,
		Abektie:     20,
		Metiq:       10,
		BealeMetiq:  BealeMetiqTikimt,
		MebajaHamer: 18,
		Nenewie:     Date{Day: 18, Month: Yekatit},
	}
	if f.Basic != wantBasic {
		t.Errorf("basic:\n got %+v\nwant %+v", f.Basic, wantBasic)
	}

	wantFasting := Fasting{
		Abiy:       Date{Day: 2, Month: Megabit},
		DebreZeit:  Date{Day: 29, Month: Megabit},
		Hosanna:    Date{Day: 20, Month: Miazia},
		Siklet:     Date{Day: 25, Month: Miazia},
		Tinsaye:    Date{Day: 27, Month: Miazia},
		RkbeKahnat: Date{Day: 21, Month: Ginbot},
		Erget:      Date{Day: 6, Month: Sene},
		Peraklitos: Date{Day: 16, Month: Sene},
		Hawariyat:  Date{Day: 17, Month: Sene},
		Dihnet:     Date{Day: 19, Month: Sene},
		Nebiyat:    Date{Day: 15, Month: Hidar},
		Filseta:    Date{Day: 1, Month: Nehase},
		Gehad:      Date{Day: 10, Month: Tir},
	}
	if f.Fasting != wantFasting {
		t.Errorf("fasting:\n got %+v\nwant %+v", f.Fasting, wantFasting)
	}
}

func TestDate_AddDays(t *testing.T) {
	tests := []struct {
		name string
		in   Date
		add  int
		want Date
	}{
		{"abiy_from_nenewie", Date{Day: 18, Month: Yekatit}, 14, Date{Day: 2, Month: Megabit}},
		{"exact_month_end", Date{Day: 1, Month: Tir}, 29, Date{Day: 30, Month: Tir}},
		{"cross_month", Date{Day: 30, Month: Tir}, 1, Date{Day: 1, Month: Yekatit}},
		{"never_day_zero", Date{Day: 18, Month: Yekatit}, 12, Date{Day: 30, Month: Yekatit}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.in.AddDays(tt.add)
			if got != tt.want {
				t.Errorf("AddDays(%d): got %v want %v", tt.add, got, tt.want)
			}
			if !got.IsValid() {
				t.Errorf("result not valid: %v", got)
			}
		})
	}
}

func TestNewFestival_invalid(t *testing.T) {
	if _, err := NewFestival(-1); err == nil {
		t.Fatal("expected error for negative year")
	}
}

func TestNewFestival_2006_metiq30(t *testing.T) {
	// wenber=0 → (0*19)%30=0 → normalized to መጥቅ 30, በዓለ-መጥቅ መስከረም ፴.
	f, err := NewFestival(2006)
	if err != nil {
		t.Fatalf("NewFestival: %v", err)
	}
	if f.Basic.Metiq != 30 {
		t.Errorf("metiq: got %d want 30", f.Basic.Metiq)
	}
	if f.Basic.BealeMetiq != BealeMetiqMeskerem {
		t.Errorf("beale metiq: got %v want %v", f.Basic.BealeMetiq, BealeMetiqMeskerem)
	}
	if f.Basic.Nenewie != (Date{Day: 3, Month: Yekatit}) {
		t.Errorf("nenewie: got %v want 06-03", f.Basic.Nenewie)
	}
}

func TestDate_String(t *testing.T) {
	d := Date{Day: 18, Month: Yekatit}
	if got := d.String(); got != "06-18" {
		t.Errorf("String: got %q want %q", got, "06-18")
	}
}
