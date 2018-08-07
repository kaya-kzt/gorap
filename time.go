package gorap

import (
	"fmt"
	"time"
)

var timeFormat = time.RFC3339Nano

// Period represents specific term
type Period struct {
	start, end time.Time
}

// Overlap returns true if two Peridod p & p2 are overlaps
func (p *Period) Overlap(p2 *Period) bool {
	if p2.start.Before(p.end) && p.start.Before(p2.end) {
		return true
	}
	return false
}

func (p *Period) String() string {
	return fmt.Sprintln(p.start.Format(timeFormat), "->", p.end.Format(timeFormat))
}

// StringToTime fuction parses string to Time with timeFormat format
// if parse error occured, it returns current time & error
func StringToTime(s string) (time.Time, error) {
	t, err := time.Parse(s, timeFormat)
	if err != nil {
		return time.Now(), err
	}
	return t, nil
}
