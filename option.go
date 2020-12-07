package jpushclient

type Option struct {
	SendNo            int                    `json:"sendno,omitempty"`
	TimeToLive        int                    `json:"time_to_live,omitempty"`
	ApnsProduction    bool                   `json:"apns_production"`
	OverrideMsgId     int64                  `json:"override_msg_id,omitempty"`
	BigPushDuration   int                    `json:"big_push_duration,omitempty"`
	ApnsCollapseId    string                 `json:"apns_collapse_id,omitempty"`
	ThirdPartyChannel map[string]interface{} `json:"third_party_channel,omitempty"`
}

func (this *Option) SetSendNo(no int) {
	this.SendNo = no
}

func (this *Option) SetTimeToLive(timeTolive int) {
	this.TimeToLive = timeTolive
}

func (this *Option) SetOverrideMsgId(id int64) {
	this.OverrideMsgId = id
}

func (this *Option) SetApns(apns bool) {
	this.ApnsProduction = apns
}

func (this *Option) SetBigPushDuration(bigPushDuration int) {
	this.BigPushDuration = bigPushDuration
}

func (this *Option) SetApnsId(apnsCollapseId string) {
	this.ApnsCollapseId = apnsCollapseId
}

func (this *Option) SetThirdPartyChannel(thirdPartyChannel map[string]interface{}) {
	this.ThirdPartyChannel = thirdPartyChannel
}
