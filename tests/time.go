package main

import (
	"fmt"
	"time"
)

func main() {
	t0, _ := time.Parse("15:04:05", "00:00:00")
	t3, _ := time.Parse("15:04:05", "02:59:59")
	t6, _ := time.Parse("15:04:05", "05:59:59")

	fmt.Println(inTimeSpan(t0, t3, t0), inTimeSpan(t0, t3, t3), inTimeSpan(t0, t3, t6))
	//in, _ := time.Parse("15:04:05", "00:00:00")
}

func inTimeSpan(start, end, check time.Time) bool {
	return check.After(start) && check.Before(end) || start == check || end == check
}
