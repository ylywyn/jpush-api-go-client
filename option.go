package jpushclient

import "encoding/json"

type Option struct {
	SendNo            int                    `json:"sendno,omitempty"`
	TimeToLive        int                    `json:"time_to_live,omitempty"`
	ApnsProduction    bool                   `json:"apns_production"`
	OverrideMsgId     int64                  `json:"override_msg_id,omitempty"`
	BigPushDuration   int                    `json:"big_push_duration,omitempty"`
	ApnsCollapseId    string                 `json:"apns_collapse_id,omitempty"`
	ThirdPartyChannel map[string]interface{} `json:"third_party_channel,omitempty"`
}

func (this *Option) SetSendNo(no int) *Option {
	this.SendNo = no
	return this
}

func (this *Option) SetTimeToLive(timeTolive int) *Option {
	this.TimeToLive = timeTolive
	return this
}

func (this *Option) SetOverrideMsgId(id int64) *Option {
	this.OverrideMsgId = id
	return this
}

func (this *Option) SetApns(apns bool) *Option {
	this.ApnsProduction = apns
	return this
}

func (this *Option) SetBigPushDuration(bigPushDuration int) *Option {
	this.BigPushDuration = bigPushDuration
	return this
}

func (this *Option) SetApnsId(apnsCollapseId string) *Option {
	this.ApnsCollapseId = apnsCollapseId
	return this
}

func (this *Option) SetThirdPartyChannel(thirdPartyChannel map[string]interface{}) *Option {
	this.ThirdPartyChannel = thirdPartyChannel
	return this
}

func (this *Option) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *Option) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
