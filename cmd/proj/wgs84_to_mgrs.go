package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

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
			transformLLtoMGRS(record)
			fmt.Print(".")
			i += 1
		}
	}
	fmt.Printf("\nAntal linjer behandlet: %d\n", i)
}

/**
 * Transform LL to MGRS og USNG
 */

func transformLLtoMGRS(record []string) {
	location := proj.City{}
	location.BuildCity(record)
	// de to funktioner kalder undervejs LL.ToUTM()
	mgrs, _ := location.Geoloc.ToMGRS(1)
	usng, _ := location.Geoloc.ToUSNG(1)

	if strings.Compare(string(mgrs), string(usng.ToMGRS())) != 0 {
		fmt.Printf(" MGRS og USNG differ  %-18s %s %s\n", location.Name, mgrs, usng)
	}

	east := string(usng[7:12])
	north := string(usng[13:18])
	mgrsEast := string(mgrs[5:10])
	mgrsNorth := string(mgrs[10:15])
	if strings.Compare(east, record[11]) != 0 || strings.Compare(mgrsEast, record[11]) != 0 {
		fmt.Printf("East %-18s %s %s\n", location.Name, mgrs, usng)
	}
	if strings.Compare(north, record[12]) != 0 || strings.Compare(mgrsNorth, record[12]) != 0 {
		fmt.Printf("North %-18s %s %s\n", location.Name, mgrs, usng)

	}
}
