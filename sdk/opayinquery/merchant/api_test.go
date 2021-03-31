package merchant

import (
	"fmt"
	"github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"testing"
	"time"
)

func init() {
	conf.InitEnv(
		"OPAYPUB16058646510220.420473668870203",
		"OPAYPRV16058646510230.34019403186305675",
		"SrnIchuukX33koDt",
		"256620112018025",
		"sandbox",
	)

	conf.SetLog(func(a ...interface{}) {
		fmt.Println(a...)
	})
	rand.Seed(time.Now().Unix())
}

func TestApiInfoMerchantReq(t *testing.T) {
	req := InfoMerchantReq{Email:"lijian2233@163.com"}
	resp, err := ApiInfoMerchantReq(req)
	fmt.Println(resp, err)
}
