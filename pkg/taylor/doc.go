// Package taylor converts between Geographic/UTM Coordinate
/*
 Port to Go by Jackie Brundtø.

This is a simple port of the code on the Geographic/UTM Coordinate Converter (1) page from JavaScript to Go.

Using this package, you can transform between UTM and WGS84 (latitude and longitude).

Accuracy seems to be around 50 cm (I suspect rounding errors are limiting precision).

This code is provided as-is and has been minimally tested enjoy but use at your own risk!

The license is the same as the original JavaScript:

1) http://home.hiwaay.net/~taylorc/toolbox/geography/geoutm.html Not available anymore

*/
package taylor
