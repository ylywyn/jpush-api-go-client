package jpushclient

import "encoding/json"

type AdminRequest struct {
	AppName        string `json:"app_name,string"`
	AndroidPackage string `json:"android_package,string"`
	GroupName      string `json:"group_name,string"`
}

func (a *AdminRequest) SetAppName(t string) *AdminRequest {
	a.AppName = t
	return a
}

func (a *AdminRequest) SetAndroidPackage(t string) *AdminRequest {
	a.AndroidPackage = t
	return a
}

func (a *AdminRequest) SetGroupName(t string) *AdminRequest {
	a.GroupName = t
	return a
}

func (a *AdminRequest) ToJson() (string, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (a *AdminRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(a)
	if err != nil {
		return nil, err
	}
	return content, nil
}
