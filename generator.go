/**
 * User: coder.sdp@gmail.com
 * Date: 2023/7/18
 * Time: 18:34
 */

package identicon

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"image"
	"image/color"
	"image/png"
	"os"
)

type Generator struct {
	matrix [][]bool
	size   int
	color  color.RGBA
}

func NewGenerator(matrix [][]bool, size int, color color.RGBA) *Generator {
	return &Generator{
		matrix: matrix,
		size:   size,
		color:  color,
	}
}

func (g *Generator) GetBase64() string {
	img := g.getImageResource()
	buf := new(bytes.Buffer)
	png.Encode(buf, img)
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}

func (g *Generator) SaveAvatar(fileName string) error {
	img := g.getImageResource()
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()
	fileBuf := bufio.NewWriter(file)
	err = png.Encode(fileBuf, img)
	if err != nil {
		return err
	}
	return fileBuf.Flush()
}

func (g *Generator) getImageResource() *image.RGBA {
	width, height := g.size, g.size
	white := color.RGBA{R: 255, G: 255, B: 255, A: 255}
	perPixel := g.size / 5
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			row := x / perPixel
			col := y / perPixel
			if row >= len(g.matrix) {
				row = len(g.matrix) - 1
			}
			if col >= len(g.matrix[0]) {
				col = len(g.matrix[0]) - 1
			}
			if g.matrix[row][col] {
				img.Set(x, y, g.color)
			} else {
				img.Set(x, y, white)
			}
		}
	}
	return img
}
