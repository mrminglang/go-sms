package smsapi_test

import (
	"fmt"
	"github.com/mrminglang/go-sms/smsapi"
	"os"
)

func ExampleNewHandler() {
	client, err := smsapi.NewAliSMSclient("cn-hangzhou",
		"LT************yrsp1", "Is6FmeqQ***************feK7HJQVea4J7")
	if client != nil {
		fmt.Printf("%v\n", client)
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewHandler: %v\n", err)
	}
}

func ExampleSendAliSMS() {
	client, err := smsapi.NewAliSMSclient("cn-hangzhou",
		"LT************yrsp1", "Is6FmeqQ***************feK7HJQVea4J7")
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewHandler: %v\n", err)
		return
	}
	aliSMSrequest := smsapi.NewAliSMSrequest("cn-hangzhou",
		"15827088610,17720503271,13517210601", "ming", "ERR_SMS_169101455", "{\"code\":"+"666666"+"}")
	response, err := client.SendSMS(aliSMSrequest)
	if err != nil {
		fmt.Fprintf(os.Stderr, "SendAliSMS: %v\n", err)
	}
	if response.IsSuccess() {
		fmt.Printf("%d\n", response.GetHttpStatus())
		// fmt.Printf("%s\n", response.GetHttpContentBytes())
		fmt.Printf("%v\n", response.GetOriginHttpResponse())
		fmt.Printf("%s\n", response.GetHttpContentString())
	} else {
		fmt.Printf("%v\n", err)
	}
}
