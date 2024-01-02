package jpushclient

import "encoding/json"

type VOIPNotice struct {
	Key string `json:"key,omitempty"`
}

func (v *VOIPNotice) SetKey(key string) *VOIPNotice {
	v.Key = key
	return v
}

func (v *VOIPNotice) ToJson() (string, error) {
	content, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func (v *VOIPNotice) ToBytes() ([]byte, error) {
	content, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return content, nil
}
