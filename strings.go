package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("Wahid Rizka Fathurrohman", "Wahid"))
	fmt.Println(strings.Split("Wahid Rizka Fathurrohman", " "))
	fmt.Println(strings.ToLower("Wahid Rizka Fathurrohman"))
	fmt.Println(strings.ToUpper("Wahid Rizka Fathurrohman"))
	fmt.Println(strings.Trim("       Wahid Rizka Fathurrohman       ", " "))
	fmt.Println(strings.ReplaceAll("Wahid Rizka Fathurrohman Wahid Rizka", "Wahid", "Aurelia"))
}
