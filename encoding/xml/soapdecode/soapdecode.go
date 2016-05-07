package main

import (
	"encoding/xml"
	"fmt"
)

type Envelope struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Header  *SoapHeader
	Soap    *SoapBody
}

type SoapHeader struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Header"`
	Theader *CustomHeader
}

type CustomHeader struct {
	XMLName xml.Name `xml:"http://YourCorpClientHere.com/ CustomHeader"`
	Tusername *UserName
	Tpassword *Password
}

type UserName struct {
	XMLName xml.Name `xml:"http://YourCorpClientHere.com/ UserName"`
	Token string     `xml:",chardata"`
}

type Password struct {
	XMLName xml.Name `xml:"http://YourCorpClientHere.com/ Password"`
	Token string     `xml:",chardata"`
}

type SoapBody struct {
	XMLName   xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`
	Tresponse *TokenResponse
}

type TokenResponse struct {
	XMLName xml.Name `xml:"http://YourCorpClientHere.com/ GetTokenResponse"`
	Tresult *TokenResult
}

type TokenResult struct {
	XMLName xml.Name `xml:"http://YourCorpClientHere.com/ GetTokenResult"`
	Token   string   `xml:",chardata"`
}

var data = `<?xml version="1.0" encoding="utf-8"?>
<soap:Envelope xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/"
	xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
	xmlns:xsd="http://www.w3.org/2001/XMLSchema">
	<soap:Header>
		<CustomHeader xmlns="http://YourCorpClientHere.com/">
			<UserName>someName</UserName>
			<Password>somePassword</Password>
		</CustomHeader>
	</soap:Header>
	<soap:Body>
		<GetTokenResponse xmlns="http://YourCorpClientHere.com/">
			<GetTokenResult>TestToken</GetTokenResult>
		</GetTokenResponse>
	</soap:Body>
</soap:Envelope>`

func main() {
	info := new(Envelope)
	fmt.Println(data)
	err := xml.Unmarshal([]byte(data), info)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}

	fmt.Println(info)
	fmt.Println(info.Header)
	fmt.Println(info.Header.Theader)
	fmt.Println(info.Header.Theader.Tusername)
	fmt.Println(info.Header.Theader.Tusername.Token)
	fmt.Println(info.Header.Theader.Tpassword)
	fmt.Println(info.Header.Theader.Tpassword.Token)

	fmt.Println(info.Soap)
	fmt.Println(info.Soap.Tresponse)
	fmt.Println(info.Soap.Tresponse.Tresult)
	fmt.Println(info.Soap.Tresponse.Tresult.Token)
}
