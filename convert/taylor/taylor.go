// Package taylor

// Original Javascript by Chuck Taylor
// Port to C++ by Alex Hajnal
//
// *** THIS CODE USES 32-BIT FLOATS BY DEFAULT ***
// *** For 64-bit double-precision edit UTM.h: undefine FLOAT_32 and define FLOAT_64
//
// This is a simple port of the code on the Geographic/UTM Coordinate Converter (1) page from Javascript to C++.
// Using this you can easily convert between UTM and WGS84 (latitude and longitude).
// Accuracy seems to be around 50cm (I suspect rounding errors are limiting precision).
// This code is provided as-is and has been minimally tested enjoy but use at your own risk!
// The license for UTM.cpp and UTM.h is the same as the original Javascript:
// "The C++ source code in UTM.cpp and UTM.h may be copied and reused without restriction."
//
// 1) http://home.hiwaay.net/~taylorc/toolbox/geography/geoutm.html

package taylor

import "math"

var smA = 6378137.0
var smB = 6356752.314

var UTMScaleFactor = 0.9996

// DegToRad
// Converts degrees to radians.
func DegToRad(deg float64) float64 {
	return deg / 180.0 * math.Pi
}

// RadToDeg
// Converts radians to degrees.
func RadToDeg(rad float64) float64 {
	return rad / math.Pi * 180.0
}

// ArcLengthOfMeridian
// Computes the ellipsoidal distance from the equator to a point at a
// given latitude.
//
// Reference: Hoffmann-Wellenhof, B., Lichtenegger, H., and Collins, J.,
// GPS: Theory and Practice, 3rd ed.  New York: Springer-Verlag Wien, 1994.
//
// Inputs:
//     phi - Latitude of the point, in radians.
//
// Globals:
//     smA - Ellipsoid model major axis.
//     smB - Ellipsoid model minor axis.
//
// Returns:
//     The ellipsoidal distance of the point from the equator, in meters.
func ArcLengthOfMeridian(phi float64) float64 {

	/* Precalculate n */
	n := (smA - smB) / (smA + smB)

	/* Precalculate alpha */
	alpha := ((smA + smB) / 2.0) * (1.0 + (math.Pow(n, 2.0) / 4.0) + (math.Pow(n, 4.0) / 64.0))

	/* Precalculate beta */
	beta := (-3.0 * n / 2.0) + (9.0 * math.Pow(n, 3.0) / 16.0) + (-3.0 * math.Pow(n, 5.0) / 32.0)

	/* Precalculate gamma */
	gamma := (15.0 * math.Pow(n, 2.0) / 16.0) +
		(-15.0 * math.Pow(n, 4.0) / 32.0)

	/* Precalculate delta */
	delta := (-35.0 * math.Pow(n, 3.0) / 48.0) +
		(105.0 * math.Pow(n, 5.0) / 256.0)

	/* Precalculate epsilon */
	epsilon := 315.0 * math.Pow(n, 4.0) / 512.0

	/* Now calculate the sum of the series and return */
	result := alpha * (phi + (beta * math.Sin(2.0*phi)) +
		(gamma * math.Sin(4.0*phi)) + (delta * math.Sin(6.0*phi)) +
		(epsilon * math.Sin(8.0*phi)))

	return result
}

// UTMCentralMeridian
// Determines the central meridian for the given UTM zone.
//
// Inputs:
//     zone - An integer value designating the UTM zone, range [1,60].
//
// Returns:
//   The central meridian for the given UTM zone, in radians
//   Range of the central meridian is the radian equivalent of [-177,+177].
func UTMCentralMeridian(zone int) float64 {
	cmeridian := DegToRad(-183.0 + (float64(zone) * 6.0))
	return cmeridian
}

