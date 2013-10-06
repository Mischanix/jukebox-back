package main

import (
	"encoding/csv"
	"math/rand"
	"os"
	"time"
)

var (
	nameFirsts []string
	nameLasts  []string
)

func init() {
	namesPath := "names/"
	firstsFile, err := os.Open(namesPath + "first")
	if err != nil {
		panic(err)
	}
	lastsFile, err := os.Open(namesPath + "last")
	if err != nil {
		panic(err)
	}

	records, err := csv.NewReader(firstsFile).ReadAll()
	if err != nil {
		panic(err)
	}
	nameFirsts = make([]string, 0)
	for _, record := range records {
		nameFirsts = append(nameFirsts, record...)
	}

	records, err = csv.NewReader(lastsFile).ReadAll()
	if err != nil {
		panic(err)
	}
	nameLasts = make([]string, 0)
	for _, record := range records {
		nameLasts = append(nameLasts, record...)
	}

	rand.Seed(time.Now().UnixNano())
}

func FakeName() string {
	first := nameFirsts[rand.Intn(len(nameFirsts))]
	last := nameLasts[rand.Intn(len(nameLasts))]
	return first + " " + last
}
