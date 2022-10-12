package readinput

import (
	"fmt"
	"io"
	"os"
)

func CopyFile() {
	_copyFile(fmt.Sprintf(WorkPath, "source.txt"), fmt.Sprintf(WorkPath, "input.txt"))
	fmt.Println("Copy done!")
}

func _copyFile(dstName, srcName string) (written int64, err error) {

	src, err := os.Open(srcName)
	if err != nil {
		fmt.Sprintf("open file err : %s", err)
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		fmt.Sprintf("Create file err : %s", err)
		return
	}

	defer dst.Close()

	return io.Copy(dst, src)

}
