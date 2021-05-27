package app

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
)

func InitCSV() {
	path := TOPIC["csvPath"]
	if IsExist(path) {
		os.Remove(path)
	}

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
		panic(err)
	}
	defer file.Close()
	//file.Seek(0, io.SeekEnd)
	file.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM，防止中文乱码
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}

	//w.Flush()
	return
}

func WriteCSV(path string, row []string) {
	//这样可以追加写
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("can not create file, err is %+v", err)
		panic(err)
	}
	defer file.Close()
	file.Seek(0, io.SeekEnd)
	//w := csv.NewWriter(simplifiedchinese.GB18030.NewDecoder().(file))
	w := csv.NewWriter(file)

	//设置属性
	w.Comma = ','
	w.UseCRLF = true
	//row := []string{"1", "2", "3", "4", "5,6"}
	err = w.Write(row)
	if err != nil {
		log.Fatalf("can not write, err is %+v", err)
	}

	w.Flush()
	return
}

func WriteAskLog(userID int64, question string, answer string, flag int) {
	row := []string{}
	strUID := strconv.FormatInt(userID, 10)
	row = append(row, strUID)
	row = append(row, question)
	row = append(row, answer)
	row = append(row, strconv.Itoa(flag))

	path := TOPIC["csvPath"]
	//filename := time.Now().Format("20060102150405")+".csv"
	WriteCSV(path, row)
	return
}
