// Read csv file cities.csv, convert from lat/lon to UTM and back again
package main

import (
	"example.com/geografi/pkg/taylor"
	"fmt"
	"math"
)

func main() {

	const diff = 0.0000000001
	latitude := 56.095833
	longitude := 10.136111

	east, north := taylor.LatLonToUTMXY(latitude, longitude, 32)

	fmt.Printf("East\t %4f\n", east)
	fmt.Printf("North\t %4f\n", north)

	lat, lon := taylor.UTMXYToLatLon(east, north, 32, false)

	fmt.Printf("Lat\t %4f\n", taylor.RadToDeg(lat))
	fmt.Printf("Lon\t %4f\n", taylor.RadToDeg(lon))

	if math.Abs(latitude-taylor.RadToDeg(lat)) > diff {
		fmt.Printf("Konvertering af latitude overskrider acceptabel tolerance")
	}
	if math.Abs(longitude-taylor.RadToDeg(lon)) > diff {
		fmt.Printf("Konvertering af longitude overskrider acceptabel tolerance")
	}

}
