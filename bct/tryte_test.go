package bct

import (
	"testing"
)

var tryteValueTests = []struct {
	name   string
	value  int
	hi, lo Trybble
}{
	{"000010", 3, 0b_00_00_00, 0b_00_01_00},
	{"001TT0", 15, 0b_00_00_01, 0b_10_10_00},
	{"011T01", 100, 0b_00_01_01, 0b_10_00_01},
	{"1T10T1", 187, 0b_01_10_01, 0b_00_10_01},
	{"11T01T", 299, 0b_01_01_10, 0b_00_01_10},
	{"0000T0", -3, 0b_00_00_00, 0b_00_10_00},
	{"00T110", -15, 0b_00_00_10, 0b_01_01_00},
	{"0TT10T", -100, 0b_00_10_10, 0b_01_00_10},
	{"T1T01T", -187, 0b_10_01_10, 0b_00_01_10},
	{"TT10T1", -299, 0b_10_10_01, 0b_00_10_01},
}

func TestTryteFromInt(t *testing.T) {
	for _, tt := range tryteValueTests {
		t.Run(tt.name, func(t *testing.T) {
			tryte := TryteFromInt(tt.value)
			if tryte.Hi != tt.hi {
				t.Errorf("high trybble: got %06b, want %06b", tryte.Hi, tt.hi)
			}
			if tryte.Lo != tt.lo {
				t.Errorf("low trybble: got %06b, want %06b", tryte.Lo, tt.lo)
			}
		})
	}
}

func TestTryte_Value(t *testing.T) {
	for _, tt := range tryteValueTests {
		t.Run(tt.name, func(t *testing.T) {
			tryte := Tryte{Hi: tt.hi, Lo: tt.lo}
			got := tryte.Value()
			if got != tt.value {
				t.Fatalf("got %v, want %v", got, tt.value)
			}
		})
	}
}
