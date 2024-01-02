package jpushclient

import "encoding/json"

type DeviceSettingRequest struct {
	Tags   *DeviceSettingRequestTags `json:"tags"`
	Alias  string                    `json:"alias"`
	Mobile string                    `json:"mobile"`
}
type DeviceSettingEmptyTagsRequest struct {
	Tags   string `json:"tags"`
	Alias  string `json:"alias"`
	Mobile string `json:"mobile"`
}
type DeviceSettingRequestAlias struct {
	Remove []string `json:"remove,omitempty"`
}
type DeviceSettingRequestTags struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}
type DeviceBindTagsRequest struct {
	Add    []string `json:"add,omitempty"`
	Remove []string `json:"remove,omitempty"`
}

func (d *DeviceSettingRequest) SetTags(tags *DeviceSettingRequestTags) *DeviceSettingRequest {
	d.Tags = tags
	return d
}

func (d *DeviceSettingRequest) SetAlias(alias string) *DeviceSettingRequest {
	d.Alias = alias
	return d
}
func (d *DeviceSettingRequest) SetMobile(mobile string) *DeviceSettingRequest {
	d.Mobile = mobile
	return d
}

func (d *DeviceSettingRequest) ToJson() (string, error) {
	content, err := json.Marshal(d)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (d *DeviceSettingRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(d)
	if err != nil {
		return nil, err
	}
	return content, nil
}
