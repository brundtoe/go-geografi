package proj

import (
	"fmt"
	"testing"
)

func TestUTM_ToLL(t *testing.T) {

	var tests = []struct {
		utm UTM   // in
		ll  LL    // out
		err error // out
	}{
		// positive tests
		{UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 399000, Northing: 5757000}, LL{Lat: 51.954519, Lon: 7.530231}, nil},
		{UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 574126, Northing: 5815291}, LL{Lat: 52.482729, Lon: 10.091555}, nil},
		// test set from chris veness
		{UTM{ZoneNumber: 31, ZoneLetter: 'N', Easting: 166021, Northing: 0}, LL{Lat: 0.0, Lon: -0.000004}, nil},
		{UTM{ZoneNumber: 31, ZoneLetter: 'N', Easting: 277438, Northing: 110597}, LL{Lat: 0.999991, Lon: 0.999998}, nil},
		{UTM{ZoneNumber: 30, ZoneLetter: 'M', Easting: 722561, Northing: 9889402}, LL{Lat: -1.0, Lon: -1.000007}, nil},
		{UTM{ZoneNumber: 31, ZoneLetter: 'N', Easting: 448251, Northing: 5411943}, LL{Lat: 48.858293, Lon: 2.294488}, nil},    // eiffel tower
		{UTM{ZoneNumber: 56, ZoneLetter: 'H', Easting: 334873, Northing: 6252266}, LL{Lat: -33.857001, Lon: 151.214998}, nil}, // sidney o/h
		{UTM{ZoneNumber: 18, ZoneLetter: 'N', Easting: 323394, Northing: 4307395}, LL{Lat: 38.897694, Lon: -77.036503}, nil},  // white house
		{UTM{ZoneNumber: 23, ZoneLetter: 'K', Easting: 683466, Northing: 7460687}, LL{Lat: -22.951904, Lon: -43.210602}, nil}, // rio christ
		{UTM{ZoneNumber: 32, ZoneLetter: 'N', Easting: 297508, Northing: 6700645}, LL{Lat: 60.391347, Lon: 5.324893}, nil},    // bergen
		// negative tests
		{UTM{ZoneNumber: 132, ZoneLetter: 'U', Easting: 574126, Northing: 5815291}, LL{}, fmt.Errorf("invalid zone number, zone number = 132")},
	}

	for _, test := range tests {
		ll, err := test.utm.ToLL()
		function := fmt.Sprintf("utm = %s, ToLL()", test.utm)
		got := fmt.Sprintf("%s %v", ll, err)
		want := fmt.Sprintf("%s %v", test.ll, test.err)
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}

func TestUTM_ToMGRS(t *testing.T) {

	var tests = []struct {
		utm      UTM    // in
		accuracy int    // in
		mgrs     string // out
	}{
		// positive tests
		{UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973, Northing: 5756497}, 1, "32ULC9897356497"},
		{UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973, Northing: 5756497}, 10, "32ULC98975649"},
		{UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973, Northing: 5756497}, 100, "32ULC989564"},
		{UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973, Northing: 5756497}, 1000, "32ULC9856"},
		{UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973, Northing: 5756497}, 10000, "32ULC95"},
		{UTM{ZoneNumber: 23, ZoneLetter: 'K', Easting: 611733, Northing: 7800614}, 1, "23KPU1173300614"},
		// negative tests
		// nothing to do here
	}

	for _, test := range tests {
		mgrs := test.utm.ToMGRS(test.accuracy)
		function := fmt.Sprintf("utm = %s, ToMGRS(%d)", test.utm, test.accuracy)
		got := string(mgrs)
		want := test.mgrs
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}

func TestUTM_ToUSNG(t *testing.T) {
	var expected USNG = "32V NJ 94858 99060"
	utm := UTM{ZoneNumber: 32, ZoneLetter: 'V', Easting: 594858, Northing: 6399060}
	actual := utm.ToUSNG(1)
	if actual != expected {
		t.Errorf("Fejlede konvertering af UTM %v to USNG %v", utm, expected)
	}
}
