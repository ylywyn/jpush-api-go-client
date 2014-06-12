package push

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"jpush/api/common"
	"strings"
)

const (
	SUCCESS_FLAG  = "msg_id"
	HOST_NAME_SSL = "https://api.jpush.cn/v3/push"
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

func (this *PushClient) Send(builder interface{}) (string, error) {
	content, err := json.Marshal(builder)
	if err != nil {
		return "", err
	}

	return this.SendPushBytes(content)
}

func (this *PushClient) SendPushString(content string) (string, error) {
	ret, err := common.SendPostString(this.BaseUrl, content, this.AuthCode)
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
	ret, err := common.SendPostBytes(this.BaseUrl, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}
