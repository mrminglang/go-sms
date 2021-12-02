package main

import (
	"fmt"
	"github.com/mrminglang/go-sms/smsapi"
	"os"
)

func main() {
	client, err := smsapi.NewAliSMSclient("cn-hangzhou---",
		"LTAI*******yrsp1", "Is6FmeqQ***************feK7HJQVea4J7")
	if err != nil {
		fmt.Fprintf(os.Stderr, "NewHandler: %v\n", err)
		return
	}

	aliSMSrequest := smsapi.NewAliSMSrequest("cn-hangzhou",
		"15827088610,17720503271,      13517210601", "ming", "SMS_169101455", "{\"code\":"+"666666"+"}")
	response, err := client.SendSMS(aliSMSrequest)
	if response.IsSuccess() {
		fmt.Printf("%d\n", response.GetHttpStatus())
		// fmt.Printf("%s\n", response.GetHttpContentBytes())
		fmt.Printf("%v\n", response.GetOriginHttpResponse())
		fmt.Printf("%s\n", response.GetHttpContentString())
	} else {
		fmt.Printf("Error: < %v >\n", err)
	}
}

// func main() {
// 	// client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", "阿里云的accessKeyId", "accessKeySecret")
// 	client, err := dysmsapi.NewClientWithAccessKey("cn-beijing", "LTAI*******yrsp1", "Is6FmeqQ***************feK7HJQVea4J7")
// 	request := requests.NewCommonRequest()
// 	request.Method = "POST"15827088610
// 	request.Scheme = "https" // https | http
// 	request.Domain = "dysmsapi.aliyuncs.com"
// 	request.Version = "2017-05-25"
// 	request.ApiName = "SendSms"
// 	request.QueryParams["RegionId"] = "cn-hangzhou"
// 	request.QueryParams["PhoneNumbers"] = "13641337591"                  //手机号
// 	request.QueryParams["SignName"] = "ming"                               //阿里云验证过的项目名 自己设置
// 	request.QueryParams["TemplateCode"] = "SMS_169101455"                //阿里云的短信模板号 自己设置
// 	request.QueryParams["TemplateParam"] = "{\"code\":" + "777777" + "}" //短信模板中的验证码内容 自己生成   之前试过直接返回，但是失败，加上code成功。
// 	response, err := client.ProcessCommonRequest(request)
// 	// fmt.Print(client.DoAction(request, response))
// 	//  fmt.Print(response)
// 	if err != nil {
// 		fmt.Print(err.Error())
// 	}
// 	fmt.Printf("response is %#v\n", response)
// 	//json数据解析
// }
