package jpushclient

type Option struct {
	SendNo            int                               `json:"sendno,omitempty"`
	TimeLive          int                               `json:"time_to_live,omitempty"`
	ApnsProduction    bool                              `json:"apns_production"`
	OverrideMsgId     int64                             `json:"override_msg_id,omitempty"`
	BigPushDuration   int                               `json:"big_push_duration,omitempty"`
	ThirdPartyChannel map[string]map[string]interface{} `json:"third_party_channel,omitempty"`
}

func (this *Option) SetSendno(no int) {
	this.SendNo = no
}

func (this *Option) SetTimelive(timelive int) {
	this.TimeLive = timelive
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
