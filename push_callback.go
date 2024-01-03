package jpushclient

import "encoding/json"

type CallBack struct {
	Url    string         `json:"url,omitempty"`    // 数据临时回调地址，指定后以此处指定为准，仅针对这一次推送请求生效；不指定，则以极光后台配置为准
	Type   string         `json:"type,omitempty"`   // 需要回调给用户的自定义参数
	Params map[string]any `json:"params,omitempty"` // 回调数据类型，1:送达回执, 2:点击回执, 3:送达和点击回执, 8:推送成功回执, 9:成功和送达回执, 10:成功和点击回执, 11:成功和送达以及点击回执
}

func (call *CallBack) SetUrl(c string) *CallBack {
	call.Url = c
	return call
}

func (call *CallBack) SetType(t string) *CallBack {
	call.Type = t
	return call
}

func (call *CallBack) AddParams(key string, value any) *CallBack {
	if call.Params == nil {
		call.Params = make(map[string]any)
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
