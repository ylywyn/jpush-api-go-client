package jpushclient

import "encoding/json"

type CallBack struct {
	Url    string                 `json:"url,omitempty"`
	Type   string                 `json:"type,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
}

func (call *CallBack) SetUrl(c string) *CallBack {
	call.Url = c
	return call
}

func (call *CallBack) SetType(t string) *CallBack {
	call.Type = t
	return call
}

func (call *CallBack) AddParams(key string, value interface{}) *CallBack {
	if call.Params == nil {
		call.Params = make(map[string]interface{})
	}
	call.Params[key] = value
	return call
}

func (call *CallBack) ToJson() (string, error) {
	content, err := json.Marshal(call)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (call *CallBack) ToBytes() ([]byte, error) {
	content, err := json.Marshal(call)
	if err != nil {
		return nil, err
	}
	return content, nil
}
