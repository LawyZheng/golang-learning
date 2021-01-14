package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func readFileByOS() {
	fileObj, err := os.Open("./copy.go")
	if err != nil {
		fmt.Printf("Error ocurred: %s\n", err)
		return
	}

	defer fileObj.Close()

	// tmp := make([128]string, 128)
	var tmp [128]byte

	for {
		n, err := fileObj.Read(tmp[:])
		if err == io.EOF {
			fmt.Println("Finished reading.")
			return
		}
		if err != nil {
			fmt.Printf("Error ocurred: %s\n", err)
			return
		}

		fmt.Print(string(tmp[:n]))
	}
}

func readFileByBufio() {
	fileObj, err := os.Open("./copy.go")
	if err != nil {
		fmt.Printf("Error Occurred: %s\n", err)
		return
	}

	reader := bufio.NewReader(fileObj)
	// 中文
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("Error Occured: ", err)
			return
		}

		fmt.Print(string(line))
	}
}

func readFileByIoutil() {
	content, err := ioutil.ReadFile("./copy.go")
	if err != nil {
		fmt.Println("Error Occured: ", err)
		return
	}

	fmt.Print(string(content))

}

func copy(src, dst string) {
	srcFile, err := os.Open(src)
	if err != nil {
		fmt.Println("Error occurred when open the source file.")
		return
	}

	defer srcFile.Close()

	dstFile, err := os.OpenFile(dst, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0755)
	if err != nil {
		fmt.Println("Error occurred when open the destinaton file.")
		return
	}

	defer dstFile.Close()

	reader := bufio.NewReader(srcFile)
	writer := bufio.NewWriter(dstFile)

	// srcFile.Read()
	// reader.Read()

	io.Copy(writer, reader)
	writer.Flush()
	// io.Copy()

}

func main() {
	copy("./snow.go", "test.txt")
}