// FootpointLatitude
//
// Computes the footpoint latitude for use in converting transverse
// Mercator coordinates to ellipsoidal coordinates.
//
// Reference: Hoffmann-Wellenhof, B., Lichtenegger, H., and Collins, J.,
//   GPS: Theory and Practice, 3rd ed.  New York: Springer-Verlag Wien, 1994.
//
// Inputs:
//   y - The UTM northing coordinate, in meters.
//
// Returns:
//   The footpoint latitude, in radians.
func FootpointLatitude(y float64) float64 {

	/* Precalculate n (Eq. 10.18) */
	n := (smA - smB) / (smA + smB)

	/* Precalculate alpha_ (Eq. 10.22) */
	/* (Same as alpha in Eq. 10.17) */
	alpha_ := ((smA + smB) / 2.0) *
		(1 + (math.Pow(n, 2.0) / 4) + (math.Pow(n, 4.0) / 64))

	/* Precalculate y_ (Eq. 10.23) */
	y_ := y / alpha_

	/* Precalculate beta_ (Eq. 10.22) */
	beta_ := (3.0 * n / 2.0) + (-27.0 * math.Pow(n, 3.0) / 32.0) +
		(269.0 * math.Pow(n, 5.0) / 512.0)

	/* Precalculate gamma_ (Eq. 10.22) */
	gamma_ := (21.0 * math.Pow(n, 2.0) / 16.0) + (-55.0 * math.Pow(n, 4.0) / 32.0)

	/* Precalculate delta_ (Eq. 10.22) */
	delta_ := (151.0 * math.Pow(n, 3.0) / 96.0) + (-417.0 * math.Pow(n, 5.0) / 128.0)

	/* Precalculate epsilon_ (Eq. 10.22) */
	epsilon_ := 1097.0 * math.Pow(n, 4.0) / 512.0

	/* Now calculate the sum of the series (Eq. 10.21) */
	result := y_ + (beta_ * math.Sin(2.0*y_)) +
		(gamma_ * math.Sin(4.0*y_)) + (delta_ * math.Sin(6.0*y_)) +
		(epsilon_ * math.Sin(8.0*y_))

	return result
}

// MapLatLonToXY
// Converts a latitude/longitude pair to x and y coordinates in the
// Transverse Mercator projection.  Note that Transverse Mercator is not
// the same as UTM a scale factor is required to convert between them.
//
// Reference: Hoffmann-Wellenhof, B., Lichtenegger, H., and Collins, J.,
// GPS: Theory and Practice, 3rd ed.  New York: Springer-Verlag Wien, 1994.
//
// Inputs:
//    phi - Latitude of the point, in radians.
//    lambda - Longitude of the point, in radians.
//    lambda0 - Longitude of the central meridian to be used, in radians.
//
// Outputs:
//    x - The x coordinate of the computed point.
//    y - The y coordinate of the computed point.
//
// Returns:
//    The function does not return a value.
func MapLatLonToXY(phi float64, lambda float64, lambda0 float64) (float64, float64) {
	/* Precalculate ep2 */
	ep2 := (math.Pow(smA, 2.0) - math.Pow(smB, 2.0)) / math.Pow(smB, 2.0)

	/* Precalculate nu2 */
	nu2 := ep2 * math.Pow(math.Cos(phi), 2.0)

	/* Precalculate n */
	n := math.Pow(smA, 2.0) / (smB * math.Sqrt(1.0+nu2))

	/* Precalculate t */
	t := math.Tan(phi)
	t2 := t * t
	//let tmp = (t2 * t2 * t2) - math.Pow(t, 6.0) // Unused

	/* Precalculate l */
	l := lambda - lambda0

	/* Precalculate coefficients for l**n in the equations below
	   so a normal human being can read the expressions for easting
	   and northing
	   -- l**1 and l**2 have coefficients of 1.0 */

	l3coef := 1.0 - t2 + nu2

	l4coef := 5.0 - t2 + 9.0*nu2 + 4.0*(nu2*nu2)

	l5coef := 5.0 - 18.0*t2 + (t2 * t2) + 14.0*nu2 - 58.0*t2*nu2

	l6coef := 61.0 - 58.0*t2 + (t2 * t2) + 270.0*nu2 - 330.0*t2*nu2

	l7coef := 61.0 - 479.0*t2 + 179.0*(t2*t2) - (t2 * t2 * t2)

	l8coef := 1385.0 - 3111.0*t2 + 543.0*(t2*t2) - (t2 * t2 * t2)

	/* Calculate easting (x) */
	x := n*math.Cos(phi)*l + (n / 6.0 * math.Pow(math.Cos(phi), 3.0) * l3coef * math.Pow(l, 3.0)) +
		(n / 120.0 * math.Pow(math.Cos(phi), 5.0) * l5coef * math.Pow(l, 5.0)) +
		(n / 5040.0 * math.Pow(math.Cos(phi), 7.0) * l7coef * math.Pow(l, 7.0))

	/* Calculate northing (y) */
	y := ArcLengthOfMeridian(phi) + (t / 2.0 * n * math.Pow(math.Cos(phi), 2.0) * math.Pow(l, 2.0)) +
		(t / 24.0 * n * math.Pow(math.Cos(phi), 4.0) * l4coef * math.Pow(l, 4.0)) +
		(t / 720.0 * n * math.Pow(math.Cos(phi), 6.0) * l6coef * math.Pow(l, 6.0)) +
		(t / 40320.0 * n * math.Pow(math.Cos(phi), 8.0) * l8coef * math.Pow(l, 8.0))

	return x, y
}

