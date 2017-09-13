package jpushclient

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	SUCCESS_FLAG  = "msg_id"
	HOST_NAME_SSL = "https://api.jpush.cn/v3/push"
	HOST_SCHEDULE = "https://api.jpush.cn/v3/schedules"
	HOST_REPORT   = "https://report.jpush.cn/v3/received"
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

func (this *PushClient) Send(data []byte) (string, error) {
	return this.SendPushBytes(data)
}
func (this *PushClient) CreateSchedule(data []byte) (string, error) {
	// this.BaseUrl = HOST_SCHEDULE
	return this.SendScheduleBytes(data, HOST_SCHEDULE)
}
func (this *PushClient) DeleteSchedule(id string) (string, error) {
	// this.BaseUrl = HOST_SCHEDULE
	return this.SendDeleteScheduleRequest(id, HOST_SCHEDULE)
}
func (this *PushClient) GetSchedule(id string) (string, error) {
	// GET https://api.jpush.cn/v3/schedules/{schedule_id}
	// this.BaseUrl = HOST_SCHEDULE
	return this.SendGetScheduleRequest(id, HOST_SCHEDULE)

}
func (this *PushClient) GetReport(msg_ids string) (string, error) {
	// this.BaseUrl = HOST_REPORT
	return this.SendGetReportRequest(msg_ids, HOST_REPORT)
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

func (this *PushClient) SendScheduleBytes(content []byte, url string) (string, error) {
	ret, err := SendPostBytes2(url, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "schedule_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}

}

func (this *PushClient) SendGetReportRequest(msg_ids string, url string) (string, error) {
	return Get(url).SetBasicAuth(this.AppKey, this.MasterSecret).Param("msg_ids", msg_ids).String()
}

func UnmarshalResponse(rsp string) (map[string]interface{}, error) {
	mapRs := map[string]interface{}{}
	if len(strings.TrimSpace(rsp)) == 0 {
		return mapRs, nil
	}
	err := json.Unmarshal([]byte(rsp), &mapRs)
	if err != nil {
		return nil, err
	}
	if _, ok := mapRs["error"]; ok {
		return nil, fmt.Errorf(rsp)
	}
	return mapRs, nil
}

func (this *PushClient) SendDeleteScheduleRequest(schedule_id string, url string) (string, error) {
	rsp, err := Delete(strings.Join([]string{url, schedule_id}, "/")).Header("Authorization", this.AuthCode).String()
	if err != nil {
		return "", err
	}
	_, err = UnmarshalResponse(rsp)
	if err != nil {
		return "", err
	}
	return rsp, nil
}
func (this *PushClient) SendGetScheduleRequest(schedule_id string, url string) (string, error) {
	rsp, err := Get(strings.Join([]string{url, schedule_id}, "/")).Header("Authorization", this.AuthCode).String()
	if err != nil {
		return "", err
	}
	_, err = UnmarshalResponse(rsp)
	if err != nil {
		return "", err
	}
	return rsp, nil
}
