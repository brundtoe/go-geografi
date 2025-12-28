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

	fmt.Printf("%-18s %-25s %-25s\n", "City", "City UTM", "Transformed UTM")
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
			fromWgs84toUtm(record)
		}
	}

}

func fromWgs84toUtm(record []string) {
	location := mgrs.City{}
	location.BuildCity(record)
	utm := location.Geoloc.ToUTM()
	fmt.Printf("%-18s %25s %25s", location.Name, location.Utm, utm)

	if strings.Compare(location.Utm.String(), utm.String()) != 0 {
		fmt.Print(" Transformed UTM differs")
	}
	fmt.Println("")
}
