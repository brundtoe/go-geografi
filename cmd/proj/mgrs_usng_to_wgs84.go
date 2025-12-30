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
	fmt.Println("MGRS: Konverterer MGRS og USNG til WGS84")
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
			FromMgrsToLL(record)
			i += 1
			fmt.Print(".")
		}
	}
	fmt.Printf("\nAntal linjer behandler: %d\n", i)
}

// FromMgrsToLL converts MGRS and USNG to WGS84
/*
The accuracy of the conversion is limited by 0.0001 degrees.
*/
func FromMgrsToLL(record []string) {
	location := proj.City{}
	location.BuildCity(record)
	//fixme foranstillede nuller for east og north
	//to build usng location
	//milGrid := location.Utm.ToMGRS(1)
	mgrsLL, _, _ := location.Mgrs.ToLL()
	//usng := location.Utm.ToUSNG(1)
	usngLL, _, _ := location.Usng.ToLL()

	/**
	 * Check if milGrid og usng conversion to Wgs84 is the same
	 */
	if mgrsLL != usngLL {
		fmt.Printf("\n%-18s %18s %18s %20s %20s\n", location.Name, location.Mgrs, location.Usng, mgrsLL, location.Geoloc)
	}
	/**
	 * Check if milGrid conversion to Wgs84 is the same as the master wgs84 location i city object
	 * The difference is due to milGrid and usng being integers
	 */
	if math.Abs(mgrsLL.Lon-location.Geoloc.Lon) > 0.0001 || math.Abs(mgrsLL.Lat-location.Geoloc.Lat) > 0.0001 {
		fmt.Printf("\n%-18s %18s %18s %20s %20s\n", location.Name, location.Mgrs, location.Usng, mgrsLL, location.Geoloc)
	}

}
