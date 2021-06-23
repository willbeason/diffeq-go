package main

import (
	"github.com/willbeason/diffeq-go/pkg/models"
)

const (
	Width  = 2560
	Height = 1440
)

func work(y0, yp0, t0, h float64, n int, out chan int) {
	t := t0
	for i := 0; i < n; i++ {

		t += h
	}
}

func main() {
	spring := models.DuffingOscillator{
		Delta:     0.02,
		Alpha:     1.0,
		Beta:      5.0,
		Gamma:     12.0,
		Frequency: 0.5,
	}

	results := make([]int64, Width*Height)

}
