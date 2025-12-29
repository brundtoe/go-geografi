//Package mgrs
/*
Purpose:
- MGRS/UTMREF <-> UTM <-> Lon Lat

Description:
- Package for converting coordinates between WGS84 Lon Lat, UTM and MGRS/UTMREF.

Releases:
- v0.1.0 - 2019/05/09 : initial release
- v0.2.0 - 2019/05/10 : coord formatting changed

Author:
- Klaus Tockloth

Copyright and license:
- Copyright (c) 2019 Klaus Tockloth
- MIT license

Permission is hereby granted, free of charge, to any person obtaining a copy of this software
and associated documentation files (the Software), to deal in the Software without restriction,
including without limitation the rights to use, copy, modify, merge, publish, distribute,
sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or
substantial portions of the Software.

The software is provided 'as is', without warranty of any kind, express or implied, including
but not limited to the warranties of merchantability, fitness for a particular purpose and
noninfringement. In no event shall the authors or copyright holders be liable for any claim,
damages or other liability, whether in an action of contract, tort or otherwise, arising from,
out of or in connection with the software or the use or other dealings in the software.

Contact (eMail):
- freizeitkarte@googlemail.com

Remarks:
- This library is a partial port from "github.com/proj4js/mgrs" (JavaScript).
- Possible coordinate conversions:
  UTM -> Lon Lat
  UTM -> MGRS
  Lon Lat -> UTM
  Lon Lat -> MGRS
  MGRS -> UTM
  MGRS -> Lon Lat
- Build library:
  go install
- Test library:
  go test
  go test -cover
  go test -coverprofile=c.out + go tool cover -html=c.out
- Document library:
  godoc
  view document in browser (http://localhost:6060)

Links:
- https://github.com/proj4js/mgrs
- https://gist.github.com/tmcw/285949

*/
/*
Package mgrs (coordinate conversion) provides methods for converting coordinates between WGS84 Lon Lat, UTM and MGRS/UTMREF.

Supported conversions:

utm.ToLL()   : converts from UTM to LL
utm.ToMGRS() : converts from UTM to MGRS
utm.ToUSNG   : converts from UTM to USNG
ll.ToUTM()   : converts from LL to UTM
ll.ToMGRS()  : converts from LL to MGRS
mgrs.ToUTM() : converts from MGRS to UTM
mgrs.ToLL()  : converts from MGRS to LL
usng.ToLL	 : converts from USNG to LL
usng.ToMGRS	 : converts from USNG to MGRS
usng.toUTM   : converts from USNG to UTM

Data objects:

UTM  : ZoneNumber ZoneLetter Easting Northing
LL   : Longitude Latitude
MGRS : String
USNG : string

Abbreviations:

Lon    : Longitude
Lat    : Latitude
MGRS   : Military Grid Reference System (same as UTMREF)
USNG   : United Tastes National Grid samt as MGRS formated with spaces
UTM    : Universal Transverse Mercator
UTMREF : UTM Reference System (same as MGRS)
WGS84  : World Geodetic System 1984 (same as EPSG:4326)
*/
package proj
