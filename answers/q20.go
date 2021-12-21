package answers

import "fmt"

type Algorithm map[int]bool

type Image map[Vector]int // Vector imported from Q9

func (image Image) Bounds() (int, int, int, int) {
	minx := 99999
	maxx := -99999
	miny := 99999
	maxy := -99999
	for elem := range image {
		if elem.x < minx {
			minx = elem.x
		}
		if elem.x > maxx {
			maxx = elem.x
		}
		if elem.y < miny {
			miny = elem.y
		}
		if elem.y > maxy {
			maxy = elem.y
		}
	}
	return minx, maxx, miny, maxy
}

func (image Image) LitPixels() int {
	count := 0
	for _, j := range image {
		count += j
	}
	return count
}

func (image Image) SurroundingNumber(x int, y int) int {
	// Gets the nine surrouinding values, and converts into int
	return image[Vector{x: x - 1, y: y - 1}]*256 +
		image[Vector{x: x, y: y - 1}]*128 +
		image[Vector{x: x + 1, y: y - 1}]*64 +
		image[Vector{x: x - 1, y: y}]*32 +
		image[Vector{x: x, y: y}]*16 +
		image[Vector{x: x + 1, y: y}]*8 +
		image[Vector{x: x - 1, y: y + 1}]*4 +
		image[Vector{x: x, y: y + 1}]*2 +
		image[Vector{x: x + 1, y: y + 1}]*1
}

func (image Image) Print() {
	minX, maxX, minY, maxY := image.Bounds()
	// img := [][]byte
	for y := minY; y <= maxY; y++ {
		row := []byte{}
		for x := minX; x <= maxX; x++ {
			if image[Vector{x, y}] == 1 {
				row = append(row, '#')
			} else {
				row = append(row, '.')
			}
		}
		fmt.Println(string(row))
	}
}

func (image Image) ApplyAlgorithm(algo Algorithm, step int) Image {
	newImage := Image{}
	minX, maxX, minY, maxY := image.Bounds()

	//Start wide, and shrink with each step
	increment := 1
	if step%2 == 1 {
		increment = 4
	}

	for x := minX - increment; x <= maxX+increment; x++ {
		for y := minY - increment; y <= maxY+increment; y++ {
			imgNumber := image.SurroundingNumber(x, y)
			isColoured := algo[imgNumber]
			if isColoured {
				newImage[Vector{x, y}] = 1
			}
		}
	}
	return newImage
}

func (image *Image) Clip() {
	// Removes the outer border
	minX, maxX, minY, maxY := image.Bounds()
	for vec := range *image {
		if vec.x <= minX+1 || vec.x >= maxX-1 || vec.y <= minY+1 || vec.y >= maxY-1 {
			delete(*image, vec)
		}
	}
}

func ParseImageData(data []string) (Algorithm, Image) {
	algo := Algorithm{}
	for idx, char := range data[0] {
		algo[idx] = char == '#'
	}

	image := Image{}

	for y := 0; y < len(data)-2; y++ {
		row := data[y+2]
		fmt.Println(row)
		for x := 0; x < len(row); x++ {
			if row[x] == '#' {
				image[Vector{x: x, y: y}] = 1
			}
		}
	}

	return algo, image
}

func Day20() []int {
	data := ReadInputAsStr(20)
	algo, image := ParseImageData(data)
	return []int{q20part1(algo, image), q20part2(algo, image)}
}

func q20part1(algo Algorithm, image Image) int {
	image = image.ApplyAlgorithm(algo, 1)
	image.Print()
	image = image.ApplyAlgorithm(algo, 2)
	image.Print()
	image.Clip()
	return image.LitPixels()
}

func q20part2(algo Algorithm, image Image) int {
	for i := 1; i <= 50; i++ {
		image = image.ApplyAlgorithm(algo, i)
		// Because it flashes on
		if i%2 == 0 {
			image.Clip()
		}
	}
	image.Print()
	return image.LitPixels()
}
