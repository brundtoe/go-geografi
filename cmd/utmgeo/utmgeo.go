package main

import (
	"example.com/geografi/convert"
	"fmt"
	"math"
)

func main() {

	const diff = 0.0000000001
	latitude := 56.095833
	longitude := 10.136111

	east, north := convert.LatLonToUTMXY(latitude, longitude, 32)

	fmt.Printf("East\t %4f\n", east)
	fmt.Printf("North\t %4f\n", north)

	lat, lon := convert.UTMXYToLatLon(east, north, 32, false)

	fmt.Printf("Lat\t %4f\n", convert.RadToDeg(lat))
	fmt.Printf("Lon\t %4f\n", convert.RadToDeg(lon))

	if math.Abs(latitude-convert.RadToDeg(lat)) > diff {
		fmt.Printf("Konvertering af latitude overskrider acceptabel tolerance")
	}
	if math.Abs(longitude-convert.RadToDeg(lon)) > diff {
		fmt.Printf("Konvertering af longitude overskrider acceptabel tolerance")
	}

}
