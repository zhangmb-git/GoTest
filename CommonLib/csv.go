package commonlib

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func InitCSV(path string) {

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

//
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
