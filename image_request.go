package jpushclient

import "encoding/json"

type ImageRequest struct {
	ImageType       int    `json:"image_type"`
	ImageUrl        string `json:"image_url,omitempty"`
	JiGuangImageUrl string `json:"jiguang_image_url,omitempty"`
	XiaomiImageUrl  string `json:"xiaomi_image_url,omitempty"`
	OppoImageUrl    string `json:"oppo_image_url,omitempty"`
	HuaweiImageUrl  string `json:"huawei_image_url,omitempty"`
	FcmImageUrl     string `json:"fcm_image_url,omitempty"`
}

func (this *ImageRequest) SetImageType(c int) *ImageRequest {
	this.ImageType = c
	return this
}

func (this *ImageRequest) SetImageUrl(t string) *ImageRequest {
	this.ImageUrl = t
	return this
}

func (this *ImageRequest) setJiGuangImageUrl(t string) *ImageRequest {
	this.JiGuangImageUrl = t
	return this
}

func (this *ImageRequest) SetXiaomiImageUrl(t string) *ImageRequest {
	this.XiaomiImageUrl = t
	return this
}

func (this *ImageRequest) SetOppoImageUrl(t string) *ImageRequest {
	this.OppoImageUrl = t
	return this
}

func (this *ImageRequest) SetHuaweiImageUrl(t string) *ImageRequest {
	this.HuaweiImageUrl = t
	return this
}

func (this *ImageRequest) SetFcmImageUrl(t string) *ImageRequest {
	this.FcmImageUrl = t
	return this
}

func (this *ImageRequest) ToJson() (string, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (this *ImageRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(this)
	if err != nil {
		return nil, err
	}
	return content, nil
}
