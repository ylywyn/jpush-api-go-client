package jpushclient

type SmsMessage struct {
	Content      string `json:"content"`
	DelayTime    int    `json:"delay_time,int,omitempty"`
	SignId       int    `json:"signid,int,omitempty"`
	TempId       int64  `json:"temp_id,long,omitempty"`
	TempPara     string `json:"temp_para,long,omitempty"`
	ActiveFilter bool   `json:"active_filter,bool,omitempty"`
}

func (this *SmsMessage) SetContent(c string) {
	this.Content = c
}

func (this *SmsMessage) SetDelayTime(delayTime int) {
	this.DelayTime = delayTime
}

func (this *SmsMessage) SetSignId(signId int) {
	this.SignId = signId
}

func (this *SmsMessage) SetTempId(tempId int64) {
	this.TempId = tempId
}

func (this *SmsMessage) SetTempPara(t string) {
	this.TempPara = t
}

func (this *SmsMessage) SetActiveFilter(t bool) {
	this.ActiveFilter = t
}
