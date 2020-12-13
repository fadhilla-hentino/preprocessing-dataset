package main

import (
	"fadhilla-hentino/preprocessing-dataset/lib/csv"
	"fmt"
	"time"
)

const (
	BASE_PATH        = "/home/fadhil/documents/Fadhil/Binus/THESIS _ JOURNAL/BIA/Dataset/"
	SOURCE_FILE      = BASE_PATH + "chat_dataset_3M.csv"
	DESTINATION_FILE = BASE_PATH + "clean_chat_dataset_3M.csv"
)

func writeCleanData(records [][]string, destination string) {
	fmt.Println("writing clean data")
	csv.WriteCsvFile(records, destination, true)
}

func main() {
	startTime := time.Now()
	fmt.Printf("start removing unused fields from dataset -- %s \n", startTime.Format(time.RFC1123Z))

	// read source file
	records := csv.ReadCsvFile(SOURCE_FILE)

	// write clean dataset
	writeCleanData(records, DESTINATION_FILE)

	endTIme := time.Now()
	fmt.Printf("finish -- %s \n", endTIme.Format(time.RFC1123Z))
	duration := time.Since(startTime)
	fmt.Printf("duration -- %s \n", duration.String())
}
