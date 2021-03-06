// Modulet indeholder rutiner til konvertering mellem lat/lon og UTM, samt fra MGRS til UTM

// Original Javascript by Chuck Taylor
// Port to Go by Jackie Brundtø
//
//
// This is a simple port of the code on the Geographic/UTM Coordinate Converter (1) page from Javascript to Go.
// Using this you can easily pkg between UTM and WGS84 (latitude and longitude).
// Accuracy seems to be around 50cm (I suspect rounding errors are limiting precision).
// This code is provided as-is and has been minimally tested enjoy but use at your own risk!
// The license is the same as the original Javascript:
//
package main
