// Package main - a number of commands using geographic transformations
/*
Each command reads a csv file with city names and coordinates,
converts each line to a City object.

The City object [pkg/github.com/brundtoe/go-geografi/pkg/proj.City] contains the original coordinates,
which are used as inout for the transformation and the expected results.

For each transformation, the city object values are displayed if the transformation exceeds the allowed error margin.

Transformations
	- wgs84_to_utm: transform WGS84 coordinates to UTM coordinates
	- utm_to_wgs84: transform UTM coordinates to WGS84 coordinates
	- mgrs_usng_to_wgs84: transform MGRS and USNG coordinates to WGS84 coordinates
	- wgs84_to_mgrs: transform WGS84 coordinates to MGRS coordinates

*/
package main
