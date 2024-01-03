package jpushclient

import "encoding/json"

type SmsMessage struct {
	DelayTime    int         `json:"delay_time"`              // 单位为秒，不能超过24小时。设置为0，表示立即发送短信。该参数仅对 android 和 iOS 平台有效，Winphone 平台则会立即发送短信。
	SignId       int         `json:"signid,omitempty"`        // 签名ID，该字段为空则使用应用默认签名。
	TempId       int64       `json:"temp_id"`                 // 短信补充的内容模板 ID。没有填写该字段即表示不使用短信补充功能。
	TempPara     interface{} `json:"temp_para,omitempty"`     // 短信模板中的参数。
	ActiveFilter bool        `json:"active_filter,omitempty"` // active_filter 字段用来控制是否对补发短信的用户进行活跃过滤，默认为 true ，做活跃过滤；为 false，则不做活跃过滤；
}

func (m *SmsMessage) SetDelayTime(delayTime int) *SmsMessage {
	m.DelayTime = delayTime
	return m
}

func (m *SmsMessage) SetSignId(signId int) *SmsMessage {
	m.SignId = signId
	return m
}

func (m *SmsMessage) SetTempId(tempId int64) *SmsMessage {
	m.TempId = tempId
	return m
}

func (m *SmsMessage) SetTempPara(t string) *SmsMessage {
	m.TempPara = t
	return m
}

func (m *SmsMessage) SetActiveFilter(t bool) *SmsMessage {
	m.ActiveFilter = t
	return m
}

func (m *SmsMessage) ToJson() (string, error) {
	content, err := json.Marshal(m)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (m *SmsMessage) ToBytes() ([]byte, error) {
	content, err := json.Marshal(m)
	if err != nil {
		return nil, err
	}
	return content, nil
}
