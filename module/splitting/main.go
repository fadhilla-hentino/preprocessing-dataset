package main

import (
	"fadhilla-hentino/preprocessing-dataset/lib/csv"
	"fmt"
	"time"
)

const (
	BASE_PATH                 = "/home/fadhil/documents/Fadhil/Binus/THESIS _ JOURNAL/BIA/Dataset/"
	SOURCE_FILE               = BASE_PATH + "clean_chat_dataset_3M.csv"
	TRAINING_DESTINATION_FILE = BASE_PATH + "clean_training_dataset_2M.csv"
	TESTING_DESTINATION_FILE  = BASE_PATH + "clean_testing_dataset_1M.csv"
)

func splitDataset(records [][]string, destTraining, destTesting string) {
	fmt.Println("writing training data")
	training := records[1:2000001]
	csv.WriteCsvFile(training, destTraining, false)

	fmt.Println("writing testing data")
	testing := records[2000001:]
	csv.WriteCsvFile(testing, destTesting, false)
}

func main() {
	startTime := time.Now()
	fmt.Printf("start splitting dataset -- %s \n", startTime.Format(time.RFC1123Z))

	// read source file
	records := csv.ReadCsvFile(SOURCE_FILE)

	// split dataset into training and testing
	splitDataset(records, TRAINING_DESTINATION_FILE, TESTING_DESTINATION_FILE)

	endTIme := time.Now()
	fmt.Printf("finish -- %s \n", endTIme.Format(time.RFC1123Z))
	duration := time.Since(startTime)
	fmt.Printf("duration -- %s \n", duration.String())
}
