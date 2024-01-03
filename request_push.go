package jpushclient

import (
	"encoding/json"
)

type PushRequest struct {
	Cid             string  `json:"cid,omitempty"`
	Platform        any     `json:"platform"`
	Audience        any     `json:"audience"`
	Notification    any     `json:"notification,omitempty"`
	Message         any     `json:"message,omitempty"`
	InAppMessage    bool    `json:"inapp_message,omitempty"`
	SmsMessage      any     `json:"sms_message,omitempty"`
	Notification3rd any     `json:"notification_3rd,omitempty"`
	LiveActivity    any     `json:"live_Activity,omitempty"`
	CallBack        any     `json:"callback,omitempty"`
	Options         *Option `json:"options,omitempty"`
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
func (push *PushRequest) SetCid(cid string) *PushRequest {
	push.Cid = cid
	return push
}

func (push *PushRequest) SetPlatform(pf *Platform) *PushRequest {
	push.Platform = pf.Os
	return push
}

func (push *PushRequest) SetAudience(ad *Audience) *PushRequest {
	push.Audience = ad.Object
	return push
}

func (push *PushRequest) SetMessage(m *Message) *PushRequest {
	push.Message = m
	return push
}

func (push *PushRequest) SetInAppMessage(b bool) *PushRequest {
	push.InAppMessage = b
	return push
}

func (push *PushRequest) SetSmsMessage(m *SmsMessage) *PushRequest {
	push.SmsMessage = m
	return push
}

func (push *PushRequest) SetNotice(notice *Notice) *PushRequest {
	push.Notification = notice
	return push
}

func (push *PushRequest) SetNotice3rd(notice *Notice3rd) *PushRequest {
	push.Notification3rd = notice
	return push
}
func (push *PushRequest) SetLiveActivity(liveActivity *LiveActivity) *PushRequest {
	push.LiveActivity = liveActivity
	return push
}
func (push *PushRequest) SetCallback(callback *CallBack) *PushRequest {
	push.CallBack = callback
	return push
}

func (push *PushRequest) SetOptions(o *Option) *PushRequest {
	push.Options = o
	return push
}

func (push *PushRequest) ToJson() (string, error) {
	content, err := json.Marshal(push)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (push *PushRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(push)
	if err != nil {
		return nil, err
	}
	return content, nil
}
