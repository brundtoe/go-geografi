// Package cities

// Indlæsning af filen cites.csv og print til stdout
package main

import (
	"bufio"
	"fmt"
	"log"

	"github.com/brundtoe/go-geografi/pkg/utils"
)

func main() {
	filename := "geografi/cities.csv"
	fp, err := utils.OpenDataFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		fmt.Printf("Datafilen %s lukkes", filename)
		if err = fp.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(fp)

	var tx string
	for scanner.Scan() {
		tx = scanner.Text()
		fmt.Println(tx)
	}

}
