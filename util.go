package wxpay

import (
	"bytes"
	"crypto/tls"
	"encoding/pem"
	"encoding/xml"
	"log"
	"strconv"
	"strings"
	"time"

	"golang.org/x/crypto/pkcs12"
)

func XmlToMap(xmlStr string) Params {

	params := make(Params)
	decoder := xml.NewDecoder(strings.NewReader(xmlStr))

	var (
		key   string
		value string
	)

	inBetween := false
	for t, err := decoder.Token(); err == nil; t, err = decoder.Token() {
		switch token := t.(type) {
		case xml.StartElement:
			key = token.Name.Local
			if key == "xml" {
				continue
			}
			inBetween = true
		case xml.CharData:
			if !inBetween {
				continue
			}
			value = string(token)
			params[key] = value
		case xml.EndElement:
			inBetween = false
		}

	}

	return params
}

func MapToXml(params Params) string {
	var buf bytes.Buffer
	buf.WriteString(`<xml>`)
	for k, v := range params {
		buf.WriteString(`<`)
		buf.WriteString(k)
		buf.WriteString(`><![CDATA[`)
		buf.WriteString(v)
		buf.WriteString(`]]></`)
		buf.WriteString(k)
		buf.WriteString(`>`)
	}
	buf.WriteString(`</xml>`)

	return buf.String()
}

// 用时间戳生成随机字符串
func nonceStr() string {
	return strconv.FormatInt(time.Now().UTC().UnixNano(), 10)
}

// 将Pkcs12转成Pem
func pkcs12ToPem(p12 []byte, password string) tls.Certificate {

	blocks, err := pkcs12.ToPEM(p12, password)

	defer func() {
		if x := recover(); x != nil {
			log.Print(x)
		}
	}()

	if err != nil {
		panic(err)
	}

	var pemData []byte
	for _, b := range blocks {
		pemData = append(pemData, pem.EncodeToMemory(b)...)
	}

	cert, err := tls.X509KeyPair(pemData, pemData)
	if err != nil {
		panic(err)
	}
	return cert
}
