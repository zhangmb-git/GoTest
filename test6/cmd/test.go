package cmd

import (
	"flag"
	"fmt"
	"github.com/e421083458/golang_common/log"
	"os"
	"path/filepath"
	"sync"
	"time"
	"encoding/json"
)

var (
	hostname string
	once     sync.Once
)



func Test() {
	var1 := fmt.Sprintf("http://%d", 1)
	var2, var3 := os.Hostname()
	fmt.Println(var1)
	fmt.Println(var2)
	fmt.Println(var3)

	flag.StringVar(&hostname, "hostname", hostname, "machine hostname")
	fmt.Println("hostname:" + hostname)

}

func Test1() {
	var name string
	var age int

	// 这里参数需要传入指针类型，获取写入数据
	if _, err := fmt.Scan(&name, &age); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("name: %s, age: %d", name, age)

}

//flag
func test3() {

	//var ip = flag.Int("flagname", 1234, "help message for flagname")
	//fmt.Println("ip has value ", *ip)

	//r := filepath.FromSlash("/a/b/c/d.txt")

	m, _ := filepath.Glob("D://work/*")
	fmt.Println(m)

}

func onces() {
	fmt.Println("onces")
}
func onced() {
	fmt.Println("onced")
}

func test4() {
	for i, v := range make([]string, 10) {
		once.Do(onces)
		fmt.Println("count:", v, "---", i)
	}
	for i := 0; i < 10; i++ {

		//go func() {
		once.Do(onced)
		fmt.Println("213")
		//}()
	}
	time.Sleep(3000)
}

type Study  struct{
	Name  string
}
func Test5(){
	//v:="C0:com.internet.ic.client.model.param.item.ItemBatchUpdateDTO\221\nitemDOList`q\023java.util.ArrayListC0/com.internet.ic.client.model.domain.item.ItemDO\310O\023categoryAliasDOList\017threeCategoryDO\020depotStoreDOList\020itemPropertyList\vitemFeature\bitemTags\ritemSkuDOList\017itemLogisticsDO\voperateTime\aendDate\tstartDate\vgmtModified\ngmtCreated\apicUrls\016suportCampaign\017businessPattern\aoutCode\noperateUid\tcreateUid\fsupplierName\006tpCode\016minOrderAmount\ramountPerUnit\vsubItemType\023transportTemplateId\022operateValidStatus\bhashCode\017categoryAliasId\vwarehouseId\bshopType\006shopId\ngroupPrice\aisgroup\vaddressName\bareaCode\bcityCode\fprovinceCode\borderNum\abrandId\nsellerType\004code\004days\blatitude\tlongitude\005sales\006domain\aoutType\005outId\tisPublish\visCanDelete\nisCanClose\fisCanPublish\aoptions\006status\aversion\apayType\abarCode\fdepositPrice\tcostPrice\roriginalPoint\016originalCredit\roriginalPrice\005point\006credit\005price\tdetailUrl\006source\frealStockNum\bstockNum\vdescription\aoneWord\bsubTitle\005title\bitemType\005spuId\bsellerId\ncategoryId\016rootCategoryId\002idaNNNw\220C05com.internet.ic.client.model.param.item.ItemSkuPVPair\226\005vType\004vTxt\003vId\005pType\004pTxt\003pIdb\220\004范德萨发\340\220\002品名\370(b\220\005Q345B\340\220\002材质\357b\220\005Q325B\340\220\002规格\370\020b\220\005上海范德萨\340\220\002钢厂\370\021b\220\001吨\370\177\220\004计价单位\370&b\220\003协议品\340\220\004质量等级\370\026b\220\n2019-07-30\340\220\004出厂日期\370$NNNNJ\000\000\001q\265\030\276\260NNJ\000\000\001v\034u\t\350K\001\216\213lH\001\066\060\245|o_w-1679157ff6c94d3a98d3a338789e2a18.gif|o_w-8fee3aa3b0ee40339293df7ba4443a94.jpg|o_w-80d8d23bbb7540a6aaf7429dc4024ecd.jpg|o_w-0f370062978740e7ac6ff10a828f6b33.jpg|Z\221\221N\340\340NN\340\340\221\340\222Ip\210v\367\340\341\232\354\340\220N\006\063\061\060\061\061\063\006\063\061\060\061\060\060\006\063\061\060\060\060\060\220\340\340N\220[[\310A\313\350\221<)\006TFTF\341\222<:\023\220N\340\340\340\340\340\340\340=\206\253N\220\220\311\023NN\v@-（）*&……%#@\025弹簧钢 Q345B Q325B 上海范德萨\313 \340<%\334Y\000\234E\017Y\000\234D\261<)\006"

	res := &Study{}
	test,err := json.Marshal(res)
	if err !=nil {
		log.Fatal("error")
	}

	fmt.Println(test)
	json.Unmarshal([]byte(test), &res)
	fmt.Println(res)

}
