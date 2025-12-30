package proj

import (
	"fmt"
	"log"
)

/**
 * Examples tests checker fmt.Printf("%s -> %s\n", utm, ll) mod linjen efter // Output
 */

func ExampleUTM_ToLL() {

	utm := UTM{ZoneNumber: 32, ZoneLetter: 'U', Easting: 489092.85, Northing: 6101296.46}
	ll, err := utm.ToLL()
	if err != nil {
		log.Fatalf("error <%v> at utm.ToLL()", err)
	}
	fmt.Printf("Bredebro: %s -> %s\n", utm, ll)
	// Output:
	// Bredebro: 32U 489092.85 6101296.46 -> 55.058337 8.829244
}

func ExampleUTM_ToMGRS() {

	utm := UTM{ZoneNumber: 32, ZoneLetter: 'V', Easting: 594857.92, Northing: 6399059.92}
	accuracy := 1 // meters
	mgrs := utm.ToMGRS(accuracy)
	fmt.Printf("Skagen: %s -> %s\n", utm, mgrs)
	// Output:
	// Skagen: 32V 594857.92 6399059.92 -> 32VNJ9485799059
}

func ExampleUTM_ToUSNG() {

	utm := UTM{ZoneNumber: 32, ZoneLetter: 'V', Easting: 594857.92, Northing: 6399059.92}
	accuracy := 1 // meters
	mgrs := utm.ToUSNG(accuracy)
	fmt.Printf("Skagen: %s -> %s\n", utm, mgrs)
	// Output:
	// Skagen: 32V 594857.92 6399059.92 -> 32V NJ 94857 99059
}

func ExampleLL_ToUTM() {

	ll := LL{Lon: 9.639181, Lat: 56.982376}
	utm := ll.ToUTM()
	fmt.Printf("Nibe: %s -> %s\n", ll, utm)
	// Output:
	// Nibe: 56.982376 9.639181 -> 32V 538846.91 6315605.80
}

func ExampleLL_ToMGRS() {

	ll := LL{Lon: 10.236330, Lat: 55.098237}
	accuracy := 10 // meters
	mgrs, err := ll.ToMGRS(accuracy)
	if err != nil {
		log.Fatalf("error <%v> at ll.ToMGRS()", err)
	}
	fmt.Printf("Svendborg: %s -> %s (accuracy %d meters)\n", ll, mgrs, accuracy)
	// Output:
	// Svendborg: 55.098237 10.236330 -> 32UNG78890642 (accuracy 10 meters)
}

func ExampleLL_ToUSNG() {

	ll := LL{Lon: 10.236330, Lat: 55.098237}
	accuracy := 10 // meters
	mgrs, err := ll.ToUSNG(accuracy)
	if err != nil {
		log.Fatalf("error <%v> at ll.ToMGRS()", err)
	}
	fmt.Printf("Svendborg: %s -> %s (accuracy %d meters)\n", ll, mgrs, accuracy)
	// Output:
	// Svendborg: 55.098237 10.236330 -> 32U NG 7889 0642 (accuracy 10 meters)
}

func ExampleMGRS_ToUTM() {

	mgrs := MGRS("33UVB98231797")
	utm, accuracy, err := mgrs.ToUTM()
	if err != nil {
		log.Fatalf("error <%v> at mgrs.ToUTM()", err)
	}
	fmt.Printf("Gudhjem: %s -> %s (accuracy %d meters)\n", mgrs, utm, accuracy)
	// Output:
	// Gudhjem: 33UVB98231797 -> 33U 498230.00 6117970.00 (accuracy 10 meters)
}

func ExampleMGRS_ToLL() {

	mgrs := MGRS("33UUB162700")
	ll, accuracy, err := mgrs.ToLL()
	if err != nil {
		log.Fatalf("error <%v> at mgrs.ToLL()", err)
	}
	fmt.Printf("Roskilde: %s (with accuracy %d meters) -> %s\n", mgrs, accuracy, ll)
	// Output:
	// Roskilde: 33UUB162700 (with accuracy 100 meters) -> 55.641059 12.079514
}

func ExampleMGRS_ToUSNG() {

	mgrs := MGRS("33UVB98231797")
	usng := mgrs.ToUSNG()
	fmt.Printf("Gudhjem: %s -> %s\n", mgrs, usng)
	// Output:
	// Gudhjem: 33UVB98231797 -> 33U VB 9823 1797
}

func ExampleUSNG_ToMGRS() {
	usng := USNG("33U VB 98232 17975")
	mgrs := usng.ToMGRS()
	fmt.Printf("Gudhjem: %s -> %s\n", usng, mgrs)
	// Output:
	// Gudhjem: 33U VB 98232 17975 -> 33UVB9823217975
}

func ExampleUSNG_ToUTM() {
	usng := USNG("32V MJ 81303 12511")
	utm, accuracy, err := usng.ToUTM()
	if err != nil {
		log.Fatalf("error <%v> at usng.ToUTM()", err)
	}
	fmt.Printf("Thisted: %s -> %s (accuracy %d meters)\n", usng, utm, accuracy)
	// Output:
	// Thisted: 32V MJ 81303 12511 -> 32V 481303.00 6312511.00 (accuracy 1 meters)
}

func ExampleUSNG_ToLL() {
	usng := USNG("32V MJ 81303 12511")
	ll, accuracy, err := usng.ToLL()
	if err != nil {
		log.Fatalf("error <%v> at usng.ToLL()", err)
	}
	fmt.Printf("Thisted: %s (with accuracy %d meters) -> %s\n", usng, accuracy, ll)
	// Output:
	// Thisted: 32V MJ 81303 12511 (with accuracy 1 meters) -> 56.955828 8.692583
}
