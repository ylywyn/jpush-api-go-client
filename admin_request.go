package jpushclient

import "encoding/json"

type AdminRequest struct {
	AppName        string `json:"app_name,string"`
	AndroidPackage string `json:"android_package,string"`
	GroupName      string `json:"group_name,string"`
}

func (this *AdminRequest) SetAppName(t string) *AdminRequest {
	this.AppName = t
	return this
}

func (this *AdminRequest) SetAndroidPackage(t string) *AdminRequest {
	this.AndroidPackage = t
	return this
}

func (this *AdminRequest) SetGroupName(t string) *AdminRequest {
	this.GroupName = t
	return this
}

func (this *AdminRequest) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *AdminRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
