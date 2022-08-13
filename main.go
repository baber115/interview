package main

import (
	"os"
	"strconv"
)

func main() {
	number := make(map[int]int, 10)
	number = map[int]int{
		1:   30,
		31:  60,
		61:  90,
		91:  120,
		121: 150,
		151: 180,
		181: 194,
	}
	for min, max := range number {
		path := "./" + strconv.Itoa(min) + "-" + strconv.Itoa(max) + "/"
		os.MkdirAll(path, os.ModePerm)
		for i := min; i <= max; i++ {
			fileName := strconv.Itoa(i) + ".go"
			path := path + fileName
			file, _ := os.OpenFile(path, os.O_CREATE, 0644)
			defer file.Close()
			file.WriteString("package main\n")
		}
	}
}
