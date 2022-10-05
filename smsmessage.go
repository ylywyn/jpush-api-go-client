package jpushclient

import "encoding/json"

type SmsMessage struct {
	Content      string `json:"content"`
	DelayTime    int    `json:"delay_time,int,omitempty"`
	SignId       int    `json:"signid,int,omitempty"`
	TempId       int64  `json:"temp_id,long,omitempty"`
	TempPara     string `json:"temp_para,long,omitempty"`
	ActiveFilter bool   `json:"active_filter,bool,omitempty"`
}

func (this *SmsMessage) SetContent(c string) *SmsMessage {
	this.Content = c
	return this
}

func (this *SmsMessage) SetDelayTime(delayTime int) *SmsMessage {
	this.DelayTime = delayTime
	return this
}

func (this *SmsMessage) SetSignId(signId int) *SmsMessage {
	this.SignId = signId
	return this
}

func (this *SmsMessage) SetTempId(tempId int64) *SmsMessage {
	this.TempId = tempId
	return this
}

func (this *SmsMessage) SetTempPara(t string) *SmsMessage {
	this.TempPara = t
	return this
}

func (this *SmsMessage) SetActiveFilter(t bool) *SmsMessage {
	this.ActiveFilter = t
	return this
}

func (this *SmsMessage) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *SmsMessage) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
