package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"

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
		location := mgrs.City{}
		// first line contains field names
		if record[1] != "City" {
			location.BuildCity(record)
			milgrid := location.CityToMgrs()
			usng := location.CityToUsng()
			fmt.Printf("%-18s %s %s\n", location.Name, milgrid, usng)
		}
	}
}
