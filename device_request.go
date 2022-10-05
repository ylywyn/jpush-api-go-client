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

func (this *DeviceSettingRequest) SetTags(tags *DeviceSettingRequestTags) *DeviceSettingRequest {
	this.Tags = tags
	return this
}

func (this *DeviceSettingRequest) SetAlias(alias string) *DeviceSettingRequest {
	this.Alias = alias
	return this
}
func (this *DeviceSettingRequest) SetMobile(mobile string) *DeviceSettingRequest {
	this.Mobile = mobile
	return this
}

func (this *DeviceSettingRequest) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *DeviceSettingRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
