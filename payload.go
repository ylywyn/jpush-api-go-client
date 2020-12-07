package jpushclient

import (
	"encoding/json"
)

type PushRequest struct {
	Cid          string      `json:"cid,omitempty"`
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	SmsMessage   interface{} `json:"sms_message,omitempty"`
	Options      *Option     `json:"options,omitempty"`
}

func NewPushRequest() *PushRequest {
	pl := &PushRequest{}
	o := &Option{}
	o.ApnsProduction = false
	pl.Options = o
	return pl
}
func (this *PushRequest) SetCid(cid string) {
	this.Cid = cid
}
func (this *PushRequest) SetPlatform(pf *Platform) {
	this.Platform = pf.Os
}

func (this *PushRequest) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *PushRequest) SetOptions(o *Option) {
	this.Options = o
}

func (this *PushRequest) SetMessage(m *Message) {
	this.Message = m
}
func (this *PushRequest) SetSmsMessage(m *SmsMessage) {
	this.SmsMessage = m
}
func (this *PushRequest) SetNotice(notice *Notice) {
	this.Notification = notice
}

func (this *PushRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
