package identicon

import (
	"crypto/md5"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"math/rand"
)

// Rand returns a reference of rand.Rand giving seeds.
func Rand(seeds []byte) *rand.Rand {
	var seed int64
	for _, s := range seeds {
		seed = seed << 8
		seed += int64(s)
	}
	return rand.New(rand.NewSource(seed))
}

// A Data provides to construct identicon.
type Data struct {
	hash          [md5.Size]byte
	step          int
	width, height int
}

// NewData returns a reference of Data.
func NewData(b []byte) *Data {
	return &Data{
		hash:   md5.Sum(b),
		step:   5,
		width:  50,
		height: 50,
	}
}

// NewDataString returns a reference of Data.
func NewDataString(str string) *Data {
	return NewData([]byte(str))
}

// Color returns a reference of image.Uniform belonging to color.RGBA.
func (d Data) Color() *image.Uniform {
	r := Rand(d.hash[:8])
	size := len(d.hash)
	return &image.Uniform{color.RGBA{
		uint8(d.hash[r.Intn(size)])<<2 | 0x30,
		uint8(d.hash[r.Intn(size)])<<2 | 0x30,
		uint8(d.hash[r.Intn(size)])<<2 | 0x30,
		uint8(d.hash[r.Intn(size)])<<2 | 0xf3,
	}}
}

// Draw draws identicon in img.
func (d *Data) Draw(img *image.RGBA) error {
	// Color
	fill := d.Color()

	// Points
	r := Rand(d.hash[8:])
	xhalf := d.step >> 1
	half := d.step * xhalf
	full := d.step * d.step

	points := make([]bool, full)
	for i, lim := 0, 0; i < full && lim < half; i++ {
		x := r.Intn(xhalf)
		y := r.Intn(d.step)
		for _, idx := range [2]int{x + y*d.step, d.step - x - 1 + y*d.step} {
			if !points[idx] {
				points[idx] = true
				lim++
			}
		}
	}

	if d.step%2 == 1 {
		for i := 0; i < d.step; i++ {
			y := r.Intn(d.step)
			idx := xhalf + y*d.step
			points[idx] = true
		}
	}

	bounds := img.Bounds()
	draw.Draw(img, bounds, &image.Uniform{color.White}, image.ZP, draw.Src)

	for idx, p := range points {
		if !p {
			continue
		}

		i, j := idx%d.step, idx/d.step
		x, y := i*d.width, j*d.height
		draw.Draw(img, image.Rect(x, y, x+d.width, y+d.height), fill, image.ZP, draw.Src)
	}

	return nil
}

// Encode writes the identicon image to w in PNG format.
func (d *Data) Encode(w io.Writer) error {
	img := image.NewRGBA(image.Rect(0, 0, d.step*d.width, d.step*d.height))
	if err := d.Draw(img); err != nil {
		return err
	}

	return png.Encode(w, img)
}
