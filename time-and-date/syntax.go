package main

import (
	"fmt"
	"time"
)

func addToTime() {
	t := time.Now()

	//Add 1 hours
	newT := t.Add(time.Hour * 1)
	fmt.Printf("Adding 1 hour\n: %s\n", newT)

	//Add 15 min
	newT = t.Add(time.Minute * 15)
	fmt.Printf("Adding 15 minute\n: %s\n", newT)

	//Add 10 sec
	newT = t.Add(time.Second * 10)
	fmt.Printf("Adding 10 sec\n: %s\n", newT)

	//Add 100 millisecond
	newT = t.Add(time.Millisecond * 10)
	fmt.Printf("Adding 100 millisecond\n: %s\n", newT)

	//Add 1000 microsecond
	newT = t.Add(time.Microsecond * 10)
	fmt.Printf("Adding 1000 microsecond\n: %s\n", newT)

	//Add 10000 nanosecond
	newT = t.Add(time.Nanosecond * 1000)
	fmt.Printf("Adding 1000 nanosecond\n: %s\n", newT)

	//Add 1 year 2 month 4 day
	newT = t.AddDate(1, 2, 4)
	fmt.Printf("Adding 1 year 2 month 4 day\n: %s\n", newT)
}

func subtractToTime() {
	fmt.Println("\nSubtract to time")
	t := time.Now()

	//Subtract 1 hours
	newT := t.Add(-time.Hour * 1)
	fmt.Printf("Subtracting 1 hour\n: %s\n", newT)

	//Subtract 15 min
	newT = t.Add(-time.Minute * 15)
	fmt.Printf("Subtracting 15 minute\n: %s\n", newT)

	//Subtract 10 sec
	newT = t.Add(-time.Second * 10)
	fmt.Printf("Subtracting 10 sec\n: %s\n", newT)

	//Subtract 100 millisecond
	newT = t.Add(-time.Millisecond * 10)
	fmt.Printf("Subtracting 100 millisecond\n: %s\n", newT)

	//Subtract 1000 microsecond
	newT = t.Add(-time.Microsecond * 10)
	fmt.Printf("Subtracting 1000 microsecond\n: %s\n", newT)

	//Subtract 10000 nanosecond
	newT = t.Add(-time.Nanosecond * 1000)
	fmt.Printf("Subtracting 1000 nanosecond\n: %s\n", newT)

	//Subtract 1 year 2 month 4 day
	newT = t.AddDate(-1, -2, -4)
	fmt.Printf("Subtracting 1 year 2 month 4 day\n: %s\n", newT)
}

func timeParse() {
	fmt.Println("\nTime Parse")
	// Parse YYYY-MM-DD
	timeT, _ := time.Parse("2006-01-02", "2020-01-29")
	fmt.Println(timeT)

	// Parse YY-MM-DD
	timeT, _ = time.Parse("06-01-02", "20-01-09")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD
	timeT, _ = time.Parse("2006-Jan-02", "2020-Jan-29")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05", "2020-Jan-29 Wednesday 12:19:25")
	fmt.Println(timeT)

	//Parse YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	timeT, _ = time.Parse("2006-Jan-02 Monday 03:04:05 PM MST -07:00", "2020-Jan-29 Wednesday 12:19:25 AM IST +05:30")
	fmt.Println(timeT)
}

// time.Format function can be used to format time to a string representation. The signature function is:
// func (t Time) Format(layout string)
func timeFormatting() {
	fmt.Println("\nTime Formatting")
	now := time.Now()

	//Format YYYY-MM-DD
	fmt.Printf("YYYY-MM-DD: %s\n", now.Format("2006-01-02"))

	//Format YY-MM-DD
	fmt.Printf("YY-MM-DD: %s\n", now.Format("06-01-02"))

	//Format YYYY-#{MonthName}-DD
	fmt.Printf("YYYY-#{MonthName}-DD: %s\n", now.Format("2006-Jan-02"))

	//Format HH:MM:SS
	fmt.Printf("HH:MM:SS: %s\n", now.Format("03:04:05"))

	//Format HH:MM:SS Millisecond
	fmt.Printf("HH:MM:SS Millisecond: %s\n", now.Format("03:04:05 .999"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS: %s\n", now.Format("2006-Jan-02 Monday 03:04:05"))

	//Format YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset
	fmt.Printf("YYYY-#{MonthName}-DD WeekDay HH:MM:SS PM Timezone TimezoneOffset: %s\n", now.Format("2006-Jan-02 Monday 03:04:05 PM MST -07:00"))
}

// Time Diff
// time package has a method "Sub" which can be used to get the difference between two different time values. The signature functino is:
// func (t Time) Sub(u Time) Duration
func timeDiff() {
	fmt.Println("\nTime Diff")
	currentTime := time.Now()
	oldTime := time.Date(2020, 1, 2, 0, 0, 0, 0, time.UTC)
	dif := currentTime.Sub(oldTime)
	fmt.Printf("Different is : %s\n", dif)
}

// Time Conversion
func timeConversion() {
	fmt.Println("\nTime Conversion")
	tNow := time.Now()

	//time.Time to Unix Timestamp
	tUnix := tNow.Unix()
	fmt.Printf("timeUnix %d\n", tUnix)

	//Unit Timestamp to time.Time
	timeT := time.Unix(tUnix, 0)
	fmt.Printf("time.Time: %s\n", timeT)
}

func conversionTimeBetweenDifferentLocation() {
	fmt.Println("\nTime Conversion With Different Location")
	now := time.Now()

	loc, _ := time.LoadLocation("UTC")
	fmt.Printf("UTC Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Europe/Berlin")
	fmt.Printf("Berlin Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("America/New_York")
	fmt.Printf("New York Time: %s\n", now.In(loc))

	loc, _ = time.LoadLocation("Asia/Dubai")
	fmt.Printf("Dubai Time: %s\n", now.In(loc))
}
