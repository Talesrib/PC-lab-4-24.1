package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// read a file from a filepath and return a slice of bytes
func readFile(filePath string) ([]byte, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file %s: %v", filePath, err)
		return nil, err
	}
	return data, nil
}

type Sums struct {
	Sum int
	Path string
}

// sum all bytes of a file
func sum(filePath string, s chan Sums) (int, error) {
	data, err := readFile(filePath)
	if err != nil {
		return 0, err
		}
		
	_sum := 0
	for _, b := range data {
		_sum += int(b)
	}
	fmt.Printf("%s: %v\n", filePath, _sum)
	s <- Sums{Sum:_sum, Path: filePath}
	return _sum, nil
}

// print the totalSum for all files and the files with equal sum
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <file1> <file2> ...")
		return
	}

	var totalSum int64
	sums := make(map[int][]string)
	s := make(chan Sums)
	nt := 0
	for _, path := range os.Args[1:] {
		go sum(path, s)
		nt+= 1
	}

	for i := 0; i < nt; i++ {
		x := <-s
		totalSum += int64(x.Sum)
		sums[x.Sum] = append(sums[x.Sum], x.Path)
	}

	fmt.Println(totalSum)

	for sum, files := range sums {
		if len(files) > 1 {
			fmt.Printf("Sum %d: %v\n", sum, files)
		}
	}
}
