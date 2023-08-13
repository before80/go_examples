// Note:
// This code is from https://gobyexample.com.
package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	f, err := os.CreateTemp(".", "sample")
	check(err)

	fmt.Println("Temp file name:", f.Name()) // Temp file name: .\sample548580023

	//defer os.Remove(f.Name())

	_, err = f.Write([]byte{1, 2, 3, 4})
	check(err)

	dname, err := os.MkdirTemp(".", "sampledir")
	check(err)
	fmt.Println("Temp dir name:", dname) // Temp dir name: .\sampledir1004283032

	//defer os.RemoveAll(dname)

	fname := filepath.Join(dname, "file1")
	err = os.WriteFile(fname, []byte{1, 2}, 0666)
	check(err)
}
