package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func main() {

	f, err := os.Open("/Users/david/dev/Golang/src/golang_ml/data/iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)

	reader.FieldsPerRecord = -1

	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(rawCSVData)
}
