package main

import (
	"fmt"
	"time"
)

func main() {
	timeStr := "2019-03-19T18:50:45Z" // time string retrieved from database
	formattedTime := formatZuluTimeInProperTimeFormat(timeStr)
	fmt.Println(formattedTime) // prints "Tuesday, 19-Mar-19 13:50:45 CDT"
}

// function to convert zulu(UTC) time string to proper date time format in specified time zone
func formatZuluTimeInProperTimeFormat(timeStr string) string {
	// parsing the timeStr to convert to time.Time object so that you can switch it to the required timezone. RFC3339 should be used because the timeString format is closest to RFC3339 constant
	baseTime, _ := time.Parse(time.RFC3339, timeStr) // baseTime = "2019-03-19 18:50:45 +0000 UTC"

	location, _ := time.LoadLocation("America/Chicago")
	// baseTime has be to be time.Time object. It cannot be string.
	timeInCDT := baseTime.In(location)   // timeInCDT = 2019-03-19 13:50:45 -0500 CDT
	return timeInCDT.Format(time.RFC850) // format time into any constant you require
}

// below are the constants provided by time package
const (
	ANSIC       = "Mon Jan _2 15:04:05 2006"
	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
	RFC822      = "02 Jan 06 15:04 MST"
	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
	RFC3339     = "2006-01-02T15:04:05Z07:00"
	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
	Kitchen     = "3:04PM"
	// Handy time stamps.
	Stamp      = "Jan _2 15:04:05"
	StampMilli = "Jan _2 15:04:05.000"
	StampMicro = "Jan _2 15:04:05.000000"
	StampNano  = "Jan _2 15:04:05.000000000"
)
