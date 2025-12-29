package proj

import (
	"fmt"
	"math"
)

// UTM defines coordinate in Universal Transverse Mercator
type UTM struct {
	ZoneNumber int
	ZoneLetter byte
	Easting    float64
	Northing   float64
}

/*
String returns stringified UTM object.
*/
func (utm UTM) String() string {

	return fmt.Sprintf("%d%c %.2f %.2f", utm.ZoneNumber, utm.ZoneLetter, utm.Easting, utm.Northing)
}

/*
ToLL converts UTM to Lon Lat.
*/
func (utm UTM) ToLL() (LL, error) {

	zoneNumber := utm.ZoneNumber
	zoneLetter := utm.ZoneLetter
	UTMEasting := utm.Easting
	UTMNorthing := utm.Northing

	// check the ZoneNummber is valid
	if zoneNumber < 0 || zoneNumber > 60 {
		return LL{}, fmt.Errorf("invalid zone number, zone number = %v", zoneNumber)
	}

	k0 := 0.9996
	a := 6378137.0           //ellip.radius;
	eccSquared := 0.00669438 //ellip.eccsq;
	e1 := (1 - math.Sqrt(1-eccSquared)) / (1 + math.Sqrt(1-eccSquared))

	// remove 500,000 meters offset for longitude
	x := UTMEasting - 500000.0
	y := UTMNorthing

	// We must know somehow if we are in the Northern or Southern hemisphere, this is the only time we use the letter.
	// So even if the Zone letter isn't exactly correct it should indicate the hemisphere correctly.
	if zoneLetter < 'N' {
		y -= 10000000.0 // remove 10,000,000 meters offset used
		// for southern hemisphere
	}

	// there are 60 zones with zone 1 being at West -180 to -174
	LongOrigin := (zoneNumber-1)*6 - 180 + 3 // +3 puts origin in middle of zone

	eccPrimeSquared := (eccSquared) / (1 - eccSquared)

	M := y / k0
	mu := M / (a * (1 - eccSquared/4 - 3*eccSquared*eccSquared/64 - 5*eccSquared*eccSquared*eccSquared/256))

	phi1Rad := mu + (3*e1/2-27*e1*e1*e1/32)*math.Sin(2*mu) + (21*e1*e1/16-55*e1*e1*e1*e1/32)*math.Sin(4*mu) + (151*e1*e1*e1/96)*math.Sin(6*mu)

	N1 := a / math.Sqrt(1-eccSquared*math.Sin(phi1Rad)*math.Sin(phi1Rad))
	T1 := math.Tan(phi1Rad) * math.Tan(phi1Rad)
	C1 := eccPrimeSquared * math.Cos(phi1Rad) * math.Cos(phi1Rad)
	R1 := a * (1 - eccSquared) / math.Pow(1-eccSquared*math.Sin(phi1Rad)*math.Sin(phi1Rad), 1.5)
	D := x / (N1 * k0)

	lat := phi1Rad - (N1*math.Tan(phi1Rad)/R1)*(D*D/2-(5+3*T1+10*C1-4*C1*C1-9*eccPrimeSquared)*D*D*D*D/24+(61+90*T1+298*C1+45*T1*T1-252*eccPrimeSquared-3*C1*C1)*D*D*D*D*D*D/720)
	lat = radToDeg(lat)

	lon := (D - (1+2*T1+C1)*D*D*D/6 + (5-2*C1+28*T1-3*C1*C1+8*eccPrimeSquared+24*T1*T1)*D*D*D*D*D/120) / math.Cos(phi1Rad)
	lon = float64(LongOrigin) + radToDeg(lon)

	ll := LL{}
	ll.Lat = lat
	ll.Lon = lon

	return ll, nil
}

func (utm UTM) buildGrid(accuracy int, format string) string {

	digits := 0
	// meters to number of digits
	switch accuracy {
	case 1:
		digits = 5
	case 10:
		digits = 4
	case 100:
		digits = 3
	case 1000:
		digits = 2
	case 10000:
		digits = 1
	default:
		digits = 5
	}

	// prepend with leading zeroes
	seasting := "00000" + fmt.Sprintf("%.0f", math.Floor(utm.Easting))
	snorthing := "00000" + fmt.Sprintf("%.0f", math.Floor(utm.Northing))

	east := seasting[len(seasting)-5 : len(seasting)-5+digits]
	north := snorthing[len(snorthing)-5 : len(snorthing)-5+digits]
	kmkv := get100kID(utm.Easting, utm.Northing, utm.ZoneNumber)
	return fmt.Sprintf(format,
		utm.ZoneNumber,
		string(utm.ZoneLetter),
		kmkv,
		east,
		north)

}

/*
ToMGRS converts UTM to MGRS/UTM.
accuracy holds the wanted accuracy in meters. Possible values are 1, 10, 100, 1000 or 10.000 meters.
*/
func (utm UTM) ToMGRS(accuracy int) MGRS {
	return MGRS(utm.buildGrid(accuracy, "%d%s%s%s%s"))

}

/*
ToUSNG converts UTM to USNG.
accuracy holds the wanted accuracy in meters. Possible values are 1, 10, 100, 1000 or 10000 meters.
*/
func (utm UTM) ToUSNG(accuracy int) USNG {

	return USNG(utm.buildGrid(accuracy, "%d%s %s %s %s"))
}
