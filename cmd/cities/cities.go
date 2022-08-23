// Package cities

// Indlæsning af filen $HOME/geogrgafi/cites.csv og print til stdout
package main

import (
	"bufio"
	"fmt"
	"github.com/brundtoe/go-geografi/pkg/utils"
	"log"
	"os"
)

func main() {
	part := "geografi/cities.csv"
	filename, err := utils.GetDataPath(part, "PLATFORM")
	if err != nil {
		fmt.Printf("Det er ikke muligt at finde path filen %s: %s", part, err)
		os.Exit(1)
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var tx string
	for scanner.Scan() {
		tx = scanner.Text()
		fmt.Println(tx)
	}

}
