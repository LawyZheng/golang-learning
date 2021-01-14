package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func cat(r *bufio.Reader) {
	for {
		line, err := r.ReadString('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Printf("Reading failed. Error: %s\n", err)
			return
		}

		writer := bufio.NewWriter(os.Stdout)
		writer.WriteString(string(line))
		writer.Flush()
		// fmt.Fprintf(os.Stdout, "%s", line)
	}

}

func main() {

	flag.Parse()
	//命令行未传参，从标准输入读取
	if flag.NArg() == 0 {
		cat(bufio.NewReader(os.Stdin))
	}

	for i := 0; i < flag.NArg(); i++ {
		fileObj, err := os.Open(flag.Arg(i))
		if err != nil {
			fmt.Printf("Reading from %s file failed. Error: %s\n", flag.Arg(i), err)
			continue
		}
		cat(bufio.NewReader(fileObj))
	}

}
