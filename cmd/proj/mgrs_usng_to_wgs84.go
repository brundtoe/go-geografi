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
			fromMgrsToLL(record)
		}
	}

}

func fromMgrsToLL(record []string) {
	location := proj.City{}
	location.BuildCity(record)
	//fixme foranstillede nuller for east og north
	//to build usng location
	//milGrid := location.Utm.ToMGRS(1)
	mgrsLL, _, _ := location.Mgrs.ToLL()
	//usng := location.Utm.ToUSNG(1)
	usngLL, _, _ := location.Usng.ToLL()
	fmt.Printf("%-18s %18s %18s %20s %20s", location.Name, location.Mgrs, location.Usng, mgrsLL, location.Geoloc)

	/**
	 * Check if milGrid og usng conversion to Wgs84 is the same
	 */
	if mgrsLL != usngLL {
		fmt.Printf("MilGrid and USNG differ for %s", location.Name)
	}
	/**
	 * Check if milGrid conversion to Wgs84 is the same as the master wgs84 location i city object
	 * The difference is due to milGrid and usng being integers
	 */
	if math.Abs(mgrsLL.Lon-location.Geoloc.Lon) > 0.0001 || math.Abs(mgrsLL.Lat-location.Geoloc.Lat) > 0.0001 {
		fmt.Printf("MilGrid and Geolocation differ for %s", location.Name)
	}

	fmt.Println()
}
