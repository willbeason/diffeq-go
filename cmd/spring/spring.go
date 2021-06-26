package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/rand"
	"os"
	"sync"
	"time"

	"github.com/willbeason/diffeq-go/pkg/equations"
	"github.com/willbeason/diffeq-go/pkg/models"
	"github.com/willbeason/diffeq-go/pkg/solvers/order2"
)

const (
	Width  = 2560
	Height = 1440

	MinX = -0.5
	MaxX = 2.5
	MinY = -5.0
	MaxY = 4.0
)

var (
	minY  = 100.0
	maxY  = -100.0
	minYP = 100.0
	maxYP = -100.0
)

func toPixel(y, yp float64) int {
	// if y > MaxX {
	// 	panic(y)
	// }
	// if y < MinX {
	// 	panic(y)
	// }
	// if yp > MaxY {
	// 	panic(yp)
	// }
	// if yp < MinY {
	// 	panic(yp)
	// }

	// minY = math.Min(y, minY)
	// maxY = math.Max(y, maxY)
	// minYP = math.Min(yp, minYP)
	// maxYP = math.Max(yp, maxYP)

	px := (y - MinX) * float64(Width) / (MaxX - MinX)
	py := (yp - MinY) * float64(Height) / (MaxY - MinY)

	return int(py)*Width + int(px)
}

func work(eq equations.SecondOrder, solver order2.Solver, t0, y0, yp0, h float64, n int, out chan int) (float64, float64) {
	t := t0
	y := y0
	yp := yp0

	for i := 0; i < n; i++ {
		y, yp = order2.Solve(solver, eq, t, y, yp, t+h, 1000)
		t += h

		if out != nil {
			out <- toPixel(y, yp)
		}
	}
	return y, yp
}

func reduce(in chan int, out []int) {
	for i := range in {
		if i >= len(out) || i < 0 {
			continue
		}
		out[i]++
	}
}

func main() {
	spring := models.DuffingOscillator{
		Delta:     0.02,
		Alpha:     1.0,
		Beta:      5.0,
		Gamma:     12.061735,
		Frequency: 0.5,
	}
	fmt.Println(spring.Gamma)

	results := make(chan int, 1000)

	wg := sync.WaitGroup{}
	nWorkers := 8
	wg.Add(nWorkers)

	startYs := make([]float64, nWorkers)
	startYPs := make([]float64, nWorkers)
	for i := 0; i < nWorkers; i++ {
		y0 := rand.Float64() - 0.5
		startYs[i], startYPs[i] = work(spring.Acceleration, order2.RK4, 0.0, y0, 0.0, 2*math.Pi/spring.Frequency, 1000, nil)
	}

	for i := 0; i < nWorkers; i++ {
		y0 := startYs[i]
		yp0 := startYPs[i]
		go func() {
			for n := 0; n < 30; n++ {
				y0, yp0 = work(spring.Acceleration, order2.RK4, 0.0, y0, yp0, 2*math.Pi/spring.Frequency, 1000, results)
			}
			wg.Done()
		}()
	}

	wg2 := sync.WaitGroup{}
	wg2.Add(1)
	counts := make([]int, Width*Height)

	go func() {
		reduce(results, counts)
		wg2.Done()
	}()

	wg.Wait()
	close(results)
	wg2.Wait()

	img := image.NewRGBA64(image.Rect(0, 0, Width, Height))
	maxCount := 0
	for _, c := range counts {
		if c > maxCount {
			maxCount = c
		}
	}
	fmt.Println(maxCount)
	fmt.Println(minY, maxY, minYP, maxYP)
	for i, c := range counts {
		y := math.MaxUint16 * c * 4 / (maxCount + 1)
		if y > math.MaxUint16 {
			y = math.MaxUint16
		}
		img.Set(i%Width, i/Width, color.Gray16{Y: uint16(y)})
	}

	out, err := os.Create(fmt.Sprintf("out-%d.png", time.Now().Unix()))
	if err != nil {
		panic(err)
	}

	err = png.Encode(out, img)
	if err != nil {
		panic(err)
	}
}
