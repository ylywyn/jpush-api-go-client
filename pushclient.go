package jpushclient

import (
	"encoding/base64"
	"errors"
	"strings"
)

const (
	SUCCESS_FLAG  = "msg_id"
	HOST_NAME_SSL = "https://api.jpush.cn/v3/push"
	HOST_CREATE_SCHEDULE = "https://api.jpush.cn/v3/schedules"
	HOST_REPORT = "https://report.jpush.cn/v3/received"
	BASE64_TABLE  = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
)

var base64Coder = base64.NewEncoding(BASE64_TABLE)

type PushClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewPushClient(secret, appKey string) *PushClient {
	//base64
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{secret, appKey, auth, HOST_NAME_SSL}
	return pusher
}

func NewSchedulePushClient(secret, appKey string) *PushClient{
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{MasterSecret:secret, AppKey:appKey, AuthCode:auth}
	return pusher
}
func NewReportClient(secret, appKey string) *PushClient{
	auth := "Basic " + base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &PushClient{MasterSecret:secret, AppKey:appKey, AuthCode:auth}
	return pusher
}


func (this *PushClient) Send(data []byte) (string, error) {
	return this.SendPushBytes(data)
}
func (this *PushClient) CreateSchedule(data []byte) (string, error) {
	this.BaseUrl = HOST_CREATE_SCHEDULE
	return this.SendScheduleBytes(data)
}
func (this *PushClient) GetReport(msg_ids string) (string, error) {
	this.BaseUrl = HOST_REPORT
	return this.SendGetReportRequest(msg_ids)
}
func (this *PushClient) SendPushString(content string) (string, error) {
	ret, err := SendPostString(this.BaseUrl, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}

func (this *PushClient) SendPushBytes(content []byte) (string, error) {
	//ret, err := SendPostBytes(this.BaseUrl, content, this.AuthCode)
	ret, err := SendPostBytes2(this.BaseUrl, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}


func (this *PushClient) SendScheduleBytes(content []byte) (string, error) {
	ret, err := SendPostBytes2(this.BaseUrl, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "schedule_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}

}

func (this *PushClient) SendGetReportRequest(msg_ids string) (string, error){
	return Get(this.BaseUrl).SetBasicAuth(this.AppKey, this.MasterSecret).Param("msg_ids", msg_ids).String()

}
