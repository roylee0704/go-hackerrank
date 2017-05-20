package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// time represents an instant in time w/ necessary parts.
type time struct {
	hh   int
	mm   int
	ss   int
	isPM bool
}

// String implements Stringer interface that outputs 24hours time format.
func (t *time) String() string {
	if t.isPM && t.hh >= 1 && t.hh < 12 {
		t.hh += 12
	} else if !t.isPM && t.hh == 12 {
		t.hh -= 12
	}

	return fmt.Sprintf("%02d:%02d:%02d", t.hh, t.mm, t.ss)
}

// parse parses 12hours time format
func parse(t12h string) *time {
	parts := strings.Split(t12h, ":")
	hh, _ := strconv.Atoi(parts[0])
	mm, _ := strconv.Atoi(parts[1])
	ss, _ := strconv.Atoi(parts[2][:2])
	ampm := parts[2][2:]
	isPM := false

	if ampm == "PM" {
		isPM = true
	}

	return &time{
		hh:   hh,
		mm:   mm,
		ss:   ss,
		isPM: isPM,
	}
}

func main() {
	var t12h string
	fmt.Fscan(os.Stdin, &t12h)
	t := parse(t12h)
	fmt.Println(t)
}
