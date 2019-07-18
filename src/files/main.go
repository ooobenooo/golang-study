package main

import (
	"fmt"
	"os"
)

func main() {
	os.Mkdir("test", 0777)
	os.MkdirAll("test/t1/t11", 0777)

	os.Remove("test/t1/t11")
	os.RemoveAll("test")

	fout, err := os.Create("test.txt") // 只能写，不能读
	if err != nil {
		fmt.Println(err)
	}
	defer fout.Close()
	fout.WriteString("hello world!\n")
	fout.Write([]byte("I'm Ben\n"))

	fi, err := os.Open("test.txt") // 只读
	if err != nil {
		fmt.Println(err)
	}
	defer fi.Close()
	buf := make([]byte, 1024)
	for {
		i, _ := fi.Read(buf)
		if i == 0 { // end of file
			break
		}
		os.Stdout.Write(buf[:i])
	}
	os.Remove("test.txt")
}
