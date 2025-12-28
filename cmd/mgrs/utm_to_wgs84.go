package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"

	"github.com/brundtoe/go-geografi/pkg/mgrs"
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
		}
	}

}

func fromUtmToWgs84(record []string) {
	location := mgrs.City{}
	location.BuildCity(record)
	ll, _ := location.Utm.ToLL()
	fmt.Printf("%20s %-18s", ll.String(), location.Name)

	if math.Abs(ll.Lon-location.Geoloc.Lon) > 0.000001 {
		fmt.Printf("%10.6f", location.Geoloc.Lon)
	}
	if math.Abs(ll.Lat-location.Geoloc.Lat) > 0.000001 {
		fmt.Printf("Lat %10.6f", location.Geoloc.Lat)
	}
	fmt.Println()
}
