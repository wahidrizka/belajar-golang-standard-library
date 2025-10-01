package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	_ = writer.Write([]string{"wahid", "rizka", "fathurrohman"})
	_ = writer.Write([]string{"aurelia", "syahwa", "erwana"})
	_ = writer.Write([]string{"aurel", "fathur", "rohman"})

	writer.Flush()
}
