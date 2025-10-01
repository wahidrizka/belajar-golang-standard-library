package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {
	csvString := "wahid,rizka,fathurrohman\n" +
		"aurelia,syahwa,erwana\n" +
		"aurel fathur rohman"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		fmt.Println(record)
	}
}
