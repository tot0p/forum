package tools

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().Unix())
}

func CreateImg() []byte {
	width := 420
	height := 420

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	c := color.RGBA{uint8(rand.Intn(150) + 10), uint8(rand.Intn(150) + 10), uint8(rand.Intn(150 + 10)), 0xff}
	white := color.RGBA{200, 200, 200, 0xff}

	for x := 0; x < width; x += 30 {
		for y := 0; y < height; y += 30 {
			if y == 0 || y >= height-30 || x == 0 || x >= width-30 {
				drawCarre(x, y, white, img)
			} else {
				temp := rand.Intn(3)
				switch temp {
				case 0:
					drawCarre(x, y, white, img)
				default:
					drawCarre(x, y, c, img)
				}
			}
		}
	}

	var buff bytes.Buffer

	png.Encode(&buff, img)

	// encodedString := base64.StdEncoding.EncodeToString(buff.Bytes())

	return buff.Bytes()
}

func drawCarre(startx, starty int, c color.Color, img *image.RGBA) {
	for x := startx; x < startx+30; x++ {
		for y := starty; y < starty+30; y++ {
			img.Set(x, y, c)
		}
	}
}

// func CompresseImg(s []byte) ([]byte, error) {

// 	// reader := base64.NewDecoder(base64.StdEncoding, strings.NewReader(string(s)))
// 	// m, formatString, err := image.Decode(reader)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// bounds := m.Bounds()
// 	// fmt.Println(bounds, formatString)

// 	// //Encode from image format to writer

// 	// err = png.Encode(f, m)
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// 	return
// 	// }

// 	var buff bytes.Buffer
// 	// b, err := base64.StdEncoding.DecodeString(string(s))
// 	// if err != nil {
// 	// 	return nil, errors.New("error decode base 64")
// 	// }
// 	_, err := buff.Write(s)
// 	if err != nil {
// 		return nil, errors.New("error write bytes")
// 	}
// 	img, _, err := image.Decode(&buff)
// 	if err != nil {
// 		return nil, errors.New("error decode img")
// 	}
// 	img = resize(img, 412, 412)

// 	var buff2 bytes.Buffer

// 	png.Encode(&buff2, img)

// 	// encodedString := base64.StdEncoding.EncodeToString(buff2.Bytes())

// 	return buff2.Bytes(), nil
// }

// func resize(img image.Image, width, height int) {
// 	xMax := img.Bounds().Max.X
// 	xMin := img.Bounds().Min.X
// 	yMax := img.Bounds().Max.Y
// 	yMin := img.Bounds().Min.Y
// 	if xMax-xMin != yMax-yMin {
// 		imgRogn := image.NewRGBA(image.Rect(0, 0, width, yMax-yMin))
// 		t := xMax - xMin - width/2
// 		for x := t; x < t+width; x++ {
// 			for y := yMin; x < yMin; y++ {

// 			}
// 		}
// 	}
// 	// imgRes := image.NewRGBA(image.Rect(0, 0, width, height))

// }
