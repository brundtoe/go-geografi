package proj

import (
	"fmt"
	"math"
)

// setOriginColumnLetters defines the column letters (for easting) of the lower left value, per set.
const setOriginColumnLetters = "AJSAJS"

// setOriginRowLetters defines the row letters (for northing) of the lower left value, per set.
const setOriginRowLetters = "AFAFAF"

// character constants
const (
	charA = 65 // character 'A'
	charI = 73 // character 'I'
	charO = 79 // character 'O'
	charV = 86 // character 'V'
	charZ = 90 // character 'Z'
)

/*
degToRad converts from degrees to radians.
del holds the angle in degrees.
*/
func degToRad(deg float64) float64 {

	return (deg * (math.Pi / 180.0))
}

/*
radToDeg converts from radians to degrees.
rad holds the angle in radians.
*/
func radToDeg(rad float64) float64 {

	return (180.0 * (rad / math.Pi))
}

/*
getLetterDesignator calculates the MGRS letter designator for the given latitude.
lat holds lat the latitude in WGS84 to get the letter designator for.
*/
func getLetterDesignator(lat float64) byte {

	// This is here as an error flag to show that the Latitude is outside MGRS limits
	LetterDesignator := 'Z'

	if (84 >= lat) && (lat >= 72) {
		LetterDesignator = 'X'
	} else if (72 > lat) && (lat >= 64) {
		LetterDesignator = 'W'
	} else if (64 > lat) && (lat >= 56) {
		LetterDesignator = 'V'
	} else if (56 > lat) && (lat >= 48) {
		LetterDesignator = 'U'
	} else if (48 > lat) && (lat >= 40) {
		LetterDesignator = 'T'
	} else if (40 > lat) && (lat >= 32) {
		LetterDesignator = 'S'
	} else if (32 > lat) && (lat >= 24) {
		LetterDesignator = 'R'
	} else if (24 > lat) && (lat >= 16) {
		LetterDesignator = 'Q'
	} else if (16 > lat) && (lat >= 8) {
		LetterDesignator = 'P'
	} else if (8 > lat) && (lat >= 0) {
		LetterDesignator = 'N'
	} else if (0 > lat) && (lat >= -8) {
		LetterDesignator = 'M'
	} else if (-8 > lat) && (lat >= -16) {
		LetterDesignator = 'L'
	} else if (-16 > lat) && (lat >= -24) {
		LetterDesignator = 'K'
	} else if (-24 > lat) && (lat >= -32) {
		LetterDesignator = 'J'
	} else if (-32 > lat) && (lat >= -40) {
		LetterDesignator = 'H'
	} else if (-40 > lat) && (lat >= -48) {
		LetterDesignator = 'G'
	} else if (-48 > lat) && (lat >= -56) {
		LetterDesignator = 'F'
	} else if (-56 > lat) && (lat >= -64) {
		LetterDesignator = 'E'
	} else if (-64 > lat) && (lat >= -72) {
		LetterDesignator = 'D'
	} else if (-72 > lat) && (lat >= -80) {
		LetterDesignator = 'C'
	}

	return byte(LetterDesignator)
}

/*
get100kID gets the two-letter 100k designator for a given UTM easting, northing and zone number value.
*/
func get100kID(easting, northing float64, zoneNumber int) string {

	setParm := get100kSetForZone(zoneNumber)
	setColumn := int(math.Floor(easting / 100000))
	setRow := int(math.Floor(northing/100000)) % 20

	return getLetter100kID(setColumn, setRow, setParm)
}

/*
get100kSetForZone gets the MGRS 100K set for a given UTM zone number.
*/
func get100kSetForZone(i int) int {

	// UTM zones are grouped, and assigned to one of a group of 6 sets.
	numberOf100kSets := 6

	setParm := i % numberOf100kSets
	if setParm == 0 {
		setParm = numberOf100kSets
	}

	return setParm
}

