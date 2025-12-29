package proj

import (
	"strconv"
)

const (
	name = iota + 1
	zip
	municipality
	region
	population
	lat
	lon
	zone
	belt
	kmKv
	east
	north
	easting
	northing
)

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

func (city *City) BuildCity(koord []string) {
	city.Name = koord[name]
	city.Zip = koord[zip]
	city.Municipality = koord[municipality]
	city.Region = koord[region]
	city.Population, _ = strconv.ParseInt(koord[population], 10, 64)
	city.Geoloc.Lat, _ = strconv.ParseFloat(koord[lat], 64)
	city.Geoloc.Lon, _ = strconv.ParseFloat(koord[lon], 64)
	zoneNumber, _ := strconv.ParseInt(koord[zone], 10, 64)
	zoneLetter := []byte(koord[belt])
	city.Utm.ZoneNumber = int(zoneNumber)
	city.Utm.ZoneLetter = zoneLetter[0]
	city.Utm.Easting, _ = strconv.ParseFloat(koord[easting], 64)
	city.Utm.Northing, _ = strconv.ParseFloat(koord[northing], 64)
	city.kmKv = koord[kmKv]
	city.Usng = city.Utm.ToUSNG(1)
	city.Mgrs = city.Utm.ToMGRS(1)
	city.East, _ = strconv.ParseInt(koord[east], 10, 64)
	city.North, _ = strconv.ParseInt(koord[north], 10, 64)
}
