package utils

import (
	"context"
	"fmt"
	"github.com/go-pay/gopay"
	"github.com/go-pay/gopay/alipay"
	"github.com/go-pay/gopay/pkg/xlog"
	"io"
	"os"
	"testing"
)

var (
	AppId                = "9021000133661127"
	AlipayPublicCertPath = "/Users/abing/Develop/GO/GOPATH/src/github.com/computeshare-server/configs/alipay/alipayPublicCert.crt"
	AlipayRootCertPath   = "/Users/abing/Develop/GO/GOPATH/src/github.com/computeshare-server/configs/alipay/alipayRootCert.crt"
	AppPublicCertPath    = "/Users/abing/Develop/GO/GOPATH/src/github.com/computeshare-server/configs/alipay/appPublicCert.crt"
	AppPrivateKey        = "MIIEpQIBAAKCAQEAvC/YFxa9wBTqtknQudbunVcqPswA8IjztVk1vu8gg3/3LYkVLVV+5mz3lOdMzvdgOUWir2aDEIftqWLkewrEPjNl1oxrDHsgzHOIadvD/urHEHBomPLa1mL5jDnvWoOt3imtQ7/68Zqh6YeeaNebJuJezrtftnRpHeIJU/pwv+58pOC4slprRwZiHD8xhMo0Ac4k+BEmFIvpQKolsJzhwqMXSM/27RQQBLD0AOuP+Xig9xqsVCFKjGVy0TksPh2N3+Vg9gFLEeWXp46IoLuSyKmWgszbkizQXYJYEd1tpMAoWAT8+Twbd3JrHFJRuceMdKNhH5MrYVVzr+lVP/5S6wIDAQABAoIBAQCafYSe39KC4CkX4V1zU1ms3aQWzY9v6F56YHRY/SLqEaGwRjuWbfux238dRQ2vUzIxklNOP5AVgCHBzyfXQy18CuHt0nUHWAXDEo3IqdwraD5n5oBYH/c7zSEk12MCIdwY6IoPdoni0pldiDKpy2bZ8zj16qlWthadq7UQtibdyNii7nXL30c6pLFZtfruohXN1CKpGbTHabWaaNmXAeUcnlwIsB/x4wjU2IyP4x/OkxiKH1F5IP4ONJWiv0XJVu0AoLuutZrTkombmu81G8IguFBZc10frR2v26VK85eVYHGn2ghfjpn9rhOxsvvufMDNsEhCD0Dljfag/K/0dnlBAoGBAOgyGbXs8PL3aEvPQ1HVGFEE3nwtE163SiB5tZnpg6rPqNWIu43Vt7V28qxbAfaka5BqvzUqvw29w8zxJwO+Y27Py1eEjj6EDp5VWiN0PoJMemulmPdhXBfWEigLdF1K+QoediNBe9tI0ERuKjwQm99RtZOKTNHF9xRVzQoqqPJ7AoGBAM96vzQ1BkGrGUuvh3T3fbdATYk7Es0PLDwAOP9OpIY31L/sReSEwZ+iNNjiivXGlivYTuMCM948u5wAHiVzZmKX3RLRxw1iymqL17/a6Nm5wMqbaGjVoNZjBbGinhO5ch9WO3n3SYRdmxypHoNNLpxY8r4IHujc0XP/fUUrDq5RAoGBAK1g+6HEkwZZhx9UXEg2fjnBEQBYCCapZkhpmRicLRzWgpMvzhxSgSKW9w2LxyOMEmV7z9q7WH3CPPpC94bsGwMbVhK7gBteMCw6P1xJ6IMS6DaJKHdP05xY1KXiJ7mhyeOsazpYI0vSSvFN8Wt4W2udb6ADnOjmkWA10SzOPekRAoGAGqGKwUy9DNjFWm2hWWYbANaEHUUz2JZF+z2cE0ko3QF1kVzma7qOj8rmNaB+baWbyOzu0zyJI9Fv00pAlFYHokgT3EKI5jL0AFEK1Flo4elx0Y+5CxvAlUQIixBwcd+vwk3zPGh9SzwomLksbXYOsYrJPZMVRJZzXAxDLAmmzMECgYEA05LDgU7PoNY4EYa5BULzHzXBlqaXsXHoYWs2QAQwC1fuCnevGC3LS4AQiO/bsCUT6QO7kjsynkJvdcJPtUwFnRrNa6H/QKHfJ04DaDhSD1tWN8RToa+TM9asCROxL4YfJ9VdqJcBuhdDDU9lX3r47bUxQK89iumWbsf8cEatkYc="
)

