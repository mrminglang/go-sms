package aliyun

import (
	"errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/dysmsapi"
)

type AliSMSclient struct {
	*dysmsapi.Client
}

// NewClientWithAccessKey 通过accessKeyID、accessKeySecret实例化阿里云短信发送服务的客户端
func NewClientWithAccessKey(regionID, accessKeyID,
	accessKeySecret string) (*AliSMSclient, error) {
	if len(regionID) == 0 {
		return nil, errors.New("Aliyun SMS param [RegionId] is not empty.")
	}
	if len(accessKeyID) == 0 {
		return nil, errors.New("Aliyun SMS param [accessKeyId] is not empty.")
	}
	if len(accessKeySecret) == 0 {
		return nil, errors.New("Aliyun SMS param [accessKeySecret] is not empty.")
	}
	client, err := dysmsapi.NewClientWithAccessKey(regionID, accessKeyID, accessKeySecret)
	return &AliSMSclient{
		Client: client,
	}, err
}

// SendSMS 用于在阿里云短信服务提供商下的发送短信服务，事先在阿里云商户后台配置短信签名及短信内容模板是必须的
func (client *AliSMSclient) SendSMS(request *AliSMSrequest) (*AliSMSresponse, error) {
	response, err := client.ProcessCommonRequest(request.AliSMScommonRequest)

	return &AliSMSresponse{
		CommonResponse: response,
	}, err
}
