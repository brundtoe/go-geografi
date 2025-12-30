package proj

import (
	"fmt"
	"strings"
)

// USNG defines coordinate in USNG format
/*
USNG is a string composed of utm zone, utm zone letter, 100-km grid letter, easting and northing

 For the city of Roskilde: 33U UB 162 700 with accuracy of 100 meters
		- 33 is the utm zone number
		- U is the utm zone letter
		- UB is the 100-km grid letter
		- 162 is the easting
		- 700 is the northing

Se also
 - [MGRS]
 - https://en.wikipedia.org/wiki/Military_Grid_Reference_System

*/
type USNG string

// String returns the stringified USNG object
func (usng USNG) String() string {
	return string(usng)
}

/*
ToLL converts USNG to latitude longitude.
*/
func (usng USNG) ToLL() (LL, int, error) {

	mgrs := usng.ToMGRS()

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

// ToMGRS converts USNG to MGRS
func (usng USNG) ToMGRS() MGRS {
	return MGRS(strings.Replace(string(usng), " ", "", 3))
}

// ToUTM converts USNG to UTM
func (usng USNG) ToUTM() (UTM, int, error) {
	mgrs := usng.ToMGRS()
	utm, accuracy, err := mgrs.ToUTM()
	if err != nil {
		return UTM{}, 0, err
	}
	return utm, accuracy, nil
}
