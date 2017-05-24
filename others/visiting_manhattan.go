// resource: https://www.hackerrank.com/contests/booking-womenintech/challenges/visiting-manhattan
package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	var width, height, l, h, x, y int

	fmt.Fscanf(os.Stdin, "%d %d %d %d", &width, &height, &l, &h)

	landmarks := make([]landmark, l)
	for i := 0; i < l; i++ {
		fmt.Fscanf(os.Stdin, "%d %d", &x, &y)
		landmarks[i] = landmark{x, y}
	}

	k := manhanttan{
		landmarks: landmarks,
	}

	trips := make(chan item)
	var wg sync.WaitGroup
	go func() {
		wg.Wait()
		close(trips)
	}()
	for i := 0; i < h; i++ {
		fmt.Fscanf(os.Stdin, "%d %d", &x, &y)
		wg.Add(1)
		go k.bestHotel(i, hotel{x, y}, trips, &wg)
	}

	minDist := 10000000
	minHotelIndex := 0

	for t := range trips {
		if t.dist < minDist {
			minDist = t.dist
			minHotelIndex = t.i
		}
	}

	fmt.Print(minHotelIndex + 1)
}

type manhanttan struct {
	hotels    []hotel
	landmarks []landmark
}

type item struct {
	i    int
	dist int
}

// bestHotel is a gofunc
func (m manhanttan) bestHotel(i int, h hotel, trips chan item, wg *sync.WaitGroup) {
	var dist int
	dist = h.dist(m.landmarks)
	trips <- item{i, dist}
	wg.Done()
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
