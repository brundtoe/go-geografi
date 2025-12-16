package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

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
			transform(record)
		}
	}

}

func transform(record []string) {
	location := mgrs.City{}
	location.BuildCity(record)
	milgrid := location.CityToMgrs()
	usng := location.CityToUsng()
	fmt.Printf("%-18s %s %s", location.Name, milgrid, usng)

	east := string(usng[7:12])
	north := string(usng[13:18])
	mgrsEast := string(milgrid[5:10])
	mgrsNorth := string(milgrid[10:15])
	if strings.Compare(east, record[11]) != 0 || strings.Compare(mgrsEast, record[11]) != 0 {
		fmt.Print(" East")
	}
	if strings.Compare(north, record[12]) != 0 || strings.Compare(mgrsNorth, record[12]) != 0 {
		fmt.Print(" North")
	}
	fmt.Println("")
}
