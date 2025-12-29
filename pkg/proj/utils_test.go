package proj

import (
	"fmt"
	"math"
	"testing"
)

const _DIFF = 0.0000001

func TestDegToRad(t *testing.T) {

	var tests = []struct {
		deg float64
		rad float64
	}{
		{0.0, 0.0},
		{30.0, math.Pi / 6.0},
		{60.0, math.Pi / 3.0},
		{-90.0, -math.Pi / 2.0},
		{90.0, math.Pi / 2.0},
		{180.0, math.Pi},
		{270.0, 3.0 * math.Pi / 2.0},
	}

	for _, test := range tests {
		got := degToRad(test.deg)
		want := test.rad
		function := fmt.Sprintf("degToRad %f -> %f", test.deg, test.rad)
		if math.Abs(got-want) > _DIFF {
			t.Errorf("\n%s ==> %f != %f\n", function, got, want)
		}
	}
}

func TestRadToDeg(t *testing.T) {

	var tests = []struct {
		rad float64
		deg float64
	}{
		{0.0, 0.0},
		{math.Pi / 6.0, 30.0},
		{math.Pi / 3.0, 60.0},
		{-math.Pi / 2.0, -90.0},
		{math.Pi / 2.0, 90.0},
		{math.Pi, 180.0},
		{3.0 * math.Pi / 2.0, 270.0},
	}

	for _, test := range tests {

		got := radToDeg(test.rad)
		want := test.deg
		function := fmt.Sprintf("radToDeg %f -> %f", test.rad, test.deg)
		if math.Abs(got-want) > _DIFF {
			t.Errorf("%s ==> %f != %f ", function, got, want)
		}
	}
}

func TestGetLetterDesignator(t *testing.T) {
	var tests = []struct {
		lat        float64
		zoneLetter byte
	}{{39.693611, 'S'},
		{43.889722, 'T'},
		{55.130067, 'U'},
		{57.723661, 'V'},
		{67.208333, 'W'},
		//negativ test transformation kun gyldig for latitude mellem 80 og -80
		{-80.723661, 'Z'},
		{84.723661, 'Z'},
		{-84.723661, 'Z'},
	}

	for _, test := range tests {

		got := getLetterDesignator(test.lat)
		want := test.zoneLetter
		function := fmt.Sprintf("getLetterDesignator(%f)", test.lat)
		if got != want {
			t.Errorf("%s ==> %c != %c", function, got, want)
		}
	}
}

func TestGet100kID(t *testing.T) {
	var tests = []struct {
		easting    float64
		northing   float64
		zoneNumber int
		kmkv       string
	}{
		{594857.92, 6399059.92, 32, "NJ"},
		{489092.85, 6101296.46, 32, "MG"},
		{498549.89, 6221228.79, 32, "MH"},
		{687827.95, 6135417.54, 32, "PG"},
		{334434.79, 6187906.23, 33, "UB"},
		{331063.21, 6201629.08, 33, "UC"},
		{509102.23, 6110009.07, 33, "WB"},
	}
	for _, test := range tests {
		got := get100kID(test.easting, test.northing, test.zoneNumber)
		want := test.kmkv
		function := fmt.Sprintf("get100kID (%f, %f, %d)", test.easting, test.northing, test.zoneNumber)
		if got != want {
			t.Errorf("%s ==> %s != %s", function, got, want)
		}
	}
}
