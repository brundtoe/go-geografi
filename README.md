# Geografi

Projektet indeholder tre moduler

- geoutm, er en justeret kopi af https://github.com/klaus-tockloth/coco, som er en delvis portering til golang af https://github.com/proj4js/mgrs
- taylor, der er en Go implementering af Chuck taylors oprindelige WGS84 konvertering mellem lat/lon og UTM
- utmabs, der er en omregning fra Military Grid System (MGRS) til UTM.

## Installation

Clone projektet https://github.com/brundtoe/go-geografi.git

## test

Modulet taylor
```shell
    cd pkg/taylor
    go test -v 
```
Modulet utmabs

```shell
    cd pkg/utmabs
    go test -v 
```

Modulet geoutm
```shell
```cd pkg/utmgeo
  go test -v

## Eksempler

Eksempler findes i mappen cmd.

```bash
    go run <absoulte path to project root>/cmd/cities

```

Filen cities.go

    Demo af indlæsning af filen cities.csv

Filen utmabs.go anvender ED50 modellen til konverteringen
    
    Konvertering af MGRS 32V NJ 948757 99059 til UTM 32V 5948757 6399059 

Filen utmgeo.go anvender pkg/taylor

    Konvertering fra LatLonToUTMXX og tilbage med UTMXYToLatLon med visning af evt difference

Filen to_mgrs.go anvender pkg/geoutm

    Konvertering af lat lon til MGRS opg USNG


## Dokumentation

Start dokumentationsserveren 
```bash
    godoc -http=:8000
```
Dokumentationen findes under Third party


