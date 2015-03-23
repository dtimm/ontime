// OnTime, a time library for converting between different time structures
//   and extracting interesting information regarding events that people
//   historically did calendar stuff with.

package main

import (
	"fmt"
	"flag"
	"time"
)

func main() {
	outTime := flag.String("ot", "std", "Output time clock. standard, binary, or metric.")
	//outForm := flag.String("of", " ", "Output time format.")
	//inForm := flag.String("if", " ", "Input time format.")
	flag.Parse()

	switch *outTime {
		case "std", "standard", "s":
			fmt.Printf("%d", todaySecStd() / 1000000000)
		case "met", "metric", "m":
			PrintMetTime()
		case "bin", "b", "binary", "hex":
			PrintHexTime()
	}
}

func todaySecStd () int64 {
	now := time.Now()
	return int64((now.Second() + 60 * now.Minute() + 3600 * now.Hour()) * 1000000000 + now.Nanosecond())
}

func todaySecMet () int64 {
	return todaySecStd() * 1000 / 864
}

func todaySecBin () int64 {
	return todaySecStd() * 65536 / 86400
}

func BinToStd(binSecs int64) int64 {
	return binSecs * 86400 / 65536
}

func MetToStd(metSecs int64) int64 {
	return metSecs * 864 / 1000
}

func BinToMet(binSecs int64) int64 {
	return binSecs * 100000 / 65536
}

func MetToBin(metSecs int64) int64 {
	return metSecs * 65536 / 100000
}

func PrintHexTime() {
	fmt.Printf("%X", todaySecBin() / 1000000000)
}

func PrintMetTime() {
	now := todaySecMet() / 1000000000
	hour := now / 10000
	minute := now % 10000 / 100
	second := now % 100
	fmt.Printf("%2d:%2d:%2d", hour, minute, second)
}
