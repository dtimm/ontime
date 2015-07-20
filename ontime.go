// OnTime, a time library for converting between different time structures
//   and extracting interesting information regarding events that people
//   historically did calendar stuff with.

package main

import (
	"flag"
	"fmt"
	"time"
)

var now time.Time
var outFmt string
var inFmt string
var inTime string

func main() {
	outTime := flag.String("ot", "std", "Output time clock. standard, binary,"+
		"or metric.")
	outForm := flag.String("of", " ", "Output time format.")
	inTm := flag.String("it", "", "Input time.")
	inForm := flag.String("if", " ", "Input time format.")
	flag.Parse()

	inTime = *inTm
	outFmt = *outForm
	inFmt = *inForm

	if inTime != "" {
		now, _ = time.Parse(inFmt, inTime)
	} else {
		now = time.Now()
	}

	switch *outTime {
	case "std", "standard", "s":
		fmt.Printf("%02d:%02d", now.Hour(), now.Minute())
	case "met", "metric", "m":
		PrintMetTime()
	case "bin", "b", "binary", "hex":
		PrintHexTime()
	}
}

func stdToday() int64 {
	return int64((now.Second()+60*now.Minute()+3600*now.Hour())*1000000000 + now.Nanosecond())
}

func metToday() int64 {
	return stdToday() * 1000 / 864
}

func binToday() int64 {
	return stdToday() * 65536 / 86400
}

func binHour() int {
	return int(binToday()/1000000000) / 4096
}

func binMinute() int {
	return int(binToday()/1000000000) % 4096 / 256
}

func binSecond() int {
	return int(binToday()/1000000000) % 256
}

func binNanosecond() int64 {
	return binToday() % 65536
}

func metHour() int {
	return int(metToday() / 10000000000000)
}

func metMinute() int {
	return int(metToday()/1000000000) % 10000 / 100
}

func metSecond() int {
	return int(metToday()/1000000000) % 100
}

func metNanosecond() int64 {
	return metToday() / 1000000000000000
}

// Print binary time HM
func PrintHexTime() {
	fmt.Printf("%X%X", binHour(), binMinute())
}

// Print metric time HH:MM
func PrintMetTime() {
	fmt.Printf("%02d:%02d", metHour(), metMinute())
}
