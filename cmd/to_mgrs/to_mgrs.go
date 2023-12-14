package main

import (
	"encoding/csv"
	"fmt"
	"github.com/brundtoe/go-geografi/pkg/geoutm"
	"github.com/brundtoe/go-geografi/pkg/utils"
	"io"
	"log"
	"os"
	"strconv"
)

type City struct {
	name         string
	zip          string
	municipality string
	region       string
	population   int64
	geoloc       geoutm.LL
	zone         int64
	belt         string
	kmkv         string
	east         int64
	north        int64
	easting      float64
	northing     float64
}

func main() {

	part := "geografi/cities.csv"
	filename, err := utils.GetDataPath(part, "PLATFORM")
	if err != nil {
		fmt.Printf("Det er ikke muligt at finde path filen %s: %s", part, err)
		os.Exit(1)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := csv.NewReader(file)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var city City
		// first line contains field names
		if record[1] != "City" {
			city.buildCity(record)
			milgrid := city.toMgrs()
			usng := city.toUsng()
			fmt.Printf("%-18s %s %s\n", city.name, milgrid, usng)
		}
	}
}

func (city City) toMgrs() geoutm.MGRS {
	location := geoutm.LL{
		Lat: city.geoloc.Lat,
		Lon: city.geoloc.Lon,
	}
	mgrs, _ := location.ToMGRS(1)
	return mgrs

}

func (city City) toUsng() geoutm.USNG {
	location := geoutm.LL{
		Lat: city.geoloc.Lat,
		Lon: city.geoloc.Lon,
	}
	utm := location.ToUTM()
	usng := utm.ToUSNG(1)
	return usng

}

func (city *City) buildCity(koord []string) {
	city.name = koord[1]
	city.zip = koord[2]
	city.municipality = koord[3]
	city.region = koord[4]
	city.population, _ = strconv.ParseInt(koord[5], 10, 64)
	city.geoloc.Lat, _ = strconv.ParseFloat(koord[6], 64)
	city.geoloc.Lon, _ = strconv.ParseFloat(koord[7], 64)
	city.zone, _ = strconv.ParseInt(koord[8], 10, 64)
	city.belt = koord[9]
	city.kmkv = koord[10]
	city.east, _ = strconv.ParseInt(koord[11], 10, 64)
	city.north, _ = strconv.ParseInt(koord[12], 10, 64)
	city.easting, _ = strconv.ParseFloat(koord[13], 64)
	city.northing, _ = strconv.ParseFloat(koord[14], 64)

}
