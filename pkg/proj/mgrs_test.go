package proj

import (
	"fmt"
	"testing"
)

func TestMGRS_ToUTM(t *testing.T) {

	var tests = []struct {
		mgrs     MGRS  // in
		utm      UTM   // out
		accuracy int   // out
		err      error // out
	}{
		// positive tests
		{"32ULC9897356497", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973, Northing: 5756497}, 1, nil},
		{"32ULC98975649", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398970, Northing: 5756490}, 10, nil},
		{"32ULC989564", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398900, Northing: 5756400}, 100, nil},
		{"32ULC9856", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398000, Northing: 5756000}, 1000, nil},
		{"32ULC95", UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 390000, Northing: 5750000}, 10000, nil},
		{"23KPU1173300614", UTM{ZoneNumber: 23, ZoneLetter: 'K', Easting: 611733, Northing: 7800614}, 1, nil},
		{"18TWL9334507672", UTM{ZoneNumber: 18, ZoneLetter: 'T', Easting: 593345, Northing: 4507672}, 1, nil},
		{"10SGJ0683244683", UTM{ZoneNumber: 10, ZoneLetter: 'S', Easting: 706832, Northing: 4344683}, 1, nil},
		{"31UGT0037304554", UTM{ZoneNumber: 31, ZoneLetter: 'U', Easting: 700373, Northing: 5704554}, 1, nil},
		{"30NYF6799300000", UTM{ZoneNumber: 30, ZoneLetter: 'N', Easting: 767993, Northing: 0}, 1, nil},
		// negative tests
		{"", UTM{}, 0, fmt.Errorf("invalid empty mgrs string")},
	}

	for _, test := range tests {
		utm, accuracy, err := test.mgrs.ToUTM()
		function := fmt.Sprintf("mgrs = %s, ToUTM()", test.mgrs)
		got := fmt.Sprintf("%s %d %v", utm, accuracy, err)
		want := fmt.Sprintf("%s %d %v", test.utm, test.accuracy, test.err)
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}

func TestMGRS_ToLL(t *testing.T) {

	var tests = []struct {
		mgrs     MGRS  // in
		ll       LL    // out
		accuracy int   // out
		err      error // out
	}{
		// positive tests
		{"32ULC9897356497", LL{Lat: 51.949993, Lon: 7.529986}, 1, nil},
		{"33UXP04", LL{Lat: 48.205348, Lon: 16.345927}, 10000, nil},
		{"11SPA7234911844", LL{Lat: 36.236123, Lon: -115.082098}, 1, nil},
		{"23KPU1173300614", LL{Lat: -19.887498, Lon: -43.932664}, 1, nil},
		{"31UGT03734554", LL{Lat: 51.823490, Lon: 5.956335}, 10, nil},
		{"30NYF6799300000", LL{Lat: 0.0, Lon: -0.592328}, 1, nil},
		// negative tests
		{"32ULC9897356497CORRUPT", LL{}, 0, fmt.Errorf("error <uneven number of digits, mgrs = 32ULC9897356497CORRUPT> at mgrs.ToUTM()")},
	}

	for _, test := range tests {
		ll, accuracy, err := test.mgrs.ToLL()
		function := fmt.Sprintf("mgrs = %s, mgrs.ToLL()", test.mgrs)
		got := fmt.Sprintf("%s %d %v", ll, accuracy, err)
		want := fmt.Sprintf("%s %d %v", test.ll, test.accuracy, test.err)
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}

func TestMGRS_ToUSNGS(t *testing.T) {
	var tests = []struct {
		mgrs MGRS
		usng USNG
	}{
		{"32ULC9897356497", "32U LC 98973 56497"},
		{"32ULC98975649", "32U LC 9897 5649"},
		{"32ULC989564", "32U LC 989 564"},
		{"32ULC9856", "32U LC 98 56"},
		{"32ULC95", "32U LC 9 5"},
	}

	for _, test := range tests {
		got := test.mgrs.ToUSNG()
		want := test.usng

		if got != want {
			t.Errorf("mgrs.ToUSNG() = %s, want %s", got, want)
		}
	}
}
