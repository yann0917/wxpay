package wxpay

import (
	"fmt"
	"os"
	"testing"
)

func TestXmlToMap(t *testing.T) {
	xmlStr := "<xml><return_code><![CDATA[SUCCESS]]></return_code><return_msg><![CDATA[OK]]></return_msg><appid><![CDATA[wx2421b1c4370ec43b]]></appid><mch_id><![CDATA[10000100]]></mch_id><nonce_str><![CDATA[IITRi8Iabbblz1Jc]]></nonce_str><sign><![CDATA[7921E432F65EB8ED0CE9755F0E86D72F]]></sign><result_code><![CDATA[SUCCESS]]></result_code><prepay_id><![CDATA[wx201411101639507cbf6ffd8b0779950874]]></prepay_id><trade_type><![CDATA[APP]]></trade_type></xml>"
	params := XmlToMap(xmlStr)
	if params == nil {
		t.Error(params)
	}
	t.Log(params)
}

func TestMapToXml(t *testing.T) {
	params := map[string]string{"return_msg": "OK", "appid": "wx2421b1c4370ec43b", "mch_id": "10000100"}
	xmlStr := MapToXml(params)
	t.Log(xmlStr)
}

func TestNonceStr(t *testing.T) {
	t.Log(nonceStr())
}

func TestPkcs12ToPem(t *testing.T) {
	file, err := os.ReadFile("./apiclient_cert.p12")

	if err != nil {
		fmt.Println(err)
		return
	}
	cert, _ := pkcs12ToPem(file, "1540260411")
	t.Log(string(cert.Certificate[0]))
	t.Log(cert.PrivateKey)
}
