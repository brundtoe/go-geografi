package proj

import (
	"fmt"
	"strings"
)

// USNG defines coordinate in USNG format
type USNG string

// String returns the stringified USNG object
func (usng USNG) String() string {
	return string(usng)
}

/*
ToLL converts USNG/UTM to Lon Lat.
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

/*
ToMGRS converts USNG to MGRS
*/
func (usng USNG) ToMGRS() MGRS {
	return MGRS(strings.Replace(string(usng), " ", "", 3))
}

func (usng USNG) ToUTM() (UTM, int, error) {
	mgrs := usng.ToMGRS()
	return mgrs.ToUTM()
}
