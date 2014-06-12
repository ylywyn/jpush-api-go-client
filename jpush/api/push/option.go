package push

type Option struct {
	Sendno          int  `json:"sendno"`
	Timelive        int  `json:"time_to_live"`
	Apns_production bool `json:"apns_production"`
	//override_msg_id int64 `json:"override_msg_id"`
}

func (this *Option) SetSendno(no int) {
	this.Sendno = no
}

func (this *Option) SetTimelive(timelive int) {
	this.Timelive = timelive
}

//func (this *Option) SetOverrideMsgId(id int64) {
//	this.override_msg_id = id
//}

func (this *Option) SetApns(apns bool) {
	this.Apns_production = apns
}
