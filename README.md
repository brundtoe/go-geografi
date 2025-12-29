# Geografi

Projektet indeholder to packages

- proj, er en justeret kopi af https://github.com/klaus-tockloth/coco, som er en delvis portering til golang af https://github.com/proj4js/mgrs
- taylor, der er en Go implementering af Chuck taylors oprindelige WGS84 konvertering mellem lat/lon og UTM

Projetet er en refaktoreret udgave af https://github.com/klaus-tockloth/coco
- Den oprindelige fil `coco.go` er opdelt i en fil pr. type som er golang best practice og langt mere overskuelig

## Installation

Clone projektet https://github.com/brundtoe/go-geografi.git

## test

package taylor
```shell
    cd pkg/taylor
    go test -v 
```

Package proj
```shell
```cd pkg/proj
  go test -v

## Eksempler

Eksempler findes i mappen cmd.

```bash
    go run <absoulte path to project root>/cmd/cities

```

Filen cmd/cities/cities.go

    Demo af indlæsning af filen cities.csv


Filen cmd/taylor/utmgeo.go anvender pkg/taylor

    Transfomration fra LatLonToUTMXX og tilbage med UTMXYToLatLon med visning af evt difference

Filen cmd/proj/wgs84_to_utm.go anvender pkg/mgrs

    Transformation af lat lon til UTM med visning af evt difference

Filen cmd/proj/utm_to_wgs84.go anvender pkg/mgrs

    Transformation af UTM til Wgs84 med visning af evt difference

Filen cmd/proj/wgs84_to_mgrs.go anvender pkg/mgrs

    Transformation af lat lon til MGRS opg USNG med visning af evt difference

Filen cmd/proj/mgrs_usng_to_wgs84.go anvender pkg/mgrs

    Transformation af MGRS og USNG til Wgs84 med visning af difference

Der er forskelle på konverteringen fra UTM -> wgs84 og den samme fra MGRS hhv USNG fordi 
UTM easting og northing er float med to decimaler og MGSR og USNG er integeres der anvendes som float uden decimaler

## Dokumentation

Intallation af godoc
```bash
  go install golang.org/x/tools/cmd/godoc@latest
```

Start dokumentationsserveren 
```bash
    godoc -http=:8000
```
Browser til http://localhost:8000

Dokumentationen findes i menuen under Third party


