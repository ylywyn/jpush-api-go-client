package jpushclient

import "encoding/json"

type SmsMessage struct {
	DelayTime    int    `json:"delay_time,int"`
	SignId       int    `json:"signid,int,omitempty"`
	TempId       int64  `json:"temp_id,long,omitempty"`
	TempPara     string `json:"temp_para,long,omitempty"`
	ActiveFilter bool   `json:"active_filter,bool,omitempty"`
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
