package jpushclient

import "encoding/json"

type ReportStatusRequest struct {
	MsgId           int      `json:"msg_id"`
	RegistrationIds []string `json:"registration_ids"`
	Date            string   `json:"date,omitempty"`
}

func (report *ReportStatusRequest) SetMsgId(id int) *ReportStatusRequest {
	report.MsgId = id
	return report
}

func (report *ReportStatusRequest) SetRegistrationIds(ids []string) *ReportStatusRequest {
	report.RegistrationIds = ids
	return report
}

func (report *ReportStatusRequest) SetDate(d string) *ReportStatusRequest {
	report.Date = d
	return report
}

func (report *ReportStatusRequest) ToJson() (string, error) {
	content, err := json.Marshal(report)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (report *ReportStatusRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(report)
	if err != nil {
		return nil, err
	}
	return content, nil
}
