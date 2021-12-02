package smsapi

import (
	"github.com/mrminglang/go-sms/service/aliyun"
)

// NewAliSMSclient 通过accessKeyID、accessKeySecret实例化阿里云短信发送服务的客户端
func NewAliSMSclient(regionID, accessKeyID, accessKeySecret string) (*aliyun.AliSMSclient, error) {
	return aliyun.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
}

// NewAliSMSrequest 实例化一个阿里云短信发送请求数据包，call aliyun.NewAliSMScommonRequest
func NewAliSMSrequest(regionID, phones, signName, templateCode, templateParam string) *aliyun.AliSMSrequest {
	return aliyun.NewAliSMScommonRequest().SetBasicDefaultParam().SetBizParam(regionID,
		phones, signName, templateCode, templateParam)
}
