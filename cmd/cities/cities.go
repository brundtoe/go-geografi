// Package cities

// Indlæsning af filen $HOME/geogrgafi/cites.csv og print til stdout
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	file, err := os.Open("/home/jackie/dumps/geografi/cities.csv")
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
