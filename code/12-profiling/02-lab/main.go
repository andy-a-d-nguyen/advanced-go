package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	// fileName = "./sample.csv"
	fileName        = "./bald-mountain_co.csv"
	timeStampFormat = "2006-01-02T15:04:05"
	monthDayFormat  = "2006-01"
	timeColumnIndex = 1
	tempColumnIndex = 43
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

type Results []Result

func (r Results) Len() int {
	return len(r)
}
func (r Results) Less(i, j int) bool {
	return strings.Compare(r[i].Day, r[j].Day) < 0
}
func (r Results) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
func (r Results) String() string {
	b := new(strings.Builder)
	const format = "%10v%10v%10v%10v\n"
	b.WriteString(fmt.Sprintf(format, "Day", "Min", "Avg", "Max"))
	b.WriteString(strings.Repeat("-", 40) + "\n")
	for _, result := range r {
		b.WriteString(fmt.Sprintf(format, result.Day, result.MinTemp, result.AvgTemp, result.MaxTemp))
	}
	return b.String()
}

var (
	rawRecordCh = make(chan []string)
	recordCh    = make(chan Record)
	resultCh    = make(chan Result)
	errorCh     = make(chan error)
)

func main() {
	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	r := csv.NewReader(f)
	data, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(execute(data))
}

func execute(data [][]string) Results {
	go processCSV(data, rawRecordCh, recordCh, errorCh)
	go processRecords(recordCh, resultCh)

	go func(errorCh <-chan error) {
		for range errorCh {
			// no-op - just drain the channel
		}
	}(errorCh)

	for _, d := range data[1:] {
		rawRecordCh <- d
	}
	close(rawRecordCh)

	results := make(Results, 0)
	for r := range resultCh {
		results = append(results, r)
	}
	sort.Sort(results)
	return results
}

func processRecords(recordCh <-chan Record, resultCh chan<- Result) {
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

func processCSV(rawRecords [][]string, rawRecordCh <-chan []string, recordCh chan<- Record, errorCh chan<- error) {
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
