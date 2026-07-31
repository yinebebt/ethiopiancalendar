package main

import "testing"

func TestParseDate(t *testing.T) {
	tests := []struct {
		in                  string
		wantY, wantM, wantD int
		wantErr             bool
	}{
		{"2025-2-2", 2025, 2, 2, false},
		{"2025-02-02", 2025, 2, 2, false},
		{"2017-5-25", 2017, 5, 25, false},
		{"2025/2/2", 0, 0, 0, true},
		{"2025-2", 0, 0, 0, true},
		{"2025-a-2", 0, 0, 0, true},
		{"", 0, 0, 0, true},
	}
	for _, tt := range tests {
		y, m, d, err := parseDate(tt.in)
		if tt.wantErr {
			if err == nil {
				t.Errorf("parseDate(%q): want error", tt.in)
			}
			continue
		}
		if err != nil {
			t.Errorf("parseDate(%q): %v", tt.in, err)
			continue
		}
		if y != tt.wantY || m != tt.wantM || d != tt.wantD {
			t.Errorf("parseDate(%q) = %d-%d-%d, want %d-%d-%d", tt.in, y, m, d, tt.wantY, tt.wantM, tt.wantD)
		}
	}
}