/*
getLetter100kID gets the two-letter MGRS 100k designator given information translated from the UTM northing, easting and zone number.
column holds the column index as it relates to the MGRS 100k set spreadsheet, created from the UTM easting. Values are 1-8.
row holds the row index as it relates to the MGRS 100k set spreadsheet, created from the UTM northing value. Values are from 0-19.
parm holds the set block, as it relates to the MGRS 100k set spreadsheet, created from the UTM zone. Values are from 1-60.
*/
func getLetter100kID(column, row, parm int) string {

	// colOrigin and rowOrigin are the letters at the origin of the set
	index := parm - 1
	colOrigin := setOriginColumnLetters[index]
	rowOrigin := setOriginRowLetters[index]

	// colInt and rowInt are the letters to build to return
	colInt := int(colOrigin) + column - 1
	rowInt := int(rowOrigin) + row
	rollover := false

	if colInt > charZ {
		colInt = colInt - charZ + charA - 1
		rollover = true
	}

	if colInt == charI || (colOrigin < charI && colInt > charI) || ((colInt > charI || colOrigin < charI) && rollover) {
		colInt++
	}

	if colInt == charO || (colOrigin < charO && colInt > charO) || ((colInt > charO || colOrigin < charO) && rollover) {
		colInt++
		if colInt == charI {
			colInt++
		}
	}

	if colInt > charZ {
		colInt = colInt - charZ + charA - 1
	}

	if rowInt > charV {
		rowInt = rowInt - charV + charA - 1
		rollover = true
	} else {
		rollover = false
	}

	if ((rowInt == charI) || ((rowOrigin < charI) && (rowInt > charI))) || (((rowInt > charI) || (rowOrigin < charI)) && rollover) {
		rowInt++
	}

	if ((rowInt == charO) || ((rowOrigin < charO) && (rowInt > charO))) || (((rowInt > charO) || (rowOrigin < charO)) && rollover) {
		rowInt++
		if rowInt == charI {
			rowInt++
		}
	}

	if rowInt > charV {
		rowInt = rowInt - charV + charA - 1
	}

	twoLetter := string(rune(colInt)) + string(rune(rowInt))
	return twoLetter
}

/*
getEastingFromChar gets the easting value that should be added to the other, secondary easting value.
e holds the first letter from a two-letter MGRS 100k zone.
set holds the MGRS table set for the zone number.
*/
func getEastingFromChar(e byte, set int) (float64, error) {

	// colOrigin is the letter at the origin of the set for the column
	curCol := setOriginColumnLetters[set-1]
	eastingValue := 100000.0
	rewindMarker := false

	for curCol != e {
		curCol++
		if curCol == charI {
			curCol++
		}
		if curCol == charO {
			curCol++
		}
		if curCol > charZ {
			if rewindMarker {
				return -1.0, fmt.Errorf("bad character: %v", e)
			}
			curCol = charA
			rewindMarker = true
		}
		eastingValue += 100000.0
	}

	return eastingValue, nil
}

/*
getNorthingFromChar gets the northing value that should be added to the other, secondary northing value.
n holds the second letter of the MGRS 100k zone.
set holds the MGRS table set number, which is dependent on the UTM zone number.
Remark: You have to remember that Northings are determined from the equator, and the vertical
cycle of letters mean a 2000000 additional northing meters. This happens
approx. every 18 degrees of latitude. This method does *NOT* count any
additional northings. You have to figure out how many 2000000 meters need
to be added for the zone letter of the MGRS coordinate.
*/
func getNorthingFromChar(n byte, set int) (float64, error) {

	if n > 'V' {
		return 0.0, fmt.Errorf("invalid northing, char = %v", n)
	}

	// rowOrigin is the letter at the origin of the set for the column
	curRow := setOriginRowLetters[set-1]
	northingValue := 0.0
	rewindMarker := false

	for curRow != byte(n) {
		curRow++
		if curRow == charI {
			curRow++
		}
		if curRow == charO {
			curRow++
		}
		// fixing a bug making whole application hang in this loop when 'n' is a wrong character
		if curRow > charV {
			if rewindMarker { // making sure that this loop ends
				return -1.0, fmt.Errorf("bad character, char = %v", n)
			}
			curRow = charA
			rewindMarker = true
		}
		northingValue += 100000.0
	}

	return northingValue, nil
}

/*
getMinNorthing gets the minimum northing value of a MGRS zone.
zoneLetter holds the MGRS zone to get the min northing for.
*/
func getMinNorthing(zoneLetter byte) (float64, error) {

	var northing float64

	switch zoneLetter {
	case 'C':
		northing = 1100000.0
	case 'D':
		northing = 2000000.0
	case 'E':
		northing = 2800000.0
	case 'F':
		northing = 3700000.0
	case 'G':
		northing = 4600000.0
	case 'H':
		northing = 5500000.0
	case 'J':
		northing = 6400000.0
	case 'K':
		northing = 7300000.0
	case 'L':
		northing = 8200000.0
	case 'M':
		northing = 9100000.0
	case 'N':
		northing = 0.0
	case 'P':
		northing = 800000.0
	case 'Q':
		northing = 1700000.0
	case 'R':
		northing = 2600000.0
	case 'S':
		northing = 3500000.0
	case 'T':
		northing = 4400000.0
	case 'U':
		northing = 5300000.0
	case 'V':
		northing = 6200000.0
	case 'W':
		northing = 7000000.0
	case 'X':
		northing = 7900000.0
	default:
		northing = -1.0
	}

	if northing >= 0.0 {
		return northing, nil
	}

	return northing, fmt.Errorf("Invalid zone letter: %v", zoneLetter)
}
