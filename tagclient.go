package jpushclient

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

const (
	// 查询设备的别名与标签
	TAG_HOST_NAME_SSL = "https://device.jpush.cn/v3/"
)

type TagClient struct {
	MasterSecret string
	AppKey       string
	AuthCode     string
	BaseUrl      string
}

func NewTagClient(appKey, secret string) *TagClient {
	auth := "Basic " + Base64Coder.EncodeToString([]byte(appKey+":"+secret))
	pusher := &TagClient{secret, appKey, auth, TAG_HOST_NAME_SSL}
	return pusher
}

func (this *TagClient) HttpGet(url string) (string, error) {
	rsp, err := Get(url).Header("Authorization", this.AuthCode).String()
	if err != nil {
		return "", err
	}
	_, err = UnmarshalResponse(rsp)
	if err != nil {
		return "", err
	}
	return rsp, nil
}

// GetAllTag 获取全部tag
func (this *TagClient) GetAllTag() (string, error) {
	return this.HttpGet(fmt.Sprintf(this.BaseUrl + "tags"))
}

// GetTag 读取当前设备绑定的全部tag
func (this *TagClient) GetTag(registrationId string) (string, error) {
	return this.HttpGet(fmt.Sprintf(this.BaseUrl+"devices/%s", registrationId))
}

// MvRegIdToTag 添加指定设备到标签 add/remove 最多各支持 1000 个
func (this *TagClient) MvRegIdToTag(tag string, addRegIds, rmRegIds []string) (string, error) {
	type reqRegIds struct {
		Add    []string `json:"add,omitempty"`
		Remove []string `json:"remove,omitempty"`
	}

	type req struct {
		RegistrationIds reqRegIds `json:"registration_ids"`
	}

	if len(addRegIds) == 0 && len(rmRegIds) == 0 {
		return "", nil
	}

	postData := req{
		RegistrationIds: reqRegIds{
			Add:    addRegIds,
			Remove: rmRegIds,
		},
	}

	data, err := json.Marshal(postData)
	if err != nil {
		return "", nil
	}
	return this.SendPushBytes(this.BaseUrl+"tags/"+tag, data)
}

// SendPushBytes 推送
func (this *TagClient) SendPushBytes(url string, content []byte) (string, error) {
	ret, err := SendPostBytes2(url, content, this.AuthCode)
	if err != nil {
		return ret, err
	}
	if strings.Contains(ret, "msg_id") {
		return ret, nil
	} else {
		return "", errors.New(ret)
	}
}
