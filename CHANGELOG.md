# Changelog

## 15. december 2025

Oprettet changelog

Ændringer
- easting og northing er i citites.csv ændret til float med to decimaler
- testdata justeret i forhold hertil

Test af eksempler
- ref https://go.dev/blog/examples
- Eksample tests er en del af fokumentationen når der er anført i selve sourcekoden
- Det ønskede testresultat er anført nederst i funktionen på linien efter // Output:
- go sammenligner dette med fmt.Printf("%s -> %s\n", ll, utm)


```
func ExampleLL_ToUTM() {

	ll := LL{Lon: -115.08209766, Lat: 36.23612346}
	utm := ll.ToUTM()
	fmt.Printf("%s -> %s\n", ll, utm)
	// Output:
	// 36.236123 -115.082098 -> 11S 672349.00 4011844.00
}
```
 
