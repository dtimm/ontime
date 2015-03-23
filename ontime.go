// OnTime, a time library for converting between different time structures
//   and extracting interesting information regarding events that people
//   historically did calendar stuff with.

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("standard: %d\n", todaySecStd())
	fmt.Printf("metric:   %d\n", todaySecMet())
	fmt.Printf("binary:   %d\n", todaySecBin())
}

func todaySecStd () int {
	now := time.Now()
	return now.Second() + 60 * now.Minute() + 3600 * now.Hour()
}

func todaySecMet () int {
	return todaySecStd() * 1000 / 864
}

func todaySecBin () int {
	return todaySecStd() * 65536 / 86400
}

func BinToStd(binSecs int) int {
	return binSecs * 86400 / 65536
}

func MetToStd(metSecs int) int {
	return metSecs * 864 / 1000
}

func BinToMet(binSecs int) int {
	return binSecs * 100000 / 65536
}

func MetToBin(metSecs int) int {
	return metSecs * 65536 / 100000
}

