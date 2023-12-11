// Read csv file cities.csv, convert from lat/lon to UTM and back again
package main

import (
	"fmt"
	"github.com/brundtoe/go-geografi/pkg/coco"
	"github.com/brundtoe/go-geografi/pkg/taylor"
	"math"
	"strconv"
)

func main() {

	const diff = 0.0000000001
	latitude := 57.723661
	longitude := 10.592629

	east, north := taylor.LatLonToUTMXY(latitude, longitude, 32)

	fmt.Printf("East\t %4f\n", east)
	fmt.Printf("North\t %4f\n", north)

	location := coco.LL{
		Lat: latitude,
		Lon: longitude,
	}
	utm := location.ToUTM()

	fmt.Println("UTM..", utm)
	fmt.Println("MGRS..", utm.ToMGRS(1))
	fmt.Println("USNG..", utm.ToUSNG(1))

	lat, lon := taylor.UTMXYToLatLon(east, north, 32, false)

	fmt.Printf("Lat\t %4f\n", taylor.RadToDeg(lat))
	fmt.Printf("Lon\t %4f\n", taylor.RadToDeg(lon))

	if math.Abs(latitude-taylor.RadToDeg(lat)) > diff {
		fmt.Printf("Konvertering af latitude overskrider acceptabel tolerance")
	}
	if math.Abs(longitude-taylor.RadToDeg(lon)) > diff {
		fmt.Printf("Konvertering af longitude overskrider acceptabel tolerance")
	}

	milgrid := coco.MGRS("32VNJ9485799059")

	utmsted, _, _ := milgrid.ToUTM()
	println("UTM", utmsted.String())
	mgrs := coco.USNG("32V NJ 94857 99059").ToMGRS()
	usngUtm, _, _ := coco.USNG("32V NJ 94857 99059").ToUTM()
	println("USNG UTM..", usngUtm.String())
	llSted, accuracy, _ := mgrs.ToLL()
	latt := fmt.Sprintf("%f", llSted.Lat)
	long := fmt.Sprintf("%f", llSted.Lon)
	println(latt, long, strconv.Itoa(accuracy))
}
