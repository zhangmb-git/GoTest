package app

import (
	"container/list"
	"fmt"
	"gotools/AIPangMao/logger"
	"log"
	"strconv"
	"strings"
	"sync"
)

var mapRequest map[int64]*AskRequest
var RequestChan chan *AskRequest   // 发送数据通道，无缓冲区
var ResponseChan chan *AskResponse // 发送数据通道，无缓冲区

var lock = sync.Mutex{}

func RunLoop() {
	//RequestChan = make(chan *AskRequest, 10)
	//ResponseChan = make(chan *AskResponse, 10)
	//go coHandleRequest()
	//go CoHandleResponse()
	return
}

func AskAllQuestions(asklist *list.List) {

	for item := asklist.Front(); item != nil; item = item.Next() {
		elem, ok := (item.Value).(OneDictionary)
		if ok {
			askReq := &AskRequest{
				AskData: AskReqData{
					DeviceCode: "11111",
					Platform:   "web",
					Question:   "查基价",
					UserId:     "100021",
					UserLevel:  "NORMAL",
				},
			}
			askReq.SetQuestion(elem)
			strPath := ""
			SearchFlag := 0
			dfsFlag, err := strconv.Atoi(TOPIC["dfsFlag"])
			if err != nil {
				panic(err)
				return
			}
			if dfsFlag == 1 {
				DFS_QuestionsWithFlag(askReq, strPath, &SearchFlag)
			} else {
				DFS_Questions(askReq, strPath)
			}

		} else {
			logger.Sugar.Error("e is not an OneDictionary")
		}
	}
	return
}

func AskQuestions(questionList *list.List) {
	for item := questionList.Front(); item != nil; item = item.Next() {
		elem, ok := (item.Value).(string)
		if ok {
			arrQuestions := strings.Split(elem, "|")
			for _, val := range arrQuestions {
				askReq := &AskRequest{
					AskData: AskReqData{
						DeviceCode: "11111",
						Platform:   "web",
						Question:   "查基价",
						UserId:     "100021",
						UserLevel:  "NORMAL",
					},
				}
				askReq.SetQuestionStr(val)
				SendRequest(askReq, elem)
			}

		} else {
			logger.Sugar.Error("e is not an OneDictionary")
		}
	}
}

