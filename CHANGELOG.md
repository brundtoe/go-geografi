# Changelog

## 29. december 2025

Ændringer:
- refaktoreret cmds - der udskrives kun en linje i output når afvigelser er større end det acceptable niveau
- Refaktoreret city.go - pseudo enum med prefix csv for at undgå shadowing i øvrige types

## Tag v2.0.0 - 29. december 2025

Ændringer
- Refaktoreret utmToMGRS og utmToUSNG ekstraheret fælles logik fra funktioktionerne til utm.buildGrid
- Refaktoreret utmgeo opdeling i en fil pr. type
- Refaktoreret package mgrs til proj

## Tag v1.0.1 - 28. december 2025

Breaking
- Class cities build City opretter to nye felter MGRS og USNG ved at kalde Utm.toMGRS og UTM.toUSNG
- Class Cities fjernet cityToMGRS og cityToUSNG

Ændringer
- geoutm.go Tilføjet en ny funktion LL toUSNG 
- ny cmd wgs84_to_utm.go 
- City.go tilføjet enum til indlæsning af data fra csv

## Tag v1.0.0 - 27. december 2025

Breaking
- fjernet den gamle mgrs utm fra Landinspektøren - er overflødig

## 25. december 2025

Ændringer:
    - go opdateret til fra 1.23 til 1.25.5
    - forenklet city.go funktionerne CityToMGRS og CityToUSNG
    - renamed cmd folder examples to show the transformation in action
    - revideret dokumentationen

## 18. december 2025

Ændringer
- package utmabs omdøbt til mgrs_to_utm som beskriver transformationen

## 16. december 2025

Ændringer:
- utmabs.go test for afvigelser i beregning af easting og northing
- to_mgrs.go test for afvigelser i beregning af east og north i USNG og implicit i MGRS

Tilføjelser:
- to_wgs84.go transformerer utm til wgs84
- milgrid_to-qgs84.go transformerer mgrs og usng til wgs84

Fejlrettelser:
- utm.toMGRS og utm.toUSNG rettet fejl i trunkering af resultatet 

## 15. december 2025

Oprettet changelog

Ændringer
- easting og northing er i cities.csv ændret til float med to decimaler
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
 
