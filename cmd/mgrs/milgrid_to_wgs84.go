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

	fmt.Printf("%-18s %18s %18s %20s %20s\n", "City", "MilGrid", "USNG", "MilGridLL", "Geolocation")
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
			transformMilGrid(record)
		}
	}

}

func transformMilGrid(record []string) {
	location := mgrs.City{}
	location.BuildCity(record)
	//fixme foranstillede nuller for east og north
	//to build usng location
	milGrid := location.Utm.ToMGRS(1)
	milGridLL, _, _ := milGrid.ToLL()
	usng := location.Utm.ToUSNG(1)
	usngLL, _, _ := usng.ToLL()
	fmt.Printf("%-18s %18s %18s %20s %20s", location.Name, milGrid, usng, milGridLL, location.Geoloc)

	if milGridLL != usngLL {
		fmt.Printf("MilGrid and USNG differ for %s", location.Name)
	}

	if math.Abs(milGridLL.Lon-location.Geoloc.Lon) > 0.0001 || math.Abs(milGridLL.Lat-location.Geoloc.Lat) > 0.0001 {
		fmt.Printf("MilGrid and Geolocation differ for %s", location.Name)
	}

	fmt.Println()
}
