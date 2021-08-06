package demo_base

import (
	"fmt"
	"testing"
	"time"
)

func TestGetTime(t *testing.T) {
	// now time
	now := time.Now()
	fmt.Println(now)
	// unix timestamp
	fmt.Println(now.Unix())
	// nano time
	fmt.Println(now.UnixNano())
	// decimal part
	fmt.Println(now.Nanosecond())
}

func TestGetTimeTypeCalendar(t *testing.T) {
	now := time.Now()
	year, month, day := now.Date()
	fmt.Printf("year:%d, month:%d, day:%d\n", year, month, day)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())

	// clock
	hour, min, sec := now.Clock()
	fmt.Printf("hour:%d, min:%d, sec:%d\n", hour, min, sec)
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())

	// week
	fmt.Println(now.Weekday())
	fmt.Println(now.YearDay())
	// timezone
	fmt.Println(now.Location())
	// day
	fmt.Println(now.YearDay())
}
