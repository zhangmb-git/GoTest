package commonlib

import (
	"bytes"
	"encoding/json"
	"gotools/AIPangMao/logger"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// 发送GET请求
// url：         请求地址
// response：    请求返回的内容
func Get(url string) string {

	// 超时时间：5秒
	client := &http.Client{Timeout: 30 * time.Second} //30s超时
	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			panic(err)
		}
	}

	return result.String()
}

// 发送POST请求
// url：         请求地址
// data：        POST请求提交的数据
// contentType： 请求体格式，如：application/json
// content：     请求放回的内容
func Post(url string, data interface{}, contentType string) string {

	// 超时时间：50秒
	client := &http.Client{Timeout: time.Second * time.Duration(50)} //30s超时
	jsonStr, err := json.Marshal(data)
	if err != nil {
		logger.Sugar.Error(err)
		panic(err)
		return "err"
	}

	//a := string(jsonStr)
	//fmt.Println(a)
	resp, err := client.Post(url, contentType, bytes.NewBuffer(jsonStr))
	if err != nil {
		logger.Sugar.Errorf("post url err,body:%s", jsonStr)
		panic(err)
		return "err"
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return string(result)
}
