package jpushclient

import "encoding/json"

type Option struct {
	SendNo            int                    `json:"sendno,omitempty"`              //推送序号
	TimeToLive        int                    `json:"time_to_live,omitempty"`        //离线消息保留时长(秒)
	OverrideMsgId     int64                  `json:"override_msg_id,omitempty"`     //要覆盖的消息 ID
	ApnsProduction    bool                   `json:"apns_production,omitempty"`     //APNs 是否生产环境
	ApnsCollapseId    string                 `json:"apns_collapse_id,omitempty"`    //更新 iOS 通知的标识符
	BigPushDuration   int                    `json:"big_push_duration,omitempty"`   //定速推送时长(分钟)
	ThirdPartyChannel map[string]interface{} `json:"third_party_channel,omitempty"` //推送请求下发通道
	Classification    int                    `json:"classification,omitempty"`
	TargetEvent       []string               `json:"target_event,omitempty"`
}

func (op *Option) SetSendNo(no int) *Option {
	op.SendNo = no
	return op
}

func (op *Option) SetTimeToLive(timeToLive int) *Option {
	op.TimeToLive = timeToLive
	return op
}

func (op *Option) SetOverrideMsgId(id int64) *Option {
	op.OverrideMsgId = id
	return op
}

func (op *Option) SetApns(apns bool) *Option {
	op.ApnsProduction = apns
	return op
}

func (op *Option) SetBigPushDuration(bigPushDuration int) *Option {
	op.BigPushDuration = bigPushDuration
	return op
}

func (op *Option) SetApnsId(apnsCollapseId string) *Option {
	op.ApnsCollapseId = apnsCollapseId
	return op
}

func (op *Option) AddThirdPartyChannel(key string, value interface{}) *Option {
	if op.ThirdPartyChannel == nil {
		op.ThirdPartyChannel = make(map[string]interface{})
	}
	op.ThirdPartyChannel[key] = value
	return op
}

func (op *Option) SetClassification(classification int) *Option {
	op.Classification = classification
	return op
}

func (op *Option) SetTargetEvent(targetEvent []string) *Option {
	op.TargetEvent = targetEvent
	return op
}

func (op *Option) ToJson() (string, error) {
	content, err := json.Marshal(op)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (op *Option) ToBytes() ([]byte, error) {
	content, err := json.Marshal(op)
	if err != nil {
		return nil, err
	}
	return content, nil
}
