package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

func main() {
	var data *ring.Ring = ring.New(5)

	// Ini cara menggunakan iterasi
	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i)
		data = data.Next()
	}

	// Ini cara manual
	// data.Value = "Value ke-1"

	// data = data.Next()
	// data.Value = "Value ke-2"

	// data = data.Next()
	// data.Value = "Value ke-3"

	// data = data.Next()
	// data.Value = "Value ke-4"

	// data = data.Next()
	// data.Value = "Value ke-5"

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
