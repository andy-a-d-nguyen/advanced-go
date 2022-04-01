package main

import (
	"fmt"
	"strings"
)

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
