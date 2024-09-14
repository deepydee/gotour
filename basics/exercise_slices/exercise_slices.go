package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {
	// Allocate two-dimensional array.
	result := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		result[y] = make([]uint8, dx)
	}

	// Do something.
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			result[y][x] = (uint8(x) * uint8(x)) + (uint8(y) * uint8(y))
		}
	}

	return result
}

func main() {
	pic.Show(Pic)
}
