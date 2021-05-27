package app

import (
	"encoding/json"
	"fmt"
	"gotools/AIPangMao/logger"
)

type ResultCallback func(res response)

type request struct {
	url          string
	data         map[string]string
	needResponse bool
	seqNum       int64
}

type AskRequest struct {
	NeedResponse bool
	SeqNum       int64
	AskData      AskReqData
}

func (req *AskRequest) ConstructBody(rspBody *AskResBodyData) {
	req.AskData.Question = rspBody.AnswerInfos[0].Answer
	return
}

func (req *AskRequest) SetQuestion(ans OneDictionary) {
	req.AskData.Question = ans.StandardQuestion
	//req.AskData.UserId = ans.
	return
}

func (req *AskRequest) SetQuestionStr(ans string) {
	req.AskData.Question = ans
	return
}

type AskReqData struct {
	DeviceCode string `json:"deviceCode"`
	Platform   string `json:"platform"`
	Question   string `json:"question"`
	UserId     string `json:"userId"`
	UserLevel  string `json:"userLevel"`
}

type response struct {
	body         string
	data         map[string]string
	answer       string
	isLast       int
	isContinue   int
	needResponse bool
	seqNum       int64
}

type AskResponse struct {
	NeedResponse bool
	SeqNum       int64
	BodyData     AskResBodyData
	success      bool
}

func (res *AskResponse) ParseBody(body []byte) {

	a := string(body)
	if err := json.Unmarshal(body, &res.BodyData); err != nil {
		fmt.Printf("Unmarshal err, %v  %s\n", err, a)
		logger.Sugar.Errorf("unmarshal body err,body:%s err:%s", a, err.Error())
		res.BodyData.ResCode = 1
		res.BodyData.ErrorMsg = err.Error()
		res.success = false
		return
	}
	res.success = true
	//fmt.Println(res.BodyData)
	return
}

type AskResBodyData struct {
	UserInfo    UserInfo
	AnswerInfos []AnswerInfo
	ResCode     int
	ErrorMsg    string
}

type UserInfo struct {
	UserId     string
	DeviceCode string
}

type ExpandInfo struct {
	Text string
	Url  string
}

type CardInfo struct {
	ButtonName     string
	CardCode       string
	CardName       string
	LoginCheck     int
	OperationCode  string
	OperationValue string
	Tips           []string
	TitleContext   string
}

type AnswerInfo struct {
	Question         string
	StandardQuestion string
	QuestionId       string
	IsHit            int
	RelatedQuestions []string
	SessionId        string
	AnswerCoded      string
	Answer           string
	IsChart          int
	ChartId          string
	IsContinuous     int
	IsLast           int
	Tips             []string
	AnswerSource     string
	ExpandInfo       ExpandInfo
	IsCard           int
	CardCode         string
	CardInfo         CardInfo
}
