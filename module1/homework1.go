package main

import (
	"fmt"
)

func main() {
	var array1 = [5]string{"I", "am", "stupid", "and", "week"}
	slice1 := array1[:]
	for i, v := range slice1 {
		if v == "stupid" {
			slice1[i] = "smart"
		} else if v == "week" {
			slice1[i] = "srong"
		} else {
			continue
		}
	}
	fmt.Println(slice1)
}