// MapXYToLatLon
// Converts x and y coordinates in the Transverse Mercator projection to
// a latitude/longitude pair.  Note that Transverse Mercator is not
// the same as UTM a scale factor is required to convert between them.
//
// Reference: Hoffmann-Wellenhof, B., Lichtenegger, H., and Collins, J.,
//   GPS: Theory and Practice, 3rd ed.  New York: Springer-Verlag Wien, 1994.
//
// Inputs:
//   x - The easting of the point, in meters.
//   y - The northing of the point, in meters.
//   lambda0 - Longitude of the central meridian to be used, in radians.
//
// Outputs:
//   phi    - Latitude in radians.
//   lambda - Longitude in radians.
//
// Returns:
//   The function does not return a value.
//
// Remarks:
//   The local variables Nf, nuf2, tf, and tf2 serve the same purpose as
//   N, nu2, t, and t2 in MapLatLonToXY, but they are computed with respect
//   to the footpoint latitude phif.
//
//   x1frac, x2frac, x2poly, x3poly, etc. are to enhance readability and
//   to optimize computations.
func MapXYToLatLon(x float64, y float64, lambda0 float64) (float64, float64) {

	/* Get the value of phif, the footpoint latitude. */
	phif := FootpointLatitude(y)

	/* Precalculate ep2 */
	ep2 := (math.Pow(smA, 2.0) - math.Pow(smB, 2.0)) / math.Pow(smB, 2.0)

	/* Precalculate cos (phif) */
	cf := math.Cos(phif)

	/* Precalculate nuf2 */
	nuf2 := ep2 * math.Pow(cf, 2.0)

	/* Precalculate Nf and initialize Nfpow */
	Nf := math.Pow(smA, 2.0) / (smB * math.Sqrt(1+nuf2))
	Nfpow := Nf

	/* Precalculate tf */
	tf := math.Tan(phif)
	tf2 := tf * tf
	tf4 := tf2 * tf2

	/* Precalculate fractional coefficients for x**n in the equations
	   below to simplify the expressions for latitude and longitude. */
	x1frac := 1.0 / (Nfpow * cf)

	Nfpow *= Nf /* now equals Nf**2) */
	x2frac := tf / (2.0 * Nfpow)

	Nfpow *= Nf /* now equals Nf**3) */
	x3frac := 1.0 / (6.0 * Nfpow * cf)

	Nfpow *= Nf /* now equals Nf**4) */
	x4frac := tf / (24.0 * Nfpow)

	Nfpow *= Nf /* now equals Nf**5) */
	x5frac := 1.0 / (120.0 * Nfpow * cf)

	Nfpow *= Nf /* now equals Nf**6) */
	x6frac := tf / (720.0 * Nfpow)

	Nfpow *= Nf /* now equals Nf**7) */
	x7frac := 1.0 / (5040.0 * Nfpow * cf)

	Nfpow *= Nf /* now equals Nf**8) */
	x8frac := tf / (40320.0 * Nfpow)

	/* Precalculate polynomial coefficients for x**n.
	   -- x**1 does not have a polynomial coefficient. */
	x2poly := -1.0 - nuf2

	x3poly := -1.0 - 2*tf2 - nuf2

	x4poly := 5.0 + 3.0*tf2 + 6.0*nuf2 - 6.0*tf2*nuf2 - 3.0*(nuf2*nuf2) - 9.0*tf2*(nuf2*nuf2)

	x5poly := 5.0 + 28.0*tf2 + 24.0*tf4 + 6.0*nuf2 + 8.0*tf2*nuf2

	x6poly := -61.0 - 90.0*tf2 - 45.0*tf4 - 107.0*nuf2 + 162.0*tf2*nuf2

	x7poly := -61.0 - 662.0*tf2 - 1320.0*tf4 - 720.0*(tf4*tf2)

	x8poly := 1385.0 + 3633.0*tf2 + 4095.0*tf4 + 1575*(tf4*tf2)

	/* Calculate latitude */
	phi := phif + x2frac*x2poly*(x*x) + x4frac*x4poly*math.Pow(x, 4.0) +
		x6frac*x6poly*math.Pow(x, 6.0) + x8frac*x8poly*math.Pow(x, 8.0)

	/* Calculate longitude */
	lambda := lambda0 + x1frac*x + x3frac*x3poly*math.Pow(x, 3.0) +
		x5frac*x5poly*math.Pow(x, 5.0) + x7frac*x7poly*math.Pow(x, 7.0)

	return phi, lambda
}