func DFS_Questions(req *AskRequest, path string) {

	askResp := SendRequest(req, path)

	newpath := path + req.AskData.Question + "|"

	if askResp.BodyData.ResCode == 0 {
		if len(askResp.BodyData.AnswerInfos) > 0 && askResp.BodyData.AnswerInfos[0].IsContinuous == 1 && askResp.BodyData.AnswerInfos[0].IsLast == 0 {
			for _, val := range askResp.BodyData.AnswerInfos[0].Tips {
				askRequest := &AskRequest{}
				askRequest.AskData = AskReqData{
					DeviceCode: "11111",
					Platform:   "web",
					Question:   "查基价",
					UserId:     "100021",
					UserLevel:  "NORMAL",
				}
				askRequest.SetQuestionStr(val)
				newpath = path + req.AskData.Question + "|"
				DFS_Questions(askRequest, newpath)
			}

			if len(askResp.BodyData.AnswerInfos[0].Tips) == 0 {
				userID, err := strconv.ParseInt(req.AskData.UserId, 10, 64)
				if err != nil {
					panic(err)
					return
				}
				answer := askResp.BodyData.AnswerInfos[0].Answer
				if TOPIC["flag"] == "1" {
					WriteAskLog(userID, newpath, answer, 2)
				} else {
					InsertAskLog(userID, newpath, answer, 2)
				}
			}

		} else {
			if len(askResp.BodyData.AnswerInfos) > 0 && askResp.BodyData.AnswerInfos[0].IsLast == 1 {
				//logger.Sugar.Infof("用户[%s]查询成功，问题[%s]，收到回复[%s]", askResp.BodyData.UserInfo.UserId,
				//	askResp.BodyData.AnswerInfos[0].Question, askResp.BodyData.AnswerInfos[0].Answer)
				answer := askResp.BodyData.AnswerInfos[0].Answer
				if askResp.BodyData.AnswerInfos[0].IsChart == 1 {
					answer = "图表"
				}
				userID, err := strconv.ParseInt(askResp.BodyData.UserInfo.UserId, 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				if TOPIC["flag"] == "1" {
					WriteAskLog(userID, newpath, answer, 1)
				} else {
					InsertAskLog(userID, newpath, answer, 1)
				}
			}
		}
	} else {
		//logger.Sugar.Errorf("用户[%s]查询失败，问题[%s], 错误码[%d]错误原因[%s]", askResp.BodyData.UserInfo.UserId,
		//	req.AskData.Question, askResp.BodyData.ResCode, askResp.BodyData.ErrorMsg)
		userID, err := strconv.ParseInt(req.AskData.UserId, 10, 64)
		if err != nil {
			panic(err)
			return
		}

		answer := fmt.Sprintf("错误码[%d]错误原因[%s]", askResp.BodyData.ResCode, askResp.BodyData.ErrorMsg)
		if TOPIC["flag"] == "1" {
			WriteAskLog(userID, newpath, answer, 0)
		} else {
			InsertAskLog(userID, newpath, answer, 0)
		}

	}
	RestoreQuestion(path)
	return
}

func DFS_QuestionsWithFlag(req *AskRequest, path string, flag *int) {
	if *flag == 1 {
		return
	}
	askResp := SendRequest(req, path)
	if askResp.success == false {
		return
	}
	//bWriteLog := false
	//answer := ""
	newpath := path + req.AskData.Question + "|"

	if askResp.BodyData.ResCode == 0 {
		if len(askResp.BodyData.AnswerInfos) > 0 && askResp.BodyData.AnswerInfos[0].IsContinuous == 1 && askResp.BodyData.AnswerInfos[0].IsLast == 0 {
			for _, val := range askResp.BodyData.AnswerInfos[0].Tips {
				askRequest := &AskRequest{}
				askRequest.AskData = AskReqData{
					DeviceCode: "11111",
					Platform:   "web",
					Question:   "查基价",
					UserId:     "100021",
					UserLevel:  "NORMAL",
				}
				askRequest.SetQuestionStr(val)
				newpath = path + req.AskData.Question + "|"
				DFS_QuestionsWithFlag(askRequest, newpath, flag)
			}

			if len(askResp.BodyData.AnswerInfos[0].Tips) == 0 {
				userID, err := strconv.ParseInt(req.AskData.UserId, 10, 64)
				if err != nil {
					panic(err)
					return
				}
				answer := askResp.BodyData.AnswerInfos[0].Answer
				if TOPIC["flag"] == "1" {
					WriteAskLog(userID, newpath, answer, 2)
				} else {
					InsertAskLog(userID, newpath, answer, 2)
				}
			}

		} else {
			if len(askResp.BodyData.AnswerInfos) > 0 && askResp.BodyData.AnswerInfos[0].IsLast == 1 {
				//logger.Sugar.Infof("用户[%s]查询成功，问题[%s]，收到回复[%s]", askResp.BodyData.UserInfo.UserId,
				//	askResp.BodyData.AnswerInfos[0].Question, askResp.BodyData.AnswerInfos[0].Answer)
				*flag = 1
				//newpath := path + req.AskData.Question + "|"
				//newpath = strings.TrimRight(newpath, "|")
				answer := askResp.BodyData.AnswerInfos[0].Answer
				if askResp.BodyData.AnswerInfos[0].IsChart == 1 {
					answer = "图表"
				}

				userID, err := strconv.ParseInt(askResp.BodyData.UserInfo.UserId, 10, 64)
				if err != nil {
					log.Fatal(err)
				}

				if TOPIC["flag"] == "1" {
					WriteAskLog(userID, newpath, answer, 1)
				} else {
					InsertAskLog(userID, newpath, answer, 1)
				}
				return
			}
		}
	} else {
		//logger.Sugar.Errorf("用户[%s]查询失败，问题[%s], 错误码[%d]错误原因[%s]", req.AskData.UserId,
		//req.AskData.Question, askResp.BodyData.ResCode, askResp.BodyData.ErrorMsg)
		userID, err := strconv.ParseInt(req.AskData.UserId, 10, 64)
		if err != nil {
			panic(err)
			return
		}

		answer := fmt.Sprintf("错误码[%d]错误原因[%s]", askResp.BodyData.ResCode, askResp.BodyData.ErrorMsg)
		if TOPIC["flag"] == "1" {
			WriteAskLog(userID, newpath, answer, 0)
		} else {
			InsertAskLog(userID, newpath, answer, 0)
		}
	}

	RestoreQuestion(path)
	return
}

func RestoreQuestion(path string) {
	logger.Sugar.Infof("restore path is %s", path)
	if path == "" {
		return
	}

	path = strings.TrimRight(path, "|")
	arrQuestions := strings.Split(path, "|")
	for _, val := range arrQuestions {
		askReq := &AskRequest{
			AskData: AskReqData{
				DeviceCode: "11111",
				Platform:   "web",
				Question:   "查基价",
				UserId:     "100021",
				UserLevel:  "NORMAL",
			},
		}
		askReq.SetQuestionStr(val)
		SendRequest(askReq, path)
	}
	return
}

func SendRequest(req *AskRequest, path string) *AskResponse {
	// request :=&AskRequest{}
	logger.Sugar.Infof("path is %s,question is %s", path, req.AskData.Question)
	url := ""
	env, err := strconv.Atoi(TOPIC["env"])
	if err != nil {
		panic(err)
		return nil
	}
	switch env {
	case 0:
		url = "http://zgim.zhaogangrentest.com/sqa/sendMsg"
		break
	case 1:
		url = "http://zgim.zhaogangrenmit.com/sqa/sendMsg"
		break
	case 2:
		url = "http://zgim.zhaogangrenuat.com/sqa/sendMsg"
		break
	case 3:
		url = "http://zgim.zhaogangrenprd.com/sqa/sendMsg"
		break
	default:
		url = "http://zgim.zhaogangrentest.com/sqa/sendMsg"
		break
	}

	contentType := "application/json;charset=utf-8"
	resp := Post(url, req.AskData, contentType)
	//fmt.Println(resp)
	askResp := &AskResponse{}
	askResp.ParseBody([]byte(resp))

	return askResp
}

func CoHandleResponse() {
	// for {
	// 	// 引入chan，比sleep，锁，状态判断优雅
	// 	rsp := <-ResponseChan

	// 	for _, val := range rsp.BodyData.AnswerInfos[0].Tips {
	// 		askRequest := &AskRequest{}
	// 		askRequest.AskData = AskReqData{
	// 			DeviceCode: "11111",
	// 			Platform:   "web",
	// 			Question:   "查基价",
	// 			UserId:     "100021",
	// 			UserLevel:  "NORMAL",
	// 		}
	// 		askRequest.SetQuestionStr(val)
	// 		RequestChan <- askRequest
	// 	}
	// }
}

func coHandleRequest() {
	//for {
	// 引入chan，比sleep，锁，状态判断优雅
	//req := <-RequestChan
	// 加入到响应确认列表
	// if req.NeedResponse {
	// 	mapRequest[req.SeqNum] = req
	// }
	//lock.Lock()
	//SendRequest(req)
	//lock.Unlock()
	//}
}
