package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func TestFile() {
	filePath := "D:/test/a.dat"
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("read file err ", err.Error())
		return
	}
	var newfile []byte
	for _, val := range data {
		val ^= 0x2a
		newfile = append(newfile, val)
	}

	file, err := os.Create("D:/a.jpg")
	defer file.Close()
	file.Write(newfile)
	return

	//fmt.Println(string(data))
}

