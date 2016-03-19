package jpushclient

import (
	"encoding/json"
)

type PayLoad struct {
	Platform     interface{} `json:"platform"`
	Audience     interface{} `json:"audience"`
	Notification interface{} `json:"notification,omitempty"`
	Message      interface{} `json:"message,omitempty"`
	Options      *Option     `json:"options,omitempty"`
}

func NewPushPayLoad() *PayLoad {
	pl := &PayLoad{}
	o := &Option{}
	o.ApnsProduction = false
	pl.Options = o
	return pl
}

func (this *PayLoad) SetPlatform(pf *Platform) {
	this.Platform = pf.Os
}

func (this *PayLoad) SetAudience(ad *Audience) {
	this.Audience = ad.Object
}

func (this *PayLoad) SetOptions(o *Option) {
	this.Options = o
}

func (this *PayLoad) SetMessage(m *Message) {
	this.Message = m
}

func (this *PayLoad) SetNotice(notice *Notice) {
	this.Notification = notice
}

func (this *PayLoad) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
