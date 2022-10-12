package readinput

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

var WorkPath = "E:\\workspace\\gowork\\src\\github.com\\jc0803kevin\\go-study\\readinput\\%s"

func IoUtil() {

	buf, err := ioutil.ReadFile(fmt.Sprintf(WorkPath, "products.txt"))

	if err != nil {
		fmt.Fprintf(os.Stderr, "File Error: %s\n", err)
		// panic(err.Error())
	}

	fmt.Printf("%s\n", string(buf))

	err = ioutil.WriteFile(fmt.Sprintf(WorkPath, "products_copy.txt"), buf, 0644) // oct, not hex
	if err != nil {
		panic(err.Error())
	}

}

func FileInput() {

	inputFile, inputError := os.Open(fmt.Sprintf(WorkPath, "input.txt"))
	if inputError != nil {
		fmt.Printf("An error occurred on opening the inputfile\n" +
			"Does the file exist?\n" +
			"Have you got acces to it?\n")
		return // exit the function on error
	}

	// 关闭文件
	defer inputFile.Close()

	inputReader := bufio.NewReader(inputFile)
	for {
		inputString, readerError := inputReader.ReadString('\n')

		fmt.Printf("The input was: %s", inputString)
		// 读取完了
		if readerError == io.EOF {
			return
		}

	}
}
