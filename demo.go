package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// type location struct {
	// 	Lat, Long float64
	// }
	type location struct {
		Lat  float64 `json:"lat"`
		Long float64 `json:"long"`
	}

	curiosity := location{3.14, 5.36}

	bytes, err := json.Marshal(curiosity)
	exitOnError(err)

	fmt.Println(string(bytes))
}

func exitOnError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
