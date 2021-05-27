package main

import (
	"gotools/AIPangMao/app"
	"gotools/AIPangMao/logger"
	"runtime"
)

//const ProfessionalExcel = "excel/找钢网生产环境(专业知识点-实例)(20191128134824).xlsx"
const ProfessionalExcel = "diff_test/找钢网生产环境(自动问答明细)/处理数据/汇总统计.xlsx"
const BaseConversationExcel = "./excel/找钢网生产环境(专业知识点-实例)(20191128134824).xlsx"
const excelPath = "D:\\test\\src\\gotools\\AIPangMao\\excel\\找钢网生产环境(专业知识点-实例)(20191128134824).xlsx"

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	logger.InitLogger("log/info.log", "DEBUG")
	app.RunLoop()
	app.LoadConfig() //要提前加载配置
	app.InitCSV()

	const testType = 1
	switch testType {
	case 1:
		//app.StartTestFromY(BaseConversationExcel)

		listQuestion := app.LoadQuestionsWithData()
		//log.Print(listQuestion)
		app.AskAllQuestions(listQuestion)

		//path := "D:\\a.csv" //app.TOPIC["path"]
		//row := []string{"1", "2", "3", "4", "5,6"}
		//app.WriteAskLog(1234, "abc", "abc", 1)
		logger.Sugar.Info("启动测试")
		break
	case 2:
		logger.Sugar.Info("启动测试，脚本自动从excel读取用户问题，分别从云问和小i处获取答案对比，输出不一样的答案")
		//app.StartTestFromYAndI(ProfessionalExcel)
		break
	case 3:
		logger.Sugar.Info("启动测试，脚本自动从excel读取数据类问题测试，包括追问和图标类答案")
		//app.StartTestFromZg(BaseConversationExcel)
		break
	}
	//time.Sleep(time.Second * 10)
	logger.Sugar.Info("测试结束")
}
