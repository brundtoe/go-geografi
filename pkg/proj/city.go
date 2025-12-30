package proj

import (
	"strconv"
)

// City csv columns used to map csv record to City struct
const (
	_Name = iota + 1
	_Zip
	_Municipality
	_Region
	_Population
	_Lat
	_Lon
	_Zone
	_Belt
	_KmKv
	_East
	_North
	_Easting
	_Northing
)

// City
/*
City struct is used to store city data which are loaded from a csv file.

The fields
 - Geoloc
 - Utm
 - Usng
 - Mgrs
are used in the transformations.

When a Geolog is transformed to UTM the values in UTM are used to validate the results within an expected error margin.


*/
type City struct {
	Name         string
	Zip          string
	Municipality string
	Region       string
	Population   int64
	Geoloc       LL
	Utm          UTM
	kmKv         string
	Usng         USNG
	Mgrs         MGRS
	East         int64
	North        int64
}

// BuildCity builds a City struct from a csv record
func (city *City) BuildCity(koord []string) {
	city.Name = koord[_Name]
	city.Zip = koord[_Zip]
	city.Municipality = koord[_Municipality]
	city.Region = koord[_Region]
	city.Population, _ = strconv.ParseInt(koord[_Population], 10, 64)
	city.Geoloc.Lat, _ = strconv.ParseFloat(koord[_Lat], 64)
	city.Geoloc.Lon, _ = strconv.ParseFloat(koord[_Lon], 64)
	zoneNumber, _ := strconv.ParseInt(koord[_Zone], 10, 64)
	zoneLetter := []byte(koord[_Belt])
	city.Utm.ZoneNumber = int(zoneNumber)
	city.Utm.ZoneLetter = zoneLetter[0]
	city.Utm.Easting, _ = strconv.ParseFloat(koord[_Easting], 64)
	city.Utm.Northing, _ = strconv.ParseFloat(koord[_Northing], 64)
	city.kmKv = koord[_KmKv]
	city.Usng = city.Utm.ToUSNG(1)
	city.Mgrs = city.Utm.ToMGRS(1)
	city.East, _ = strconv.ParseInt(koord[_East], 10, 64)
	city.North, _ = strconv.ParseInt(koord[_North], 10, 64)
}
