/**
 * User: coder.sdp@gmail.com
 * Date: 2023/7/18
 * Time: 16:28
 */

package identicon

import "fmt"

var i *Identicon

func init() {
	i = New()
}

type Identicon struct {
	matrix *Matrix
}

func New() *Identicon {
	i := new(Identicon)
	i.matrix = NewMatrix()
	return i
}

func SaveAvatar(id string, size int, fileName string) error {
	return i.saveAvatar(id, size, fileName)
}
func (i *Identicon) saveAvatar(id string, size int, fileName string) (err error) {
	matrix, color := i.matrix.GetMatrix(id)
	return NewGenerator(matrix, size, color).SaveAvatar(fileName)
}

func GetAvatarDataUri(id string, size int) string {
	return i.getAvatarDataUri(id, size)
}
func (i *Identicon) getAvatarDataUri(id string, size int) string {
	matrix, color := i.matrix.GetMatrix(id)
	return fmt.Sprintf("data:image/png;base64,%s", NewGenerator(matrix, size, color).GetBase64())
}
