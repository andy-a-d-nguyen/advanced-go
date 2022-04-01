package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"sync"
	"time"
)

const (
	// fileName = "./sample.csv"
	fileName = "./bald-mountain_co.csv"
)

type Record struct {
	Time time.Time
	Temp int
}

type Result struct {
	Day     string
	MinTemp int
	AvgTemp int
	MaxTemp int
}

var (
	rawRecordCh = make(chan []string)
	recordCh    = make(chan Record)
	errorCh     = make(chan error)
)

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	_, err = r.Read() // read headers from file and throw away
	if err != nil {
		log.Fatal(err)
	}

	go processCSV(rawRecordCh, recordCh, errorCh)
	resultCh := processRecords(recordCh)

	go func(errorCh <-chan error) {
		for err := range errorCh {
			log.Println(err)
		}
	}(errorCh)

	for {
		rawRecord, err := r.Read()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			errorCh <- err
			continue
		}
		rawRecordCh <- rawRecord
	}
	close(rawRecordCh)

	results := make(Results, 0)
	for r := range resultCh {
		results = append(results, r)
	}
	sort.Sort(results)
	fmt.Println(results)
}

func processCSV(rawRecordCh <-chan []string, recordCh chan<- Record, errorCh chan<- error) {
	wg := new(sync.WaitGroup)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go createRecords(rawRecordCh, recordCh, errorCh, wg)
	}
	go func(wg *sync.WaitGroup) {
		wg.Wait()
		close(recordCh)
	}(wg)
}

const (
	timeStampFormat = "2006-01-02T15:04:05"
	timeColumnIndex = 1
	tempColumnIndex = 43
)

func createRecords(in <-chan []string, out chan<- Record, errCh chan<- error, wg *sync.WaitGroup) {
	for rawRecord := range in {
		t, err := time.Parse(timeStampFormat, rawRecord[timeColumnIndex])
		if err != nil {
			errCh <- fmt.Errorf("Failed to parse time for raw value: %v.\n%w", rawRecord[timeColumnIndex], err)
			continue
		}
		temp, err := strconv.Atoi(rawRecord[tempColumnIndex])
		if err != nil {
			errCh <- fmt.Errorf("Failed to parse Temp for raw value: %v\n %w", rawRecord[tempColumnIndex], err)
			continue
		}
		out <- Record{Time: t, Temp: temp}
	}
	wg.Done()
}

func processRecords(recordCh <-chan Record) <-chan Result {
	resultCh := make(chan Result)
	go func() {
		dayChannels := make(map[string]chan Record)
		wg := new(sync.WaitGroup)
		for r := range recordCh {
			dayString := r.Time.Format("2006-01-02")
			if _, ok := dayChannels[dayString]; !ok {
				wg.Add(1)
				dayChannels[dayString] = make(chan Record)
				go processDay(dayString, dayChannels[dayString], resultCh, wg)
			}
			dayChannels[dayString] <- r
		}
		for _, ch := range dayChannels {
			close(ch)
		}
		wg.Wait()
		close(resultCh)
	}()
	return resultCh
}

func processDay(day string, in <-chan Record, out chan<- Result, wg *sync.WaitGroup) {
	records := make([]Record, 0)
	for record := range in {
		records = append(records, record)
	}

	minTemp, maxTemp, avgTemp := math.MaxInt16, math.MinInt16, 0

	for _, r := range records {
		if r.Temp < minTemp {
			minTemp = r.Temp
		}
		if maxTemp < r.Temp {
			maxTemp = r.Temp
		}
		avgTemp += r.Temp
	}
	avgTemp /= len(records)

	out <- Result{Day: day, MinTemp: minTemp, MaxTemp: maxTemp, AvgTemp: avgTemp}

	wg.Done()
}
