package taylor

import (
	"fmt"
	"math"
	"testing"
)

const diff = 0.0000000001

func TestDegToRad(t *testing.T) {
	actual := DegToRad(60.0)
	expected := math.Pi / 3.0
	if math.Abs(actual-expected) > diff {
		t.Error("Omregning fra Degree to Radians failed")
	}
}

func TestRadToDeg(t *testing.T) {

	actual := RadToDeg(math.Pi / 3.0)
	expected := 60.0
	if math.Abs(actual-expected) > diff {
		t.Error("Omregning fra Radians to Degrees failed")
	}

}

func TestArcLengthOfMeridian(t *testing.T) {
	actual := ArcLengthOfMeridian(1.0074679407550426)
	expected := 6400505.3577935519
	if math.Abs(actual-expected) > diff {
		t.Error("Failed ArcLengthOfMeridian")
	}
}

func TestUTMCentralMeridian(t *testing.T) {
	actual := UTMCentralMeridian(32)
	expected := 0.15707963267948966
	if math.Abs(actual-expected) > diff {
		t.Error("Failed UTMCentralMeridian")
	}
}

func TestFootpointLatitude(t *testing.T) {
	actual := FootpointLatitude(6401620.5711179469)
	expected := 1.0076427063607092

	if math.Abs(actual-expected) > diff {
		t.Error("Failed FootpointLatitude")
	}
}

func TestMapLatLonToXY(t *testing.T) {
	phi := 1.0074679407550426
	lambda := 0.18487625249223444
	lambda0 := 0.15707963267948966
	x, y := MapLatLonToXY(phi, lambda, lambda0)
	east := 94895.878789380513
	north := 6401620.5711179469

	if math.Abs(east-x) > diff {
		t.Error("east calculation failed")
	}
	if math.Abs(north-y) > diff {
		t.Error("north calculation failed")
	}
}

func TestMapXYToLatLon(t *testing.T) {
	x := 94895.878789380469
	y := 6401620.5711179469
	cmeridian := 0.15707963267948966
	lat, lon := MapXYToLatLon(x, y, cmeridian)
	latitude := 1.0074679407550504
	longitude := 0.18487625249224049

	if math.Abs(lat-latitude) > diff {
		t.Errorf("Difference between lat %4f og latitude %4f er for stor", lat, latitude)
	}
	if math.Abs(lon-longitude) > diff {
		t.Errorf("Difference between lon %4f og longitued %4f er for stor", lon, longitude)
	}
}

func TestLatLonToUTMXY(t *testing.T) {
	lat := 57.723661
	lon := 10.592629
	zone := 32
	x, y := LatLonToUTMXY(lat, lon, zone)
	east := 594857.9204378647
	north := 6399059.9228895

	if math.Abs(east-x) > diff {
		t.Errorf("Differencen mellem beregnet easting %4f og forventet %4f er for stor", x, east)
	}
	if math.Abs(north-y) > diff {
		t.Errorf("Differencen mellem beregnet northing %4f og forventet %4f er for stor", y, north)
	}
}

func TestUTMXYToLatLon(t *testing.T) {
	east := 594857.9204378647
	north := 6399059.9228895
	lat, lon := UTMXYToLatLon(east, north, 32, false)
	latitude := DegToRad(57.723661)
	longitude := DegToRad(10.592629)

	if math.Abs(latitude-lat) > diff {
		t.Errorf("Differencen mellem beregnet latitude %4f og forventet %4f er for stor", lat, latitude)
	}
	if math.Abs(longitude-lon) > diff {
		t.Errorf("Differencen mellem beregnet longitude %4f og forventet %4f er for stor", lon, longitude)
	}
}

func ExampleDegToRad() {
	rad := DegToRad(60.0)
	fmt.Printf("Radians %1.6f", rad)
	// Output: Radians 1.047198
}
