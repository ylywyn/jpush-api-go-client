package jpushclient

import "encoding/json"

type Message struct {
	MsgContent  string         `json:"msg_content"`            // 消息内容本身
	Title       string         `json:"title,omitempty"`        // 消息标题
	ContentType string         `json:"content_type,omitempty"` // 消息内容类型
	Extras      map[string]any `json:"extras,omitempty"`       // JSON 格式的可选参数
}

func (msg *Message) SetMsgContent(c string) *Message {
	msg.MsgContent = c
	return msg
}

func (msg *Message) SetTitle(title string) *Message {
	msg.Title = title
	return msg
}

func (msg *Message) SetContentType(t string) *Message {
	msg.ContentType = t
	return msg
}

func (msg *Message) AddExtras(key string, value any) *Message {
	if msg.Extras == nil {
		msg.Extras = make(map[string]any)
	}
	msg.Extras[key] = value
	return msg
}

func (msg *Message) ToJson() (string, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (msg *Message) ToBytes() ([]byte, error) {
	content, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return content, nil
}
