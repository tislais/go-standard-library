package main

import (
	"fmt"
	"time"
)

const (
	layoutISO = "2006-01-02"
	layoutUS  = "January 2, 2006"
	layoutEU  = "2 January, 2006"
)

func main() {
	t := time.Now()

	fmt.Println(t.Format(time.ANSIC))
	// output: Tue Mar 29 10:50:06 2022

	fmt.Println(t.Format(time.RFC3339))
	// output: 2022-03-29T10:50:55-07:00

	fmt.Println(t.Format(time.UnixDate))
	// output: Tue Mar 29 10:51:37 PDT 2022

	fmt.Println(t.Format(time.Kitchen))
	// output: 10:52AM

	// Create your own time format.
	// Use the following reference time!
	// Mon Jan 2 15:04:05 MST 2006
	fmt.Println(t.Format("Monday, January 2 in the year of our lord 2006"))
	// output: Tue Mar 29 10:54:15 PDT 2022
	// this is the weirdest feature of go yet

	startDate := time.Date(2018, 07, 04, 9, 00, 00, 00, time.UTC)
	fmt.Println(startDate.Format("Monday, Jan 2 2006 at 15:04:05 2006"))

	fmt.Println(startDate.Format(layoutISO))
	fmt.Println(startDate.Format(layoutUS))
	fmt.Println(startDate.Format(layoutEU))
}
