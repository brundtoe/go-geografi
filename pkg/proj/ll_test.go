package proj

import (
	"fmt"
	"testing"
)

func TestLL_ToUTM(t *testing.T) {

	var tests = []struct {
		ll  LL  // in
		utm UTM // out
	}{
		// positive tests
		{LL{Lat: 51.95, Lon: 7.53}, UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 398973.96, Northing: 5756497.74}},
		{LL{Lat: 52.482728, Lon: -1.908445}, UTM{ZoneNumber: 30, ZoneLetter: 'U', Easting: 574125.98, Northing: 5815290.89}},
		{LL{Lat: -19.887495, Lon: -43.932663}, UTM{ZoneNumber: 23, ZoneLetter: 'K', Easting: 611733.14, Northing: 7800614.37}},
		{LL{Lat: 60.0, Lon: 4.0}, UTM{ZoneNumber: 32, ZoneLetter: 'V', Easting: 221288.77, Northing: 6661953.04}},  // Norway 31->32
		{LL{Lat: 75.0, Lon: 8.0}, UTM{ZoneNumber: 31, ZoneLetter: 'X', Easting: 644293.43, Northing: 8329692.65}},  // Svalbard 32->31
		{LL{Lat: 75.0, Lon: 10.0}, UTM{ZoneNumber: 33, ZoneLetter: 'X', Easting: 355706.57, Northing: 8329692.65}}, // Svalbard 32->33
		{LL{Lat: 75.0, Lon: 22.0}, UTM{ZoneNumber: 35, ZoneLetter: 'X', Easting: 355706.57, Northing: 8329692.65}}, // Svalbard 34->35
		{LL{Lat: 75.0, Lon: 32.0}, UTM{ZoneNumber: 35, ZoneLetter: 'X', Easting: 644293.43, Northing: 8329692.65}}, // Svalbard 36->35
		{LL{Lat: 75.0, Lon: 34.0}, UTM{ZoneNumber: 37, ZoneLetter: 'X', Easting: 355706.57, Northing: 8329692.65}}, // Svalbard 36->37
		// negative tests
		// nothing to do here
	}

	for _, test := range tests {
		utm := test.ll.ToUTM()
		function := fmt.Sprintf("ll = %s, ToUTM()", test.ll)
		got := utm.String()
		want := test.utm.String()
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}

func TestLL_ToMGRS(t *testing.T) {

	var tests = []struct {
		ll       LL     // in
		accuracy int    // in
		mgrs     string // out
		err      error  // out
	}{
		// positive tests
		{LL{Lat: 51.95, Lon: 7.53}, 1, "32ULC9897356497", nil},
		{LL{Lat: 51.95, Lon: 7.53}, 100, "32ULC989564", nil},
		{LL{Lat: -19.887495, Lon: -43.932663}, 1, "23KPU1173300614", nil},
		{LL{Lat: 0.0, Lon: -0.592328}, 1, "30NYF6799300000", nil},
		// negative tests
		{LL{Lat: 51.95, Lon: 188.53}, 100, "", fmt.Errorf("invalid longitude, lon = 188.53")},
		{LL{Lat: 51.95, Lon: -188.53}, 100, "", fmt.Errorf("invalid longitude, lon = -188.53")},
		{LL{Lat: 99.95, Lon: 7.53}, 100, "", fmt.Errorf("invalid latitude, lat = 99.95")},
		{LL{Lat: -99.95, Lon: 7.53}, 100, "", fmt.Errorf("invalid latitude, lat = -99.95")},
		{LL{Lat: 88.95, Lon: 7.53}, 100, "", fmt.Errorf("polar regions below 80°S and above 84°N not supported, lat = 88.95")},
		{LL{Lat: -88.95, Lon: 7.53}, 100, "", fmt.Errorf("polar regions below 80°S and above 84°N not supported, lat = -88.95")},
	}

	for _, test := range tests {
		mgrs, err := test.ll.ToMGRS(test.accuracy)
		function := fmt.Sprintf("ll = %s, ll.ToMGRS(%d)", test.ll, test.accuracy)
		got := fmt.Sprintf("%s %v", mgrs, err)
		want := fmt.Sprintf("%s %v", test.mgrs, test.err)
		if got != want {
			t.Errorf("\n%s -> %s != %s\n", function, got, want)
		}
	}
}
