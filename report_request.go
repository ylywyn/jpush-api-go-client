package jpushclient

import "encoding/json"

type ReportStatusRequest struct {
	MsgId           int      `json:"msg_id,int"`
	RegistrationIds []string `json:"registration_ids"`
	Date            string   `json:"date,omitempty"`
}

func (this *ReportStatusRequest) SetMsgId(id int) *ReportStatusRequest {
	this.MsgId = id
	return this
}

func (this *ReportStatusRequest) SetRegistrationIds(ids []string) *ReportStatusRequest {
	this.RegistrationIds = ids
	return this
}

func (this *ReportStatusRequest) SetDate(d string) *ReportStatusRequest {
	this.Date = d
	return this
}

func (this *ReportStatusRequest) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *ReportStatusRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
