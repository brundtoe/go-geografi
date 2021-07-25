# Geografi

Projektet indeholder to moduler

- taylor, der er en Go implementering af Chuck taylors oprindelige WGS84 konvertering mellem lat/lon og UTM
- utmabs, der er en omregning fra Military Grid System (MGRS) til UTM.

## Installation

Clone projektet https://github.com/brundtoe/go-geografi.git

## test

Modulet taylor
```bash
    cd pkg/taylor
    go test -v 
```
Modulet utmabs

```bash
    cd pkg/utmabs
    go test -v 
```

## Eksempler

Eksempler findes i mappen cmd.

```bash
    go run <absoulte path to project root>/cmd/cities

```

## Dokumentation

Start dokumentationsserveren 
```bash
    godoc -http=:8000
```
Dokumentationen findes under Third party


