package jpushclient

import (
	"encoding/json"
)

type PushRequest struct {
	Cid             string      `json:"cid,omitempty"`
	Platform        interface{} `json:"platform"`
	Audience        interface{} `json:"audience"`
	Notification    interface{} `json:"notification,omitempty"`
	Message         interface{} `json:"message,omitempty"`
	InAppMessage    bool        `json:"inapp_message,omitempty"`
	SmsMessage      interface{} `json:"sms_message,omitempty"`
	Notification3rd interface{} `json:"notification_3rd,omitempty"`
	CallBack        interface{} `json:"callback,omitempty"`
	Options         *Option     `json:"options,omitempty"`
}

type PushSingleRequest struct {
	PushList PushList
}

type PushList struct {
	Cid struct {
		Platform     Platform    `json:"platform"`
		Target       string      `json:"target"`
		Notification *Notice     `json:"notification,omitempty"`
		Message      *Message    `json:"message,omitempty"`
		SmsMessage   *SmsMessage `json:"sms_message,omitempty"`
		Options      *Option     `json:"options,omitempty"`
	}
}

func NewPushRequest() *PushRequest {
	pl := &PushRequest{}
	o := &Option{}
	o.ApnsProduction = false
	pl.Options = o
	return pl
}
func (this *PushRequest) SetCid(cid string) *PushRequest {
	this.Cid = cid
	return this
}

func (this *PushRequest) SetPlatform(pf *Platform) *PushRequest {
	this.Platform = pf.Os
	return this
}

func (this *PushRequest) SetAudience(ad *Audience) *PushRequest {
	this.Audience = ad.Object
	return this
}

func (this *PushRequest) SetMessage(m *Message) *PushRequest {
	this.Message = m
	return this
}

func (this *PushRequest) SetInAppMessage(b bool) *PushRequest {
	this.InAppMessage = b
	return this
}

func (this *PushRequest) SetSmsMessage(m *SmsMessage) *PushRequest {
	this.SmsMessage = m
	return this
}

func (this *PushRequest) SetNotice(notice *Notice) *PushRequest {
	this.Notification = notice
	return this
}

func (this *PushRequest) SetNotice3rd(notice *Notice3rd) *PushRequest {
	this.Notification3rd = notice
	return this
}

func (this *PushRequest) SetCallback(callback *CallBack) *PushRequest {
	this.CallBack = callback
	return this
}

func (this *PushRequest) SetOptions(o *Option) *PushRequest {
	this.Options = o
	return this
}

func (this *PushRequest) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *PushRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
