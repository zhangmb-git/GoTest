package pkg

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gotools/AIPangMao/logger"
	"io/ioutil"
	"net/http"
	"time"
)

const
(
	zgRobot      = "http://zgim.zhaogangrenmit.com/sqa/sendMsg"
	zgUserId     = 10091009
	zgDeviceCode = 111111
)

type ZgAnswerJsonInfo struct {
	Question string `json:"question"`

	StandardQuestion string   `json:"standardQuestion"`
	QuestionId       string   `json:"questionId"`
	IsHit            int      `json:"isHit"`
	Answer           string   `json:"answer"`
	IsChart          int      `json:"isChart"`
	ChartId          int      `json:"chartId"`
	IsContinuous     int      `json:"isContinuous"`
	IsLast           int      `json:"isLast"`
	Tips             []string `json:"tips"`
	AnswerSource     string   `json:"answerSource"`
}

type ZgAnswerJson struct {
	Answers  []*ZgAnswerJsonInfo `json:"answers"`
	ResCode  int                 `json:"resCode"`
	ErrorMsg string              `json:"errorMsg"`
}

// 从找钢FAQ API获取答案
func GetAnswerFromZG(question string) (*ZgAnswerJson, error) {
	reqJson :=
		fmt.Sprintf(`{"deviceCode":"%d","platform":"web","question":"%s","userId":"%d", "userLevel":"NORMAL"}`, zgDeviceCode, question, zgUserId)
	client := http.Client{Timeout: time.Second * time.Duration(20)}

	res, err := client.Post(zgRobot, "application/json", bytes.NewBuffer([]byte(reqJson)))
	if err != nil {
		logger.Sugar.Errorf("post error:%s", err.Error())
		return nil, err
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}
	defer res.Body.Close()

	answer := &ZgAnswerJson{}
	err = json.Unmarshal(data, answer)
	if err != nil {
		logger.Sugar.Error(err)
		return nil, err
	}

	return answer, nil
}
