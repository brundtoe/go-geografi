# Geografi

Projektet indeholder tre packages

- mgrs, er en justeret kopi af https://github.com/klaus-tockloth/coco, som er en delvis portering til golang af https://github.com/proj4js/mgrs
- taylor, der er en Go implementering af Chuck taylors oprindelige WGS84 konvertering mellem lat/lon og UTM
- mgrs_to_utm, der er en omregning fra Military Grid System (MGRS) til UTM. Er porterete fra python/geografi

## Installation

Clone projektet https://github.com/brundtoe/go-geografi.git

## test

package taylor
```shell
    cd pkg/taylor
    go test -v 
```
Package mgrs_to_utm

```shell
    cd pkg/mgrs_to_utm
    go test -v 
```

Package mgrs
```shell
```cd pkg/mgrs
  go test -v

## Eksempler

Eksempler findes i mappen cmd.

```bash
    go run <absoulte path to project root>/cmd/cities

```

Filen cities.go

    Demo af indlæsning af filen cities.csv

Filen mgrs_to_utm.go anvender ED50 modellen til konverteringen
    
    Transformation af MGRS 32V NJ 948757 99059 til UTM 32V 5948757 6399059 

Filen utmgeo.go anvender pkg/taylor

    Transfomration fra LatLonToUTMXX og tilbage med UTMXYToLatLon med visning af evt difference

Filen to_mgrs.go anvender pkg/geoutm

    Transformation af lat lon til MGRS opg USNG med visning af evt difference

Filen to_wgs84.go anvender pkg/geoutm

    Transformation af UTM til Wgs84 med visning af evt difference

Filen milgrid_to_wgs84.go anvender pkg/geoutm

    Transformation af MGRS og USNG til Wgs84 med visning af difference

Der er forskelle på konverteringen fra UTM -> wgs84 og den samme fra MGSR hhv UNSG fordi 
UTM earsting og northing er float med to decimaler og MGSR og USNG er float uden decimaler


## Dokumentation

Start dokumentationsserveren 
```bash
    godoc -http=:8000
```
Dokumentationen findes under Third party


