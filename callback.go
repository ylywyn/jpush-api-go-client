package jpushclient

import "encoding/json"

type CallBack struct {
	Url    string                 `json:"url,omitempty"`
	Type   string                 `json:"type,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
}

func (this *CallBack) SetUrl(c string) *CallBack {
	this.Url = c
	return this
}

func (this *CallBack) SetType(t string) *CallBack {
	this.Type = t
	return this
}

func (this *CallBack) AddParams(key string, value interface{}) *CallBack {
	if this.Params == nil {
		this.Params = make(map[string]interface{})
	}
	this.Params[key] = value
	return this
}

func (this *CallBack) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *CallBack) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
