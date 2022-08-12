package main

import (
	"embed"
	"fmt"
	"io/fs"
	"io/ioutil"
)

//go:embed hello.txt
var A string

//go:embed hello.txt
var b []byte

//go:embed hello.txt
//go:embed hello2.txt
var f embed.FS

func main() {
	fmt.Println(A)
	fmt.Println(b)

	data, _ := f.ReadFile("hello.txt")
	fmt.Println(string(data))
	data, _ = f.ReadFile("hello2.txt")
	fmt.Println(string(data))

	helloFile, _ := f.Open("hello.txt")
	stat, _ := helloFile.Stat()
	fmt.Println(stat.Name(), stat.Size())

	dirEntries, _ := f.ReadDir("p")
	for _, de := range dirEntries {
		fmt.Println(de.Name(), de.IsDir())
	}

	ps, _ := fs.Sub(f, "p")
	hi, _ := ps.Open("q/hi.txt")
	data1, _ := ioutil.ReadAll(hi)
	fmt.Println(string(data1))
}
