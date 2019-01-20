package sms

import (
	"encoding/json"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

type AlibabaSendRequest struct {
	RegionId          string
	AccessKeyId       string
	AccessSecret      string
	AlibabaSendClient *sdk.Client
}
type AlibabaSendResponse struct {
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	BizId     string `json:"BizId"`
	Code      string `json:"Code"`
}

func NewAlibabaSendRequest(accessKeyId, accessSecret string, regionIdArr ...string) *AlibabaSendRequest {
	var regionId = "default"
	if len(regionIdArr) > 0 {
		regionId = regionIdArr[0]
	}
	return &AlibabaSendRequest{
		RegionId:     regionId,
		AccessKeyId:  accessKeyId,
		AccessSecret: accessSecret,
	}
}

func (m *AlibabaSendRequest) AlibabaSendSms(signName, templateCode, phoneNumber string, templateParam map[string]interface{}) (aibabaSendResponse *AlibabaSendResponse, err error) {
	templateParamJson, err := json.Marshal(templateParam)
	if err != nil {
		return nil, err
	}
	return m.AlibabaSend("SendSms", signName, templateCode, phoneNumber, string(templateParamJson))
}
func (m *AlibabaSendRequest) AlibabaSendBatchSms(templateCode string, signName, phoneNumber []string, templateParam []map[string]interface{}) (aibabaSendResponse *AlibabaSendResponse, err error) {
	signNameJson, err := json.Marshal(signName)
	if err != nil {
		return nil, err
	}
	phoneNumberJson, err := json.Marshal(phoneNumber)
	if err != nil {
		return nil, err
	}
	templateParamJson, err := json.Marshal(templateParam)
	if err != nil {
		return nil, err
	}
	return m.AlibabaSend("SendBatchSms", string(signNameJson), templateCode, string(phoneNumberJson), string(templateParamJson))
}
func (m *AlibabaSendRequest) AlibabaSend(apiName, signName, templateCode, phoneNumber, templateParam string) (aibabaSendResponse *AlibabaSendResponse, err error) {
	if m.AlibabaSendClient == nil {
		client, err := sdk.NewClientWithAccessKey(m.RegionId, m.AccessKeyId, m.AccessSecret)
		if err != nil {
			return nil, err
		}
		m.AlibabaSendClient = client
	}
	request := requests.NewCommonRequest()
	request.Method = "POST"
	request.Domain = "dysmsapi.aliyuncs.com"
	request.Version = "2017-05-25"
	request.ApiName = apiName
	request.QueryParams["TemplateCode"] = templateCode
	if apiName == "SendSms" {
		request.QueryParams["PhoneNumbers"] = phoneNumber    //18010460987  支持逗号分割 上线为1000个手机号
		request.QueryParams["SignName"] = signName           // 然宇信息
		request.QueryParams["TemplateParam"] = templateParam //{name:\"nanjishidu\"}
	} else if apiName == "SendBatchSms" {
		request.QueryParams["PhoneNumberJson"] = phoneNumber     //["18010460987","18010460987"]
		request.QueryParams["SignNameJson"] = signName           //[{name:\"然宇信息\"},{name:\"然宇信息\"}]
		request.QueryParams["TemplateParamJson"] = templateParam //"[{name:\"nanjishidu\"},{name:\"nanjishidu\"}]"
	} else {

	}
	var response *responses.CommonResponse
	response, err = m.AlibabaSendClient.ProcessCommonRequest(request)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response.GetHttpContentBytes(), &aibabaSendResponse)
	return aibabaSendResponse, err
}
