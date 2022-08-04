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

	hostType := os.Getenv("PLATFORM")
	var filename string
	if hostType == "kvm" {
		filename = "/nfs/data/geografi/cities.csv"
	} else if hostType == "none" {
		filename = "/home/projects/devops/data/geografi/cities.csv"
	} else {
		fmt.Println("Denne hosttype %v er ikke understøttet", hostType)
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
