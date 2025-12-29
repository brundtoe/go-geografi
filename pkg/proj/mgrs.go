package proj

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// MGRS defines coordinate in MGRS
type MGRS string

// String returns the stringified MGRS object
func (mgrs MGRS) String() string {
	return string(mgrs)
}

/*
ToLL converts MGRS/UTM to Lon Lat.
*/
func (mgrs MGRS) ToLL() (LL, int, error) {

	utm, accuracy, err := mgrs.ToUTM()
	if err != nil {
		return LL{}, 0, fmt.Errorf("error <%v> at mgrs.ToUTM()", err)
	}

	ll, err := utm.ToLL()
	if err != nil {
		return LL{}, 0, fmt.Errorf("error <%v> at utm.ToLL(), utm = %#v", err, utm)
	}

	return ll, accuracy, nil
}

func (mgrs MGRS) ToUSNG() USNG {
	zoneGrid := mgrs[0:3]
	kmkv := mgrs[3:5]
	rest := mgrs[5:]
	digits := len(rest) / 2
	east := mgrs[5 : 5+digits]
	north := mgrs[5+digits:]
	fmt.Printf("digits = %d\n", digits)
	return USNG(zoneGrid + " " + kmkv + " " + east + " " + north)

}

/*
ToUTM converts MGRS to UTM.
*/
func (mgrs MGRS) ToUTM() (UTM, int, error) {

	mgrsTmp := string(mgrs)
	if mgrs == "" {
		return UTM{}, 0, fmt.Errorf("invalid empty mgrs string")
	}

	mgrsTmp = strings.ToUpper(mgrsTmp)

	sb := ""
	i := 0

	// get Zone number
	re := regexp.MustCompile("[A-Z]")
	for !re.MatchString(string(mgrsTmp[i])) {
		if i >= 2 {
			return UTM{}, 0, fmt.Errorf("bad conversion, mgrs = %s", mgrs)
		}
		sb += string(mgrsTmp[i])
		i++
	}

	zoneNumberTmp, err := strconv.ParseInt(sb, 10, 0)
	if err != nil {
		return UTM{}, 0, fmt.Errorf("error <%v> at strconv.ParseInt(), string = %v", err, sb)
	}
	zoneNumber := int(zoneNumberTmp)

	// A good MGRS string has to be 4-5 digits long, ##AAA/#AAA at least.
	if i == 0 || i+3 > len(mgrsTmp) {
		return UTM{}, 0, fmt.Errorf("bad conversion, mgrs = %s", mgrs)
	}

	zoneLetter := mgrsTmp[i]
	i++

	// Should we check the zone letter here? Why not.
	if zoneLetter <= 'A' || zoneLetter == 'B' || zoneLetter == 'Y' || zoneLetter >= 'Z' || zoneLetter == 'I' || zoneLetter == 'O' {
		return UTM{}, 0, fmt.Errorf("zone letter %v not handled, mgrs = %s", zoneLetter, mgrs)
	}

	hunK := mgrsTmp[i : i+2]
	i += 2

	set := get100kSetForZone(zoneNumber)

	east100k, err := getEastingFromChar(hunK[0], set)
	if err != nil {
		return UTM{}, 0, fmt.Errorf("error <%v> at getEastingFromChar()", err)
	}

	north100k, err := getNorthingFromChar(hunK[1], set)
	if err != nil {
		return UTM{}, 0, fmt.Errorf("error <%v> at getNorthingFromChar()", err)
	}

	// We have a bug where the northing may be 2.000.000 too low. How do we know when to roll over?
	minNorthing, err := getMinNorthing(zoneLetter)
	if err != nil {
		return UTM{}, 0, fmt.Errorf("error <%v> at getMinNorthing()", err)
	}

	for north100k < minNorthing {
		north100k += 2000000
	}

	// calculate the char index for easting/northing separator
	remainder := len(mgrsTmp) - i

	if remainder%2 != 0 {
		return UTM{}, 0, fmt.Errorf("uneven number of digits, mgrs = %s", mgrs)
	}

	sep := remainder / 2

	sepEasting := 0.0
	sepNorthing := 0.0
	accuracy := 0.0
	if sep > 0 {
		accuracy = 100000.0 / math.Pow(10, float64(sep))

		sepEastingString := mgrsTmp[i : i+sep]
		tmpEasting, err := strconv.ParseFloat(sepEastingString, 64)
		if err != nil {
			return UTM{}, 0, fmt.Errorf("error <%v> at strconv.ParseFloat(), easting string = %v", err, sepEastingString)
		}
		sepEasting = tmpEasting * accuracy

		sepNorthingString := mgrsTmp[i+sep:]
		tmpNorthing, err := strconv.ParseFloat(sepNorthingString, 64)
		if err != nil {
			return UTM{}, 0, fmt.Errorf("error <%v> at strconv.ParseFloat(), northing string = %v", err, sepNorthingString)
		}
		sepNorthing = tmpNorthing * accuracy
	}

	easting := sepEasting + east100k
	northing := sepNorthing + north100k

	utm := UTM{}
	utm.ZoneNumber = zoneNumber
	utm.ZoneLetter = zoneLetter
	utm.Easting = easting
	utm.Northing = northing

	return utm, int(accuracy), nil
}
