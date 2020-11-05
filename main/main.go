package main

import (
	"echo/config"
	"echo/main/worker"
	"echo/utils"
	"fmt"
	"log"
	"math"
	"sync"
	"time"
)

var dataHeaders = make([]string, 0)

func main() {
	start := time.Now()

	//koneksi ke database
	db, err := config.OpenDbConnection()
	if err != nil {
		log.Fatal(err.Error())
	}

	//membuka dan membaca file csv
	csvReader, csvFile, err := utils.OpenCsvFile()
	if err != nil {
		log.Fatal(err.Error())
	}
	defer csvFile.Close()

	jobs := make(chan []interface{}, 0)
	wg := new(sync.WaitGroup)

	// Fungsi Menjalankan Workers
	go worker.DispatchWorkers(db, jobs, wg)
	// Fungsi Baca CSV dan Pengiriman Jobs ke Worker
	worker.ReadCsvFilePerLineThenSendToWorker(csvReader, jobs, wg)

	wg.Wait()

	duration := time.Since(start)
	fmt.Println("done in", int(math.Ceil(duration.Seconds())), "seconds")
}
