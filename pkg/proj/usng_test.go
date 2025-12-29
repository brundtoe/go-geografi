package proj

import (
	"math"
	"testing"
)

const diff float64 = 0.000001

func TestUSNG_ToMGRS(t *testing.T) {
	expected := "32VNJ9485899060"
	var usng USNG = "32V NJ 94858 99060"
	mgrs := usng.ToMGRS()
	if string(mgrs) != expected {
		t.Errorf("konvertering af USNG %s to MGSR %s fejlede", string(usng), string(mgrs))
	}
}

func TestUSNG_ToUTM(t *testing.T) {
	var usng USNG = "32V NJ 94858 99060"
	expected := UTM{ZoneNumber: 32, ZoneLetter: 'V', Easting: 594858, Northing: 6399060}

	utm, _, _ := usng.ToUTM()
	if utm != expected {
		t.Error("hovsa USNG_ToUTM fejlede")
	}
}

func TestUSNG_ToLL(t *testing.T) {
	var usng USNG = "32V NJ 94858 99060"
	expected := LL{Lat: 57.723661, Lon: 10.592630}
	actual, _, _ := usng.ToLL()
	if math.Abs(actual.Lat-expected.Lat) > diff {
		t.Errorf("Hovsa USNG_ToLL fejlede f.s.v.a. Latitude %f %f", actual.Lat, expected.Lat)
	}
	if math.Abs(actual.Lon-expected.Lon) > diff {
		t.Errorf("Hovsa USNG_ToLL fejlede f.s.v.a. Longitude %f %f", actual.Lon, expected.Lon)
	}
}
