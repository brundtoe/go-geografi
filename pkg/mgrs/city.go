package mgrs

import "strconv"

type City struct {
	Name         string
	Zip          string
	Municipality string
	Region       string
	Population   int64
	Geoloc       LL
	Utm          UTM
	Zone         int64
	Easting      float64
	Northing     float64
	Belt         string
	Kmkv         string
	East         int64
	North        int64
}

func (city *City) CityToMgrs() MGRS {
	location := city.Geoloc
	mgrs, _ := location.ToMGRS(1)
	return mgrs

}

func (city *City) CityToUsng() USNG {
	location := LL{
		Lat: city.Geoloc.Lat,
		Lon: city.Geoloc.Lon,
	}
	utm := location.ToUTM()
	usng := utm.ToUSNG(1)
	return usng

}

func (city *City) BuildCity(koord []string) {
	city.Name = koord[1]
	city.Zip = koord[2]
	city.Municipality = koord[3]
	city.Region = koord[4]
	city.Population, _ = strconv.ParseInt(koord[5], 10, 64)
	city.Geoloc.Lat, _ = strconv.ParseFloat(koord[6], 64)
	city.Geoloc.Lon, _ = strconv.ParseFloat(koord[7], 64)
	zoneNumber, _ := strconv.ParseInt(koord[8], 10, 64)
	city.Utm.ZoneNumber = int(zoneNumber)
	zoneLetter := []byte(koord[9])
	city.Utm.ZoneLetter = zoneLetter[0]
	city.Utm.Easting, _ = strconv.ParseFloat(koord[13], 64)
	city.Utm.Northing, _ = strconv.ParseFloat(koord[14], 64)
	city.Kmkv = koord[10]
	city.East, _ = strconv.ParseInt(koord[11], 10, 64)
	city.North, _ = strconv.ParseInt(koord[12], 10, 64)

}
