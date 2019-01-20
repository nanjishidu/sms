package sms

import "testing"

var (
	accessKeyId   = ""
	accessSecret  = ""
	signName      = ""
	templateCode  = ""
	phoneNumber   = ""
	templateParam = map[string]interface{}{"name": "nanjishidu"}
)

func TestAlibabaSendSms(t *testing.T) {
	m := NewAlibabaSendRequest(accessKeyId, accessSecret)
	alibabaSendResponse, err := m.AlibabaSendSms(signName, templateCode, phoneNumber, templateParam)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(alibabaSendResponse)
	if alibabaSendResponse.Code != "OK" {
		t.FailNow()
	}
}