func TestClient(t *testing.T) {
	alipayPublicCert, _ := os.Open(AlipayPublicCertPath)
	alipayRootCert, _ := os.Open(AlipayRootCertPath)
	appPublicCert, _ := os.Open(AppPublicCertPath)
	alipayPublicCertContent, _ := io.ReadAll(alipayPublicCert)
	alipayRootContent, _ := io.ReadAll(alipayRootCert)
	appPublicContent, _ := io.ReadAll(appPublicCert)
	// 初始化支付宝客户端
	// appid：应用ID
	// privateKey：应用私钥，支持PKCS1和PKCS8
	// isProd：是否是正式环境，沙箱环境请选择新版沙箱应用。
	client, err := alipay.NewClient(AppId, AppPrivateKey, false)
	if err != nil {
		xlog.Error(err)
		return
	}
	// 打开Debug开关，输出日志，默认关闭
	client.DebugSwitch = gopay.DebugOn
	// 设置支付宝请求 公共参数
	//    注意：具体设置哪些参数，根据不同的方法而不同，此处列举出所有设置参数
	client.SetLocation(alipay.LocationShanghai). // 设置时区，不设置或出错均为默认服务器时间
							SetCharset(alipay.UTF8).                                         // 设置字符编码，不设置默认 utf-8
							SetSignType(alipay.RSA2).                                        // 设置签名类型，不设置默认 RSA2
							SetReturnUrl("http://localhost:9999").                           // 设置返回URL，付款结束后跳转的url
							SetNotifyUrl("https://34c099660e.yicp.fun/v1/alipay/pay/notify") // 设置异步通知URL

	// 自动同步验签（只支持证书模式）
	// 传入 alipayCertPublicKey_RSA2.crt 内容
	client.AutoVerifySign(alipayPublicCertContent)

	// 公钥证书模式，需要传入证书，以下两种方式二选一
	// 证书路径
	//err := client.SetCertSnByPath("appCertPublicKey.crt", "alipayRootCert.crt", "alipayCertPublicKey_RSA2.crt")
	// 证书内容
	err = client.SetCertSnByContent(appPublicContent, alipayRootContent, alipayPublicCertContent)
	if err != nil {
		panic(err)
	}
	// 初始化 BodyMap
	bm := make(gopay.BodyMap)
	bm.Set("subject", "共享算力积分购买"). // 标题
					Set("out_trade_no", "abing_order_id6666671"). // 订单号，支付成功后会返回
					Set("total_amount", "2").                     // 订单金额
					Set("timeout_express", "5m").                 // 支付超时时间
					Set("product_code", "FAST_INSTANT_TRADE_PAY") // 必填 具体参考文档

	aliRsp, err := client.TradePagePay(context.Background(), bm)
	if err != nil {
		xlog.Error("err:", err)
		return
	}
	fmt.Println(aliRsp)
}

func TestNotify(t *testing.T) {
	// 解析异步通知的参数
	// req：*http.Request
	//notifyReq, err := alipay.ParseNotifyToBodyMap(c.Request) // c.Request 是 gin 框架的写法
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//// value：url.Values
	//notifyReq, err = alipay.ParseNotifyByURLValues()
	//if err != nil {
	//	xlog.Error(err)
	//	return
	//}
	//ok, err = alipay.VerifySignWithCert("alipayPublicCert.crt content", notifyReq)
}
