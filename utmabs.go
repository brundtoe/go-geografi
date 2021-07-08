package main

import (
	"example.com/convert"
	"example.com/utmabs/morestrings"
	"fmt"
	"log"
	"strings"
)

/*
Som rust udgaven hånderes kmkv angive med 1 m nøjagtighed dvs 5 cifre
*/

func main() {
	fmt.Println("Hello Jackie")

	//koord := Mgrs{"32", "U", "NG", "08600", "77000", "Somewhere"}

	koordCsv := "32;U;NG;08600;77000;Somewhere"

	ka := strings.Split(koordCsv, ";")

	koord := convert.Mgrs{ka[0], ka[1], ka[2], ka[3], ka[4], ka[5]}

	result, err := convert.UtmAbs(koord)

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

	res := morestrings.ReverseRunes("jensen")
	fmt.Println("Jensen ", res)
	fmt.Println("Message", morestrings.Hello())
}
