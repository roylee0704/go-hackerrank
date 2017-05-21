// resource: https://www.hackerrank.com/contests/booking-womenintech/challenges/visiting-manhattan
package main

import "fmt"

func main() {

	// var x, y, l, h int
	//
	// fmt.Fscanf(os.Stdin, "%d %d %d %d", &x, &y, &l, &h)
	//
	// fmt.Print(x, y, l, h)

	p1 := &point{4, 4}
	p2 := &point{1, 1}

	fmt.Print(p1.dist(p2))

	// fmt.Println(dist(1, 2, 1, 1))
	// fmt.Println(dist(1, 2, 2, 2))

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
