package main

import (
	"encoding/csv"
	"fmt"
	"github.com/brundtoe/go-geografi/pkg/geoutm"
	"github.com/brundtoe/go-geografi/pkg/utils"
	"io"
	"log"
	"os"
)

func main() {

	part := "geografi/cities.csv"
	filename, err := utils.GetDataPath(part, "PLATFORM")
	if err != nil {
		fmt.Printf("Det er ikke muligt at finde path filen %s: %s", part, err)
		os.Exit(1)
	}
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	r := csv.NewReader(file)
	r.Comma = ';'

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		location := geoutm.City{}
		// first line contains field names
		if record[1] != "City" {
			location.BuildCity(record)
			milgrid := location.CityToMgrs()
			usng := location.CityToUsng()
			fmt.Printf("%-18s %s %s\n", location.Name, milgrid, usng)
		}
	}
}
