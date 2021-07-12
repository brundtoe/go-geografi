// Read a csv formated file cities.csv and convert from MGRS to UTM
package main

import (
	"encoding/csv"
	"example.com/geografi/pkg/utmabs"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/home/jackie/dumps/geografi/cities.csv")
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
		fmt.Println(record[8], record[9], record[10], record[11], record[12], record[1])
		if record[8] != "Zone" {
			transform(record)
		}
	}

}
func transform(ka []string) {
	koord := utmabs.Mgrs{Zone: ka[8], Belt: ka[9], Kmkv: ka[10], East: ka[11], North: ka[12], Town: ka[1]}

	result, err := utmabs.UtmAbs(koord)

	log.SetPrefix("Error: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("zone %d\t", result.Zone)
	fmt.Printf("belt %s\t", result.Belt)
	fmt.Printf("east %s\t", koord.East)
	fmt.Printf("north %s\t", koord.North)
	fmt.Printf("easting %.0f\t", result.Easting)
	fmt.Printf("northing %.0f\n", result.Northing)
}
