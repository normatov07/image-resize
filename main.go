package main

import (
	"bufio"
	"image"
	"image/color"
	"image/jpeg"
	_ "image/jpeg"
	_ "image/png"
	"log"
	"os"

	"golang.org/x/image/draw"
)

const (
	SIZE = 200
)

func main() {
	// TestImg()
	// grid := loadImage("img.jpg")
	// log.Println("xlen: ", len(grid), "ylen:", len(grid[0]))
	// // flip(grid)
	// grid = resize(grid, SIZE)
	// log.Println("xlen: ", len(grid), "ylen:", len(grid[0]))
	// resizeDraw(5)
	// grid = crop(grid, 200)
	// data, _ := json.Marshal(grid)
	// reader := bytes.NewReader(data)
	// saveImage("newImage.jpg", grid, SIZE)
	dir, _ := os.Getwd()
	log.Println(dir)
}

func loadImage(imagePath string) (grid [][]color.Color) {
	file, err := os.Open(imagePath)

	if err != nil {
		log.Println("Cannot read file:", err)
	}
	defer file.Close()
	rdr := bufio.NewReader(file)

	img, _, err := image.Decode(rdr)
	if err != nil {
		log.Println("Can not decode file: ", err)
	}
	size := img.Bounds().Size()
	for i := 0; i < size.X; i++ {
		var y []color.Color
		for j := 0; j < size.Y; j++ {
			y = append(y, img.At(i, j))
			// json, err := json.Marshal(img.At(i, j))
			// if err != nil {
			// 	log.Println("marshal err:", err)

			// }
			// log.Println(string(json))
		}
		grid = append(grid, y)
	}

	return
}

func saveImage(filePath string, grid [][]color.Color, size int) {
	var topX, topY, xlen, ylen int
	if len(grid) > size {
		topX = (len(grid) - size) / 2
		xlen = topX + size
	} else {
		topX = 0
		xlen = len(grid)
	}

	if len(grid[0]) > size {
		topY = (len(grid[0]) - size) / 2
		ylen = topY + size
	} else {
		topY = 0
		ylen = len(grid[0])
	}
	rect := image.Rect(topX, topY, xlen, ylen)
	img := image.NewNRGBA(rect)

	for x := topX; x < xlen; x++ {
		for y := topY; y < ylen; y++ {
			img.Set(x, y, grid[x][y])
		}
	}
	file, err := os.Create(filePath)
	if err != nil {
		log.Println("Can not create file: ", err)
	}
	defer file.Close()

	jpeg.Encode(file, img.SubImage(img.Rect), nil)
}

func flip(grid [][]color.Color) {
	for x := 0; x < len(grid); x++ {
		col := grid[x]
		for y := 0; y < len(col)/2; y++ {
			z := len(col) - y - 1
			col[y], col[z] = col[z], col[y]
		}
	}
}

func resize(grid [][]color.Color, size float64) (resized [][]color.Color) {
	var scale float64
	if len(grid) < len(grid[0]) {
		scale = float64(len(grid)) / size
	} else {
		scale = float64(len(grid[0])) / size
	}

	var xlen, ylen int = int(float64(len(grid)) / scale), int(float64(len(grid[0])) / scale)
	resized = make([][]color.Color, xlen)
	for i := 0; i < len(resized); i++ {
		resized[i] = make([]color.Color, ylen)
	}

	for x := 0; x < xlen; x++ {
		for y := 0; y < ylen; y++ {
			// xp := int(math.Floor(float64(x) / scale))
			// yp := int(math.Floor(float64(y) / scale))
			xp := int(scale * float64(x))
			yp := int(scale * float64(y))
			resized[x][y] = grid[xp][yp]

		}
	}

	return
}

func crop(grid [][]color.Color, size int) (resized [][]color.Color) {

	var xlen, ylen int = size, size
	resized = make([][]color.Color, xlen)
	for i := 0; i < len(resized); i++ {
		resized[i] = make([]color.Color, ylen)
	}
	topX := (len(grid[0]) - size) / 2
	topY := (len(grid) - size) / 2
	for x := topX; x < xlen; x++ {
		for y := topY; y < ylen; y++ {
			// xp := int(math.Floor(float64(x) / scale))
			// yp := int(math.Floor(float64(y) / scale))
			// xp := int(scale * float64(x))
			// yp := int(scale * float64(y))
			resized[x][y] = grid[x][y]

		}
	}

	return
}

func resizeDraw(scale int) {
	input, err := os.Open("img.jpg")
	defer input.Close()
	if err != nil {
		log.Println("err: ", err)
	}
	output, err := os.Create("your_image_resized.jpg")
	if err != nil {
		log.Println("mage create err: ", err)
	}
	defer output.Close()
	// Decode the image (from PNG to image.Image):
	src, err := jpeg.Decode(input)
	if err != nil {
		log.Println("decode err : ", err)
	}
	// Set the expected size that you want:
	dst := image.NewRGBA(image.Rect(0, 0, src.Bounds().Max.X/scale, src.Bounds().Max.Y/scale))

	draw.NearestNeighbor.Scale(dst, dst.Rect, src, src.Bounds(), draw.Over, nil)

	// Encode to `output`:
	jpeg.Encode(output, dst, nil)
}
