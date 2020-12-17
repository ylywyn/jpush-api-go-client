package jpushclient

import "encoding/json"

type Message struct {
	MsgContent  string                 `json:"msg_content"`
	Title       string                 `json:"title,omitempty"`
	ContentType string                 `json:"content_type,omitempty"`
	Extras      map[string]interface{} `json:"extras,omitempty"`
}

func (this *Message) SetMsgContent(c string) *Message {
	this.MsgContent = c
	return this
}

func (this *Message) SetTitle(title string) *Message {
	this.Title = title
	return this
}

func (this *Message) SetContentType(t string) *Message {
	this.ContentType = t
	return this
}

func (this *Message) AddExtras(key string, value interface{}) *Message {
	if this.Extras == nil {
		this.Extras = make(map[string]interface{})
	}
	this.Extras[key] = value
	return this
}

func (this *Message) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *Message) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
