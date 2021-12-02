package aliyun

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
)

type AliSMSrequest struct {
	AliSMScommonRequest *requests.CommonRequest
}

func NewAliSMScommonRequest() *AliSMSrequest {
	return &AliSMSrequest{
		AliSMScommonRequest: requests.NewCommonRequest(),
	}
}

func (r *AliSMSrequest) SetBasicDefaultParam() *AliSMSrequest {
	r.AliSMScommonRequest.Method = "POST"
	r.AliSMScommonRequest.Scheme = "https"
	r.AliSMScommonRequest.Domain = "dysmsapi.aliyuncs.com"
	r.AliSMScommonRequest.Version = "2017-05-25"
	r.AliSMScommonRequest.ApiName = "SendSms"
	return r
}

func (r *AliSMSrequest) SetBizParam(regionID, phones, signName, templateCode,
	templateParam string) *AliSMSrequest {
	r.AliSMScommonRequest.QueryParams["RegionId"] = regionID
	r.AliSMScommonRequest.QueryParams["PhoneNumbers"] = phones
	r.AliSMScommonRequest.QueryParams["SignName"] = signName
	r.AliSMScommonRequest.QueryParams["TemplateCode"] = templateCode
	r.AliSMScommonRequest.QueryParams["TemplateParam"] = templateParam
	return r
}
