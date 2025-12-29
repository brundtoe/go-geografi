package proj

import (
	"strconv"
)

const (
	csvName = iota + 1
	csvZip
	csvMunicipality
	csvRegion
	csvPopulation
	csvLat
	csvLon
	csvZone
	csvBelt
	csvKmKv
	csvEast
	csvNorth
	csvEasting
	csvNorthing
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
	city.Name = koord[csvName]
	city.Zip = koord[csvZip]
	city.Municipality = koord[csvMunicipality]
	city.Region = koord[csvRegion]
	city.Population, _ = strconv.ParseInt(koord[csvPopulation], 10, 64)
	city.Geoloc.Lat, _ = strconv.ParseFloat(koord[csvLat], 64)
	city.Geoloc.Lon, _ = strconv.ParseFloat(koord[csvLon], 64)
	zoneNumber, _ := strconv.ParseInt(koord[csvZone], 10, 64)
	zoneLetter := []byte(koord[csvBelt])
	city.Utm.ZoneNumber = int(zoneNumber)
	city.Utm.ZoneLetter = zoneLetter[0]
	city.Utm.Easting, _ = strconv.ParseFloat(koord[csvEasting], 64)
	city.Utm.Northing, _ = strconv.ParseFloat(koord[csvNorthing], 64)
	city.kmKv = koord[csvKmKv]
	city.Usng = city.Utm.ToUSNG(1)
	city.Mgrs = city.Utm.ToMGRS(1)
	city.East, _ = strconv.ParseInt(koord[csvEast], 10, 64)
	city.North, _ = strconv.ParseInt(koord[csvNorth], 10, 64)
}
