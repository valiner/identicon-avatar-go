/**
 * User: coder.sdp@gmail.com
 * Date: 2023/7/18
 * Time: 16:36
 */

package identicon

import (
	"crypto/md5"
	"image/color"
)

const (
	b7 = 2 << 6
)

type Matrix struct {
}

func NewMatrix() *Matrix {
	return new(Matrix)
}

// GetMatrix 获取图片矩阵
func (m *Matrix) GetMatrix(id string) ([][]bool, color.RGBA) {

	hash := md5.Sum([]byte(id))
	color := color.RGBA{
		R: hash[0],
		G: hash[1],
		B: hash[2],
		A: 255,
	}
	arrayBool := make([]bool, len(hash))
	for i := 0; i < len(hash); i++ {
		if hash[i] > b7 {
			arrayBool[i] = true
		}
	}
	leftMatrix := make([][]bool, 5)
	for i := 0; i < len(leftMatrix); i++ {
		leftMatrix[i] = make([]bool, 3)
	}

	for i := 0; i < len(arrayBool)-1; i++ {
		leftMatrix[i/3][i%3] = arrayBool[i]
	}
	// 扩展矩阵
	res := make([][]bool, 5)
	for i := range res {
		res[i] = make([]bool, 5)
		copy(res[i], leftMatrix[i])
	}
	// 镜像反转
	for i := range res {
		for j := 0; j < len(res[i])/2; j++ {
			k := len(res[i]) - j - 1
			res[i][k] = res[i][j]
		}
	}
	return res, color
}
