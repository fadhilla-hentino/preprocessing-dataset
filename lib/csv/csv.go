package csv

import (
	"encoding/csv"
	"log"
	"os"
)

func ReadCsvFile(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func WriteCsvFile(records [][]string, destination string, isClean bool) {
	file, err := os.Create(destination)
	if err != nil {
		log.Fatal("Unable to create CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if isClean {
		for i := 0; i < len(records); i++ {
			record := records[i][3]
			if err := writer.Write([]string{record}); err != nil {
				log.Fatal("Unable to create CSV file at index : "+string(rune(i)), err)
			}
		}
	} else {
		for i := 0; i < len(records); i++ {
			record := records[i]
			if err := writer.Write(record); err != nil {
				log.Fatal("Unable to create CSV file at index : "+string(rune(i)), err)
			}
		}
	}
}
