// Read a csv formated file cities.csv and convert from MGRS to UTM
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"

	"github.com/brundtoe/go-geografi/pkg/utils"
	"github.com/brundtoe/go-geografi/pkg/utmabs"
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
		fmt.Println("Kilde: ", record[8], record[9], record[10], record[11], record[12], record[1])
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

	fmt.Printf("Result: %2d ", result.Zone)
	fmt.Printf("%.0f ", result.Easting)
	fmt.Printf("%.0f\n", result.Northing)
	fmt.Println("----------------")
}
