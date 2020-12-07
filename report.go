package jpushclient

type ReportStatusRequest struct {
	MsgId           int      `json:"msg_id,int"`
	RegistrationIds []string `json:"registration_ids"`
	Date            string   `json:"date,omitempty"`
}
