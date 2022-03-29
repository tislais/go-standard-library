package main

import (
	"fmt"
	"time"
)

func main() {

	first := time.Now()
	fmt.Printf("It is currently %v \n", first.Format("15:04:05"))
	time.Sleep(2 * time.Second)
	second := time.Now()
	fmt.Printf("It is now %v \n", second.Format("15:04:05"))

	// Mon Jan 2 15:04:05 MST 2006
	today := time.Now()
	fmt.Printf("Today is %v\n", today.Format("Monday, Jan 2 2006"))
	startDate := time.Date(2018, 07, 04, 9, 00, 00, 0, time.UTC)
	elapsed := time.Since(startDate)

	fmt.Printf("Hours: %.0f Minutes: %.0f Seconds: %.0f\n", elapsed.Hours(), elapsed.Minutes(), elapsed.Seconds())

	// challenge: find the day six months from now
	future := today.AddDate(0, 6, 0)
	past := today.AddDate(0, -6, 0)
	future2 := today.Add(6 * time.Hour)

	fmt.Printf("In six months it will be %v\n", future.Format("Monday, January 2"))
	fmt.Printf("Six months ago it was %v\n", past.Format("Monday, January 2"))
	fmt.Printf("In six hours it will be %v\n", future2.Format("15:04"))

	// set deadlines
	bedtime := time.Date(2022, 3, 29, 23, 0, 0, 0, time.Local)

	fmt.Printf("There are %.0f hours until bed time\n", time.Until(bedtime).Hours())
}
