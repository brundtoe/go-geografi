package utmabs

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Mgrs struct {
	Zone  string
	Belt  string
	Kmkv  string
	East  string
	North string
	Town  string
}

type Utm struct {
	Zone     int
	Belt     string
	Easting  float64
	Northing float64
}

func UtmAbs(koord Mgrs) (Utm, error) {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recover from ", r)
		}
	}()

	var result Utm
	const alfastr = "ABCDEFGHJKLMNPQRSTUVWXYZ"

	var midt = [10]float64{0442.1322, 1326.5688, 2211.5090, 3097.2460,
		3984.0123, 4871.9620, 5761.1579, 6651.5667, 7543.0601, 8435.4253}

	zone, err := strconv.Atoi(koord.Zone)
	if err != nil {
		convErr := fmt.Sprintf("Konverteringsfejl zone = %s - kan ikke konverteres til heltal", koord.Zone)
		return result, errors.New(convErr)
	}

	if zone < 1 || zone > 60 {
		convErr := fmt.Sprintf("Konverteringsfejl zone = %d - skal være mellem 1 og 60", zone)
		return result, errors.New(convErr)
	}

	belt := koord.Belt
	if belt < "N" || belt > "X" || belt == "O" {
		convErr := fmt.Sprintf("Belt = %s - skal være mellem N og X og må ikke være O", belt)
		return result, errors.New(convErr)
	}
	vbelt := strings.Index(alfastr, koord.Belt) + 1

	opdelt := strings.Split(koord.Kmkv, "")
	if len(opdelt) != 2 {
		convErr := fmt.Sprintf("Kmkv %s - skal indeholde to bogstaver", koord.Kmkv)
		return result, errors.New(convErr)
	}

	first := opdelt[0]
	second := opdelt[1]
	if second > "V" {
		convErr := fmt.Sprintf("Andet bogstav %s i kmkv skal være mindre end V", second)
		return result, errors.New(convErr)
	}

	utmg1, err := strconv.ParseFloat(koord.East, 64)
	if err != nil {
		convErr := fmt.Sprintf("East %s - kan ikke konverteres til et tal", koord.East)
		return result, errors.New(convErr)
	}
	utmg1 = utmg1 / 10.0

	utmg2, err := strconv.ParseFloat(koord.North, 64)
	if err != nil {
		convErr := fmt.Sprintf("North %s - kan ikke konverteres til et tal", koord.North)
		return result, errors.New(convErr)

	}
	utmg2 = utmg2 / 10.0

	//Index returnerer -1 hvis ej fundet
	val1 := strings.Index(alfastr, first) + 1
	val2 := strings.Index(alfastr, second) + 1

	if val1 == 0 || val2 == 0 {
		convErr := fmt.Sprintf("kmkv skal være indenfor de gyldige værdier")
		return result, errors.New(convErr)
	}

	ewutm := float64((val1-8*((zone+2)%3))*100) + utmg1/100.0

	if ewutm >= 900 || ewutm < 100 {
		convErr := fmt.Sprintf("East kan ikke beregnes for den angivne zone og kmkv")
		return result, errors.New(convErr)
	}

	nsutm := float64((val2+74+5*(zone%2))*100) + utmg2/100.0
	temp := math.Round((nsutm - midt[vbelt-1-12] + 500) / 2000)
	nsutm = nsutm - temp*2000
	ewutm = ewutm * 1000
	nsutm = nsutm * 1000
	return Utm{zone, belt, ewutm, nsutm}, err
}
