package gorap

import (
	"testing"
	"time"
)

var (
	t1 = time.Date(2018, time.April, 1, 12, 0, 1, 12, time.UTC)
	t2 = time.Date(2018, time.April, 20, 18, 0, 1, 12, time.UTC)
	t3 = time.Date(2018, time.May, 1, 12, 0, 1, 12, time.UTC)
	t4 = time.Date(2018, time.March, 2, 18, 0, 1, 12, time.UTC)
)

func TestOverlap(t *testing.T) {
	p1 := Period{start: t1, end: t3}
	p2 := Period{start: t4, end: t2}
	result := p1.Overlap(&p2)
	expected := true
	if result != expected {
		t.Fatal("test failed")
	}
}
