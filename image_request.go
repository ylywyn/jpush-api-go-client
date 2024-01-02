package jpushclient

import "encoding/json"

type ImageRequest struct {
	ImageType       int    `json:"image_type"`
	ImageUrl        string `json:"image_url,omitempty"`
	JiGuangImageUrl string `json:"jiguang_image_url,omitempty"`
	XiaomiImageUrl  string `json:"xiaomi_image_url,omitempty"`
	OppoImageUrl    string `json:"oppo_image_url,omitempty"`
	HuaweiImageUrl  string `json:"huawei_image_url,omitempty"`
	HonorImageUrl   string `json:"honor_image_url,omitempty"`
	FcmImageUrl     string `json:"fcm_image_url,omitempty"`
}

func (img *ImageRequest) SetImageType(c int) *ImageRequest {
	img.ImageType = c
	return img
}

func (img *ImageRequest) SetImageUrl(t string) *ImageRequest {
	img.ImageUrl = t
	return img
}

func (img *ImageRequest) setJiGuangImageUrl(t string) *ImageRequest {
	img.JiGuangImageUrl = t
	return img
}

func (img *ImageRequest) SetXiaomiImageUrl(t string) *ImageRequest {
	img.XiaomiImageUrl = t
	return img
}

func (img *ImageRequest) SetOppoImageUrl(t string) *ImageRequest {
	img.OppoImageUrl = t
	return img
}

func (img *ImageRequest) SetHuaweiImageUrl(t string) *ImageRequest {
	img.HuaweiImageUrl = t
	return img
}

func (img *ImageRequest) SetHonorImageUrl(t string) *ImageRequest {
	img.HonorImageUrl = t
	return img
}

func (img *ImageRequest) SetFcmImageUrl(t string) *ImageRequest {
	img.FcmImageUrl = t
	return img
}

func (img *ImageRequest) ToJson() (string, error) {
	content, err := json.Marshal(img)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (img *ImageRequest) ToBytes() ([]byte, error) {
	content, err := json.Marshal(img)
	if err != nil {
		return nil, err
	}
	return content, nil
}
