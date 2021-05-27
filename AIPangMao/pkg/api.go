package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gotools/AIPangMao/logger"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

//const robotI = "http://10.0.9.101:8000/robot-zhaogang/http/RobotService?question=%s&platform=web&userId=yingchun.xu&brand=%s&location=sh"
const robotI = "http://10.80.16.101:8000/robot-zhaogang/http/RobotService?question=%s&platform=web&userId=yingchun.xu&brand=%s&location=sh"

const robotYGetToken = "http://faqrobot.zhaogang.com/token/getToken?appId=qT8zVKOqyR3NM86axA&secret=vMkHtUysrT6CBF4FA4EF"
const robotYGetAnswer = "http://faqrobot.zhaogang.com/servlet/apichat/v4?question=%s?&s=aq&access_token=%s&sysNum=%s&original=true&sourceId=%d&clientId=123&userId=123"
const robotYOffline = "http://faqrobot.zhaogang.com/servlet/apichat/v4?access_token=%s&sysNum=%s&s=offline&sourceId=%d&clientId=123&userId=123"

var robotYAccessToken = ""
var sysNum = ""

type JsonText struct {
	Content string `json:"content"`
}

type RebotGuideText struct {
	Question string `json:"question"`
	Seq      int    `json:"seq"`
}

type RebotGuide struct {
	List []RebotGuideText
}

type RobotReply struct {
	MsgType string   `json:"msgtype"`
	Text    JsonText `json:"text"`

	Guide      RebotGuide `json:"guide"`
	AnswerType int        `json:"answerType"`
}

type JsonData struct {
	Reply []RobotReply `json:"robotReply"`
}

type HttpAccessToken struct {
	Status      int    `json:"status"`
	AccessToken string `json:"access_token"`
	SysNum      string `json:"sysNum"`
}

type HttpResponse struct {
	Status int      `json:"status"`
	Data   JsonData `json:"data"`
}

// 从云问获取答案
// role:1代表？、2代表？、3代表？
func GetAnswerFromY(question string, role int) (string, string, error) {
	// 获取accessToken
	if robotYAccessToken == "" {
		res, err := http.Get(robotYGetToken)
		if err != nil {
			return "", "", err
		}
		if res.StatusCode != 200 {
			return "", "", errors.New(res.Status)
		}
		data, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", "", err
		}
		token := &HttpAccessToken{}
		err = json.Unmarshal(data, token)
		if err != nil {
			return "", "", err
		}
		robotYAccessToken = token.AccessToken
		sysNum = token.SysNum
	}

	// 只需要对question做url编码
	urlStr := fmt.Sprintf(robotYGetAnswer, url.QueryEscape(question), robotYAccessToken, sysNum, role)
	res, err := http.Get(urlStr)
	if err != nil {
		return "", "", err
	}

	if res.StatusCode != 200 {
		return "", "", errors.New(res.Status)
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", "", err
	}

	// 解析结果
	jsonData := &HttpResponse{}
	err = json.Unmarshal(data, jsonData)
	if err != nil {
		return "", "", err
	}
	if len(jsonData.Data.Reply) == 0 {
		return "", "", errors.New("not have any answer，jsonData.Data.Reply[] len is 0")
	}

	// 建议问当成错误，直接返回问题列表
	if jsonData.Data.Reply[0].AnswerType == -1 {
		guildQuestion := "建议问："
		tempList := jsonData.Data.Reply[0].Guide.List
		for i := range tempList {
			guildQuestion += fmt.Sprintf(" %d.%s", tempList[i].Seq, tempList[i].Question)
		}
		return "建议问", "", errors.New(guildQuestion)
	}

	if jsonData.Data.Reply[0].MsgType == "text" {
		// 处理答案中带''的问题
		str := strings.ReplaceAll(jsonData.Data.Reply[0].Text.Content, "''", "")
		return "普通答案", delEscape(str), nil
	}
	return "普通答案", "", errors.New(string(data))
}

func delEscape(str string) string {
	var buf bytes.Buffer
	for _, c := range str {
		if c == '\n' {
			continue
		} else if c == '\t' {
			continue
		} else if c == '\r' {
			continue
		}
		buf.WriteRune(c)
	}
	return buf.String()
}

// 从小i机器人处获取答案
func GetAnswerFromI(question string, role int) (string, error) {
	url := fmt.Sprintf(robotI, question, "VIP")
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", nil
	}

	type tempResult struct {
		Content string `json:"content"`
	}

	result := &tempResult{}
	err = json.Unmarshal(data, result)
	if err != nil {
		return "", nil
	}

	// 对结果处理
	// 原始：<HTML>&lt;TEXT&gt;${E_001_001|煤炭采选|钢铁行业|水泥建材}&lt;/TEXT&gt;&lt;URL&gt;http://m.data.eastmoney.com/zjlx/list/bk#hy&lt;/URL&gt;</HTML>\r\n\r\n<br>以上答案是否解决了您的问题：</br>[link submit=\"faqvote:001544580018804618280050568f2589 1 QbnJ18q98MH3tq/H6b/2Pw== vfHM7EG5ydfKvfDB97avx+m/9j8= PEhUTUw+Jmx0O1RFWFQmZ3Q7JHtFXzAwMV8wMDF8w7rMv7LJ0aF8uNbM+tDQ0rV8y67E4L2ossR9Jmx0Oy9URVhUJmd0OyZsdDtVUkwmZ3Q7aHR0cDovL20uZGF0YS5lYXN0bW9uZXkuY29tL3pqbHgvbGlzdC9iayNoeSZsdDsvVVJMJmd0OzwvSFRNTD4=\"]解决[/link] [link submit=\"faqvote:001544580018804618280050568f2589 2 QbnJ18q98MH3tq/H6b/2Pw== vfHM7EG5ydfKvfDB97avx+m/9j8= PEhUTUw+Jmx0O1RFWFQmZ3Q7JHtFXzAwMV8wMDF8w7rMv7LJ0aF8uNbM+tDQ0rV8y67E4L2ossR9Jmx0Oy9URVhUJmd0OyZsdDtVUkwmZ3Q7aHR0cDovL20uZGF0YS5lYXN0bW9uZXkuY29tL3pqbHgvbGlzdC9iayNoeSZsdDsvVVJMJmd0OzwvSFRNTD4=\"]未解决[/link]\r\n
	// 去掉\r\n\r\n后面的内容
	index := strings.Index(result.Content, "[link submit=")
	subIndex := len("\r\n\r\n<br>以上答案是否解决了您的问题：</br>")
	if index >= subIndex {
		return delEscape(result.Content[0 : index-subIndex]), nil
	} else {
		return delEscape(result.Content), nil
	}
}

// 手动下线云问机器人（对方API同一个会话调用次数有限制，大概300次左右需要下线一次）
func OfflineY(role int) {
	_, err := http.Get(fmt.Sprintf(robotYOffline, robotYAccessToken, sysNum, role))
	if err != nil {
		logger.Sugar.Error(err)
	} else {
		robotYAccessToken = ""
	}
}
