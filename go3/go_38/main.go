package ch3

import (
	"image"
	"image/color"
	"image/jpeg"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
	"strconv"
)

const prec = 100

func init() {
	for i := range colors {
		colors[i] = interp(i, len(colors), len(rainbow)-1)
	}
}

func main() {
	ex8_Rat()
	ex8_BigFloat()
	//ex8_complex64()
	//ex8_complex128()
}

func ex8_Rat() {
	const width, height = 400, 400
	var (
		xmin = big.NewRat(-2, 1)
		ymin = big.NewRat(-2, 1)
		xmax = big.NewRat(2, 1)
		ymax = big.NewRat(2, 1)
	)
	if len(os.Args) == 4 {
		xmin64, _ := strconv.ParseFloat(os.Args[1], 64)
		ymin64, _ := strconv.ParseFloat(os.Args[2], 64)
		xmax64, _ := strconv.ParseFloat(os.Args[3], 64)

		xmin.SetFloat64(xmin64)
		ymin.SetFloat64(ymin64)
		xmax.SetFloat64(xmax64)
		t := new(big.Rat)
		ymax.Add(ymin, t.Sub(xmax, xmin))
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y, ydiff := new(big.Rat), new(big.Rat)
		t1, t2, t3 := new(big.Rat), new(big.Rat), new(big.Rat)

		// y := float64(py)/float64(height)*(ymax-ymin) + ymin
		y.Add(t1.Mul(t2.Quo(new(big.Rat).SetInt64(int64(py)), new(big.Rat).SetInt64(height)), t3.Sub(ymax, ymin)), ymin)

		// ydiff := (ymax - ymin) / float64(4*height)
		ydiff.Quo(t1.Sub(ymax, ymin), t2.Mul(big.NewRat(4, 1), new(big.Rat).SetInt64(height)))

		for px := 0; px < width; px++ {
			x, xdiff := new(big.Rat), new(big.Rat)

			// x := float64(px)/float64(width)*(xmax-xmin) + xmin
			x.Add(t1.Mul(t2.Quo(new(big.Rat).SetInt64(int64(px)), new(big.Rat).SetInt64(width)), t3.Sub(xmax, xmin)), xmin)

			// xdiff := (xmax - xmin) / float64(4*width)
			xdiff.Quo(t1.Sub(xmax, xmin), t2.Mul(new(big.Rat).SetInt64(4), new(big.Rat).SetInt64(width)))

			var colVec []color.Color
			colVec = append(colVec, mandelbrotRat(t1.Sub(x, xdiff), t2.Sub(y, ydiff)))
			colVec = append(colVec, mandelbrotRat(t1.Sub(x, xdiff), t2.Add(y, ydiff)))
			colVec = append(colVec, mandelbrotRat(t1.Add(x, xdiff), t2.Sub(y, ydiff)))
			colVec = append(colVec, mandelbrotRat(t1.Add(x, xdiff), t2.Add(y, ydiff)))

			// Image point (px, py) represents complex value z.
			img.Set(px, py, average(colVec))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	f, _ := os.Create("ch3/ex8_Rat.jpg")

	jpeg.Encode(f, img, nil)
	f.Close()
}

func mandelbrotRat(zRe *big.Rat, zIm *big.Rat) color.Color {
	const iterations = 10
	const contrast = 7

	vRe, vIm := new(big.Rat), new(big.Rat)
	for n := uint8(0); n < iterations; n++ {
		t1, t2, t3 := new(big.Rat), new(big.Rat), new(big.Rat)

		// Real part
		vReNew := new(big.Rat)
		t1.Mul(vRe, vRe)
		t2.Mul(vIm, vIm)
		t3.Sub(t1, t2)
		vReNew.Add(t3, zRe)

		// Imaginary part
		vImNew := new(big.Rat)
		t1.Mul(vRe, vIm)
		t2.Mul(big.NewRat(2, 1), t1)
		vImNew.Add(t2, zIm)

		vRe.Set(vReNew)
		vIm.Set(vImNew)

		// Absolute value squared, hence comparing to 4 instead of 2
		t3.Add(t1.Mul(vRe, vRe), t2.Mul(vIm, vIm))

		if cmp := t3.Cmp(big.NewRat(4, 1)); cmp == 1 {
			return colors[255-contrast*n]
		}
	}
	return color.Black
}

func ex8_BigFloat() {
	const width, height = 400, 400
	var (
		xmin = big.NewFloat(-2).SetPrec(prec)
		ymin = big.NewFloat(-2).SetPrec(prec)
		xmax = big.NewFloat(2).SetPrec(prec)
		ymax = big.NewFloat(2).SetPrec(prec)
	)
	if len(os.Args) == 4 {
		xmin.Parse(os.Args[1], 10)
		ymin.Parse(os.Args[2], 10)
		xmax.Parse(os.Args[3], 10)
		t := new(big.Float).SetPrec(prec)
		t.Sub(xmax, xmin)
		ymax.Add(ymin, t)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y, ydiff := new(big.Float).SetPrec(prec), new(big.Float).SetPrec(prec)
		t1, t2, t3 := new(big.Float).SetPrec(prec), new(big.Float).SetPrec(prec), new(big.Float).SetPrec(prec)

		// y := float64(py)/float64(height)*(ymax-ymin) + ymin
		y.Add(t1.Mul(t2.Quo(big.NewFloat(float64(py)).SetPrec(prec), big.NewFloat(height).SetPrec(prec)), t3.Sub(ymax, ymin)), ymin)

		// ydiff := (ymax - ymin) / float64(4*height)
		ydiff.Quo(t1.Sub(ymax, ymin), t2.Mul(big.NewFloat(4).SetPrec(prec), big.NewFloat(height).SetPrec(prec)))

		for px := 0; px < width; px++ {
			x, xdiff := new(big.Float).SetPrec(prec), new(big.Float).SetPrec(prec)

			// x := float64(px)/float64(width)*(xmax-xmin) + xmin
			x.Add(t1.Mul(t2.Quo(big.NewFloat(float64(px)).SetPrec(prec), big.NewFloat(width).SetPrec(prec)), t3.Sub(xmax, xmin)), xmin)

			// xdiff := (xmax - xmin) / float64(4*width)
			xdiff.Quo(t1.Sub(xmax, xmin), t2.Mul(big.NewFloat(4).SetPrec(prec), big.NewFloat(width).SetPrec(prec)))

			var colVec []color.Color
			colVec = append(colVec, mandelbrot_BigFloat(t1.Sub(x, xdiff), t2.Sub(y, ydiff)))
			colVec = append(colVec, mandelbrot_BigFloat(t1.Sub(x, xdiff), t2.Add(y, ydiff)))
			colVec = append(colVec, mandelbrot_BigFloat(t1.Add(x, xdiff), t2.Sub(y, ydiff)))
			colVec = append(colVec, mandelbrot_BigFloat(t1.Add(x, xdiff), t2.Add(y, ydiff)))

			// Image point (px, py) represents complex value z.
			img.Set(px, py, average(colVec))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	f, _ := os.Create("ch3/ex8_Float.jpg")

	jpeg.Encode(f, img, nil)
	f.Close()
}

func mandelbrot_BigFloat(zRe *big.Float, zIm *big.Float) color.Color {
	const iterations = 200
	const contrast = 7

	vRe, vIm := new(big.Float).SetPrec(prec), new(big.Float).SetPrec(prec)
	for n := uint8(0); n < iterations; n++ {
		t1, t2, t3 := new(big.Float).SetPrec(prec), new(big.Float).SetPrec(prec), new(big.Float).SetPrec(prec)

		// Real part
		vReNew := new(big.Float).SetPrec(prec)
		t1.Mul(vRe, vRe)
		t2.Mul(vIm, vIm)
		t3.Sub(t1, t2)
		vReNew.Add(t3, zRe)

		// Imaginary part
		vImNew := new(big.Float).SetPrec(prec)
		t1.Mul(vRe, vIm)
		t2.Mul(big.NewFloat(2).SetPrec(prec), t1)
		vImNew.Add(t2, zIm)

		vRe.Set(vReNew)
		vIm.Set(vImNew)

		// Absolute value
		t4 := new(big.Float).SetPrec(prec)
		t4.Sqrt(t3.Add(t1.Mul(vRe, vRe), t2.Mul(vIm, vIm)))

		if cmp := t4.Cmp(big.NewFloat(2).SetPrec(prec)); cmp == 1 {
			return colors[255-contrast*n]
		}
	}
	return color.Black
}

func ex8_complex64() {
	const width, height = 1000, 1000
	xmin64, ymin64, xmax64, ymax64 := -2.0, -2.0, 2.0, 2.0
	if len(os.Args) == 4 {
		xmin64, _ = strconv.ParseFloat(os.Args[1], 32)
		ymin64, _ = strconv.ParseFloat(os.Args[2], 32)
		xmax64, _ = strconv.ParseFloat(os.Args[3], 32)
		ymax64 = ymin64 + (xmax64 - xmin64)
	}
	xmin, ymin, xmax, ymax := float32(xmin64), float32(ymin64), float32(xmax64), float32(ymax64)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float32(py)/float32(height)*(ymax-ymin) + ymin
		ydiff := (ymax - ymin) / float32(4*height)
		for px := 0; px < width; px++ {
			x := float32(px)/float32(width)*(xmax-xmin) + xmin
			xdiff := (xmax - xmin) / float32(4*width)

			var colVec []color.Color
			colVec = append(colVec, mandelbrot_complex64(complex(x-xdiff, y-ydiff)))
			colVec = append(colVec, mandelbrot_complex64(complex(x-xdiff, y+ydiff)))
			colVec = append(colVec, mandelbrot_complex64(complex(x+xdiff, y-ydiff)))
			colVec = append(colVec, mandelbrot_complex64(complex(x+xdiff, y+ydiff)))

			// Image point (px, py) represents complex value z.
			img.Set(px, py, average(colVec))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	f, _ := os.Create("ch3/ex8_cmplx64.jpg")

	jpeg.Encode(f, img, nil)
	f.Close()
}

func mandelbrot_complex64(z complex64) color.Color {
	const iterations = 200
	const contrast = 7

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 {
			return colors[255-contrast*n]
		}
	}
	return color.Black
}

func ex8_complex128() {
	const width, height = 1000, 1000
	xmin, ymin, xmax, ymax := -2.0, -2.0, 2.0, 2.0
	if len(os.Args) == 4 {
		xmin, _ = strconv.ParseFloat(os.Args[1], 64)
		ymin, _ = strconv.ParseFloat(os.Args[2], 64)
		xmax, _ = strconv.ParseFloat(os.Args[3], 64)
		ymax = ymin + (xmax - xmin)
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/float64(height)*(ymax-ymin) + ymin
		ydiff := (ymax - ymin) / float64(4*height)
		for px := 0; px < width; px++ {
			x := float64(px)/float64(width)*(xmax-xmin) + xmin
			xdiff := (xmax - xmin) / float64(4*width)

			var colVec []color.Color
			colVec = append(colVec, mandelbrot_complex128(complex(x-xdiff, y-ydiff)))
			colVec = append(colVec, mandelbrot_complex128(complex(x-xdiff, y+ydiff)))
			colVec = append(colVec, mandelbrot_complex128(complex(x+xdiff, y-ydiff)))
			colVec = append(colVec, mandelbrot_complex128(complex(x+xdiff, y+ydiff)))

			// Image point (px, py) represents complex value z.
			img.Set(px, py, average(colVec))
		}
	}
	png.Encode(os.Stdout, img) // NOTE: ignoring errors
	f, _ := os.Create("ch3/ex8_cmplx128.jpg")

	jpeg.Encode(f, img, nil)
	f.Close()
}

func mandelbrot_complex128(z complex128) color.Color {
	const iterations = 200
	const contrast = 7

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 {
			return colors[255-contrast*n]
		}
	}
	return color.Black
}
