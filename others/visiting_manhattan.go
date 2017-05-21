// resource: https://www.hackerrank.com/contests/booking-womenintech/challenges/visiting-manhattan
package main

import "fmt"

func main() {

	// var x, y, l, h int
	//
	// fmt.Fscanf(os.Stdin, "%d %d %d %d", &x, &y, &l, &h)
	//
	// fmt.Print(x, y, l, h)

	h1 := &hotel{4, 4}
	l1 := landmark{1, 1}
	l2 := landmark{2, 2}
	fmt.Print(h1.dist([]landmark{l1, l2}))
}

type hotel point
type landmark point

func (h *hotel) dist(landmarks []landmark) int {

	var sum int
	var p2 point
	p1 := point(*h)

	for _, l := range landmarks {
		p2 = point(l)
		sum += p1.dist(&p2)
	}
	return sum
}

type point struct {
	x int
	y int
}

func (p *point) dist(p2 *point) int {
	return abs(p.x-p2.x) + abs(p.y-p2.y)
}

func abs(val int) int {
	if val < 0 {
		val = val * -1
	}
	return val
}
