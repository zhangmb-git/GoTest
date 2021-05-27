package app

import (
	"container/list"
	"encoding/csv"
	"github.com/Luxurioust/excelize"
	"gotools/AIPangMao/logger"
	"gotools/AIPangMao/pkg"
	"os"
	"regexp"
	"strings"
)

var excelFd *os.File
var csvWriter *csv.Writer

// 加载数据类知识问题
func loadQuestionsWithData(fileName string) *list.List {
	xlsx, err := excelize.OpenFile(fileName)
	if err != nil {
		panic(err)
		return nil
	}

	knowledge := list.New()

	//for _, sheet := range xlFile.Sheets { // 循环sheet
	rows,err := xlsx.GetRows("sheet1" )
	count := len(rows)
	logger.Sugar.Infof("共 %d 行", count)

	for _, row := range rows {
		// 标准问题 列：11
		// 权限 列：16
		// 标准答案 列：17
		if len(row.Cells) <= 17 {
			continue
		}

		question := row.Cells[11]
		role := row.Cells[15]
		answer := row.Cells[16]
		enableFlag := row.Cells[10] // 禁用标准

		// 禁用类的知识不验证
		if enableFlag.Value == "禁用" {
			continue
		}

		// 去掉标题行+空行
		if question.Value == "" || question.Value == "标准问题" {
			continue
		}

		// string类型
		if question.Type() != xlsx.CellTypeString || answer.Type() != xlsx.CellTypeString ||
			role.Type() != xlsx.CellTypeString {
			continue
		}

		// 正则匹配数据类知识
		// (\$).*({.*})
		// 测试地址：https://tool.lu/regex/，勾选单行
		// 样本1：<HTML>螺纹钢${B_005_002|10}年10合约目前盘面利润为${B_003_001}；<br>${B_005_002|01}年01合约目前盘面利润为${B_003_002}；<br>${B_005_002|05}年05合约目前盘面利润为${B_003_003}。<div></div></HTML>
		// 样本2：${A_048_002|time:$[@时间]}
		// 样本3：${A_043_002|city:$[城市|全国|省份|省名集合|区域]|fct_name:$[钢厂分类|钢厂简介集合]|okw:$[钢厂检修关键词集合|@时间]|prd:$[钢厂检修品类集合]}

		// 对excel中的\n作特殊处理，否则答案匹配不上
		formatValue := strings.ReplaceAll(answer.Value, "\\n", "\n")
		// 没有正则严谨，但是结果也没错
		//if !strings.Contains(formatValue, "${") {
		//	continue
		//}
		match, _ := regexp.MatchString(`(\$).*({.*})`, formatValue)
		if !match {
			continue
		}

		logger.Sugar.Infof("find q=%s,a=%s", question.Value, formatValue)

		//logger.Sugar.Infof("标准问题：%s\t标准答案：%s", question, answer)
		knowledge.PushBack(OneDictionary{
			StandardQuestion: question.Value,
			Role:             role.Value,
			Answer:           formatValue,
		})
	}

	return knowledge
}

/*
func StartTestFromZg(excelFilePath string) {
	knowledgeList := loadQuestionsWithData(excelFilePath)

	err := initExcel()
	if err != nil {
		return
	}
	defer excelFd.Close()

	passedCount := 0
	totalCount := knowledgeList.Len()
	curIndex := 0

	for i := knowledgeList.Front(); i != nil; i = i.Next() {
		curIndex++

		dic, ok := i.Value.(OneDictionary)
		if ok {
			continue
		}
		time.Sleep(100 * time.Microsecond)
		question := dic.StandardQuestion

		answer, err := pkg.GetAnswerFromZG(question)
		if err != nil {
			logger.Sugar.Errorf("标准问题:%s,接口调用错误:%s", question, err.Error())
			continue
		}

		if answer == nil || answer.ResCode != 0 {
			logger.Sugar.Infof("q=%s error:%s", question, answer.ErrorMsg)
			writeExcel(question, answer.ErrorMsg)
			continue
		}

		if len(answer.Answers) < 1 {
			logger.Sugar.Infof("q=%s 没有答案", question)
			writeExcel(question, "没有答案")
			continue
		}

		if curIndex%10 == 0 {
			//logger.Sugar.Infof("测试中，目前通过率=%.2f%s [%d/%d]", math.Ceil((float64(passedCount)*100)/float64(curIndex)), "%", passedCount, curIndex)
		}

		info := answer.Answers[0]
		// 多轮会话
		if info.IsContinuous == 1 {
			a := testMultiRoundSession(info)
			writeExcel2(question, a)
		}

		//if checkAnswer(question.Answer, answer) {
		//	passedCount++
		//	//logger.Sugar.Infof("[%d/%d]测试通过，标准问题:%s", curIndex, totalCount, question.StandardQuestion)
		//} else {
		//	logger.Sugar.Infof("[%d/%d]测试不通过,标准问题:\t%s\t期望答案：%s，实际答案：\t%s", curIndex, totalCount,
		//		question.StandardQuestion, question.Answer, answer)
		//}
	}
	logger.Sugar.Infof("测试完成，通过率=%.2f%s [%d/%d]", math.Ceil((float64(passedCount)*100)/float64(curIndex)), "%", passedCount, curIndex)
}
*/

//func initExcel() error {
//	const fileName = "outpu.csv"
//	_ := os.Remove(fileName)
//	excelFd, _ = os.Create(fileName)
//	csvWriter = csv.NewWriter(excelFd)
//	return nil
//}

func writeExcel(q, a string) {
	err := csvWriter.Write([]string{q, a})
	if err != nil {
		logger.Sugar.Warn(err)
	}
}

func writeExcel2(q string, a []string) {
	tempArr := make([]string, len(a)+1)
	tempArr = append(tempArr, q)
	num := copy(tempArr, a)
	if num <= 0 {
		logger.Sugar.Warn("copy failed")
	}

	err := csvWriter.Write(tempArr)
	if err != nil {
		logger.Sugar.Warn(err)
	}
}

// 递归-多伦会话问题（回答）
//func testMultiRoundSession(question string, json *pkg.ZgAnswerJsonInfo) []string {
//	answer, err := pkg.GetAnswerFromZG(question)
//	if err != nil {
//		logger.Sugar.Errorf("标准问题:%s,接口调用错误:%s", question, err.Error())
//	}
//
//	if answer == nil || answer.ResCode != 0 {
//		logger.Sugar.Infof("q=%s error:%s", question, answer.ErrorMsg)
//		writeExcel(question, answer.ErrorMsg)
//	}
//
//	if len(answer.Answers) < 1 {
//		logger.Sugar.Infof("q=%s 没有答案", question)
//		writeExcel(question, "没有答案")
//	}
//
//	// 追问没有完
//	if json.IsContinuous == 1 {
//		testMultiRoundSession()
//	}
//
//	return nil
//}

// 图标类问题测试
func testChartAnswer(json *pkg.ZgAnswerJsonInfo) []string {
	return nil
}
