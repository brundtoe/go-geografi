package proj

import (
	"fmt"
	"math"
)

// LL defines coordinate in Longitude / Latitude
type LL struct {
	Lat float64
	Lon float64
}

/*
String returns the stringified LL object.

The order of the coordinates is latitude, longitude in degrees (order according to ISO-6709, precision 0.11 meter).

	For the city of Holstebro: 56.366667 8.616667
*/
func (ll LL) String() string {

	return fmt.Sprintf("%.6f %.6f", ll.Lat, ll.Lon)
}

func (ll LL) validateLL() (string, error) {
	if ll.Lon < -180 || ll.Lon > 180 {
		return "", fmt.Errorf("invalid longitude, lon = %v", ll.Lon)
	}
	if ll.Lat < -90 || ll.Lat > 90 {
		return "", fmt.Errorf("invalid latitude, lat = %v", ll.Lat)
	}
	if ll.Lat < -80 || ll.Lat > 84 {
		return "", fmt.Errorf("polar regions below 80°S and above 84°N not supported, lat = %v", ll.Lat)
	}
	return "", nil
}

/*
ToMGRS converts latitude longitude to MGRS

The accuracy holds the wanted accuracy in meters. Possible values are 1, 10, 100, 1000 or 10.000 meters.
*/
func (ll LL) ToMGRS(accuracy int) (MGRS, error) {

	str, err := ll.validateLL()
	if err != nil {
		return MGRS(str), err
	}
	utm := ll.ToUTM()
	mgrs := utm.ToMGRS(accuracy)

	return mgrs, nil
}

/*
ToUSNG converts latitude longitude to USNG.

The accuracy holds the wanted accuracy in meters. Possible values are 1, 10, 100, 1000 or 10.000 meters.
*/
func (ll LL) ToUSNG(accuracy int) (USNG, error) {
	str, err := ll.validateLL()
	if err != nil {
		return USNG(str), err
	}
	utm := ll.ToUTM()
	return utm.ToUSNG(accuracy), nil
}

/*
ToUTM converts latitude longitude to UTM.
*/
func (ll LL) ToUTM() UTM {

	Lat := ll.Lat
	Long := ll.Lon
	a := 6378137.0           //ellip.radius;
	eccSquared := 0.00669438 //ellip.eccsq;
	k0 := 0.9996
	LatRad := degToRad(Lat)
	LongRad := degToRad(Long)

	ZoneNumber := 0 // (int)
	ZoneNumber = int(math.Floor((Long+180)/6) + 1)

	// make sure the longitude 180.00 is in Zone 60
	if Long == 180 {
		ZoneNumber = 60
	}

	// Special zone for Norway
	if Lat >= 56.0 && Lat < 64.0 && Long >= 3.0 && Long < 12.0 {
		ZoneNumber = 32
	}

	// special zones for Svalbard
	if Lat >= 72.0 && Lat < 84.0 {
		if Long >= 0.0 && Long < 9.0 {
			ZoneNumber = 31
		} else if Long >= 9.0 && Long < 21.0 {
			ZoneNumber = 33
		} else if Long >= 21.0 && Long < 33.0 {
			ZoneNumber = 35
		} else if Long >= 33.0 && Long < 42.0 {
			ZoneNumber = 37
		}
	}

	LongOrigin := (ZoneNumber-1)*6 - 180 + 3 // +3 puts origin in middle of zone
	LongOriginRad := degToRad(float64(LongOrigin))

	eccPrimeSquared := eccSquared / (1 - eccSquared)

	N := a / math.Sqrt(1-eccSquared*math.Sin(LatRad)*math.Sin(LatRad))
	T := math.Tan(LatRad) * math.Tan(LatRad)
	C := eccPrimeSquared * math.Cos(LatRad) * math.Cos(LatRad)
	A := math.Cos(LatRad) * (LongRad - LongOriginRad)

	M := a * ((1-eccSquared/4-3*eccSquared*eccSquared/64-5*eccSquared*eccSquared*eccSquared/256)*LatRad - (3*eccSquared/8+3*eccSquared*eccSquared/32+45*eccSquared*eccSquared*eccSquared/1024)*math.Sin(2*LatRad) + (15*eccSquared*eccSquared/256+45*eccSquared*eccSquared*eccSquared/1024)*math.Sin(4*LatRad) - (35*eccSquared*eccSquared*eccSquared/3072)*math.Sin(6*LatRad))

	UTMEasting := (k0*N*(A+(1-T+C)*A*A*A/6.0+(5-18*T+T*T+72*C-58*eccPrimeSquared)*A*A*A*A*A/120.0) + 500000.0)

	UTMNorthing := (k0 * (M + N*math.Tan(LatRad)*(A*A/2+(5-T+9*C+4*C*C)*A*A*A*A/24.0+(61-58*T+T*T+600*C-330*eccPrimeSquared)*A*A*A*A*A*A/720.0)))
	if Lat < 0.0 {
		UTMNorthing += 10000000.0 // 10.000.000 meters offset for the Southern Hemisphere
	}

	utm := UTM{}
	utm.ZoneNumber = ZoneNumber
	utm.ZoneLetter = getLetterDesignator(Lat)
	utm.Easting = UTMEasting
	utm.Northing = UTMNorthing

	return utm
}
