package smsapi_test

import (
	"encoding/json"
	"github.com/mrminglang/go-sms/smsapi"
	"strings"
	"testing"
)

func TestNewAliSMSclient(t *testing.T) {
	client, err := smsapi.NewAliSMSclient("cn-hangzhou", "LT************yrsp1", "Is6FmeqQ***************feK7HJQVea4J7")
	if client == nil {
		t.Errorf("NewAliSMSclient is <nil> want not empty\n")
	}
	if err != nil {
		t.Errorf("NewHandler: %v\n", err)
	}
}

func TestSendAliSMS_ErrTemplateCode(t *testing.T) {
	client, err := smsapi.NewAliSMSclient("cn-hangzhou",
		"LT************yrsp1", "Is6FmeqQ***************feK7HJQVea4J7")
	if err != nil {
		t.Errorf("NewHandler: %v\n", err)
		return
	}
	aliSMSrequest := smsapi.NewAliSMSrequest("cn-hangzhou",
		"15827088610,17720503271,13517210601", "ming", "ERR_SMS_169101455", "{\"code\":"+"666666"+"}")
	response, err := client.SendSMS(aliSMSrequest)
	if err != nil {
		t.Errorf("SendAliSMS: %v\n", err)
	}
	if response.IsSuccess() {
		// t.Logf("%d\n", response.GetHttpStatus())
		// fmt.Printf("%s\n", response.GetHttpContentBytes())
		// t.Logf("%v\n", response.GetOriginHttpResponse())
		// t.Logf("%s\n", response.GetHttpContentString())
		type content struct {
			Message   string
			RequestId string
			Code      string
		}
		c := &content{}
		err := json.Unmarshal([]byte(response.GetHttpContentString()), c)
		if err != nil {
			t.Errorf("%v\n", err.Error())
		}
		if c.Message != "模板不合法(不存在或被拉黑)" {
			t.Fatalf("got <%s>,  want <模板不合法(不存在或被拉黑)>\n", c.Message)
		}
		if c.Code != "isv.SMS_TEMPLATE_ILLEGAL" {
			t.Fatalf("got <%s>,  want <isv.SMS_TEMPLATE_ILLEGAL>\n", c.Code)
		}
	} else {
		t.Errorf("%v\n", err)
	}
}

func TestSendAliSMS_ErrKey(t *testing.T) {
	client, err := smsapi.NewAliSMSclient("cn-hangzhou",
		"LT************yrsp1-", "Is6FmeqQ***************feK7HJQVea4J7-")
	if err != nil {
		t.Errorf("NewHandler: %v\n", err)
		return
	}
	aliSMSrequest := smsapi.NewAliSMSrequest("cn-hangzhou",
		"15827088610,17720503271,13517210601", "ming", "ERR_SMS_169101455", "{\"code\":"+"666666"+"}")
	response, err := client.SendSMS(aliSMSrequest)
	if err != nil {
		if !strings.Contains(err.Error(), "ErrorCode: InvalidAccessKeyId.NotFound") {
			t.Errorf("SendAliSMS: %s\n", err.Error())
		}
	}
	if response.IsSuccess() {
		// t.Logf("%d\n", response.GetHttpStatus())
		// fmt.Printf("%s\n", response.GetHttpContentBytes())
		// t.Logf("%v\n", response.GetOriginHttpResponse())
		// t.Logf("%s\n", response.GetHttpContentString())
		type content struct {
			Message   string
			RequestId string
			Code      string
		}
		c := &content{}
		err := json.Unmarshal([]byte(response.GetHttpContentString()), c)
		if err != nil {
			t.Errorf("%v\n", err)
		}
		if c.Message != "模板不合法(不存在或被拉黑)" {
			t.Fatalf("got <%s>,  want <模板不合法(不存在或被拉黑)>\n", c.Message)
		}
		if c.Code != "isv.SMS_TEMPLATE_ILLEGAL" {
			t.Fatalf("got <%s>,  want <isv.SMS_TEMPLATE_ILLEGAL>\n", c.Code)
		}
	} else {
		if !strings.Contains(err.Error(), "ErrorCode: InvalidAccessKeyId.NotFound") {
			t.Errorf("SendAliSMS: %s\n", err.Error())
		}
	}
}
