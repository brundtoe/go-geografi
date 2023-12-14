// Read csv file cities.csv, convert from lat/lon to UTM and back again
package main

import (
	"encoding/csv"
	"fmt"
	"github.com/brundtoe/go-geografi/pkg/taylor"
	"github.com/brundtoe/go-geografi/pkg/utils"
	"io"
	"log"
	"math"
	"strconv"
)

func main() {

	filename := "geografi/cities.csv"
	fp, err := utils.OpenDataFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Printf("Datafilen %s lukkes", filename)
		if err = fp.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := csv.NewReader(fp)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// fmt.Println("Kilde: ", record[1], record[8], record[9], record[10], record[11], record[12])
		if record[8] != "Zone" {
			transform(record)
		}
	}
}

func transform(koord []string) {

	const diff = 0.0000000001
	latitude, _ := strconv.ParseFloat(koord[6], 64)
	longitude, _ := strconv.ParseFloat(koord[7], 64)

	zone, _ := strconv.ParseInt(koord[8], 10, 64)
	east, north := taylor.LatLonToUTMXY(latitude, longitude, int(zone))

	lat, lon := taylor.UTMXYToLatLon(east, north, int(zone), false)

	// fmt.Printf("Lat\t %4f\n", taylor.RadToDeg(lat))
	// fmt.Printf("Lon\t %4f\n", taylor.RadToDeg(lon))

	fmt.Printf("%18s East\t %4f North\t %4f\n", koord[1], east, north)

	if math.Abs(latitude-taylor.RadToDeg(lat)) > diff {
		fmt.Printf("%s Konvertering af latitude overskrider acceptabel tolerance\n", koord[1])
	}
	if math.Abs(longitude-taylor.RadToDeg(lon)) > diff {
		fmt.Printf("%s Konvertering af longitude overskrider acceptabel tolerance\n", koord[1])
	}
}
