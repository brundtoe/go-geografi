package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"

	"github.com/brundtoe/go-geografi/pkg/proj"
	"github.com/brundtoe/go-geografi/pkg/utils"
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
	i := 0
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		// The first line contains field names
		if record[1] != "City" {
			fromUtmToWgs84(record)
			i += 1
			fmt.Print(".")
		}
	}
	fmt.Printf("\nAntal linjer behandler: %d\n", i)
}

func fromUtmToWgs84(record []string) {
	location := proj.City{}
	location.BuildCity(record)
	ll, _ := location.Utm.ToLL()

	if math.Abs(ll.Lon-location.Geoloc.Lon) > 0.000001 {
		fmt.Printf("\nLongitude: %10.6f\n", location.Geoloc.Lon)
	}
	if math.Abs(ll.Lat-location.Geoloc.Lat) > 0.000001 {
		fmt.Printf("\nLatitude %10.6f\n", location.Geoloc.Lat)
	}

}
