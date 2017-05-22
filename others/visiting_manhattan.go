// resource: https://www.hackerrank.com/contests/booking-womenintech/challenges/visiting-manhattan
package main

import (
	"fmt"
	"os"
)

func main() {

	var width, height, l, h, x, y int

	fmt.Fscanf(os.Stdin, "%d %d %d %d", &width, &height, &l, &h)

	landmarks := make([]landmark, l)
	for i := 0; i < l; i++ {
		fmt.Fscanf(os.Stdin, "%d %d", &x, &y)
		landmarks[i] = landmark{x, y}
	}

	hotels := make([]hotel, h)
	for i := 0; i < h; i++ {
		fmt.Fscanf(os.Stdin, "%d %d", &x, &y)
		hotels[i] = hotel{x, y}
	}

	k := manhanttan{
		hotels:    hotels,
		landmarks: landmarks,
	}
	fmt.Println(k.bestHotel())
}

type manhanttan struct {
	hotels    []hotel
	landmarks []landmark
}

func (m manhanttan) bestHotel() int {

	var res, dist int
	var min = 10000000 // quick hack

	for i, h := range m.hotels {
		dist = h.dist(m.landmarks)

		if dist < min {
			min = dist
			res = i
		}
	}
	return res + 1 // 1-based index
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
