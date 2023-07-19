/**
 * User: coder.sdp@gmail.com
 * Date: 2023/7/18
 * Time: 17:32
 */

package identicon

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	fmt.Println(GetAvatarDataUri("sdp", 125))
}

func TestGetAvatarDataUri(t *testing.T) {
	type args struct {
		id   string
		size int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test base64",
			args: args{
				id:   "sdp",
				size: 125,
			},
			want: "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAH0AAAB9CAIAAAAA4vtyAAABWUlEQVR4nOzdsQlCMRRAUZFsITidYzmd4ByxtLWI3PD/OQOEcHnlIxlzzst+7s/3qqNej9uqoxa61hc4Kd0bujd0b+je0L2he0P3hu4N3Ru6N3Rv6N7QvaF7Q/eG7g3dG7o3dG/o3tC9oXtD98ZYuCHE78x7Q/eG7g3dG7o3dG/o3tC9oXtD94buDd0bujd0b+je0L2he0P3hu4N3Ru6N3Rv6N7QvTHqC/zdnk83mfeG7g3dG7o3dG/o3tC9oXtD94buDd0bujd0b+je0L2he0P3hu4N3Ru6N3Rv6N7QvaF74/h7S/6J40v3hu4N3Ru6N3Rv6N7QvaF7Q/eG7g3dG7o3dG/o3tC9oXtD94buDd0bujd0b+je0L0x9lzrOfzvdea9oXtD94buDd0bujd0b+je0L2he0P3hu4N3Ru6N3Rv6N7QvaF7Q/eG7g3dG7o3dG/o3vgEAAD//7/GEALp3+eeAAAAAElFTkSuQmCC",
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetAvatarDataUri(tt.args.id, tt.args.size); got != tt.want {
				t.Errorf("GetAvatarDataUri() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveAvatar(t *testing.T) {
	type args struct {
		id       string
		size     int
		fileName string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "sdp",
			args: args{
				id:       "sdp",
				size:     123,
				fileName: "test.png",
			},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := SaveAvatar(tt.args.id, tt.args.size, tt.args.fileName); (err != nil) != tt.wantErr {
				t.Errorf("SaveAvatar() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
