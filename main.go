package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func readCsvFile(filePath string) [][]string {
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

func writeCsvFile(records [][]string, destination string, isClean bool) {
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

func splitDataset(records [][]string, destTraining, destTesting string, done chan struct{}) {
	fmt.Println("writing training data")
	training := records[1:2000001]
	writeCsvFile(training, destTraining, false)

	fmt.Println("writing testing data")
	testing := records[2000001:]
	writeCsvFile(testing, destTesting, false)

	done <- struct{}{}
}

func writeCleanData(records [][]string, destTraining, destTesting string, done chan struct{}) {
	fmt.Println("writing clean training data")
	training := records[1:2000001]
	writeCsvFile(training, destTraining, true)

	fmt.Println("writing clean testing data")
	testing := records[2000001:]
	writeCsvFile(testing, destTesting, true)

	done <- struct{}{}
}

const (
	basePath = "./"
)

func main() {
	startTime := time.Now()
	fmt.Printf("start -- %s \n", startTime.Format(time.RFC1123Z))

	splitChan := make(chan struct{}, 1)
	cleanChan := make(chan struct{}, 1)

	// read source file
	records := readCsvFile(basePath + "chat_dataset_3M.csv")

	// split dataset into training and testing
	go splitDataset(records, basePath+"training_dataset_2M.csv", basePath+"testing_dataset_1M.csv", splitChan)

	// write clean training and testing dataset
	go writeCleanData(records, basePath+"clean_training_dataset_2M.csv", basePath+"clean_testing_dataset_1M.csv", cleanChan)

	// waiting all process to finish
	<-splitChan
	<-cleanChan

	endTIme := time.Now()
	fmt.Printf("finish -- %s \n", endTIme.Format(time.RFC1123Z))
	duration := time.Since(startTime)
	fmt.Printf("duration -- %s \n", duration.String())
}
