package proj

import (
	"fmt"
	"testing"
)

func TestUSNG_ToMGRS(t *testing.T) {
	expected := "32VNJ9485899060"
	var usng USNG = "32V NJ 94858 99060"
	mgrs := usng.ToMGRS()
	if string(mgrs) != expected {
		t.Errorf("konvertering af USNG %s to MGSR %s fejlede", string(usng), string(mgrs))
	}
}

func TestUSNG_ToUTM(t *testing.T) {

	var tests = []struct {
		usng     USNG  // in
		utm      UTM   // out
		accuracy int   // out
		err      error // out
	}{
		// positive tests
		{"32U LC 98973 56497", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973, Northing: 5756497}, 1, nil},
		{"32U LC 9897 5649", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398970, Northing: 5756490}, 10, nil},
		{"32U LC 989 564", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398900, Northing: 5756400}, 100, nil},
		{"32U LC 98 56", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398000, Northing: 5756000}, 1000, nil},
		{"32U LC 9 5", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 390000, Northing: 5750000}, 10000, nil},
		{"23K PU 11733 00614", UTM{ZoneNumber: 23, ZoneLetter: 'K', Easting: 611733, Northing: 7800614}, 1, nil},
		{"18T WL 93345 07672", UTM{ZoneNumber: 18, ZoneLetter: 'T', Easting: 593345, Northing: 4507672}, 1, nil},
		{"10S GJ 06832 44683", UTM{ZoneNumber: 10, ZoneLetter: 'S', Easting: 706832, Northing: 4344683}, 1, nil},
		{"31U GT 00373 04554", UTM{ZoneNumber: 31, ZoneLetter: 'U', Easting: 700373, Northing: 5704554}, 1, nil},
		{"30NYF6799300000", UTM{ZoneNumber: 30, ZoneLetter: 'N', Easting: 767993, Northing: 0}, 1, nil},
		// negative tests - der skal refereres til mgrs i error message da det er mgrs string der fejler i mgrs.ToUTM som kaldes af usng.ToUTM
		{"", UTM{}, 0, fmt.Errorf("invalid empty mgrs string")},
	}

	for _, test := range tests {
		utm, accuracy, err := test.usng.ToUTM()
		function := fmt.Sprintf("usng = %s, ToUTM()", test.usng)
		got := fmt.Sprintf("%s %d %v", utm, accuracy, err)
		want := fmt.Sprintf("%s %d %v", test.utm, test.accuracy, test.err)
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}

func TestUSNG_ToLL(t *testing.T) {

	var tests = []struct {
		usng     USNG  // in
		ll       LL    // out
		accuracy int   // out
		err      error // out
	}{
		// positive tests
		{"32U LC 98973 56497", LL{Lat: 51.949993, Lon: 7.529986}, 1, nil},
		{"33U XP 0 4", LL{Lat: 48.205348, Lon: 16.345927}, 10000, nil},
		{"11S PA 72349 11844", LL{Lat: 36.236123, Lon: -115.082098}, 1, nil},
		{"23K PU 11733 00614", LL{Lat: -19.887498, Lon: -43.932664}, 1, nil},
		{"31U GT 0373 4554", LL{Lat: 51.823490, Lon: 5.956335}, 10, nil},
		{"30N YF 67993 00000", LL{Lat: 0.0, Lon: -0.592328}, 1, nil},
		// negative tests
		// der skal refereres til mgrs i error message da det er mgrs string der fejler i mgrs.ToUTM som kaldes af usng.ToUTM
		{"32U LC 98973 56497CORRUPT", LL{}, 0, fmt.Errorf("error <uneven number of digits, mgrs = 32ULC9897356497CORRUPT> at mgrs.ToUTM()")},
	}

	for _, test := range tests {
		ll, accuracy, err := test.usng.ToLL()
		function := fmt.Sprintf("mgrs = %s, mgrs.ToLL()", test.usng)
		got := fmt.Sprintf("%s %d %v", ll, accuracy, err)
		want := fmt.Sprintf("%s %d %v", test.ll, test.accuracy, test.err)
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}
