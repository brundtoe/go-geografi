// Read a csv formated file cities.csv and convert from MGRS to UTM
package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"math"
	"strconv"

	"github.com/brundtoe/go-geografi/pkg/mgrs_to_utm"
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
		if record[8] != "Zone" {
			transform(record)
		}
	}

}
func transform(ka []string) {
	koord := mgrs_to_utm.Mgrs{Zone: ka[8], Belt: ka[9], Kmkv: ka[10], East: ka[11], North: ka[12], Town: ka[1]}

	result, err := mgrs_to_utm.UtmAbs(koord)

	log.SetPrefix("Error: ")
	log.SetFlags(0)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("Kilde: ", ka[8], " ", ka[13], " ", ka[14])

	fmt.Printf(" Result: %2d %.0f %.0f %-16s", result.Zone, result.Easting, result.Northing, ka[1])

	easting, err := strconv.ParseFloat(ka[13], 64)
	northing, _ := strconv.ParseFloat(ka[14], 64)

	if math.Abs(easting-result.Easting) >= 1 {
		fmt.Print(" East ")
	}
	if math.Abs(northing-result.Northing) >= 1 {
		fmt.Print(" North ")
	}

	fmt.Println()
}
