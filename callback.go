package jpushclient

type CallBack struct {
	Url    string                 `json:"url,omitempty"`
	Type   string                 `json:"type,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
}

func (this *CallBack) SetUrl(c string) {
	this.Url = c
}

func (this *CallBack) SetType(t string) {
	this.Type = t
}

func (this *CallBack) AddParams(key string, value interface{}) {
	if this.Params == nil {
		this.Params = make(map[string]interface{})
	}
	this.Params[key] = value
}
