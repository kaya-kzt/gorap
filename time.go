package handy

import (
	"fmt"
	"time"
)

var timeFormat = time.RFC3339

// Period represents specific term
type Period struct {
	Start, End time.Time
}

// Overlap returns true if two Peridod p & p2 are overlaps
func (p *Period) Overlap(p2 *Period) bool {
	if p2.Start.Before(p.End) && p.Start.Before(p2.End) {
		return true
	}
	return false
}

func (p *Period) String() string {
	return fmt.Sprintln(p.Start.Format(timeFormat), "->", p.End.Format(timeFormat))
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
