package cashier

import (
	"fmt"
	conf "github.com/opay-services/opay-sdk-golang/sdk/conf"
	"math/rand"
	"testing"
	"time"
)

var mConf *conf.OpayMerchantConf
func init() {
	mConf = conf.InitEnv(
		"OPAYPUB16058646510220.420473668870203",
		"OPAYPRV16058646510230.34019403186305675",
		"SrnIchuukX33koDt",
		"256620112018025",
		"sandbox",
		"NGN",
	)

	mConf.SetLog(func(a ...interface{}) {
		fmt.Println(a...)
	})
	rand.Seed(time.Now().Unix())
}

func TestApiCashierInitialize(t *testing.T) {
	req := CashierInitializeReq{}
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.MchShortName = "lijian_test"
	req.ProductName = "goods"
	req.ProductDesc = "test goods"
	req.UserPhone = "+2349876543210"
	req.UserRequestIp = "123.123.123.123"
	req.Amount = "100"
	req.Currency = "NGN"
	req.PayTypes = []string{"BalancePayment", "BonusPayment", "OWealth"}
	req.PayMethods = []string{"account", "qrcode", "bankCard", "bankAccount"}
	req.ExpireAt = "10"
	req.CallbackUrl = "http://localhost:8080"
	req.ReturnUrl = "http://localhost:8080"
	rsp, err := ApiCashierInitializeReq(&req, mConf)
	fmt.Println(rsp, err)
	//fmt.Println(fmt.Sprintf("%+v and err :%+v", resp, err))
}

func TestApiCashierStatusReq(t *testing.T) {
	req := CashierStatusReq{Reference:"testlijian_1616638639862310000"}
	rsp, err :=ApiCashierStatusReq(&req, mConf)
	fmt.Println(rsp, err)

}

func TestApiCashierCloseReq(t *testing.T) {
	//only init status order can be closed
	req := CashierCloseReq{Reference:"testlijian_1616589761615889000"}
	rsp, err :=ApiCashierCloseReq(&req, mConf)
	fmt.Println(rsp, err)

}

func TestApiCashierRefundByBankAccountReq(t *testing.T) {
	req := CashierRefundByBankAccountReq{}
	req.OriginalReference = "testlijian_1616595268369294000"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Amount = "1"
	req.Currency = "NGN"
	req.RefundType = "refund2bankaccount"
	req.Reason = "test"
	req.BankCode = "033"
	req.Country = "NG"
	req.BankAccountNumber = "2070043686"

	rsp, err :=ApiCashierRefundByBankAccountReq(&req, mConf)
	fmt.Println(rsp, err)
}

func TestApiCashierRefundByOpayMerchantAccountReq(t *testing.T) {
	req := CashierRefundByOpayMerchantAccountReq{}
	req.OriginalReference = "testlijian_1616595268369294000"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Amount = "1"
	req.Currency = "NGN"
	req.RefundType = "refund2opayaccount"
	req.Reason = "test"
	req.Country = "NG"
	req.Receiver = MerchantReceiver{
		MerchantId:"256620112018024",
		Type:"MERCHANT",
	}
	rsp, err :=ApiCashierRefundByOpayMerchantAccountReq(&req, mConf)
	fmt.Println(rsp, err)
}

func TestApiCashierRefundByOpayUserAccountReq(t *testing.T) {
	req := CashierRefundByOpayUserAccountReq{}
	req.OriginalReference = "testlijian_1616595268369294000"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Amount = "1"
	req.Currency = "NGN"
	req.RefundType = "refund2opayaccount"
	req.Reason = "test"
	req.Country = "NG"
	req.Receiver = UserReceiver{
		PhoneNumber:"+2348160564736",
		Type:"USER",
	}

	rsp, err :=ApiCashierRefundByOpayUserAccountReq(&req, mConf)
	fmt.Println(rsp, err)
}

func TestApiCashierRefundByOriginReq(t *testing.T) {
	req := CashierRefundByOriginReq{}
	req.OriginalReference = "testlijian_1616595268369294000"
	req.Reference = fmt.Sprintf("testlijian_%v",time.Now().UnixNano())
	req.Amount = "1"
	req.Currency = "NGN"
	req.RefundType = "refundoriginal"
	req.Reason = "test"
	req.Country = "NG"

	rsp, err :=ApiCashierRefundByOriginReq(&req, mConf)
	fmt.Println(rsp, err)
}

func TestApiCashierRefundStatusReq(t *testing.T) {
	req := CashierRefundStatusReq{OrderNo:"210324250539103459"}
	rsp, err :=ApiCashierRefundStatusReq(&req, mConf)
	fmt.Println(rsp, err)
}