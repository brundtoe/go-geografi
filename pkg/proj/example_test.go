package proj

import (
	"fmt"
	"log"
)

/**
 * Examples tests checker fmt.Printf("%s -> %s\n", utm, ll) mod linjen efter // Output
 */

func ExampleUTM_ToLL() {

	utm := UTM{ZoneNumber: 23, ZoneLetter: 'K', Easting: 611733.14, Northing: 7800614.37}
	ll, err := utm.ToLL()
	if err != nil {
		log.Fatalf("error <%v> at utm.ToLL()", err)
	}
	fmt.Printf("%s -> %s\n", utm, ll)
	// Output:
	// 23K 611733.14 7800614.37 -> -19.887495 -43.932663
}

func ExampleLL_ToUTM() {

	ll := LL{Lon: -115.08209766, Lat: 36.23612346}
	utm := ll.ToUTM()
	fmt.Printf("%s -> %s\n", ll, utm)
	// Output:
	// 36.236123 -115.082098 -> 11S 672349.00 4011844.00
}

func ExampleUTM_ToMGRS() {

	utm := UTM{ZoneNumber: 31, ZoneLetter: 'U', Easting: 700373, Northing: 5704554}
	accuracy := 1 // meters
	mgrs := utm.ToMGRS(accuracy)
	fmt.Printf("%s -> %s\n", utm, mgrs)
	// Output:
	// 31U 700373.00 5704554.00 -> 31UGT0037304554
}

func ExampleMGRS_ToUTM() {

	mgrs := MGRS("32ULC989564")
	utm, accuracy, err := mgrs.ToUTM()
	if err != nil {
		log.Fatalf("error <%v> at mgrs.ToUTM()", err)
	}
	fmt.Printf("%s -> %s (accuracy %d meters)\n", mgrs, utm, accuracy)
	// Output:
	// 32ULC989564 -> 32U 398900.00 5756400.00 (accuracy 100 meters)
}

func ExampleLL_ToMGRS() {

	ll := LL{Lon: -88.53, Lat: 51.95}
	accuracy := 10 // meters
	mgrs, err := ll.ToMGRS(accuracy)
	if err != nil {
		log.Fatalf("error <%v> at ll.ToMGRS()", err)
	}
	fmt.Printf("%s -> %s (accuracy %d meters)\n", ll, mgrs, accuracy)
	// Output:
	// 51.950000 -88.530000 -> 16UCC94855658 (accuracy 10 meters)
}

func ExampleMGRS_ToLL() {

	mgrs := MGRS("11SPA7234911844")
	ll, accuracy, err := mgrs.ToLL()
	if err != nil {
		log.Fatalf("error <%v> at mgrs.ToLL()", err)
	}
	fmt.Printf("%s (with accuracy %d meters) -> %s\n", mgrs, accuracy, ll)
	// Output:
	// 11SPA7234911844 (with accuracy 1 meters) -> 36.236123 -115.082098
}