// LatLonToUTMXY
// Converts a latitude/longitude pair to x and y coordinates in the
// Universal Transverse Mercator projection.
//
// Inputs:
//   lat - Latitude of the point, in degrees.
//   lon - Longitude of the point, in degrees.
//   zone - UTM zone to be used for calculating values for x and y.
//          If zone is less than 1 or greater than 60, the routine
//          will determine the appropriate zone from the value of lon.
//
// Outputs:
//   x - The x coordinate (easting) of the computed point. (in meters)
//   y - The y coordinate (northing) of the computed point. (in meters)
//
// Returns:
//   The UTM zone used for calculating the values of x and y.
func LatLonToUTMXY(lat float64, lon float64, zone int) (float64, float64) {

	if (zone < 1) || (zone > 60) {
		zone = int((lon+180.0)/6) + 1
	}
	x, y := MapLatLonToXY(DegToRad(lat), DegToRad(lon), UTMCentralMeridian(zone))

	/* Adjust easting and northing for UTM system. */
	x = x*UTMScaleFactor + 500000.0
	y = y * UTMScaleFactor
	if y < 0.0 {
		y = y + 10000000.0
	}
	return x, y
}

// UTMXYToLatLon
//
// Converts x and y coordinates in the Universal Transverse Mercator
// projection to a latitude/longitude pair.
//
// Inputs:
// x - The easting of the point, in meters.
// y - The northing of the point, in meters.
// zone - The UTM zone in which the point lies.
// southhemi - True if the point is in the southern hemisphere
//               false otherwise.
//
// Outputs:
// lat - The latitude of the point, in radians.
// lon - The longitude of the point, in radians.
//
// Returns:
// The function does not return a value.
func UTMXYToLatLon(x float64, y float64, zone int, southHemi bool) (float64, float64) {

	x -= 500000.0
	x /= UTMScaleFactor

	/* If in southern hemisphere, adjust y accordingly. */
	if southHemi {
		y -= 10000000.0
	}
	y /= UTMScaleFactor

	cmeridian := UTMCentralMeridian(zone)
	lat, lon := MapXYToLatLon(x, y, cmeridian)

	return lat, lon
}
